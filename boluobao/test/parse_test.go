package test

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"testing"
	"world.misaki.go/crawler/boluobao"
)

func TestBookChapterParse(t *testing.T) {
	document, _ := goquery.NewDocument("https://book.sfacg.com/Novel/589345/779722/6960479/")
	e := &colly.HTMLElement{DOM: document.Selection}
	type args struct {
		html *colly.HTMLElement
		book *boluobao.Book
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "https://book.sfacg.com/Novel/589345/779722/6960479/",
			args: args{
				html: e,
				book: &boluobao.Book{
					Contents: make([]boluobao.BookContent, 0, 100),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boluobao.BookChapterParse(tt.args.html, tt.args.book)
		})
	}
}

func TestBookDetailParse(t *testing.T) {
	type args struct {
		html *colly.HTMLElement
		book *boluobao.Book
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boluobao.BookDetailParse(tt.args.html, tt.args.book)
		})
	}
}
