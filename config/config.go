package config

type Config struct {
	Version  int           `json:"version"`
	DataDir  string        `json:"dataDir"`
	LogLevel string        `json:"logLevel"` //info,warn,debug.
	ChainUrl string        `json:"chainUrl"` //chain url.
	RPC      RPCConfig     `json:"rpc"`
	Partners []*PartnerCfg `json:"partners"`
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
	Name      string `json:"name"`
	ID        string `json:"id"`
	SonataUrl string `json:"sonataUrl"`
}
