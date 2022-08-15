package config


type Config struct{
	AbsoluteLogPath string

	//  gRPC
	GRPCNetworkType string
	GRPCHost string
	GRPCPort string
	
	
	// REST Gateway
	RESTHost string
	RESTPort string
	
	
	//  for local deployment
	DBHost string
	DBName string
	DBUser   string
	DBPass string
	DBPort string
	DBEnableTLS bool
	
	//  gorm 
	DBDialect string
	
	// auth
	TokenKey string
	TokenIssuer string
	AccessTokenAliveTime int
	RefreshTokenAliveTime int
	OuthCookieAliveTime int
	
	// password reset time
	PasswordResetInterval string
	
	// OTP
	OTPLength int
	OTPAliveTime int
	
	//aws
	OTPSender string
	
	// rest endpoints
	LoginUI string
	
	//  file location
	TemplatesHtml string
	
	
	//  google auth endpoints and secret
	GoogleLoginEnpoint string
	GoogleCallbackEnpoint string

	GoogleRedirectURl string
	GoogleClientId string
	GoogleClientSecret string
	
	
	

// facebook login configs
	FacebookLoginEndPoint string
	FacebookCallbackEndPoint string
	FacebookClientId string
	FacebookClientSecret string
	FacebookRedirectUrl string
	
	
	
//  aws config for OTP

	AWSAceesKeyId string
	AWSSecretAccessKey string
	AWSRegion string
	AWSSMTPUser string
	AWSSMTPPassword string
	AWSHost string
	AWSPort string
	AWSSenderEmail string	
	
}

func NewConfig() *Config {
	return &Config{}
}