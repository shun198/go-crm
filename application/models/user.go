package models

import (
	"time"
)

// https://gorm.io/docs/models.html
type User struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	EmployeeNumber string `gorm:"uniqueIndex;not null;size:8" json:"employeeNumber"`
	Email          string `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Name           string `gorm:"not null;size:255" json:"name"`
	Password       string `gorm:"not null;size:255" json:"-"`
	IsVerified     bool   `gorm:"not null" json:"is_verified"`
	Role           uint8  `gorm:"not null" json:"role"`
	IsActive       bool   `gorm:"not null" json:"is_active"`
	// CreatedAt and UpdatedAt are special fields that GORM automatically populates with the current time when a record is created or updated
	CreatedAt   time.Time `gorm:"- autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"- autoUpdateTime" json:"-"`
	CreatedByID *uint     `gorm:"" json:"-"`
	CreatedBy   *User     `gorm:"not null" json:"-"`
	UpdatedByID *uint     `gorm:"" json:"-"`
	UpdatedBy   *User     `gorm:"not null" json:"-"`
}
