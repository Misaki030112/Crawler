package main

import (
	"flag"
	"log"
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
	switch module {
	case "SF":
		//Temporarily simple default implementation, will be expanded later
		log.Printf("input args module:%s ; storage:%s ; targetUrl:%s", module, storage, targetUrl)
		b := sf.NewBaoCrawler()
		b.StoragePath = storage
		book := b.CrawlOneBook(targetUrl)
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
