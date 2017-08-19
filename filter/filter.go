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
 *     Initial: 2017/08/19     Tang Xiaoji
 */

package filter

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"Haku/middleware/jwt"
	"Haku/utils"
	"Haku/general"
	"Haku/general/errorcode"
)

func MustLogin(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
	    admin := c.Get("user").(*jwtgo.Token) // 获取 token
	    claims := admin.Claims.(jwtgo.MapClaims)

	    aId := claims[jwt.ClaimAdminId].(string)
	    adminName := claims[jwt.ClaimAdminName]

	    adminId, err := utils.SerialConvertInt64(aId)
	    if err != nil {
		    return general.NewErrorWithMessage(errorcode.ErrInternalServer, err.Error())
	    }

	    c.Set(jwt.ClaimAdminId, adminId)
	    c.Set(jwt.ClaimAdminName, adminName)

	    return next(c)
    }
}
