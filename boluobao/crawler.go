package boluobao

import (
	"github.com/gocolly/colly/v2"
)

const (
	DOMAIN          = "book.sfacg.com"
	DEFAULT_CACHE   = "./cache"
	DEFAULT_STORAGE = "./storage"
)

type BaoCrawler struct {
	MainCollector *colly.Collector
	// the other collector
	Deputies map[string]*colly.Collector

	targetDomain string

	cacheDir string

	storageDir string

	deputiesCount uint16
}

func (b BaoCrawler) TargetDomain() string {
	return b.targetDomain
}

func (b BaoCrawler) CacheDir() string {
	return b.cacheDir
}

func (b BaoCrawler) StorageDir() string {
	return b.storageDir
}

func (b BaoCrawler) DeputiesCount() uint16 {
	return b.deputiesCount
}

func NewBaoCrawler() *BaoCrawler {
	c := &BaoCrawler{}
	c.MainCollector = colly.NewCollector(
		colly.AllowedDomains(DOMAIN),
		colly.CacheDir(DEFAULT_CACHE),
		colly.IgnoreRobotsTxt(),
	)
	c.Deputies = make(map[string]*colly.Collector)
	c.Deputies["BookAllocator"] = c.MainCollector.Clone()
	c.Deputies["BookContent"] = c.MainCollector.Clone()

	c.targetDomain = DOMAIN
	c.cacheDir = DEFAULT_CACHE
	c.storageDir = DEFAULT_STORAGE
	c.deputiesCount = 0
	return c
}
