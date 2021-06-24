package reservedserver

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/service"
	"github.com/engelsjk/faadb/internal/utils"
)

type ReservedService struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewReserved(dataPath, dbPath string, reload bool) (*ReservedService, error) {
	name := "reserved"
	numFields := 13

	r := &ReservedService{Name: name}

	r.codes = initCodes()

	var err error
	r.svc, err = service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
		Reload:    reload,
	}, r.DecodeLine)
	if err != nil {
		return nil, err
	}

	if err := r.svc.CreateIndexJSON("nnumber", "*", "nnumber"); err != nil {
		return nil, err
	}
	if err := r.svc.CreateIndexJSON("serial_number", "*", "serial_number"); err != nil {
		return nil, err
	}
	if err := r.svc.CreateIndexJSON("registrant_name", "*", "registrant.name"); err != nil {
		return nil, err
	}
	if err := r.svc.CreateIndexJSON("registrant_street_1", "*", "registrant.street_1"); err != nil {
		return nil, err
	}
	if err := r.svc.CreateIndexJSON("registrant_state", "*", "registrant.state"); err != nil {
		return nil, err
	}

	// index: registrant.state?

	return r, nil
}

func (r *ReservedService) DecodeLine(line []string) (string, string, error) {

	record := Record{
		NNumber: utils.ToUpper(line[0]),
		Registrant: Registrant{
			Name:    utils.ToUpper(line[1]),
			Street1: utils.ToUpper(line[2]),
			Street2: utils.ToUpper(line[3]),
			City:    utils.ToUpper(line[4]),
			State:   utils.ToUpper(line[5]),
			ZipCode: utils.ToUpper(line[6]),
		},
		ReserveDate: utils.ToUpper(line[7]),
		ReservationType: codes.Description{
			Code:        utils.ToUpper(line[8]),
			Description: codes.DecodeDescription(line[8], r.codes.ReservationType),
		},
		ExpirationNoticeDate: utils.ToUpper(line[9]),
		NNumberForChange:     utils.ToUpper(line[10]),
		PurgeDate:            utils.ToUpper(line[11]),
	} // line[34] is empty

	key := record.NNumber

	b, err := json.Marshal(record)
	return key, string(b), err
}

type Record struct {
	NNumber              string            `json:"nnumber"`
	Registrant           Registrant        `json:"registrant"`
	ReserveDate          string            `json:"reserve_date"`
	ReservationType      codes.Description `json:"reservation_type"`
	ExpirationNoticeDate string            `json:"expiration_notice_date"`
	NNumberForChange     string            `json:"nnumber_for_change"`
	PurgeDate            string            `json:"purge_date"`
}

type Registrant struct {
	Name    string `json:"name"`
	Street1 string `json:"street_1"`
	Street2 string `json:"street_2"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipcode"`
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
