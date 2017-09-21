package task

import (
	"github.com/bamzi/jobrunner"
	"Haku/service/blog"
	"github.com/jinzhu/gorm"
	"Haku/orm/cockroach"
	"Haku/orm"
)

type StarTask struct {
	conn     orm.Connection
}

var StarCount   int16

func StartJob()  {
	var  sc       []blog.Blog

	conn, err := cockroach.DbConnPool.GetConnection()
	if err != nil {
		return
	}
	defer cockroach.DbConnPool.ReleaseConnection(conn)

	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Content)
	db.Table("blog").Select("star").Scan(&sc)

	for _, k := range sc {
		StarCount += k.Star
	}

	jobrunner.Start()
	jobrunner.Schedule("@every 60s", StarTask{db})
}

func (e StarTask) Run() {
	var (
		sc          []blog.Blog
		cc          int16
	)

	db := e.conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Content)
	db.Table("blog").Select("star").Scan(&sc)

	for _, k := range sc {
		cc += k.Star
	}

	StarCount = cc
}

