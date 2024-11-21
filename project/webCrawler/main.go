package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var url string = ""

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type Parser interface {
	GetSeoData(resp *http.Response) (SeoData, error)
}

type DefaultParser struct {
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func isSitemap(urls []string) ([]string, []string) {}

func extractSitemapURLs(startURL string) []string {}

func makeRequest(url string) (*http.Response, error) {}

func scrapeUrls(urls []string, parser Parser, concurrency int) []SeoData {}

func extractUrls(response *http.Response) ([]string, error) {}

func crawlPage(url string, tokens chan struct{}) (*http.Response, error) {}

func scrapePage(url string, token chan struct{}, parser Parser) (SeoData, error) {}

func (d DefaultParser) GetSeoData(resp *http.Response) (SeoData, error) {}

func ScraperSitemap(url string, parser Parser, concurrency int) []SeoData {}

func main() {
	fmt.Println("New Start")
	p := DefaultParser{}

	results := ScraperSitemap(url, p, 10)
	for _, res := range results {

	}
}
