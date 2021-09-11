package config

import (
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetConf() gorm.Option {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: strings.NewReplacer("CP", "Cp"),
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
}
