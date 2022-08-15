package main

import (
	"flag"

	"github.com/sajanjswl/auth/config"
	"github.com/sajanjswl/auth/pkg/cmd"
	"go.uber.org/zap"
)

const service string = "auth-service"

func main() {
	cfg := config.NewConfig()

	// grpc-server configs
	flag.StringVar(&cfg.GRPCHost, "grpc-host", "localhost", "grpc host")
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "8000", "grpc port")
	flag.StringVar(&cfg.GRPCNetworkType, "grpc-network-type", "tcp", "grpc network type")

	// rest-server configs
	flag.StringVar(&cfg.RESTHost, "rest-host", "localhost", "rest host")
	flag.StringVar(&cfg.RESTPort, "rest-port", "9000", "rest port")
	flag.StringVar(&cfg.LoginUIEndPoint, "login-ui", "/v1/login/ui/", "Oauth login ui endpoint")
	flag.StringVar(&cfg.HTMLTemplate, "html-templates", "./templates", "html template path")
	flag.StringVar(&cfg.AbsoluteLogPath, "logpath", "logs.log", "application logs path")

	// database configs
	flag.StringVar(&cfg.DBHost, "db-host", "localhost", "database host")
	flag.StringVar(&cfg.DBPort, "db-port", "5432", "database port")
	flag.StringVar(&cfg.DBUserName, "db-user-name", "postgres", "db username")
	flag.StringVar(&cfg.DBPassword, "db-password", "pgpass", "db password")
	flag.StringVar(&cfg.DBName, "db-name", "postgres", "database name")
	flag.StringVar(&cfg.DBSLLMode, "db-ssl-mode", "disable", "database ssl mode switch")
	flag.StringVar(&cfg.DBDialect, "db-dialect", "postgres", "database dialect for gorm")

	// auth configs
	flag.StringVar(&cfg.TokenKey, "token-key", "something", "token key")
	flag.StringVar(&cfg.TokenIssuer, "token-issuer", "stone-tech", "issuer of token")
	flag.IntVar(&cfg.AccessTokenAliveTime, "access-token-alive-time", 1000000, "access-token-alive-time")
	flag.IntVar(&cfg.RefreshTokenAliveTime, "refresh-token-alive-time", 1000000, "refresh-token-alive-time")
	flag.IntVar(&cfg.OuthCookieAliveTime, "oauth-cookie-alive-time", 20, "oauth-cookie-alive-time")
	flag.IntVar(&cfg.PasswordResetInterval, "password-reset-interval", 10, "time interval to reset password when acc is locked")

	// otp config
	flag.IntVar(&cfg.OTPLength, "otp-length", 6, "otp length")
	flag.IntVar(&cfg.OTPAliveTime, "otp-alive-time", 6, "otp alive time")
	flag.StringVar(&cfg.OTPSender, "otp-sender", "sjnjaiswal2@gmail.com", "otp sender")

	// aws otp config
	flag.StringVar(&cfg.AWSAceesKeyId, "aws-access-key-id", "AKIATY2HYWVHFCHKYQMM", "aws access key id")
	flag.StringVar(&cfg.AWSSecretAccessKey, "aws-secret-access-key", "something", "token key")
	flag.StringVar(&cfg.AWSRegion, "aws-region", "something", "token key")
	flag.StringVar(&cfg.AWSSMTPUser, "aws-smtp-user", "something", "token key")
	flag.StringVar(&cfg.AWSSMTPPassword, "aws-smtp-password", "something", "token key")
	flag.StringVar(&cfg.AWSSenderEmail, "aws-sender-email", "something", "token key")
	flag.StringVar(&cfg.AWSHost, "aws-host", "something", "token key")
	flag.StringVar(&cfg.AWSPort, "aws-port", "something", "token key")

	// google Oauth configs
	flag.StringVar(&cfg.GoogleLoginEnpoint, "google-login-enpoint", "/auth/google/login", "Oauth google login enpoint")
	flag.StringVar(&cfg.GoogleCallbackEnpoint, "google-callback-enpoint", "/auth/google/callback", "Outh google callback enpoint")
	flag.StringVar(&cfg.GoogleRedirectURl, "google-redirect-uri", "http://localhost:9000/auth/google/callback", "Oauth google redirect URI")
	flag.StringVar(&cfg.GoogleClientId, "google-client-id", "997183094499-fidl2a9ol2gutlmjbpaadrnndbura862.apps.googleusercontent.com", "Oauth google client id")
	flag.StringVar(&cfg.GoogleClientSecret, "google-client-secret", "-7rHfJbWJW-RS50McT6g8iEG", "Oauth google client secret")

	// facebook Oauth configs
	flag.StringVar(&cfg.FacebookLoginEndPoint, "facebook-login-enpoint", "/auth/facebook/login", "Oauth facebook login enpoint")
	flag.StringVar(&cfg.FacebookCallbackEndPoint, "facebook-callback-enpoint", "/auth/facebook/callback", "Oauth facebook login enpoint")
	flag.StringVar(&cfg.FacebookClientId, "facebook-clientId", "694838257754877", "Oauth facebook client id")
	flag.StringVar(&cfg.FacebookClientSecret, "facebook-client-secret", "a5d062c0b2cf2082c50610365ff5ff57", "Oauth facebook client secret")
	flag.StringVar(&cfg.FacebookRedirectUrl, "facebook-redirect-url", "http://localhost:9000/auth/facebook/callback", "Oauth facebook redirect url")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	// logger.Info("failed to fetch URL",
	// 	// Structured context as strongly typed Field values.
	// 	zap.String("url", url),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
	cmd.RunServer(cfg, logger)

}
