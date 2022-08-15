package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sajanjswl/auth/config"
	schema "github.com/sajanjswl/auth/data"
	"github.com/sajanjswl/auth/utils"
	guuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	facebookOAuth "golang.org/x/oauth2/facebook"
)

type OAuthUserDetails struct {
	ID         string
	Name       string
	Email      string
	IsVerified string
}

const oauthFacebookUrlAPI = "https://graph.facebook.com/me?fields=id,name,email&access_token="

// GetFacebookOAuthConfig will return the config to call facebook Login
func GetFacebookOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("FAEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("FACEBOOK_REDIRECT_URL"),
		Endpoint:     facebookOAuth.Endpoint,
		Scopes:       []string{"email"},
	}
}

func (s *RestServer) oauthFacebookLogin(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	OAuth2Config := GetFacebookOAuthConfig()

	url := OAuth2Config.AuthCodeURL(oauthState)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func (s *RestServer) oauthFacebookCallback(w http.ResponseWriter, r *http.Request) {

	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")
	var OAuth2Config = GetFacebookOAuthConfig()
	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth facebook state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	facebookUserDetails, err := getUserDataFromFacebook(OAuth2Config, r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	accestoken, refreshToken := saveFacebookOAuthDataToDB(s.Db, facebookUserDetails)
	cookie := http.Cookie{
		Name:  "AccessToken, RefreshToken",
		Value: accestoken + refreshToken,
		Path:  "/",
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/v1/home", http.StatusPermanentRedirect)

}

func getUserDataFromFacebook(OAuth2Config *oauth2.Config, code string) (*OAuthUserDetails, error) {
	// Use code to get token and get user info from Google.
	//var fbUserDetails OAuthUserDetails

	var fbUserDetails OAuthUserDetails
	token, err := OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(oauthFacebookUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	decoder := json.NewDecoder(response.Body)
	decoderErr := decoder.Decode(&fbUserDetails)

	defer response.Body.Close()

	if decoderErr != nil {
		log.Error("Error occurred while getting information from Facebook")
		return nil, decoderErr
	}

	return &fbUserDetails, nil
}

func saveFacebookOAuthDataToDB(db *gorm.DB, outhUser *OAuthUserDetails) (string, string) {
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
			FacebookId:     outhUser.ID,
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
