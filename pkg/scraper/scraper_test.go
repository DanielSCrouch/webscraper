package scraper

import (
	"fmt"
	"net/http"
	"testing"
)

var testHTMLResp string = `<!DOCTYPE html>
<html>
<a href="http://www.google.com">Google</a>
<body>
	<p>
	<a href="http://www.duckduckgo.com">DuckDuckGo</a>
	</p>
</body>
</html>`

// RunTestHTTPServer - runs a simple test server on localhost
// port 8090 with endpoint /test. Server closes when test ends
// curl http://localhost:8090/test
func RunTestHTTPServer(t *testing.T) (endpoint string) {
	t.Helper()
	port := 8090

	http.HandleFunc("/test", httpTestServerEndpoint)
	// go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		// Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("listen: %s\n", err)
		}
	}()

	t.Logf("test HTTP server running on local host port %d", port)

	cleanup := func() {
		srv.Close()
	}
	t.Cleanup(cleanup)

	return fmt.Sprintf("http://localhost:%d/test", port)

}

// httpTestServeEndpoint - writes a simple HTTP response containing
// a parent and child element each with a href attribute
func httpTestServerEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, testHTMLResp)
}

func TestScraper(t *testing.T) {
	endpoint := RunTestHTTPServer(t)
	htmlScraper := NewScraper(WebPageFormat("html"))

	err := htmlScraper.Scrape(endpoint)
	if err != nil {
		t.Fatal(err)
	}

	if string(htmlScraper.rawBodyData) != testHTMLResp {
		t.Error("expected body resp does not match actual")
	}
}

// func TestHTMLParse(t *testing.T) {
// 	endpoint := RunTestHTTPServer(t)
// 	htmlScraper := NewScraper(WebPageFormat("html"))

// 	err := htmlScraper.Scrape(endpoint)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(htmlScraper.parsedHTMLData.FirstChild.Data)

// if htmlScraper.parsedHTMLData.Type == html.ElementNode && htmlScraper.parsedHTMLData.Data == "a" {
// 	// Do something with n...
// }
// for c := n.FirstChild; c != nil; c = c.NextSibling {
// 	f(c)
// }

// }
