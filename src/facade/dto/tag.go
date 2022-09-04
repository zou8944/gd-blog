package dto

type Tag struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	BlogCount int64  `json:"blog_count"`
}
