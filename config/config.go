package config

import "gopkg.in/validator.v2"

type Config struct {
	Version  int    `json:"version" validate:"nonzero"`
	DataDir  string `json:"dataDir" validate:"nonzero"`
	LogLevel string `json:"logLevel" validate:"nonzero"` //info,warn,debug.

	RPC     *RPCConfig  `json:"rpc"`
	Partner *PartnerCfg `json:"partner" validate:"nonnil"`
}

type RPCConfig struct {
	Enable bool `json:"enabled"`

	// TCP or UNIX socket address for the RPC server to listen on
	ListenAddress string `json:"listenAddress" validate:"nonzero"`

	// TCP or UNIX socket address for the gRPC server to listen on
	GRPCListenAddress  string   `json:"gRPCListenAddress" validate:"nonzero"`
	CORSAllowedOrigins []string `json:"httpCors" validate:"min=1"`
}

type PartnerCfg struct {
	BackEndURL     string            `json:"backEndURL" validate:"nonzero"`
	Implementation string            `json:"implementation"`
	Account        string            `json:"account" validate:"nonzero"`
	IsFake         bool              `json:"isFake"`
	IsPrivacy      bool              `json:"isPrivacy"`
	Extra          map[string]string `json:"extra,omitempty"`
}

type Validator func(*Config) error

func (c *Config) Verify(validators ...Validator) error {
	if err := validator.Validate(c); err != nil {
		return err
	}

	for _, validator := range validators {
		if err := validator(c); err != nil {
			return err
		}
	}
	return nil
}
