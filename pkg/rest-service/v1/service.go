package v1

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sajanjswl/auth/config"
	"github.com/sajanjswl/auth/models"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type RestServer struct {
	Db     *gorm.DB
	Mux    *http.ServeMux
	cfg    *config.Config
	logger *zap.Logger
}

func NewRestServer(db *gorm.DB, cfg *config.Config, logger *zap.Logger) *RestServer {
	return &RestServer{
		Db:     db,
		cfg:    cfg,
		logger: logger,
	}
}

type OAuthUserDetails struct {
	ID         string
	Name       string
	Email      string
	IsVerified string
}

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (s *RestServer) oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	Oauth2Config := &oauth2.Config{
		ClientID:     s.cfg.GoogleClientId,
		ClientSecret: s.cfg.GoogleClientSecret,
		RedirectURL:  s.cfg.GoogleRedirectURl,

		Endpoint: google.Endpoint,
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
	}

	url := Oauth2Config.AuthCodeURL(oauthState)
	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (s *RestServer) oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {

	Oauth2Config := &oauth2.Config{
		ClientID:     s.cfg.GoogleClientId,
		ClientSecret: s.cfg.GoogleClientSecret,
		RedirectURL:  s.cfg.GoogleRedirectURl,

		Endpoint: google.Endpoint,
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	googleUserDetails, err := getUserDataFromGoogle(Oauth2Config, r.FormValue("code"))

	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	accestoken, refreshToken, _ := s.saveGoogleOAuthDataToDB(s.Db, googleUserDetails)
	cookie := http.Cookie{
		Name:  "AccessToken, RefreshToken",
		Value: accestoken + refreshToken,
		Path:  "/",
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/v1/home", http.StatusPermanentRedirect)

}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(OAuth2Config *oauth2.Config, code string) (*OAuthUserDetails, error) {
	// Use code to get token and get user info from Google.

	var googleUserDetails OAuthUserDetails
	token, err := OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	decoder := json.NewDecoder(response.Body)
	decoderErr := decoder.Decode(&googleUserDetails)

	defer response.Body.Close()

	if decoderErr != nil {
		log.Error("Error occurred while getting information from Google")
		return nil, decoderErr
	}
	return &googleUserDetails, nil
}

func (s *RestServer) saveGoogleOAuthDataToDB(db *gorm.DB, outhUser *OAuthUserDetails) (string, string, error) {
	var accessToken, refreshToken string

	user := &models.User{}

	if err := models.GetUser(s.Db, user, outhUser.Email); err == nil {
		s.logger.Info("user  exists in database", zap.String("email", user.Email))
		return accessToken, refreshToken, nil
	}

	user.Email = outhUser.Email
	user.Name = outhUser.Name

	if err := models.CreateUser(s.Db, user); err != nil {
		s.logger.Warn("Failed to register user", zap.String("email", user.Email), zap.Error(err))

		return "", "", err
	}

	return accessToken, refreshToken, nil

}
