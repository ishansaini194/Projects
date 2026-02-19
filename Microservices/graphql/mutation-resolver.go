package graphql

import "context"

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateAccount(ctx context.Context, in AccountInput) (*Account, error) {

}

func (r *mutationResolver) CreateProduct(ctx context.Context, in AccountInput) (*Account, error) {

}

func (r *mutationResolver) CreateOrder(ctx context.Context, in AccountInput) (*Account, error) {

}
