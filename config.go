package skyscanner

import "time"

type Config struct {
	APIKey         string
	QueriesTimeout time.Duration
}
