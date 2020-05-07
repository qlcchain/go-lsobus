package config

type Config struct {
	Version  int    `json:"version"`
	DataDir  string `json:"dataDir"`
	LogLevel string `json:"logLevel"` //info,warn,debug.
	ChainUrl string `json:"chainUrl"` //chain url.
}
