package main

import (
	"flag"
	"log"
	_3qb "world.misaki.go/crawler/23qb"
	"world.misaki.go/crawler/domain"
	"world.misaki.go/crawler/sf"
)

var (
	module    string
	storage   string
	targetUrl string
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	argsAnalyze()
	log.Printf("input args module:%s ; storage:%s ; targetUrl:%s", module, storage, targetUrl)
	switch module {
	case "SF":
		//Temporarily simple default implementation, will be expanded later
		b := domain.NewBookCrawler("book.sfacg.com")
		book := sf.CrawlOneBook(b, targetUrl)
		book.ToTxtFile()
		b.CleanCache()
	case "QB":
		b := domain.NewBookCrawler("www.23qb.com")
		book := _3qb.CrawlOneBook(b, targetUrl)
		book.ToTxtFile()
		b.CleanCache()

	default:
		log.Fatal("sorry the other module still not supported ...")
	}
}

func argsAnalyze() {
	flag.StringVar(&module, "m", "SF", "the module type example: SF ....")
	flag.StringVar(&storage, "s", "storage", "the file will save position")
	flag.StringVar(&targetUrl, "t", "", "the base url")
	flag.Parse()
}
