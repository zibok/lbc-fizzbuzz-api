package httpapi

const (
	defaultAddr     = ":8080"
	defaultMaxLimit = 10000
)

type Config struct {
	Addr     string `json:"addr"`
	MaxLimit int    `json:"maxLimit"`
}

func DefaultConfig() Config {
	return Config{
		Addr:     defaultAddr,
		MaxLimit: defaultMaxLimit,
	}
}

func (cfg Config) WithDefaults() Config {
	defaults := DefaultConfig()

	if cfg.Addr == "" {
		cfg.Addr = defaults.Addr
	}

	if cfg.MaxLimit == 0 {
		cfg.MaxLimit = defaults.MaxLimit
	}

	return cfg
}
