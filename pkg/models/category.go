package models

type Category struct {
	BaseModel
	Name        string `json:"name" gorm:"uniqueIndex" `
	Description string `json:"description"`
}

type CategoryWithCount struct {
	Category
	BlogCount int64 `json:"blog_count"`
}

func AllCategory() ([]CategoryWithCount, error) {
	var rs []map[string]interface{}
	db.Model(&Category{}).Select("id, name, (select count(1) from blog_categories where category_id = categories.id) as blog_count").Order("blog_count desc").Find(&rs)
	if db.Error != nil {
		return nil, db.Error
	}
	var categories []CategoryWithCount
	for _, r := range rs {
		categories = append(categories, CategoryWithCount{
			Category: Category{
				BaseModel: BaseModel{
					ID: r["id"].(uint),
				},
				Name: r["name"].(string),
			},
			BlogCount: r["blog_count"].(int64),
		})
	}
	return categories, nil
}
