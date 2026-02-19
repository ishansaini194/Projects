package graphql

import "context"

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query *string, id *string) ([]*Product, error) {
}
