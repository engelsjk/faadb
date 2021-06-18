package master

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
)

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
	AirworthinessClassification codes.Description
	ApprovedOperation           codes.Description
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
