// Package memory is a in-memory store store
package memory

import (
	"sync"
	"time"

	"github.com/micro/go-micro/config/options"
	"github.com/micro/go-micro/store"
)

type memoryStore struct {
	options.Options

	sync.RWMutex
	values map[string]*memoryRecord
}

type memoryRecord struct {
	r *store.Record
	c time.Time
}

func (m *memoryStore) List() ([]*store.Record, error) {
	m.RLock()
	defer m.RUnlock()

	//nolint:prealloc
	var values []*store.Record

	for _, v := range m.values {
		// get expiry
		d := v.r.Expiry
		t := time.Since(v.c)

		if d > time.Duration(0) {
			// expired
			if t > d {
				continue
			}
			// update expiry
			v.r.Expiry -= t
			v.c = time.Now()
		}

		values = append(values, v.r)
	}

	return values, nil
}

func (m *memoryStore) Read(keys ...string) ([]*store.Record, error) {
	m.RLock()
	defer m.RUnlock()

	//nolint:prealloc
	var records []*store.Record

	for _, key := range keys {
		v, ok := m.values[key]
		if !ok {
			return nil, store.ErrNotFound
		}

		// get expiry
		d := v.r.Expiry
		t := time.Since(v.c)

		// expired
		if d > time.Duration(0) {
			if t > d {
				return nil, store.ErrNotFound
			}
			// update expiry
			v.r.Expiry -= t
			v.c = time.Now()
		}

		records = append(records, v.r)
	}

	return records, nil
}

func (m *memoryStore) Write(records ...*store.Record) error {
	m.Lock()
	defer m.Unlock()

	for _, r := range records {
		// set the record
		m.values[r.Key] = &memoryRecord{
			r: r,
			c: time.Now(),
		}
	}

	return nil
}

func (m *memoryStore) Delete(keys ...string) error {
	m.Lock()
	defer m.Unlock()

	for _, key := range keys {
		// delete the value
		delete(m.values, key)
	}

	return nil
}

// NewStore returns a new store.Store
func NewStore(opts ...options.Option) store.Store {
	options := options.NewOptions(opts...)

	return &memoryStore{
		Options: options,
		values:  make(map[string]*memoryRecord),
	}
}
