package models

type Comment struct {
	BaseModel
	BlogID    int32  `json:"blog_id"`
	Commenter string `json:"commenter"`
	ReplyID   int32  `json:"reply_id"`
	GValue    int32  `json:"g_value"`
	Content   string `json:"content"`
	Approved  int32  `json:"approved"`
}
