package lookup

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo, l *LookupService) {

	e.GET("/aircraft", func(c echo.Context) error {

		nnumber := c.QueryParam("n")
		registrantName := c.QueryParam("registrant_name")
		sameRegistrantName := c.QueryParam("same_registrant_name")

		// todo: clean up query param logic!
		if nnumber != "" {
			if sameRegistrantName == "true" {
				return c.JSONBlob(http.StatusOK, l.AugmentToBytes(l.GetOtherAircraftByRegistrantName(nnumber)))
			}
			return c.JSONBlob(http.StatusOK, l.AugmentToBytes(l.GetAircraftByNNumber(nnumber)))
		}
		if registrantName != "" {
			return c.JSONBlob(http.StatusOK, l.AugmentToBytes(l.GetAircraftByRegistrantName(registrantName)))
		}
		return c.JSON(http.StatusBadRequest, "request not valid")
	})

	e.GET("/dereg", func(c echo.Context) error {
		nnumber := c.QueryParam("n")
		sameSerialNumber := c.QueryParam("same_serial_number")

		if nnumber != "" {
			if sameSerialNumber == "true" {
				return c.JSONBlob(http.StatusOK, l.AugmentToBytes(l.GetOtherDeregBySerialNumber(nnumber)))
			}
			return c.JSONBlob(http.StatusOK, l.AugmentToBytes(l.GetDeregByNNumber(nnumber)))
		}
		return c.JSON(http.StatusBadRequest, "request not valid")
	})
}
