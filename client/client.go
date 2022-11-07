package client

import "context"

type Client struct {
	db func(ctx context.Context)
}
