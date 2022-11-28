package test

import (
	"os"
	"testing"
	"world.misaki.go/crawler/boluobao"
)

func TestBaoCrawler_CrawlOneBook(t *testing.T) {

	testUrl := "https://book.sfacg.com/Novel/589345/"
	b := boluobao.NewBaoCrawler()
	book := b.CrawlOneBook(testUrl)
	//by, err := json.Marshal(book)
	//if err != nil {
	//	t.Fatal("json encoded error")
	//}
	f, _ := os.Create("testFile")
	defer f.Close()
	for _, content := range book.Contents {
		f.WriteString(content.ChapterName)
		f.WriteString(content.ChapterContent)
		f.WriteString("\n============章节分割线=============\n")
	}
}
