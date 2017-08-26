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
	"Haku/orm"
	"Haku/orm/cockroach"
	"github.com/jinzhu/gorm"
	"time"
)

type Blog struct {
	Id       int64      `json:"id"`
	Title    *string    `json:"title" validate:"required"`
	Category *int8      `json:"category" validate:"required"`
	Abstract *string    `json:"abstract" validate:"required"`
	Content  *string    `json:"content" validate:"required"`
	Created  *time.Time `json:"created"`
}

func (Blog) TableName() string {
	return "blog"
}

type blogServiceProvider struct{}

var BlogService = &blogServiceProvider{}

func (b *blogServiceProvider) Create(conn orm.Connection, bo Blog) error {
	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Content)
	return db.Create(bo).Error
}

func (b *blogServiceProvider) GetList(conn orm.Connection, category int8) ([]Blog, error) {
	var (
		blog Blog
		list []Blog
	)

	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Content)

	err := db.Model(&blog).Where("category=?", category).Find(&list).Error

	return list, err
}
