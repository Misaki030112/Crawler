package _3qb

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"world.misaki.go/crawler/domain"
)

// BookDetailParse Parse the basic information of the book details page
func BookDetailParse(html *colly.HTMLElement, book *domain.Book) {
	dom := html.DOM
	book.BookName = dom.Find(".d_title").Children().Eq(0).Text()
	coverImage, exist := dom.Find("img").Eq(0).Attr("src")
	if exist {
		book.CoverImage = coverImage
	}
	book.Author = dom.Find("#count").Children().Children().Eq(1).Text()
}

// BookChapterParse Parse the Book Content of the one Book Chapter page
func BookChapterParse(html *colly.HTMLElement, book *domain.Book) {
	dom := html.DOM.Find("#mlfy_main_text")
	chapter := domain.BookContent{}
	chapter.ChapterName = dom.Children().Eq(0).Text()
	by := bytes.Buffer{}
	dom.Find("#TextContent").Find("p").Each(func(i int, e *goquery.Selection) {
		by.WriteString(e.Text())
		by.WriteString("\n")
	})
	book.CrawledChapters++
	chapter.ChapterContent = by.String()
	chapter.IsVip = false
	book.Contents = append(book.Contents, chapter)
}
