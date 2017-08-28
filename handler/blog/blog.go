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
	"github.com/labstack/echo"
	"net/http"
	"Haku/service/blog"
	"Haku/general"
	"Haku/general/errorcode"
	"Haku/orm/cockroach"
)

func Create(c echo.Context) error {
	var (
		err   error
		b     blog.BlogReq
	)

	if err = c.Bind(&b); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	if err = c.Validate(b); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	conn, err := cockroach.DbConnPool.GetConnection()
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBConnection, err.Error())
	}
	defer cockroach.DbConnPool.ReleaseConnection(conn)

	err = blog.BlogService.Create(conn, b)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBOperationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func GetList(c echo.Context) error {
	type cate struct {
		Category  int8   `json:"category"`
	}

	var (
		err error
		blogs []blog.Blog
		ca    cate
	)

	if err = c.Bind(&ca); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	if err = c.Validate(ca); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	conn, err := cockroach.DbConnPool.GetConnection()
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBConnection, err.Error())
	}
	defer cockroach.DbConnPool.ReleaseConnection(conn)

	blogs, err = blog.BlogService.GetList(conn, ca.Category)
	if err != nil {
		general.NewErrorWithMessage(errorcode.ErrInternalServer, err.Error())
	}

	return c.JSON(http.StatusOK, blogs)
}

func GetBlogDetail(c echo.Context) error {
	type id struct {
		Id     int32   `json:"id"`
	}

	var (
		err  error
		i    id
	)

	if err = c.Bind(&i); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	if err = c.Validate(i); err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInvalidParams, err.Error())
	}

	cont, err := blog.BlogService.GetBlogDetail(i.Id)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBOperationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, cont)
}
