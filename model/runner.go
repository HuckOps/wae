package model

import "gorm.io/gorm"

type Runner struct {
	gorm.Model

	UUID string `gorm:"primaryKey"`
	OS   string `gorm:"column:os"`
	Arch string `gorm:"column:arch"`
	CPU  int32  `gorm:"column:cpu"`

	Version string `gorm:"column:version"`
}
