package config

type Config struct {
	Version  int    `json:"version"`
	DataDir  string `json:"dataDir"`
	LogLevel string `json:"logLevel"` //info,warn,debug.

	RPC     *RPCConfig  `json:"rpc"`
	Partner *PartnerCfg `json:"partner"`
	Privacy *PrivacyCfg `json:"privacy"`
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
	Name           string `json:"name" validate:"nonzero"`
	SonataUrl      string `json:"sonataUrl" validate:"nonzero"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	APIToken       string `json:"token,omitempty"`
	Implementation string `json:"implementation"`
	ChainUrl       string `json:"chainUrl,omitempty"`
	Account        string `json:"account" validate:"nonzero"`
	IsFake         bool   `json:"isFake"`
}

type PrivacyCfg struct {
	Enable         bool     `json:"enabled"`
	From           string   `json:"from"`
	For            []string `json:"for"`
	PrivateGroupID string   `json:"privateGroupID"`
}
