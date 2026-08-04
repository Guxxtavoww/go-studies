package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityId uuid.UUID

func (id EntityId) IsNil() bool {
	return id == EntityId(uuid.Nil)
}

func NewEntityId() EntityId {
	return EntityId(uuid.New())
}

type BaseModel struct {
	Id        EntityId `gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Id.IsNil() {
		base.Id = NewEntityId()
	}

	return
}
