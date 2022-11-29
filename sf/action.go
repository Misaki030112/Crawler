package sf

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

func (b BaoCrawler) CrawlOneBook(bookDetailUrl string) *Book {
	log.Print("will crawl one book the book url:", bookDetailUrl)
	book := &Book{}

	b.MainCollector.OnHTML(".container", func(e *colly.HTMLElement) {
		BookDetailParse(e, book)
		bookCatelog, exist := e.DOM.Find("#BasicOperation").Children().First().Attr("href")
		if !exist {
			log.Fatal("can not find the bookCatLog Url, can not click \"点击阅读\" button ....")
		}
		bookCatelog = "https://" + e.Request.URL.Host + bookCatelog
		if b.Deputies["BookAllocator"].Visit(bookCatelog) != nil {
			log.Fatal("can not visit the bookCatLog Url : ", bookCatelog)
		}
	})

	b.Deputies["BookAllocator"].OnHTML(".catalog-list .clearfix", func(e *colly.HTMLElement) {
		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			book.AllChapters++
			bookContent, exist := s.Find("a").Attr("href")
			if !exist {
				log.Fatal("can not find the chapter button ....")
			}
			if strings.Contains(bookContent, "vip") {
				// TODO implement login Crawl to crawl some vip chapters,or Optimized to stop the execution of the entire Each when encountering a VIP
				return
			}
			bookContent = "https://" + e.Request.URL.Host + bookContent
			if b.Deputies["BookContent"].Visit(bookContent) != nil {
				log.Fatal("can not visit the bookContent Url : ", bookContent)
			}
		})
	})

	b.Deputies["BookContent"].OnHTML(".container", func(e *colly.HTMLElement) {
		BookChapterParse(e, book)
		book.StoragePath = BookPath(book.BookName, b.StoragePath)
	})

	if b.MainCollector.Visit(bookDetailUrl) != nil {
		log.Fatal("the bookDetailUrl can not visit ....")
	}

	return book
}
