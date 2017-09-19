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
	"Haku/general"
	"Haku/general/errorcode"
	"Haku/orm/cockroach"
	"Haku/service/badge"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Create(c echo.Context) error {
	var (
		err error
		ba  badge.Badge
	)

	if err = c.Bind(&ba); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	if err = c.Validate(ba); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	conn, err := cockroach.DbConnPool.GetConnection()
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBConnection, err.Error())
	}
	defer cockroach.DbConnPool.ReleaseConnection(conn)

	err = badge.BadgeService.Create(conn, ba)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBOperationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func GetBadgeByBlogId(c echo.Context) error {
	var (
		err       error
		ba        badge.BadgeById
		badgelist []badge.Badge
	)

	if err = c.Bind(&ba); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	if err = c.Validate(ba); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	conn, err := cockroach.DbConnPool.GetConnection()
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBConnection, err.Error())
	}
	defer cockroach.DbConnPool.ReleaseConnection(conn)

	i64, err := strconv.ParseInt(ba.Blogid, 10, 64)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInternalServer, err.Error())
	}

	badgelist, err = badge.BadgeService.GetBadgeByBlogId(conn, i64)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBOperationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, badgelist)
}
