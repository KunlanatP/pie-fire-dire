package main

import (
	"net/http"

	"github.com/kunlanat/pie-fire-dire/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	apiGroup := e.Group("/api")
	v1Group := apiGroup.Group("/v1")

	serv := service.BeefServiceImpl()
	ctl := beefCtl{svc: serv}
	v1Group.GET("/beef/summary", ctl.GetBeef)
	e.Logger.Fatal(e.Start(":1323"))
}

type beefCtl struct {
	svc service.BeefService
}

func (s *beefCtl) GetBeef(ctx echo.Context) error {
	counter, err := s.svc.GetBeef(ctx.Request().Context())
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Bad Request")
	}
	return ctx.JSON(http.StatusOK, counter)
}
