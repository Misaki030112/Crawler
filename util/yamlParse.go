package util

type CrawlerProperties struct {
	CrawlerType string `yaml:"crawler_type"`
}

// NewCrawlerProperties new Crawler
func NewCrawlerProperties() *CrawlerProperties {
	// TODO read yaml config from config.yml to struct
	//os.ReadFile("config.yml")
	return nil
}
