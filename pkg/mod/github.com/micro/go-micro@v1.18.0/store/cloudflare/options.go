package cloudflare

import (
	"github.com/micro/go-micro/config/options"
)

// Token sets the cloudflare api token
func Token(t string) options.Option {
	// TODO: change to store.cf.api_token
	return options.WithValue("CF_API_TOKEN", t)
}

// Account sets the cloudflare account id
func Account(id string) options.Option {
	// TODO: change to store.cf.account_id
	return options.WithValue("CF_ACCOUNT_ID", id)
}

// Namespace sets the KV namespace
func Namespace(ns string) options.Option {
	// TODO: change to store.cf.namespace
	return options.WithValue("KV_NAMESPACE_ID", ns)
}
