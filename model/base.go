package model

import (
	"errors"
	"time"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time `gorm:"comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'"`
	CreateBy  uint64    `gorm:"<-:create;comment:'创建人id'"`
	UpdateBy  uint64    `gorm:"<-;comment:'更新人id'"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	userId := tx.Statement.Context.Value("userId")
	if userId == nil {
		return errors.New("userId is nil")
	}
	userLongId := cast.ToUint64(userId)
	tx.Statement.SetColumn("CreateBy", userLongId)
	m.CreateBy = userLongId
	return nil
}

func (m *BaseModel) BeforeSave(tx *gorm.DB) error {
	userId := tx.Statement.Context.Value("userId")
	if userId == nil {
		return errors.New("userId is nil")
	}
	userLongId := cast.ToUint64(userId)
	tx.Statement.SetColumn("UpdateBy", userLongId)
	m.UpdateBy = userLongId
	return nil
}
