package boluobao

import (
	"github.com/gocolly/colly/v2"
)

// BookDetailParse Parse the basic information of the book details page
func BookDetailParse(html *colly.HTMLElement) *Book {
	book := &Book{}
	dom := html.DOM
	book.BookName = dom.Find(".title").First().Text()
	dom.Find(".summary-pic").First().Attr("src")
	dom.Find("")
}
