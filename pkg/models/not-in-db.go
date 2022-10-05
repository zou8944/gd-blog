package models

type Author struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Avatar string `json:"avatar"`
	CSDN   string `json:"csdn"`
	Github string `json:"github"`
}

type SiteInfo struct {
	Author     Author     `json:"author"`
	Statistics Statistics `json:"statistics"`
	Beian      string     `json:"beian"`
}

type Statistics struct {
	BlogCount     int64 `json:"blog_count"`
	CategoryCount int64 `json:"category_count"`
	TagCount      int64 `json:"tag_count"`
	VisitorCount  int64 `json:"visitor_count"`
	ViewCount     int64 `json:"view_count"`
}

func GetStatistics() (*Statistics, error) {
	var blogCount int64
	var categoryCount int64
	var tagCount int64
	var visitorCount int64
	var viewCount int64
	db.Model(&Blog{}).Count(&blogCount)
	db.Model(&Category{}).Count(&categoryCount)
	db.Model(&Tag{}).Count(&tagCount)
	var viewCountResult map[string]int64
	db.Model(&Blog{}).Select("sum(view_count) as total").First(&viewCountResult)
	if db.Error != nil {
		return nil, db.Error
	}
	viewCount = viewCountResult["total"]
	visitorCount = viewCount / 3
	return &Statistics{
		BlogCount:     blogCount,
		CategoryCount: categoryCount,
		TagCount:      tagCount,
		VisitorCount:  visitorCount,
		ViewCount:     viewCount,
	}, nil
}

func GetSiteInfo() (*SiteInfo, error) {
	stat, err := GetStatistics()
	if err != nil {
		return nil, err
	}
	siteInfo := SiteInfo{
		Author: Author{
			Name:   "果冻",
			Desc:   "果冻的碎碎念",
			Avatar: "https://thirdwx.qlogo.cn/mmopen/vi_32/DYAIOgq83equib0YGKeGrRww67LyZ7hSONtAW59RHDTd2JuKmSfQLEs8zWIB14hUcHibNG41zNibv5mr5QhM5QDMQ/132",
			CSDN:   "https://blog.csdn.net/zou8944",
			Github: "https://github.com/zou8944",
		},
		Statistics: *stat,
		Beian:      "粤ICP备2021024139号",
	}
	return &siteInfo, nil
}
