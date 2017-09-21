package count

import (
	"github.com/labstack/echo"
	"Haku/orm/cockroach"
	"Haku/general"
	"Haku/general/errorcode"
	"net/http"
	"Haku/service/count"
)

func Counter(c echo.Context) error {
	conn, err := cockroach.DbConnPool.GetConnection()
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrDBConnection, err.Error())
	}
	defer cockroach.DbConnPool.ReleaseConnection(conn)

	info, err := count.CountService.Counter(conn)
	if err != nil {
		return general.NewErrorWithMessage(errorcode.ErrInternalServer, err.Error())
	}

	return c.JSON(http.StatusOK, &info)
}
