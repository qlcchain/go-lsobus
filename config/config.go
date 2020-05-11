package config

type Config struct {
	Version  int           `json:"version"`
	DataDir  string        `json:"dataDir"`
	LogLevel string        `json:"logLevel"` //info,warn,debug.
	ChainUrl string        `json:"chainUrl"` //chain url.
	Partners []*PartnerCfg `json:"partners"`
}

type PartnerCfg struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	SonataUrl string `json:"sonataUrl"`
}
