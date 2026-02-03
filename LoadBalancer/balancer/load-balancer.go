package balancer

import (
	"fmt"
	"net/http"

	"github.com/ishansaini194/Projects/server"
)

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []server.Server
}

func NewLoadBalancer(port string, servers []server.Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() server.Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}

	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	target := lb.getNextAvailableServer()
	fmt.Printf("Forwarding request to %s\n", target.Address())
	target.Serve(w, r)
}

func (lb *LoadBalancer) Port() string {
	return lb.port
}
