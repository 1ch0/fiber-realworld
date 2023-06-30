package service

import (
	"context"
	"fmt"

	"github.com/1ch0/fiber-realworld/pkg/server/config"
)

var needInitData []DataInit

// DataInit the service set that needs init data
type DataInit interface {
	Init(ctx context.Context) error
}

// InitData init data
func InitData(ctx context.Context) error {
	for _, init := range needInitData {
		if err := init.Init(ctx); err != nil {
			return fmt.Errorf("database init failure %w", err)
		}
	}
	return nil
}

// InitServiceBean init all service instance
func InitServiceBean(c config.Config) []interface{} {
	// systemInfoService := NewSystemInfoService()
	// needInitData = []DataInit{systemInfoService}
	// return []interface{}{systemInfoService}
	return []interface{}{}
}