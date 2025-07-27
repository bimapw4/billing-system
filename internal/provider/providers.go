package provider

import (
	"billing/bootstrap"
)

type Provider struct {
}

func NewProvider(cfg bootstrap.Providers) Provider {
	return Provider{}
}
