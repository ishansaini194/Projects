package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/ishansaini194/Projects/utils"
)

type SimpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func NewSimpleServer(addr string) *SimpleServer {
	serverURL, err := url.Parse(addr)
	utils.HandleErr(err)

	return &SimpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

func (s *SimpleServer) Address() string {
	return s.addr
}

func (s *SimpleServer) IsAlive() bool {
	return true // can be extended later
}

func (s *SimpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}
