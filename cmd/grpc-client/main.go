package main

import (
	"context"
	"flag"
	"fmt"
	"time"
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"

	v1 "github.com/dezhab-service/pkg/api/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ExampleInternalCA-Root.crt")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
func main() {
	// get configuration

	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()


	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// transportOption = grpc.WithTransportCredentials(tlsCredentials)



	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// register(c, ctx)

	login(c, ctx)

	// otp(c, ctx)
	// verifyOTP(c, ctx)

	// reset(c, ctx)

	// requestTokens(c, ctx)

}

func register(c v1.UserServiceClient, ctx context.Context) {

	fmt.Println("I  was here")
	req := &v1.RegistrationRequest{
		ApiVersion: apiVersion,
		User: &v1.User{
			EmailID:      "sjnjaiswal1@gmail.com",
			Password:     "password1",
			FirstName:    "Sajan",
			LastName:     "Jaiswal",
			MobileNumber: "+917064274923",
		},
	}

	resp, err := c.Register(ctx, req)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp)

}

func reset(c v1.UserServiceClient, ctx context.Context) {

	req1 := &v1.ResetPasswordRequest{
		ApiVersion: apiVersion,
		EmailID:    "sjnjaiswal@gmail.com",
		Password:   "sajan4",
		OTP:        "632111",
	}
	res1, err := c.ResetPassword(ctx, req1)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}
	log.Printf("login result: <%+v>\n\n", res1)

}

func login(c v1.UserServiceClient, ctx context.Context) {

	req1 := &v1.LoginRequest{
		ApiVersion: apiVersion,
		EmailID:    "sjnjaiswal@gmail.com",
		Password:   "password",
	}
	res1, err := c.Login(ctx, req1)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}
	log.Printf("login result: <%+v>\n\n", res1)

}

func otp(c v1.UserServiceClient, ctx context.Context) {

	req1 := &v1.LoginWithOTPRequest{
		ApiVersion: apiVersion,
		EmailID:    "sjnjaiswal@gmail.com",
	}

	res1, err := c.OTP(ctx, req1)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}
	log.Printf("login result: <%+v>\n\n", res1)

}

func verifyOTP(c v1.UserServiceClient, ctx context.Context) {

	req1 := &v1.VerifyOTPRequest{
		ApiVersion: apiVersion,
		EmailID:    "sjnjaiswal@gmail.com",
		OTP:        "826559",
	}

	res1, err := c.VerifyOTP(ctx, req1)

	// res1, err := c.VerifyOTP()(ctx, req1)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}
	log.Printf("verify result: <%+v>\n\n", res1)

}

func requestTokens(c v1.UserServiceClient, ctx context.Context) {

	req1 := &v1.AccessTokenAndRefreshTokenRequest{
		ApiVersion:   apiVersion,
		AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJzam5qYWlzd2FsQGdtYWlsLmNvbSIsImV4cCI6MTU5MTY4ODc2MCwiaWF0IjoxNTkwNjg4NzYwLCJpc3MiOiJUZXNsYSIsIm5iZiI6MTU5MDY4ODc2MH0.bRsE5bz0js0DzS8vOFvZhAqMqzZHona_6PnPS3AVyRU",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJzam5qYWlzd2FsQGdtYWlsLmNvbSIsImV4cCI6MTU5MDc4ODc2MCwiaWF0IjoxNTkwNjg4NzYwLCJpc3MiOiJUZXNsYSIsIm5iZiI6MTU5MDY4ODc2MH0.V72BZLEf9AtyT0W2JRfgrUszXrVAtVs8_ImURQg1p1s",
	}

	res1, err := c.RequestTokens(ctx, req1)

	// res1, err := c.VerifyOTP()(ctx, req1)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}
	log.Printf(" requestTokens result: <%+v>\n\n", res1)

}
