package activeserver

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/service"
	"github.com/engelsjk/faadb/internal/utils"
)

type ActiveService struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewActiveService(dataPath, dbPath string, reload bool) (*ActiveService, error) {
	name := "active"
	numFields := 35

	m := &ActiveService{Name: name}

	m.codes = initCodes()

	var err error
	m.svc, err = service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
		Reload:    reload,
	}, m.DecodeLine)
	if err != nil {
		return nil, err
	}

	if err := m.svc.CreateIndexJSON("nnumber", "*", "nnumber"); err != nil {
		return nil, err
	}
	if err := m.svc.CreateIndexJSON("serial_number", "*", "serial_number"); err != nil {
		return nil, err
	}
	if err := m.svc.CreateIndexJSON("mode_s_code_hex", "*", "mode_s.code_hex"); err != nil {
		return nil, err
	}
	if err := m.svc.CreateIndexJSON("registrant_name", "*", "registrant.name"); err != nil {
		return nil, err
	}
	if err := m.svc.CreateIndexJSON("registrant_street_1", "*", "registrant.street_1"); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *ActiveService) DecodeLine(line []string) (string, string, error) {
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

type Record struct {
	NNumber              string            `json:"nnumber"`
	SerialNumber         string            `json:"serial_number"`
	Manufacturer         Manufacturer      `json:"manufacturer"`
	Registrant           Registrant        `json:"registrant"`
	LastActivityDate     string            `json:"last_activity_date"`
	CertificateIssueDate string            `json:"certification_issue_date"`
	Certification        Certification     `json:"certification"`
	AircraftType         codes.Description `json:"aircraft_type"`
	EngineType           codes.Description `json:"engine_type"`
	StatusCode           codes.Description `json:"status_code"`
	ModeS                ModeS             `json:"mode_s"`
	AirworthinessDate    string            `json:"airworthiness_date"`
	Ownership            Ownership         `json:"ownership"`
	ExpirationDate       string            `json:"expiration_date"`
	UniqueID             string            `json:"unique_id"`
	Kit                  Kit               `json:"kit"`
}

type Certification struct {
	AirworthinessClassification codes.Description `json:"airworthiness_classification"`
	ApprovedOperation           codes.Description `json:"approved_operation"`
}

type Manufacturer struct {
	AircraftModelCode string `json:"aircraft_model_code"`
	EngineModelCode   string `json:"engine_model_code"`
	Year              string `json:"year"`
}

type Registrant struct {
	Type    codes.Description `json:"type"`
	Name    string            `json:"name"`
	Street1 string            `json:"street_1"`
	Street2 string            `json:"street_2"`
	City    string            `json:"city"`
	State   string            `json:"state"`
	ZipCode string            `json:"zipcode"`
	Region  codes.Description `json:"region"`
	County  string            `json:"county"`
	Country string            `json:"country"`
}

type ModeS struct {
	Code    string `json:"code"`
	CodeHex string `json:"code_hex"`
}

type Ownership struct {
	Fractional codes.Status `json:"fractional"`
	OtherName1 string       `json:"other_name_1"`
	OtherName2 string       `json:"other_name_2"`
	OtherName3 string       `json:"other_name_3"`
	OtherName4 string       `json:"other_name_4"`
	OtherName5 string       `json:"other_name_5"`
}

type Kit struct {
	ManufacturerName string `json:"manufacturer_name"`
	ModelName        string `json:"model_name"`
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
