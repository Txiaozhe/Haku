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

package route

import (
	"github.com/labstack/echo"
	"Haku/handler"
	"Haku/filter"
	"Haku/middleware/jwt"
	"Haku/handler/admin"
	"Haku/handler/blog"
)

func InitRoute(server *echo.Echo, tokenKey string) {
	if server == nil {
		panic("[InitRouter], server couldn't be nil")
	}

	jwt.InitJWTWithToken(tokenKey)

	// admin
	server.POST("/api/v1/admin/create", admin.Create)
	server.POST("/api/v1/admin/login", admin.Login)

	// blog
	server.GET("/api/v1/blog/github/get", blog.GetListFromGitHub)
	server.GET("/api/v1/blog/github/label/get", blog.GetLabelFromGitHub)
	server.POST("/api/v1/blog/github/detail", blog.GetDetailFromGitHub)

	server.POST("/api/v1/blog/create", blog.Create, filter.MustLogin)
	server.GET("/api/v1/checkloginstatus", handler.CheckLoginStatus, filter.MustLogin)
}
