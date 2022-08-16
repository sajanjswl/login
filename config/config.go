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

	//  gorm
	DBDialect string

	// auth
	TokenKey              string
	TokenIssuer           string
	AccessTokenAliveTime  int
	RefreshTokenAliveTime int
	OuthCookieAliveTime   int

	// password reset time
	PasswordResetInterval int

	// OTP
	OTPLength    int
	OTPAliveTime int

	//aws
	OTPSender string

	// rest endpoints
	LoginUIEndPoint string

	//  file location
	HTMLTemplate string

	//  google auth endpoints and secret
	GoogleLoginEnpoint    string
	GoogleCallbackEnpoint string
	GoogleRedirectURl  string
	GoogleClientId     string
	GoogleClientSecret string

	// facebook login configs
	FacebookLoginEndPoint    string
	FacebookCallbackEndPoint string
	FacebookClientId         string
	FacebookClientSecret     string
	FacebookRedirectUrl      string

	//  aws config for OTP
	AWSAceesKeyId      string
	AWSSecretAccessKey string
	AWSRegion          string
	AWSSMTPUser        string
	AWSSMTPPassword    string
	AWSHost            string
	AWSPort            string
	AWSSenderEmail     string
}

func NewConfig() *Config {
	return &Config{}
}
