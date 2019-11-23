package anime

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/log/level"
)

// SearchHandler for the Anime Collection
func (svc *Service) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	t, _ := url.QueryUnescape(query.Get("q"))
	animes := svc.Find(t)
	for i := range animes {
		level.Debug(svc.logger).Log("var", animes[i])
	}
	fmt.Fprintf(w, "Found: %d Animes", len(animes))
	fmt.Fprintf(w, "Animes: %v", animes)
}
