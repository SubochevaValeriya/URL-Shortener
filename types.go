package urls

import "time"

type UrlInfo struct {
	//Id          int       `json:"id" db:"id"`
	OriginalURL string    `json:"original_URL" db:"original_URL"`
	ShortURL    string    `json:"short_URL" db:"short_URL"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Visits      int       `json:"visits" db:"visits"`
}
