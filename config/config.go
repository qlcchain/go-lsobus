package config

type Config struct {
	Version  int    `json:"version"`
	DataDir  string `json:"dataDir"`
	LogLevel string `json:"logLevel"` //info,warn,debug.

	RPC     RPCConfig   `json:"rpc"`
	Partner *PartnerCfg `json:"partner"`
	Privacy PrivacyCfg  `json:"privacy"`
}

type RPCConfig struct {
	Enable bool `json:"enabled"`

	// TCP or UNIX socket address for the RPC server to listen on
	ListenAddress string `json:"listenAddress"`

	// TCP or UNIX socket address for the gRPC server to listen on
	GRPCListenAddress  string   `json:"gRPCListenAddress"`
	CORSAllowedOrigins []string `json:"httpCors"`
}

type PartnerCfg struct {
	Name           string `json:"name"`
	SonataUrl      string `json:"sonataUrl"`
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	Implementation string `json:"implementation"`
	ChainUrl       string `json:"chainUrl,omitempty"`
	Account        string `json:"account"`
	IsFake         bool   `json:"isFake"`
}

type PrivacyCfg struct {
	Enable         bool     `json:"enabled"`
	From           string   `json:"from"`
	For            []string `json:"for"`
	PrivateGroupID string   `json:"privateGroupID"`
}
