package main

import (
	"fmt"
	"net/http"

	"github.com/ishansaini194/Projects/balancer"
	"github.com/ishansaini194/Projects/server"
	"github.com/ishansaini194/Projects/utils"
)

func main() {
	servers := []server.Server{
		server.NewSimpleServer("https://www.facebook.com"),
		server.NewSimpleServer("https://www.google.com"),
		server.NewSimpleServer("https://duckduckgo.com"),
	}
	lb := balancer.NewLoadBalancer("8000", servers)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.ServeProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Load Balancer started at :%s\n", lb.Port())
	err := http.ListenAndServe(":"+lb.Port(), nil)
	utils.HandleErr(err)
}
