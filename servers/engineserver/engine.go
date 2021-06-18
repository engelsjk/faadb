package engineserver

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/service"
	"github.com/engelsjk/faadb/internal/utils"
)

type EngineService struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewEngineService(dataPath, dbPath string) (*EngineService, error) {
	name := "engine"
	numFields := 7

	e := &EngineService{Name: name}

	e.codes = initCodes()

	var err error
	e.svc, err = service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
	}, e.DecodeLine)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *EngineService) DecodeLine(line []string) (string, string, error) {

	record := Record{
		ManufacturerModelCode: utils.ToUpper(line[0]),
		ManufacturerName:      utils.ToUpper(line[1]),
		ModelName:             utils.ToUpper(line[2]),
		EngineType: codes.Description{
			Code:        utils.ToUpper(line[3]),
			Description: codes.DecodeDescription(line[3], e.codes.EngineType),
		},
		Horsepower:     codes.ParseInt32(line[4]),
		PoundsOfThrust: codes.ParseInt32(line[5]),
	}

	key := record.ManufacturerModelCode

	b, err := json.Marshal(record)
	return key, string(b), err
}

type Record struct {
	ManufacturerModelCode string            `json:"manufacturer_model_code"`
	ManufacturerName      string            `json:"manufacturer_name"`
	ModelName             string            `json:"engine_model_name"`
	EngineType            codes.Description `json:"engine_type"`
	Horsepower            int32             `json:"horsepower"`
	PoundsOfThrust        int32             `json:"pounds_of_thrust"`
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
