package count

import (
	"Haku/orm"
	"github.com/jinzhu/gorm"
	"Haku/orm/cockroach"
	"Haku/task"
)

type Info struct {
	Blog    int16    `json:"blog"`
	Visits  int16    `json:"visits"`
	Stars   int16    `json:"stars"`
}

type countServiceProvider struct {}

var CountService = &countServiceProvider{}

func (c *countServiceProvider) Counter(conn orm.Connection) (Info, error) {
	var (
		blogCount   int16
		info        Info
	)

	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Content)

	err := db.Table("blog").Count(&blogCount).Error

	info = Info{
		Blog: blogCount,
		Visits: 0,
		Stars: task.StarCount,
	}

	return info, err
}
