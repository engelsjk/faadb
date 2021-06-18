package reserved

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
)

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
