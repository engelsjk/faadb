package master

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/utils"
	"github.com/engelsjk/faadb/service"
)

type Master struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewMaster(dataPath, dbPath string) (*Master, error) {
	name := "master"
	numFields := 35

	m := &Master{Name: name}

	m.codes = initCodes()

	svc, err := service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
	}, m.DecodeLine)
	if err != nil {
		return nil, err
	}

	err = svc.CreateIndexJSON("nnumber", "*", "nnumber")
	if err != nil {
		return nil, err
	}

	err = svc.CreateIndexJSON("registrant_name", "*", "registrant.name")
	if err != nil {
		return nil, err
	}

	err = svc.CreateIndexJSON("registrant_street_1", "*", "registrant.street_1")
	if err != nil {
		return nil, err
	}

	m.svc = svc

	return m, nil
}

func (m *Master) DecodeLine(line []string) (string, string, error) {
	record := Record{
		NNumber:      utils.ToUpper(line[0]),
		SerialNumber: utils.ToUpper(line[1]),
		Manufacturer: Manufacturer{
			AircraftModelCode: utils.ToUpper(line[2]),
			EngineModelCode:   utils.ToUpper(line[3]),
			Year:              utils.ToUpper(line[4]),
		},
		Registrant: Registrant{
			Type: codes.Description{
				Code:        utils.ToUpper(line[5]),
				Description: codes.DecodeDescription(line[5], m.codes.RegistrantType),
			},
			Name:    utils.ToUpper(line[6]),
			Street1: utils.ToUpper(line[7]),
			Street2: utils.ToUpper(line[8]),
			City:    utils.ToUpper(line[9]),
			State:   utils.ToUpper(line[10]),
			ZipCode: utils.ToUpper(line[11]),
			Region: codes.Description{
				Code:        utils.ToUpper(line[12]),
				Description: codes.DecodeDescription(line[12], m.codes.RegistrantRegion),
			},
			County:  utils.ToUpper(line[13]),
			Country: utils.ToUpper(line[14]),
		},
		LastActivityDate:     utils.ToUpper(line[15]),
		CertificateIssueDate: utils.ToUpper(line[16]),
		Certification:        decodeCertification(line[17], m.codes.Certification),
		AircraftType: codes.Description{
			Code:        utils.ToUpper(line[18]),
			Description: codes.DecodeDescription(line[18], m.codes.AircraftType),
		},
		EngineType: codes.Description{
			Code:        utils.ToUpper(line[19]),
			Description: codes.DecodeDescription(line[19], m.codes.EngineType),
		},
		StatusCode: codes.Description{
			Code:        utils.ToUpper(line[20]),
			Description: codes.DecodeDescription(line[20], m.codes.StatusCode),
		},
		ModeS: ModeS{
			Code:    utils.ToUpper(line[21]),
			CodeHex: utils.ToUpper(line[33]),
		},
		AirworthinessDate: utils.ToUpper(line[23]),
		Ownership: Ownership{
			Fractional: codes.Status{
				Code:   utils.ToUpper(line[22]),
				Status: codes.DecodeStatus(line[22], m.codes.FractionalOwnership),
			},
			OtherName1: utils.ToUpper(line[24]),
			OtherName2: utils.ToUpper(line[25]),
			OtherName3: utils.ToUpper(line[26]),
			OtherName4: utils.ToUpper(line[27]),
			OtherName5: utils.ToUpper(line[28]),
		},
		ExpirationDate: utils.ToUpper(line[29]),
		UniqueID:       utils.ToUpper(line[30]),
		Kit: Kit{
			ManufacturerName: utils.ToUpper(line[31]),
			ModelName:        utils.ToUpper(line[32]),
		},
	}

	key := record.NNumber

	b, err := json.Marshal(record)
	return key, string(b), err
}
