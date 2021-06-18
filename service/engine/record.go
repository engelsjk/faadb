package engine

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
)

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
