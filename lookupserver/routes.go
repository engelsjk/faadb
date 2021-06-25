package lookupserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo, l *LookupService) {

	e.GET("/aircraft/n/:n", func(c echo.Context) error {
		nNumber := c.Param("n")
		sameSerialNumber := c.QueryParam("same_serial_number")
		// todo: by same Registrant Name/Street/State

		if nNumber == "" {
			return c.JSON(http.StatusBadRequest, "nnumber required")
		}
		if sameSerialNumber == "true" {
			r, err := l.GetOtherAircraftWithSameSerialNumber(nNumber)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return c.JSONBlob(http.StatusOK, ToBytes(r))
		}

		r, err := l.GetAircraftByNNumber(nNumber)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONBlob(http.StatusOK, ToBytes(r))
	})

	e.GET("/aircraft/sn/:sn", func(c echo.Context) error {
		serialNumber := c.Param("sn")
		// todo: by same Registrant Name/Street/State

		if serialNumber == "" {
			return c.JSON(http.StatusBadRequest, "serial number required")
		}
		r, err := l.GetAircraftBySerialNumber(serialNumber)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONBlob(http.StatusOK, ToBytes(r))
	})

	///////////////////////////////////////

	e.GET("/registered/n/:n", func(c echo.Context) error {
		nNumber := c.Param("n")
		if nNumber == "" {
			return c.JSON(http.StatusBadRequest, "nnumber required")
		}
		r, err := l.GetMasterByNNumber(nNumber)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONBlob(http.StatusOK, l.AugmentToBytes(r))
	})

	e.GET("/reserved/n/:n", func(c echo.Context) error {
		nNumber := c.Param("n")
		if nNumber == "" {
			return c.JSON(http.StatusBadRequest, "nnumber required")
		}
		r, err := l.GetReservedByNNumber(nNumber)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONBlob(http.StatusOK, l.AugmentToBytes(r))
	})

	e.GET("/deregistered/n/:n", func(c echo.Context) error {
		nNumber := c.Param("n")
		if nNumber == "" {
			return c.JSON(http.StatusBadRequest, "nnumber required")
		}
		r, err := l.GetDeregByNNumber(nNumber)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONBlob(http.StatusOK, l.AugmentToBytes(r))
	})
}
