package config

type Config struct {
	//  gRPC
	GRPCNetworkType string
	GRPCHost        string
	GRPCPort        string

	// REST Gateway
	RESTHost        string
	RESTPort        string
	AbsoluteLogPath string

	//  for local deployment
	DBHost     string
	DBName     string
	DBUserName string
	DBPassword string
	DBPort     string
	DBSLLMode  string

	// rest endpoints
	LoginUIEndPoint string

	//  file location
	HTMLTemplate string

	//  google auth endpoints and secret
	GoogleLoginEnpoint    string
	GoogleCallbackEnpoint string
	GoogleRedirectURl     string
	GoogleClientId        string
	GoogleClientSecret    string
}

func NewConfig() *Config {
	return &Config{}
}
