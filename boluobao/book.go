package boluobao

type Book struct {
	CoverImage      string
	BookName        string
	WorldCount      uint32
	Contents        []BookContent
	Author          string
	CrawledChapters uint32
	AllChapters     uint32
}

type BookContent struct {
	ChapterName    string
	ChapterContent string
	IsVip          bool
}
