package Smoe

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Smoe) Post(c echo.Context) error {
	cid, _ := IsNum(c.Param("cid"))
	data := struct {
		Post     []Contents
		TestPost Contents
	}{
		s.QueryWithCid(cid),
		s.TestQueryPostWithCid(cid),
	}
	return c.Render(http.StatusOK, "testpost.template", data)
}
