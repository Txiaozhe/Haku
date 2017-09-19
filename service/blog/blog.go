/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co,Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the 'Software'), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/08/20     Tang Xiaoji
 */

package blog

import (
	"Haku/mongo"
	"Haku/orm"
	"Haku/orm/cockroach"
	"Haku/utils"
	"github.com/jinzhu/gorm"
	"time"
	"gopkg.in/mgo.v2/bson"
)

// 接收
type BlogReq struct {
	Title    *string `json:"title" validate:"required"`
	Category int8    `json:"category" validate:"required"`
	Abstract *string `json:"abstract" validate:"required"`
	Tag      *string `json:"tag" validate:"required"`
	Content  *string `json:"content" validate:"required"`
}

// roach 创建
type Blog struct {
	Id        int64      `json:"id"`
	Title     *string    `json:"title"`
	Category  int8       `json:"category"`
	Abstract  *string    `json:"abstract"`
	Tag       *string    `json:"tag"`
	Contentid int32      `json:"contentid"`
	Created   *time.Time `json:"created"`
}

// mongo 创建
type BlogContent struct {
	ContentId int32   `bson:"contentid"`
	Content   *string `bson:"content"`
}

func (Blog) TableName() string {
	return "blog"
}

type blogServiceProvider struct{}

var BlogService = &blogServiceProvider{}

func (b *blogServiceProvider) Create(conn orm.Connection, bo BlogReq) (err error) {
	// 生成 contentid
	contid := utils.GetMillNanoId()

	// cockroach 存储列表项
	now := time.Now()
	blog := &Blog{
		Title:     bo.Title,
		Category:  bo.Category,
		Abstract:  bo.Abstract,
		Tag:       bo.Tag,
		Contentid: contid,
		Created:   &now,
	}

	tx := conn.(*gorm.DB).Begin().Exec("SET DATABASE = " + cockroach.Content)

	defer func() {
		if err != nil {
			err = tx.Rollback().Error
		} else {
			err = tx.Commit().Error
		}
	}()

	if err = tx.Create(blog).Error; err != nil {
		return
	}

	// mongo 存储内容
	collect := mongo.Db.C("blog")
	con := &BlogContent{
		ContentId: contid,
		Content:   bo.Content,
	}

	if err = mongo.Insert(collect, con); err != nil {
		return
	}

	return
}

func (b *blogServiceProvider) GetList(conn orm.Connection, category int8) ([]Blog, error) {
	var (
		blog  Blog
		list []Blog
	)

	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Content)

	err := db.Model(&blog).Where("category=?", category).Find(&list).Error

	return list, err
}

func (b *blogServiceProvider) GetAllBlog(conn orm.Connection) ([]Blog, error) {
	var (
		blog Blog
		list []Blog
	)

	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Content)

	err := db.Model(&blog).Find(&list).Error

	return list, err
}

func (b *blogServiceProvider) GetBlogDetail(id int32) (cont BlogContent, err error) {
	collect := mongo.Db.C("blog")
	err = mongo.GetUniqueOne(collect, bson.M{"contentid": id}, &cont)
	return
}
