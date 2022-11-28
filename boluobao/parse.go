package boluobao

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"log"
	"regexp"
	"strconv"
	"world.misaki.go/crawler/util"
)

// BookDetailParse Parse the basic information of the book details page
func BookDetailParse(html *colly.HTMLElement, book *Book) {
	dom := html.DOM
	book.BookName = dom.Find(".title").First().Text()
	coverImage := book.BookName + ".jpg"
	if imgUrl, exist := dom.Find(".summary-pic").First().Attr("src"); exist {
		if util.FileDownloadFromUrl(imgUrl, coverImage) {
			book.CoverImage = coverImage
		}
	}
	numberReg := regexp.MustCompile("[0-9]+")
	count, err := strconv.ParseUint(numberReg.FindString(dom.Find(".text-row").Children().Eq(1).Text()), 10, 32)
	if err != nil {
		log.Println("can not find world count element")
	}
	book.WorldCount = uint32(count)
	book.Author = dom.Find(".author-name").First().Text()
	//TODO can not get AllContent Count
}

// BookChapterParse Parse the Book Content of the one Book Chapter page
func BookChapterParse(html *colly.HTMLElement, book *Book) {
	dom := html.DOM
	chapter := BookContent{}
	chapter.ChapterName = dom.Find(".article-title").Text()
	by := bytes.Buffer{}
	dom.Find("#ChapterBody").Children().Each(func(i int, e *goquery.Selection) {
		by.WriteString(e.Text())
		by.WriteString("\n")
	})
	chapter.ChapterContent = by.String()
	book.Contents = append(book.Contents, chapter)
}
