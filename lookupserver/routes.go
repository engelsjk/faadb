package lookupserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Query struct {
	NNumber           string
	SerialNumber      string
	ModeSCodeHex      string
	RegistrantName    string
	RegistrantStreet1 string
}

type Filter struct {
	AircraftModelCode               string
	RegistrantState                 string
	AirworthinessClassificationCode string
	ApprovedOperationCode           string
}

func addRoutes(e *echo.Echo, l *LookupService) {

	e.GET("/aircraft/n/:n", func(c echo.Context) error {
		nNumber := c.Param("n")
		serialNumber := c.QueryParam("serial_number")
		registrantName := c.QueryParam("registrant_name")
		registrantStreet := c.QueryParam("registrant_street")

		if nNumber == "" {
			return c.JSON(http.StatusBadRequest, "nnumber required")
		}
		if serialNumber == "same" {
			r, err := l.GetOtherAircraft(&Query{NNumber: nNumber, SerialNumber: serialNumber}, nil)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return c.JSONBlob(http.StatusOK, ToBytes(r))
		}
		if registrantName == "same" {

			registrantState := c.QueryParam("registrant_state")
			aircraftModelCode := c.QueryParam("aircraft_model_code")
			airworthinessClassificationCode := c.QueryParam("airworthiness_classification_code")
			approvedOperationCode := c.QueryParam("approved_operation_code")

			r, err := l.GetOtherAircraft(&Query{
				NNumber:        nNumber,
				RegistrantName: registrantName,
			}, &Filter{
				RegistrantState:                 registrantState,
				AircraftModelCode:               aircraftModelCode,
				AirworthinessClassificationCode: airworthinessClassificationCode,
				ApprovedOperationCode:           approvedOperationCode,
			})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return c.JSONBlob(http.StatusOK, ToBytes(r))
		}
		if registrantStreet == "same" {
			r, err := l.GetOtherAircraft(&Query{NNumber: nNumber, RegistrantStreet1: registrantStreet}, nil)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return c.JSONBlob(http.StatusOK, ToBytes(r))
		}
		r, err := l.GetAircraft(&Query{NNumber: nNumber}, nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONBlob(http.StatusOK, ToBytes(r))
	})

	e.GET("/aircraft/sn/:sn", func(c echo.Context) error {
		serialNumber := c.Param("sn")
		if serialNumber == "" {
			return c.JSON(http.StatusBadRequest, "serial number required")
		}
		r, err := l.GetAircraft(&Query{SerialNumber: serialNumber}, nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSONBlob(http.StatusOK, ToBytes(r))
	})

	e.GET("/aircraft/hex/:hex", func(c echo.Context) error {
		modeSCodeHex := c.Param("hex")
		if modeSCodeHex == "" {
			return c.JSON(http.StatusBadRequest, "mode s code hex required")
		}
		r, err := l.GetAircraft(&Query{ModeSCodeHex: modeSCodeHex}, nil)
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
