// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User User information table
type User struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:PK" json:"id"`                                                // PK
	UserName  string         `gorm:"column:user_name;not null;comment:User name" json:"user_name"`                                                // User name
	Email     string         `gorm:"column:email;not null;comment:User email" json:"email"`                                                       // User email
	Password  string         `gorm:"column:password;not null;comment:User password" json:"password"`                                              // User password
	Address   string         `gorm:"column:address;not null;comment:User wallet address" json:"address"`                                          // User wallet address
	Private   string         `gorm:"column:private;not null;comment:User wallet private key" json:"private"`                                      // User wallet private key
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:User information create time" json:"created_at"` // User information create time
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:User information update time" json:"updated_at"` // User information update time
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:User information delete time" json:"deleted_at"`                                    // User information delete time
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
