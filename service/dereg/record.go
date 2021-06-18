package dereg

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
)

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
