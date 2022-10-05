package models

type Tag struct {
	BaseModel
	Name string `json:"name" gorm:"uniqueIndex" json:"name"`
}

type TagWithCount struct {
	Tag
	BlogCount int64 `json:"blog_count"`
}

func AllTag() ([]TagWithCount, error) {
	var rs []map[string]interface{}
	db.Model(&Tag{}).Select("id, name, (select count(1) from blog_tags where tag_id = tags.id) as blog_count").Order("blog_count desc").Find(&rs)
	if db.Error != nil {
		return nil, db.Error
	}
	var tags []TagWithCount
	for _, r := range rs {
		tags = append(tags, TagWithCount{
			Tag: Tag{
				BaseModel: BaseModel{
					ID: r["id"].(uint),
				},
				Name: r["name"].(string),
			},
			BlogCount: r["blog_count"].(int64),
		})
	}
	return tags, nil
}
