package main

import (
	"fmt"
	"log"
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

func extractSitemapURLs(startURL string) []string {
	worklist := make(chan []string) 					// contains all the (batched together) URLs that is to be crawled ; by using slice of string we have achievec batching 
	toCrawl := []string{}
	var n int
	n++
	func() { worklist <- []string{startURL} }() 		//sending startingURL into channel
	for ; n > 0; n-- {
		list := <-worklist								// extracting LAST URL STORED in the channel
		for _, link := range list {
			go func(link string) {
				response, err := makeRequest(link)
				if err != nil {
					log.Printf("Error retrieving URL: %s", link)
				}
				urls, _ := extractUrls(response)
				if err != nil {
					log.Printf("Error extracting document from response, URL: %s", link)
				}
				sitemapFiles, pages := isSitemap(urls)
				if sitemapFiles != nil {
					worklist <- sitemapFiles
					n++											// increasing number of loops as new URLs are found
				}
				for _, page := range pages {
					toCrawl = append(toCrawl, page)
				}
			}(link)
		}
	}
	return toCrawl
}

func makeRequest(url string) (*http.Response, error) {
	// A new HTTP request is tailored with client timeout, type=GET and with User-agent(browser)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	if err != nil {
		return nil, err
	}
	// The new HTTP request is acutally sent 
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func isSitemap(urls []string) ([]string, []string) {}

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
