package helpers

import (
	"fmt"
	"github.com/sqids/sqids-go"
	"sync"
)

func DefaultOptions() Options {
	return Options{
		MinLength: 10,
		Alphabet:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		Blocklist: []string{},
	}
}

var (
	encoder *sqids.Sqids
	once    sync.Once
	initErr error
)

type Options struct {
	MinLength uint8
	Alphabet  string
	Blocklist []string
}

func Initialize(opts Options) error {
	once.Do(func() {
		sqidsOpts := sqids.Options{
			MinLength: opts.MinLength,
			Alphabet:  opts.Alphabet,
			Blocklist: opts.Blocklist,
		}

		var err error
		encoder, err = sqids.New(sqidsOpts)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize sqids: %w", err)
		}
	})
	return initErr
}

func GetEncoder() (*sqids.Sqids, error) {
	if encoder == nil {
		return nil, fmt.Errorf("sqids encoder not initialized")
	}
	return encoder, nil
}

func MustGetEncoder() *sqids.Sqids {
	encoder, err := GetEncoder()
	if err != nil {
		panic(err)
	}
	return encoder
}

func Encode(numbers ...uint64) (string, error) {
	encoder, err := GetEncoder()
	if err != nil {
		return "", err
	}
	return encoder.Encode(numbers)
}

func Decode(id string) ([]uint64, error) {
	encoder, err := GetEncoder()
	if err != nil {
		return nil, err
	}
	return encoder.Decode(id), nil
}
