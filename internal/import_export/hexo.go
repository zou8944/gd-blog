package import_export

import (
	"bufio"
	"gd-blog/pkg/models"
	"gopkg.in/errgo.v2/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"log"
	"reflect"
	"strings"
	"time"
)

type HexoHeader struct {
	Title      string
	Categories []string
	Tags       []string
	Date       time.Time
	Updated    time.Time
}

func readBlog(r io.Reader) (string, string) {
	scanner := bufio.NewScanner(r)
	var header string
	var content string
	inHeader := true
	for scanner.Scan() {
		line := scanner.Text()
		if inHeader {
			if strings.Trim(line, " ") == "---" {
				if len(header) > 0 {
					inHeader = false
				}
				continue
			}
			header = header + line + "\n"
		} else {
			content = content + line + "\n"
		}
	}
	content = strings.Trim(content, "\n")
	return header, content
}

func parseAsHexoHeader(headerString string) (*HexoHeader, error) {
	hexoHeader := &HexoHeader{}
	scanner := bufio.NewScanner(strings.NewReader(headerString))
	mod := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")
		switch strings.Split(line, ":")[0] {
		case "title":
			hexoHeader.Title = strings.Trim(strings.Split(line, "title:")[1], " ")
		case "date":
			t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(strings.Split(line, "date:")[1], " "))
			if err != nil {
				return nil, err
			}
			hexoHeader.Date = t
		case "updated":
			t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(strings.Split(line, "updated:")[1], " "))
			if err != nil {
				return nil, err
			}
			hexoHeader.Updated = t
		case "categories":
			value := strings.Trim(strings.Split(line, "categories:")[1], " ")
			if len(value) == 0 {
				mod = 1
			} else {
				hexoHeader.Categories = append(hexoHeader.Categories, value)
			}
		case "tags":
			value := strings.Trim(strings.Split(line, "tags:")[1], " ")
			if len(value) == 0 {
				mod = 2
			} else {
				hexoHeader.Tags = append(hexoHeader.Tags, value)
			}
		default:
			value := strings.Trim(strings.TrimPrefix(line, "-"), " ")
			if value == "null" || len(value) == 0 {
				continue
			}
			switch mod {
			case 1:
				hexoHeader.Categories = append(hexoHeader.Categories, value)
			case 2:
				hexoHeader.Tags = append(hexoHeader.Tags, value)
			}
		}
	}
	return hexoHeader, nil
}

func Import(db *gorm.DB, r io.Reader) error {
	header, content := readBlog(r)
	hexoHeader, err := parseAsHexoHeader(header)
	if err != nil {
		return err
	}

	var blog models.Blog
	db.Where(&models.Blog{Title: hexoHeader.Title}).First(&blog)
	if !reflect.DeepEqual(blog, models.Blog{}) {
		log.Printf("blog with title '%s' already exist, ignore\n", hexoHeader.Title)
		return nil
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		blog = models.Blog{
			BaseModel: models.BaseModel{
				CreatedAt: hexoHeader.Date,
				UpdatedAt: hexoHeader.Updated,
			},
			Title:        hexoHeader.Title,
			Content:      content,
			LikeCount:    0,
			CollectCount: 0,
		}
		tx.Omit(clause.Associations).Create(&blog)

		// 单独处理many to many的bug: 关联数据已存在时，会插入脏数据
		var categories []models.Category
		for _, cname := range hexoHeader.Categories {
			categories = append(categories, models.Category{
				Name:        cname,
				Description: "",
			})
		}
		tx.Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "name"}},
				DoUpdates: clause.Assignments(map[string]interface{}{"updated_at": gorm.Expr("CURRENT_TIMESTAMP")}),
			},
			clause.Returning{},
		).Create(&categories)

		var tags []models.Tag
		for _, tname := range hexoHeader.Tags {
			tags = append(tags, models.Tag{
				Name: tname,
			})
		}
		tx.Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "name"}},
				DoUpdates: clause.Assignments(map[string]interface{}{"updated_at": gorm.Expr("CURRENT_TIMESTAMP")}),
			},
			clause.Returning{},
		).Create(&tags)

		err = tx.Model(&blog).Association("Categories").Append(categories)
		if err != nil {
			return err
		}
		return tx.Model(&blog).Association("Tags").Append(tags)
	})

	if err != nil {
		log.Fatalln(err)
	}
	if db.Error != nil {
		log.Fatalln(errors.Wrap(db.Error))
	}
	return nil
}

func Export(blog models.Blog) (io.Reader, error) {
	return nil, nil
}
