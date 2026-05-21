package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model

	Name   string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Repo   string `gorm:"type:varchar(512);not null"`
	Domain string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Ref    string `gorm:"type:varchar(100);not null;default:main"`

	Creator string         `gorm:"type:varchar(64);not null"`
	Admins  datatypes.JSON `gorm:"type:json;not null;default:'[]'"`

	Status     string `gorm:"type:varchar(50);not null;default:draft;index"`
	Version    string `gorm:"type:varchar(100)"`
	LastDeploy int64  `gorm:"type:bigint;default:0"`

	Description string `gorm:"type:varchar(1024)"`
}
