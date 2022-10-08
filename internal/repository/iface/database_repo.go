package iface

import "gorm.io/gorm"

type Repo interface {
	GetDbR() *gorm.DB
	GetDbW() *gorm.DB
	DbRClose() error
	DbWClose() error
}
