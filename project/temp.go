package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery" // Import goquery for parsing HTML
)

// SeoData is a struct of useful SEO data
type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

// Parser defines the parsing interface
type Parser interface {
	GetSeoData(resp *http.Response) (SeoData, error)
}

// DefaultParser is an empty struct for implementing the default parser
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
	return userAgents[rand.Intn(len(userAgents))]
}

// Segregating pages and sitemap files (URLs containing more URLs and not actual pages)
func isSitemap(urls []string) ([]string, []string) {
	sitemapFiles := []string{}
	pages := []string{}
	for _, page := range urls {
		if strings.Contains(page, "xml") {
			fmt.Println("Found Sitemap:", page)
			sitemapFiles = append(sitemapFiles, page)
		} else {
			pages = append(pages, page)
		}
	}
	return sitemapFiles, pages
}

func extractSitemapURLs(startURL string) []string {
	var toCrawl []string
	worklist := make(chan []string) // Channel to manage URLs
	var wg sync.WaitGroup           // WaitGroup to synchronize goroutines

	// Start processing with the initial URL
	wg.Add(1)
	go func() { worklist <- []string{startURL} }()

	go func() {
		wg.Wait()
		close(worklist)
	}()

	// Process URLs from the channel
	for list := range worklist {
		for _, link := range list {
			wg.Add(1)
			go func(link string) {
				defer wg.Done()
				response, err := makeRequest(link)
				if err != nil {
					log.Printf("Error retrieving URL: %s", link)
					return
				}

				urls, err := extractUrls(response)
				if err != nil {
					log.Printf("Error extracting document from response, URL: %s", link)
					return
				}

				sitemapFiles, pages := isSitemap(urls)
				if sitemapFiles != nil {
					go func() { worklist <- sitemapFiles }()
				}
				toCrawl = append(toCrawl, pages...)
			}(link)
		}
	}
	return toCrawl
}

func makeRequest(url string) (*http.Response, error) {
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", randomUserAgent())
	return client.Do(req)
}

func extractUrls(response *http.Response) ([]string, error) {
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	results := []string{}
	doc.Find("loc").Each(func(i int, s *goquery.Selection) {
		results = append(results, s.Text())
	})
	return results, nil
}

func scrapePage(url string, tokens chan struct{}, parser Parser) (SeoData, error) {
	tokens <- struct{}{}        // Acquire token
	defer func() { <-tokens }() // Release token

	resp, err := makeRequest(url)
	if err != nil {
		return SeoData{}, err
	}
	return parser.GetSeoData(resp)
}

func (d DefaultParser) GetSeoData(resp *http.Response) (SeoData, error) {
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return SeoData{}, err
	}
	return SeoData{
		URL:             resp.Request.URL.String(),
		Title:           doc.Find("title").Text(),
		H1:              doc.Find("h1").Text(),
		MetaDescription: doc.Find("meta[name=description]").AttrOr("content", ""),
		StatusCode:      resp.StatusCode,
	}, nil
}

func ScrapeSitemap(url string, parser Parser, concurrency int) []SeoData {
	results := extractSitemapURLs(url)

	tokens := make(chan struct{}, concurrency) // Limit concurrency
	var wg sync.WaitGroup
	data := make([]SeoData, 0)

	for _, url := range results {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			result, err := scrapePage(url, tokens, parser)
			if err != nil {
				log.Printf("Error scraping URL %s: %v", url, err)
				return
			}
			data = append(data, result)
		}(url)
	}
	wg.Wait()
	return data
}

func main() {
	p := DefaultParser{}
	results := ScrapeSitemap("https://www.quicksprout.com/sitemap.xml", p, 10)
	for _, res := range results {
		fmt.Println(res)
	}
}
