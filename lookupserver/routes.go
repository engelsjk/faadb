package lookupserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo, l *LookupService) {

	e.GET("/aircraft", func(c echo.Context) error {

		nNumber := c.QueryParam("n")
		registrantName := c.QueryParam("registrant_name")
		sameRegistrantName := c.QueryParam("same_registrant_name")

		// todo: clean up query param logic!

		if nNumber != "" {
			if sameRegistrantName == "true" {
				r, _ := l.GetOtherAircraftByRegistrantName(nNumber)
				return c.JSONBlob(http.StatusOK, l.AugmentToBytes(r))
			}
			r, _ := l.GetAircraftByNNumber(nNumber)
			return c.JSONBlob(http.StatusOK, l.AugmentToBytes(r))
		}
		if registrantName != "" {
			r, _ := l.GetAircraftByRegistrantName(registrantName)
			return c.JSONBlob(http.StatusOK, l.AugmentToBytes(r))
		}
		return c.JSON(http.StatusBadRequest, "request not valid")
	})

	e.GET("/aircraft2", func(c echo.Context) error {
		nNumber := c.QueryParam("n")
		serialNumber := c.QueryParam("sn")
		sameSerialNumber := c.QueryParam("same_serial_number")

		if nNumber != "" {
			if sameSerialNumber == "true" {
				r, _ := l.GetOtherAircraftBySerialNumber2(nNumber)
				return c.JSONBlob(http.StatusOK, ToBytes(r))
			}
			r, _ := l.GetAircraftByNNumber2(nNumber)
			return c.JSONBlob(http.StatusOK, ToBytes(r))
		}
		if serialNumber != "" {
			r, _ := l.GetAircraftBySerialNumber2(serialNumber)
			return c.JSONBlob(http.StatusOK, ToBytes(r))
		}
		return c.JSON(http.StatusBadRequest, "request not valid")
	})

	e.GET("/dereg", func(c echo.Context) error {
		nNumber := c.QueryParam("n")
		sameSerialNumber := c.QueryParam("same_serial_number")

		if nNumber != "" {
			if sameSerialNumber == "true" {
				r, _ := l.GetOtherDeregBySerialNumber(nNumber)
				return c.JSONBlob(http.StatusOK, l.AugmentToBytes(r))
			}
			r, _ := l.GetDeregByNNumber(nNumber)
			return c.JSONBlob(http.StatusOK, l.AugmentToBytes(r))
		}
		return c.JSON(http.StatusBadRequest, "request not valid")
	})
}
