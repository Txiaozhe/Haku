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

package jwt

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
	"strconv"
)

const (
	tokenExpireInHour = 48

	ClaimAdminId = "id"
	ClaimAdminName  =  "name"
	ClaimExpire   =  "exp"

	respTokenKey  =  "token"
)

var (
	TokenHMACKey string
	urlMap map[string]struct{}
)

func CustomSkipper(c echo.Context) bool {
	if _, ok := urlMap[c.Request().RequestURI]; ok {
		return true
	}

	return false
}

func CustomJWT(tokenKey string) echo.MiddlewareFunc  {
    jwtConfig := middleware.JWTConfig{
        Skipper: CustomSkipper,
	    SigningKey: []byte(tokenKey),
    }

	return middleware.JWTWithConfig(jwtConfig)
}

func InitJWTWithToken(token string) {
	TokenHMACKey = token
}

func NewToken(id int64, name string) (string, string, error) {
    token := jwtgo.New(jwtgo.SigningMethodHS256)

	claims := token.Claims.(jwtgo.MapClaims)
	claims[ClaimAdminId] = strconv.FormatInt(id, 10)
	claims[ClaimAdminName] = name
	claims[ClaimExpire] = time.Now().Add(time.Hour * tokenExpireInHour).Unix()

	t, err := token.SignedString([]byte(TokenHMACKey))

	return respTokenKey, t, err
}

func GetAdmin(c echo.Context) (int64, string)  {
	return c.Get(ClaimAdminId).(int64), c.Get(ClaimAdminName).(string)
}

func init()  {
	urlMap = make(map[string]struct{})

	urlMap["/admin/login"] = struct{}{}
}
