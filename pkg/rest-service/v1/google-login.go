package v1

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	guuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sajanjswl/auth/config"
	schema "github.com/sajanjswl/auth/data"
	"github.com/sajanjswl/auth/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func getGoogleOAuthConfig() *oauth2.Config {

	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),

		Endpoint: google.Endpoint,
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
}

func (s *RestServer) oauthGoogleLogin(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	OAuth2Config := getGoogleOAuthConfig()

	url := OAuth2Config.AuthCodeURL(oauthState)
	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (s *RestServer) oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {

	OAuth2Config := getGoogleOAuthConfig()
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	googleUserDetails, err := getUserDataFromGoogle(OAuth2Config, r.FormValue("code"))

	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	accestoken, refreshToken := saveGoogleOAuthDataToDB(s.Db, googleUserDetails)
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

func readDataAndSave(db *gorm.DB, data []byte) (string, string) {

	var dat map[string]interface{}

	if err := json.Unmarshal(data, &dat); err != nil {
		log.Panic(err)
	}
	var accessToken string
	var refreshToken string
	email := fmt.Sprintf("%v", dat["email"])

	user, err := utils.FindUser(db, config.FindUserByEmail, email)

	if user != nil {

		accessToken, refreshToken, err = utils.GetAccessTokenAndRefreshtoken(user.Email)
		if err != nil {
			log.Error(err)
		}

		err := utils.SaveAccessTokenAndRefreshToken(db, user.Email, accessToken, refreshToken)
		if err != nil {

			log.Error(err)
		}
	} else {

		accessToken, refreshToken, err := utils.GetAccessTokenAndRefreshtoken(email)
		if err != nil {
			log.Error(err)
		}
		user := &schema.DbUser{
			Uuid: guuid.New(),

			Email:          email,
			CreatedAt:      time.Now().Format(time.RFC3339),
			PendingDetails: true,
			IsVerified:     true,
			IsBlocked:      false,
			GoogleId:       fmt.Sprintf("%v", dat["id"]),
			RefreshToken:   refreshToken,
			AccessToken:    accessToken,
		}

		//registering user
		registeringUser := db.Table(schema.TableUsers).Create(&user)
		if registeringUser.Error != nil {
			log.Error(registeringUser.Error)

		}

	}

	return accessToken, refreshToken

}

func saveGoogleOAuthDataToDB(db *gorm.DB, outhUser *OAuthUserDetails) (string, string) {
	var accessToken string
	var refreshToken string
	user, err := utils.FindUser(db, config.FindUserByEmail, outhUser.Email)

	if user != nil {

		accessToken, refreshToken, err = utils.GetAccessTokenAndRefreshtoken(outhUser.Email)
		if err != nil {
			log.Error(err)
		}

		err := utils.SaveAccessTokenAndRefreshToken(db, outhUser.Email, accessToken, refreshToken)
		if err != nil {

			log.Error(err)
		}
	} else {

		accessToken, refreshToken, err := utils.GetAccessTokenAndRefreshtoken(outhUser.Email)
		if err != nil {
			log.Error(err)
		}
		user := &schema.DbUser{
			Uuid: guuid.New(),
			// FirstName:      outhUser.Name,
			Email:          outhUser.Email,
			CreatedAt:      time.Now().Format(time.RFC3339),
			PendingDetails: true,
			IsVerified:     true,
			IsBlocked:      false,
			GoogleId:       outhUser.ID,
			RefreshToken:   refreshToken,
			AccessToken:    accessToken,
		}

		//registering user
		registeringUser := db.Table(schema.TableUsers).Create(&user)
		if registeringUser.Error != nil {
			log.Error(registeringUser.Error)

		}

	}

	return accessToken, refreshToken

}
