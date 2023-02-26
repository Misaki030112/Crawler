package _3qb

import (
	"github.com/gocolly/colly/v2"
	"log"
	"regexp"
	"world.misaki.go/crawler/domain"
)

func CrawlOneBook(b *domain.BookCrawler, bookDetailUrl string) *domain.Book {
	log.Print("will crawl one book the book url:", bookDetailUrl)
	book := &domain.Book{}

	b.MainCollector.OnHTML("#maininfo", func(e *colly.HTMLElement) {
		BookDetailParse(e, book)
	})

	b.MainCollector.OnHTML("#chapterList", func(e *colly.HTMLElement) {
		book.AllChapters = uint32(e.DOM.Size())
		bookContent, exist := e.DOM.Eq(0).Children().Eq(0).Children().Eq(0).Attr("href")
		if !exist {
			log.Printf("the start chapter can not visited..")
			return
		}
		bookContent = "https://" + e.Request.URL.Host + bookContent
		if b.Deputies["BookContent"].Visit(bookContent) != nil {
			log.Fatal("can not visit the bookContent Url : ", bookContent)
		}
	})

	b.Deputies["BookContent"].OnHTML("body", func(e *colly.HTMLElement) {
		BookChapterParse(e, book)
		book.StoragePath = domain.BookPath(book.BookName, b.StoragePath)
		jsCode := e.DOM.Find("script").Last().Text()
		re := regexp.MustCompile(`"([^"]+)"`)
		match := re.FindAllStringSubmatch(jsCode, 4)
		if len(match) > 1 {
			nextContext := match[1][1]
			// Interpretation whether it ends with .html
			if !(len(nextContext) >= len(".html") && nextContext[len(nextContext)-len(".html"):] == ".html") {
				log.Printf("has no next chapter...")
				return
			}
			nextContext = "https://" + e.Request.URL.Host + nextContext
			log.Printf("crawl %s", nextContext)
			for i := 0; i < 21 && b.Deputies["BookContent"].Visit(nextContext) == nil; i++ {
				log.Printf("visit bookcontent url %s fail .. retry [%d/20]", nextContext, i)
				if i == 20 {
					log.Panicf("visit bookcontent url %s fail ... retry too many", nextContext)
				}
			}
		} else {
			log.Printf("has no next chapter...")
		}
	})
	if b.MainCollector.Visit(bookDetailUrl) != nil {
		log.Fatal("the bookDetailUrl can not visit ....")
	}

	return book
}
