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
 *     Initial: 2017/09/19     Tang Xiaoji
 */

package badge

import (
	"time"
	"Haku/orm"
	"github.com/jinzhu/gorm"
	"Haku/orm/cockroach"
)

type BadgeReq struct {
	Blogid     int64      `json:"blogid" validate:"required"`
	Name       *string     `json:"name" validate:"required"`
	Avatar     *string     `json:"avatar" validate:"required"`
	Content    *string     `json:"content" validate:"required"`
	Created    *time.Time  `json:"created"`
}

func (BadgeReq) TableName() string {
	return "badge"
}

type badgeServiceProvider struct {}

var BadgeService = &badgeServiceProvider{}

func (c *badgeServiceProvider) Create(conn orm.Connection, ba BadgeReq) error {
	now := time.Now()
	badge := &BadgeReq{
		Blogid: ba.Blogid,
		Name: ba.Name,
		Avatar: ba.Avatar,
		Content: ba.Content,
		Created: &now,
	}

	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Grade)

	if err := db.Create(badge).Error; err != nil {
		return err
	}

	return nil
}
