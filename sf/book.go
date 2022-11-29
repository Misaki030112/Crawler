package sf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Book struct {
	CoverImage      string
	BookName        string
	WorldCount      uint32
	Contents        []BookContent
	Author          string
	CrawledChapters uint32
	AllChapters     uint32
	StoragePath     string
}

type BookContent struct {
	ChapterName    string
	ChapterContent string
	IsVip          bool
}

func (b Book) ToTxtFile() {
	bookPath := b.StoragePath
	if err := os.MkdirAll(filepath.Dir(bookPath), 0750); err != nil {
		log.Panic("can not create the file parent dir... bookFile : ", bookPath)
	}
	f, err := os.Create(bookPath)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Print("can not close the file ,file content may be lost.. bookFile : ", bookPath)
		}
	}(f)

	if err != nil {
		log.Panic("can not create the file ... bookFile : ", bookPath)
	}
	log.Print("will start write book content to txt file ... write File ", bookPath)

	if _, err := f.WriteString(fmt.Sprintf("书名：%s\n作者：%s\n来源：%s\n字数：%d\n已抓取章节数：%d\n总章节数：%d\n",
		b.BookName, b.Author, "SF Novel", b.WorldCount, b.CrawledChapters, b.AllChapters)); err != nil {
		log.Print("write book head information error bookName :", b.BookName)
	}
	for _, content := range b.Contents {
		if _, err := f.WriteString(fmt.Sprintf("\n%s\n%s\n\n", content.ChapterName, content.ChapterContent)); err != nil {
			log.Printf("the chapter %s write error ....", content.ChapterName)
		}
	}
	if f.Sync() != nil {
		log.Printf("the book file %s file sync fail ,may lost some content.....", bookPath)
	}
	log.Printf("the book  %s save success!", b.BookName)
}

// ToEpubFile Convert Book object to Epub format ....
func (b Book) ToEpubFile() {
	//TODO Implementation this func ....
	//reference doc https://pkg.go.dev/github.com/bmaupin/go-epub#section-readme
}

func BookPath(bookName, basePath string) string {
	return fmt.Sprintf("%s/%s.txt", basePath, bookName)
}
