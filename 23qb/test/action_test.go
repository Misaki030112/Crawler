package test

import (
	"log"
	"testing"
	_3qb "world.misaki.go/crawler/23qb"
	"world.misaki.go/crawler/domain"
)

func TestQbCrawler_CrawAction(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	testUrl := "https://www.23qb.com/book/1888/"
	b := domain.NewBookCrawler("www.23qb.com")
	book := _3qb.CrawlOneBook(b, testUrl)

	book.ToTxtFile()
	//b.CleanCache()
}
