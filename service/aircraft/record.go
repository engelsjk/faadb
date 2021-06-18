package aircraft

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
)

type Record struct {
	ManufacturerModelSeriesCode string            `json:"manufacturer_model_series_code"`
	ManufacturerName            string            `json:"manufacturer_name"`
	ModelName                   string            `json:"model_name"`
	AircraftType                codes.Description `json:"aircraft_type"`
	EngineType                  codes.Description `json:"engine_type"`
	AircraftCategoryCode        codes.Description `json:"aircraft_category_code"`
	BuilderCertificationCode    codes.Description `json:"builder_certification_code"`
	NumberOfEngines             int32             `json:"number_of_engines"`
	NumberOfSeats               int32             `json:"number_of_seats"`
	AircraftWeight              codes.Description `json:"aircraft_weight"`
	AircraftCruisingSpeed       int32             `json:"aircraft_cruising_speed"`
	TCDataSheet                 string            `json:"tc_data_sheet"`
	TCDataHolder                string            `json:"tc_data_holder"`
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
