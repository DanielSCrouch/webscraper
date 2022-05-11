package scraper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

// WebPageFormat
type WebPageFormat string

type scraper struct {
	webPageFormat  WebPageFormat
	rawBodyData    []byte
	parsedHTMLData *html.Node
}

func NewScraper(webPageFormat WebPageFormat) (s *scraper) {
	return &scraper{
		webPageFormat: webPageFormat,
	}
}

func (s *scraper) Scrape(uri string) (err error) {

	response, err := http.Get(uri)
	if err != nil {
		return errors.Wrapf(err, "error retrieving site `%s`", uri)
	}
	defer response.Body.Close()

	s.rawBodyData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "error reading response body")
	}

	s.parseData()

	return nil
}

// parseData - process raw web page body by file format
func (s *scraper) parseData() (err error) {

	switch s.webPageFormat {
	case WebPageFormat("html"):
		s.parsedHTMLData, err = html.Parse(bytes.NewReader(s.rawBodyData))
		if err != nil {
			return errors.Wrap(err, "failed to parse to html")
		}

	case WebPageFormat("xml"):
		return errors.New("web page format `xml` not implemented")

	default:
		return errors.New(fmt.Sprintf("web page format `%s` not implemented", s.webPageFormat))
	}

	return nil
}

// Dump - returns raw bytes body response
func (s *scraper) Dump() []byte {
	return s.rawBodyData
}

// ParseAttributes - returns a list of all attributes filtered by key
// Copied from https://pkg.go.dev/golang.org/x/net@v0.0.0-20220425223048-2871e0cb64e4/html#Parse
func (s *scraper) ParseAttributes(key string, elemFilter string) (attrs []string, err error) {

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == key && n.Data == elemFilter {
					attrs = append(attrs, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(s.parsedHTMLData)

	return attrs, nil
}
