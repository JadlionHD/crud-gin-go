package migration

import (
	"database/sql/driver"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type PostDB struct {
	gorm.Model
	Title     string
	Body      string
	Tags      PostTags        `gorm:"type:text[]"`
	Reactions *PostReactionDB `gorm:"embedded"`
	Views     uint
	UserID    uint
}

type PostTags []string

func (o *PostTags) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*o = strings.Split(string(bytes), ",")
	return nil
}
func (o PostTags) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}
	return strings.Join(o, ","), nil
}

type PostReactionDB struct {
	Likes    uint
	Dislikes uint
}
