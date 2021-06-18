package aircraft

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/service"
	"github.com/engelsjk/faadb/internal/utils"
)

type AircraftService struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewAircraftService(dataPath, dbPath string) (*AircraftService, error) {
	name := "aircraft"
	numFields := 14

	a := &AircraftService{Name: name}

	a.codes = initCodes()

	var err error
	a.svc, err = service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
	}, a.DecodeLine)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *AircraftService) DecodeLine(line []string) (string, string, error) {

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

type Record struct {
	ManufacturerModelSeriesCode string            `json:"manufacturer_model_series_code"`
	ManufacturerName            string            `json:"manufacturer_name"`
	ModelName                   string            `json:"model_name"`
	AircraftType                codes.Description `json:"aircraft_type"`
	EngineType                  codes.Description `json:"engine_type"`
	AircraftCategoryCode        codes.Description `json:"aircraft_category_code"`
	BuilderCertificationCode    codes.Description `json:"builder_certification_code"`
	NumberOfEngines             int32             `json:"number_of_engines"`
	NumberOfSeats               int32             `json:"number_of_seats"`
	AircraftWeight              codes.Description `json:"aircraft_weight"`
	AircraftCruisingSpeed       int32             `json:"aircraft_cruising_speed"`
	TCDataSheet                 string            `json:"tc_data_sheet"`
	TCDataHolder                string            `json:"tc_data_holder"`
}

func (r *Record) MarshalJSON() ([]byte, error) {
	type Alias Record
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	return json.Marshal(a)
}

func (r *Record) UnmarshalJSON(b []byte) error {
	type Alias Record
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	return json.Unmarshal(b, &a)
}
