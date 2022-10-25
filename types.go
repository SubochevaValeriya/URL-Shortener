package urls

import "time"

type UrlInfo struct {
	OriginalURL string    `json:"original_URL" db:"original_URL" bson:"original_URL"`
	ShortURL    string    `json:"short_URL" db:"short_URL" bson:"short_URL"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" bson:"created_at"`
	Visits      int       `json:"visits" db:"visits" bson:"visits"`
}
