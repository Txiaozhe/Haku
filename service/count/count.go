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

type Visit struct {
	Vis     int16    `json:"vis"`
}

type countServiceProvider struct {}

var CountService = &countServiceProvider{}

func (c *countServiceProvider) Counter(conn orm.Connection) (*Info, error) {
	var (
		db          *gorm.DB
		blogCount   int16
		info        Info
		v           Visit
	)

	// 统计 博客数和 star 数
	db = conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Content)
	err := db.Table("blog").Count(&blogCount).Error

	// 统计访问数
	db = conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Core)
	err = db.Table("visit").Model(&v).Update("vis", gorm.Expr("vis + ?", 1)).Limit(1).Error
	if err != nil {
		return nil, err
	}

	err = db.Table("visit").Model(&v).First(&v).Error

	info = Info{
		Blog: blogCount,
		Visits: v.Vis,
		Stars: task.StarCount,
	}

	return &info, err
}
