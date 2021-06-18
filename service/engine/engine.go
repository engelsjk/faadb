package engine

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/utils"
	"github.com/engelsjk/faadb/service"
)

type Engine struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewEngine(dataPath, dbPath string) (*Engine, error) {
	name := "engine"
	numFields := 7

	e := &Engine{Name: name}

	e.codes = initCodes()

	svc, err := service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
	}, e.DecodeLine)
	if err != nil {
		return nil, err
	}

	e.svc = svc

	return e, nil
}

func (e *Engine) DecodeLine(line []string) (string, string, error) {

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
