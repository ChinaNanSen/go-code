// Package consul is a consul implementation of kv
package consul

import (
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro/config/options"
	"github.com/micro/go-micro/store"
)

type ckv struct {
	options.Options
	client *api.Client
}

func (c *ckv) Read(keys ...string) ([]*store.Record, error) {
	var records []*store.Record

	for _, key := range keys {
		keyval, _, err := c.client.KV().Get(key, nil)
		if err != nil {
			return nil, err
		}

		if keyval == nil {
			return nil, store.ErrNotFound
		}

		records = append(records, &store.Record{
			Key:   keyval.Key,
			Value: keyval.Value,
		})
	}

	return records, nil
}

func (c *ckv) Delete(keys ...string) error {
	var err error
	for _, key := range keys {
		if _, err = c.client.KV().Delete(key, nil); err != nil {
			return err
		}
	}
	return nil
}

func (c *ckv) Write(records ...*store.Record) error {
	var err error
	for _, record := range records {
		_, err = c.client.KV().Put(&api.KVPair{
			Key:   record.Key,
			Value: record.Value,
		}, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ckv) List() ([]*store.Record, error) {
	keyval, _, err := c.client.KV().List("/", nil)
	if err != nil {
		return nil, err
	}
	if keyval == nil {
		return nil, store.ErrNotFound
	}
	var vals []*store.Record
	for _, keyv := range keyval {
		vals = append(vals, &store.Record{
			Key:   keyv.Key,
			Value: keyv.Value,
		})
	}
	return vals, nil
}

func (c *ckv) String() string {
	return "consul"
}

func NewStore(opts ...options.Option) store.Store {
	options := options.NewOptions(opts...)
	config := api.DefaultConfig()

	var nodes []string

	if n, ok := options.Values().Get("store.nodes"); ok {
		nodes = n.([]string)
	}

	// set host
	if len(nodes) > 0 {
		addr, port, err := net.SplitHostPort(nodes[0])
		if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
			port = "8500"
			config.Address = fmt.Sprintf("%s:%s", nodes[0], port)
		} else if err == nil {
			config.Address = fmt.Sprintf("%s:%s", addr, port)
		}
	}

	client, _ := api.NewClient(config)

	return &ckv{
		Options: options,
		client:  client,
	}
}
