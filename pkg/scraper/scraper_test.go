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
		<a 	href="http://www.duckduckgo.com">DuckDuckGo</a>
		</p>
	</body>
</html>`

// RunTestHTTPServer - runs a simple test server on localhost
// port 8090 with endpoint /test. Server closes when test ends
// curl http://localhost:8090/test
func RunTestHTTPServer(t *testing.T) (endpoint string) {
	t.Helper()
	port := 8090

	mux := http.NewServeMux()
	mux.HandleFunc("/test", httpTestServerEndpoint)
	// http.HandleFunc("/test", httpTestServerEndpoint)
	// go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
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

	if string(htmlScraper.Dump()) != testHTMLResp {
		t.Error("expected body resp does not match actual")
	}
}

func TestParseAttributes(t *testing.T) {
	endpoint := RunTestHTTPServer(t)
	htmlScraper := NewScraper(WebPageFormat("html"))

	err := htmlScraper.Scrape(endpoint)
	if err != nil {
		t.Fatal(err)
	}

	attrs, err := htmlScraper.ParseAttributes("href", "a")
	if err != nil {
		t.Fatal(err)
	}

	if len(attrs) != 2 {
		t.Fatalf("identified %d of 2 attributes", len(attrs))
	}

	// TODO: Assert incorrectly assumes order is maintained in implementation
	exp := "http://www.google.com"
	if attrs[0] != exp {
		t.Fatalf("expected first attribute %s received %s", exp, attrs[0])
	}

	exp = "http://www.duckduckgo.com"
	if attrs[1] != exp {
		t.Fatalf("expected first attribute %s received %s", exp, attrs[1])
	}

}
