package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model

	Name   string `gorm:"type:varchar(255);not null;uniqueIndex" json:"name"`
	Repo   string `gorm:"type:varchar(512);not null" json:"repo"`
	Domain string `gorm:"type:varchar(255);not null;uniqueIndex" json:"domain"`
	Ref    string `gorm:"type:varchar(100);default:''" json:"ref"`

	Creator string         `gorm:"type:varchar(64);not null" json:"creator"`
	Admins  datatypes.JSON `gorm:"type:json;not null" json:"admins"`

	Status     string    `gorm:"type:varchar(50);not null;default:pending;index;check:status IN ('pending', 'deploying', 'running', 'error')" json:"status"`
	Version    string    `gorm:"type:varchar(100)" json:"version"`
	LastDeploy time.Time `gorm:"type:datetime;default:null" json:"last_deploy"`

	Description string `gorm:"type:varchar(1024)" json:"description"`

	Cluster string `gorm:"type:varchar(255);not null" json:"cluster"`
}
