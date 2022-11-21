package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (sqlModel *SQLModel) Mask(dbType int) {
	uid := NewUID(uint32(sqlModel.Id), dbType, 1)
	sqlModel.FakeId = &uid
}

func (sqlModel *SQLModel) PrepareForInsert() {
	now := time.Now().UTC()
	sqlModel.Id = 0
	sqlModel.Status = 1

	sqlModel.CreatedAt = &now
	sqlModel.UpdatedAt = &now
}
