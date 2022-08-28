package import_export

import (
	"bufio"
	"gd-blog/src/repoimpl/model"
	"gopkg.in/yaml.v2"
	"io"
	"strings"
	"time"
)

type HexoHeader struct {
	Title      string
	Categories []string
	Tags       []string
	Date       time.Time
}

func Import(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	// 拆分得到结构化的头部和博客内容
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
	hexoHeader := HexoHeader{}
	err := yaml.Unmarshal([]byte(header), &hexoHeader)
	if err != nil {
		return err
	}
	// 入库: 创建分类、tag、博客，根据标题去重
	return nil
}

func Export(blog model.Blog) (io.Reader, error) {
	return nil, nil
}
