package reserved

import (
	"encoding/json"

	"github.com/engelsjk/faadb/internal/codes"
	"github.com/engelsjk/faadb/internal/utils"
	"github.com/engelsjk/faadb/service"
)

type Reserved struct {
	Name  string
	svc   *service.Service
	codes Codes
}

func NewReserved(dataPath, dbPath string) (*Reserved, error) {
	name := "reserved"
	numFields := 13

	r := &Reserved{Name: name}

	r.codes = initCodes()

	svc, err := service.NewService(service.Settings{
		Name:      name,
		NumFields: numFields,
		DataPath:  dataPath,
		DBPath:    dbPath,
	}, r.DecodeLine)
	if err != nil {
		return nil, err
	}
	err = svc.CreateIndexJSON("registrant_name", "*", "registrant.name")
	if err != nil {
		return nil, err
	}

	r.svc = svc

	return r, nil
}

func (r *Reserved) DecodeLine(line []string) (string, string, error) {

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
			Description: codes.DecodeDescription(line[9], r.codes.ReservationType),
		},
		ExpirationNoticeDate: utils.ToUpper(line[10]),
		NNumberForChange:     utils.ToUpper(line[11]),
		PurgeDate:            utils.ToUpper(line[12]),
	} // line[34] is empty

	key := record.NNumber

	b, err := json.Marshal(record)
	return key, string(b), err
}
