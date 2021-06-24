package deregserver

import (
	"encoding/json"
	"fmt"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/service"
	"github.com/engelsjk/faadb/internal/utils"
)

type DeregService struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewDeregService(dataPath, dbPath string, reload bool) (*DeregService, error) {
	name := "dereg"
	numFields := 39

	d := &DeregService{Name: name}

	d.codes = initCodes()

	var err error
	d.svc, err = service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
		Reload:    reload,
	}, d.DecodeLine)
	if err != nil {
		return nil, err
	}

	if err := d.svc.CreateIndexJSON("nnumber", "*", "nnumber"); err != nil {
		return nil, err
	}
	if err := d.svc.CreateIndexJSON("serial_number", "*", "serial_number"); err != nil {
		return nil, err
	}
	if err := d.svc.CreateIndexJSON("registrant_name", "*", "registrant.name"); err != nil {
		return nil, err
	}
	if err := d.svc.CreateIndexJSON("registrant_street_1", "*", "registrant.street_1"); err != nil {
		return nil, err
	}
	if err := d.svc.CreateIndexJSON("registrant_state", "*", "registrant.state"); err != nil {
		return nil, err
	}

	// index: registrant.state?

	return d, nil
}

func (d *DeregService) DecodeLine(line []string) (string, string, error) {

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

	// key: nnumber_serialnumber
	key := fmt.Sprintf("%s_%s", record.NNumber, record.SerialNumber)

	b, err := json.Marshal(record)
	return key, string(b), err
}

type Record struct {
	NNumber              string            `json:"nnumber"`
	SerialNumber         string            `json:"serial_number"`
	Manufacturer         Manufacturer      `json:"manufacturer"`
	StatusCode           codes.Description `json:"status_code"`
	Registrant           Registrant        `json:"registrant"`
	Certification        Certification     `json:"certification"`
	AirworthinessDate    string            `json:"airworthiness_date"`
	CancelDate           string            `json:"cancel_date"`
	ExportCountry        string            `json:"export_country"`
	LastActivityDate     string            `json:"last_activity_date"`
	CertificateIssueDate string            `json:"certification_issue_date"`
	Ownership            Ownership         `json:"ownership"`
	Kit                  Kit               `json:"kit"`
	ModeS                ModeS             `json:"mode_s"`
}

type Certification struct {
	AirworthinessClassification codes.Description
	ApprovedOperation           codes.Description
}

type Manufacturer struct {
	AircraftModelCode string `json:"aircraft_model_code"`
	EngineModelCode   string `json:"engine_model_code"`
	Year              string `json:"year"`
}

type Registrant struct {
	Type             codes.Description `json:"type"`
	Name             string            `json:"name"`
	Street1          string            `json:"street_1"`
	Street2          string            `json:"street_2"`
	City             string            `json:"city"`
	State            string            `json:"state"`
	ZipCode          string            `json:"zipcode"`
	Region           codes.Description `json:"region"`
	County           string            `json:"county"`
	Country          string            `json:"country"`
	PhysicalAddress  string            `json:"physical_address"`
	PhysicalAddress2 string            `json:"physical_address_2"`
	PhysicalCity     string            `json:"physical_city"`
	PhysicalState    string            `json:"physical_state"`
	PhysicalZipCode  string            `json:"physical_zipcode"`
	PhysicalCounty   string            `json:"physical_county"`
	PhysicalCountry  string            `json:"physical_country"`
}

type Ownership struct {
	OtherName1 string `json:"other_name_1"`
	OtherName2 string `json:"other_name_2"`
	OtherName3 string `json:"other_name_3"`
	OtherName4 string `json:"other_name_4"`
	OtherName5 string `json:"other_name_5"`
}

type Kit struct {
	ManufacturerName string `json:"manufacturer_name"`
	ModelName        string `json:"model_name"`
}

type ModeS struct {
	Code    string `json:"code"`
	CodeHex string `json:"code_hex"`
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
