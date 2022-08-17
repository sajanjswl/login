package v1

import "net/http"

// handles Oauth
func (s *RestServer) InitialRoutes() {
	s.Mux.Handle(s.cfg.LoginUIEndPoint, http.StripPrefix(s.cfg.LoginUIEndPoint, http.FileServer(http.Dir(s.cfg.HTMLTemplate))))
	s.Mux.HandleFunc(s.cfg.GoogleLoginEnpoint, s.oauthGoogleLogin)
	s.Mux.HandleFunc(s.cfg.GoogleCallbackEnpoint, s.oauthGoogleCallback)

}
