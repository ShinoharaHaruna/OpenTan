package config

type Mode string

const (
	ModeDebug   Mode = "debug"
	ModeRelease Mode = "release"
)

type Config struct {
	Host      string      `envconfig:"HOST"`
	Port      string      `envconfig:"PORT"`
	Prefix    string      `envconfig:"PREFIX"`
	Mode      Mode        `envconfig:"MODE"`
	API_KEY   string      `envconfig:"API_KEY"`
	ID        string      `envconfig:"ID"`
	Password  string      `envconfig:"PASSWORD"`
	RateLimit RateLimiter `envconfig:"RATE_LIMIT"`
}

type RateLimiter struct {
	Rate    int `envconfig:"RATE"`
	Burst   int `envconfig:"BURST"`
	MaxWait int `envconfig:"MAX_WAIT"`
}
