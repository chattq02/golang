package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID int64 `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment:'primary key is ID';"`
	RoleName string `gorm:"column:role_name;"`
	IsActive bool `gorm:"column:is_active; type:boolean;"`
	RoleNote string `gorm:"column:role_note; type:text;" `
}

func (y *Role) TableName() string { 
	return "go_db_role"
}