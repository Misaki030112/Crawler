package test

import (
	"log"
	"testing"
	"world.misaki.go/crawler/domain"
	"world.misaki.go/crawler/sf"
)

func TestBaoCrawler_CrawlOneBook(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	testUrl := "https://book.sfacg.com/Novel/589345/"
	b := domain.NewBookCrawler("book.sfacg.com")
	book := sf.CrawlOneBook(b, testUrl)
	book.ToTxtFile()
	b.CleanCache()
}
