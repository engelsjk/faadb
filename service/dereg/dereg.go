package dereg

import (
	"encoding/json"
	"fmt"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/utils"
	"github.com/engelsjk/faadb/service"
)

type Dereg struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewDereg(dataPath, dbPath string) (*Dereg, error) {
	name := "dereg"
	numFields := 39

	d := &Dereg{Name: name}

	d.codes = initCodes()

	svc, err := service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
	}, d.DecodeLine)
	if err != nil {
		return nil, err
	}

	err = svc.CreateIndexJSON("nnumber_serial", "*", "nnumber", "serial_number")
	if err != nil {
		return nil, err
	}

	err = svc.CreateIndexJSON("registrant_name", "*", "registrant.name")
	if err != nil {
		return nil, err
	}

	d.svc = svc

	return d, nil
}

func (d *Dereg) DecodeLine(line []string) (string, string, error) {

	record := Record{
		NNumber:      utils.ToUpper(line[0]),
		SerialNumber: utils.ToUpper(line[1]),
		Manufacturer: Manufacturer{
			AircraftModelCode: utils.ToUpper(line[2]),
			EngineModelCode:   utils.ToUpper(line[10]),
			Year:              utils.ToUpper(line[11]),
		},
		StatusCode: codes.Description{
			Code:        utils.ToUpper(line[3]),
			Description: codes.DecodeDescription(line[3], d.codes.StatusCode),
		},
		Registrant: Registrant{
			Type: codes.Description{
				Code:        utils.ToUpper(line[19]),
				Description: codes.DecodeDescription(line[19], d.codes.RegistrantType),
			},
			Name:    utils.ToUpper(line[4]),
			Street1: utils.ToUpper(line[5]),
			Street2: utils.ToUpper(line[6]),
			City:    utils.ToUpper(line[7]),
			State:   utils.ToUpper(line[8]),
			ZipCode: utils.ToUpper(line[9]),
			Region: codes.Description{
				Code:        utils.ToUpper(line[13]),
				Description: codes.DecodeDescription(line[13], d.codes.RegistrantRegion),
			},
			County:           utils.ToUpper(line[14]),
			Country:          utils.ToUpper(line[15]),
			PhysicalAddress:  utils.ToUpper(line[23]),
			PhysicalAddress2: utils.ToUpper(line[24]),
			PhysicalCity:     utils.ToUpper(line[25]),
			PhysicalState:    utils.ToUpper(line[26]),
			PhysicalZipCode:  utils.ToUpper(line[27]),
			PhysicalCounty:   utils.ToUpper(line[28]),
			PhysicalCountry:  utils.ToUpper(line[29]),
		},
		Certification:        decodeCertification(line[12], d.codes.Certification),
		AirworthinessDate:    utils.ToUpper(line[16]),
		CancelDate:           utils.ToUpper(line[17]),
		ExportCountry:        utils.ToUpper(line[20]),
		LastActivityDate:     utils.ToUpper(line[21]),
		CertificateIssueDate: utils.ToUpper(line[22]),
		Ownership: Ownership{
			OtherName1: utils.ToUpper(line[30]),
			OtherName2: utils.ToUpper(line[31]),
			OtherName3: utils.ToUpper(line[32]),
			OtherName4: utils.ToUpper(line[33]),
			OtherName5: utils.ToUpper(line[34]),
		},
		Kit: Kit{
			ManufacturerName: utils.ToUpper(line[35]),
			ModelName:        utils.ToUpper(line[36]),
		},
		ModeS: ModeS{
			Code:    utils.ToUpper(line[18]),
			CodeHex: utils.ToUpper(line[37]),
		},
	}

	key := fmt.Sprintf("%s_%s", record.NNumber, record.SerialNumber)

	b, err := json.Marshal(record)
	return key, string(b), err
}
