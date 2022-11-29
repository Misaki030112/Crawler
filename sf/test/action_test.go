package test

import (
	"log"
	"testing"
	"world.misaki.go/crawler/sf"
)

func TestBaoCrawler_CrawlOneBook(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	testUrl := "https://book.sfacg.com/Novel/589345/"
	b := sf.NewBaoCrawler()
	book := b.CrawlOneBook(testUrl)
	book.ToTxtFile()
	b.CleanCache()
}
