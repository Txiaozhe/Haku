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
 *     Initial: 2017/08/15     Tang Xiaoji
 */

package admin

import (
	"Haku/orm"
	"github.com/jinzhu/gorm"
	"Haku/security"
	"errors"
	"Haku/orm/cockroach"
)

type Admin struct {
	Id     int64
	Name   *string `json:"name"  validate:"required"`
	Pass   *string `json:"pass"  validate:"required"`
	Github *string `json:"github" validate:"required"`
	Email  *string `json:"email" validate:"required"`
}

type AdminLogin struct {
	Name   *string `json:"name"  validate:"required"`
	Pass   *string `json:"pass"  validate:"required"`
}

var (
	errLoginFailed = errors.New("invalid username or password.")
	errPassNotMatch = errors.New("Old password not match.")
)

func (Admin) TableName() string {
	return "admin"
}

type adminServiceProvider struct{}

var AdminService = &adminServiceProvider{}

func (a *adminServiceProvider) Create(conn orm.Connection, admin Admin) error {
	saltedPass, err := security.SaltHashGenerate(admin.Pass)
	if err != nil {
		return err
	}

	var p string = string(saltedPass)

	adm := &Admin{
		Name: admin.Name,
        Pass: &p,
        Github: admin.Github,
		Email: admin.Email,
	}

	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Core)

	return db.Create(adm).Error
}

func (a *adminServiceProvider) Login(conn orm.Connection, name, pass *string) (int64, error) {
	adm := &Admin{}

	db := conn.(*gorm.DB).Exec("SET DATABASE = " + cockroach.Core)
	err := db.Model(adm).Where("name = ?", name).First(adm).Error
	if err != nil {
		return -1, err
	}

	var p string = *(adm.Pass)

	if security.SaltHashCompare([]byte(p), pass) {
		return adm.Id, nil
	}

	return -1, errLoginFailed
}
