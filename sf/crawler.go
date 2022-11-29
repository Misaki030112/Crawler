package sf

import (
	"github.com/gocolly/colly/v2"
	"log"
	"os"
)

const (
	DOMAIN        = "book.sfacg.com"
	DEFAULT_CACHE = "cache"
	STORAGE_PATH  = "storage"
)

type BaoCrawler struct {
	MainCollector *colly.Collector
	// the other collector
	Deputies map[string]*colly.Collector

	StoragePath string

	targetDomain string

	cacheDir string

	deputiesCount uint16
}

func (b BaoCrawler) TargetDomain() string {
	return b.targetDomain
}

func (b BaoCrawler) CacheDir() string {
	return b.cacheDir
}

func (b BaoCrawler) DeputiesCount() uint16 {
	return b.deputiesCount
}

func NewBaoCrawler() *BaoCrawler {
	c := &BaoCrawler{}
	c.StoragePath = STORAGE_PATH
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
	c.deputiesCount = 0
	return c
}

func (b BaoCrawler) CleanCache() bool {
	if os.RemoveAll(b.cacheDir) != nil {
		log.Print("remove crawl cache error ....")
		return false
	}
	return true
}
