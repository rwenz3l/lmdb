package anime

import (
	"net/http"

	"github.com/go-kit/kit/log/level"
)

// CreateHandler for the Anime Collection
func (svc *Service) CreateHandler(w http.ResponseWriter, r *http.Request) {
	level.Debug(svc.logger).Log("msg", "Create Handler (TestDrive)")
	anime := &Anime{
		Titles: []Title{
			Title{
				Title:    "Shokugeki no Souma",
				Type:     "main",
				Language: "en",
			},
			Title{
				Title:    "Food Wars! Shokugeki no Soma",
				Type:     "official",
				Language: "en",
			},
			Title{
				Title:    "食戟のソーマ",
				Type:     "official",
				Language: "jp",
			},
		},
		Year: 2015,
	}
	svc.Create(*anime)
}
