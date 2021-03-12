package v1

import (
	"net/http"
	"os"
)

// handles Oauth
func (s *RestServer) InitialRoutes() {

	s.Mux.Handle(os.Getenv("LOGIN_UI"), http.StripPrefix(os.Getenv("LOGIN_UI"), http.FileServer(http.Dir(os.Getenv("TEMPLATES_HTML")))))
	s.Mux.HandleFunc(os.Getenv("GOOGLE_LOGIN_ENDPOINT"), s.oauthGoogleLogin)
	s.Mux.HandleFunc(os.Getenv("GOOGLE_CALLBACK_ENDPOINT"), s.oauthGoogleCallback)

	s.Mux.HandleFunc(os.Getenv("FACEBOOK_LOGIN_END_POINT"), s.oauthFacebookLogin)
	s.Mux.HandleFunc(os.Getenv("FACEBOOK_CALLBACK_END_POINT"), s.oauthFacebookCallback)

}
