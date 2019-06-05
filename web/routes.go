package web

import "net/http"

func (srv *server) routes() http.Handler {
	srv.router.GET("simple/version", srv.Version)

	//Declare web routing table at here.
	return srv.router
}
