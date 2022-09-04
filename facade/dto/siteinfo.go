package dto

type Author struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Avatar string `json:"avatar"`
	CSDN   string `json:"csdn"`
	Github string `json:"github"`
}

type Statistics struct {
	BlogCount     int64 `json:"blog_count"`
	CategoryCount int64 `json:"category_count"`
	TagCount      int64 `json:"tag_count"`
	VisitorCount  int64 `json:"visitor_count"`
	ViewCount     int64 `json:"view_count"`
}

type SiteInfo struct {
	Author     Author     `json:"author"`
	Statistics Statistics `json:"statistics"`
	Beian      string     `json:"beian"`
}
