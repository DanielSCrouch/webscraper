package scraper

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// RunTestHTTPServer - runs a simple test server on local host
// port 8090 with endpoint /test. Server closes when test ends
func RunTestHTTPServer(t *testing.T) {
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

}

// httpTestServeEndpoint - writes a simple HTTP response containing
// a parent and child element each with a href attribute
func httpTestServerEndpoint(w http.ResponseWriter, req *http.Request) {
	resp := `<!DOCTYPE html>
<html>
<a href="http://www.google.com">Google</a>
<body>
	<p>
	<a href="http://www.duckduckgo.com">DuckDuckGo</a>
	</p>
</body>
</html>`

	fmt.Fprint(w, resp)
}

func TestScraper(t *testing.T) {
	RunTestHTTPServer(t)

	time.Sleep(15 * time.Second)
	fmt.Println("finishing")
}
