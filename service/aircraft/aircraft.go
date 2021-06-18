package aircraft

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/utils"
	"github.com/engelsjk/faadb/service"
)

type Aircraft struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewAircraft(dataPath, dbPath string) (*Aircraft, error) {
	name := "aircraft"
	numFields := 14

	a := &Aircraft{Name: name}

	a.codes = initCodes()

	svc, err := service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
	}, a.DecodeLine)
	if err != nil {
		return nil, err
	}

	a.svc = svc

	return a, nil
}

func (a *Aircraft) DecodeLine(line []string) (string, string, error) {

	record := Record{
		ManufacturerModelSeriesCode: utils.ToUpper(line[0]),
		ManufacturerName:            utils.ToUpper(line[1]),
		ModelName:                   utils.ToUpper(line[2]),
		AircraftType: codes.Description{
			Code:        utils.ToUpper(line[3]),
			Description: codes.DecodeDescription(line[3], a.codes.AircraftType),
		},
		EngineType: codes.Description{
			Code:        utils.ToUpper(line[4]),
			Description: codes.DecodeDescription(line[4], a.codes.EngineType),
		},
		AircraftCategoryCode: codes.Description{
			Code:        utils.ToUpper(line[5]),
			Description: codes.DecodeDescription(line[5], a.codes.AircraftCategoryCode),
		},
		BuilderCertificationCode: codes.Description{
			Code:        utils.ToUpper(line[6]),
			Description: codes.DecodeDescription(line[6], a.codes.BuilderCertificationCode),
		},
		NumberOfEngines: codes.ParseInt32(line[7]),
		NumberOfSeats:   codes.ParseInt32(line[8]),
		AircraftWeight: codes.Description{
			Code:        utils.ToUpper(line[9]),
			Description: codes.DecodeDescription(line[9], a.codes.AircraftWeight),
		},
		AircraftCruisingSpeed: codes.ParseInt32(line[10]),
		TCDataSheet:           utils.ToUpper(line[11]),
		TCDataHolder:          utils.ToUpper(line[12]),
	}

	key := record.ManufacturerModelSeriesCode

	b, err := json.Marshal(record)
	return key, string(b), err
}
