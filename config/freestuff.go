package config

import (
	"context"
	"fmt"
	"sync"

	"github.com/caarlos0/env/v10"
)

var (
	freeStuffConfigInstance *FreeStuffConfig
	freeStuffConfigErr      error
	freeStuffConfigOnce     sync.Once
)

type FreeStuffConfig struct {
	APIKey        string `env:"FREESTUFF_API_KEY,required"`
	WebhookSecret string `env:"FREESTUFF_WEBHOOK_SECRET,required"`
}

// GetFreeStuffConfig gets the configuration used to interact with the FreeStuff API.
func GetFreeStuffConfig(_ context.Context) (*FreeStuffConfig, error) {
	freeStuffConfigOnce.Do(func() {
		freeStuffConfigInstance = &FreeStuffConfig{}

		if err := env.Parse(freeStuffConfigInstance); err != nil {
			freeStuffConfigErr = err
			return
		}
	})

	if freeStuffConfigErr != nil {
		return nil, fmt.Errorf("failed to build FreeStuff API configuration: %w", freeStuffConfigErr)
	}

	return freeStuffConfigInstance, nil
}
