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
	"Haku/general"
	"Haku/general/errorcode"
	"Haku/orm/cockroach"
	"github.com/labstack/echo"
	"net/http"
	"Haku/middleware/jwt"
	"Haku/service/admin"
)

func Create(c echo.Context) error {
	var (
		err   error
		a admin.Admin
	)

	if err = c.Bind(&a); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	if err = c.Validate(a); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	conn, err := cockroach.DbConnPool.GetConnection()
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBConnection, err.Error())
	}
	defer cockroach.DbConnPool.ReleaseConnection(conn)

	err = admin.AdminService.Create(conn, a)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBOperationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func Login(c echo.Context) error {
	var (
		err   error
		a     admin.AdminLogin
	)

	if err = c.Bind(&a); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	if err = c.Validate(a); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	conn, err := cockroach.DbConnPool.GetConnection()
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBConnection, err.Error())
	}
	defer cockroach.DbConnPool.ReleaseConnection(conn)

	id, err := admin.AdminService.Login(conn, a.Name, a.Pass)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBOperationFailed, err.Error())
	}

	key, token, err := jwt.NewToken(id, *a.Name)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInternalServer, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{key: token})
}
