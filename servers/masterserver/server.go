package masterserver

import (
	"context"

	"github.com/engelsjk/faadb/rpc/master"
)

type Server struct {
	master *MasterService
}

func NewServer(master *MasterService) *Server {
	return &Server{master: master}
}

func (s *Server) GetAircraft(ctx context.Context, query *master.Query) (*master.Aircraft, error) {
	var (
		bs  [][]byte
		err error
	)
	if query.NNumber != "" {
		nnumber := query.NNumber
		exact := true
		r := []rune(query.NNumber)
		if string(r[0]) == "*" {
			nnumber = string(r[1:])
			exact = false
		}
		bs, err = s.master.svc.List("nnumber", nnumber, "nnumber", exact)
	}
	if query.SerialNumber != "" {
		bs, err = s.master.svc.List("serial_number", query.SerialNumber, "serial_number", true)
	}
	if query.RegistrantName != "" {
		bs, err = s.master.svc.List("registrant_name", query.RegistrantName, "registrant.name", true)
	}
	if query.RegistrantStreet1 != "" {
		bs, err = s.master.svc.List("registrant_street_1", query.RegistrantName, "registrant.street_1", true)
	}
	if query.RegistrantState != "" {
		bs, err = s.master.svc.List("registrant_state", query.RegistrantState, "registrant.state", true)
	}
	if err != nil {
		return nil, err
	}
	return bytesToAircraft(bs)
}

func bytesToA(b []byte) (*master.A, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &master.A{
		NNumber:                                  record.NNumber,
		SerialNumber:                             record.SerialNumber,
		ManufacturerAircraftModelCode:            record.Manufacturer.AircraftModelCode,
		ManufacturerEngineModelCode:              record.Manufacturer.EngineModelCode,
		ManufacturerYear:                         record.Manufacturer.Year,
		RegistrantType:                           record.Registrant.Type.Description,
		RegistrantName:                           record.Registrant.Name,
		RegistrantStreet1:                        record.Registrant.Street1,
		RegistrantStreet2:                        record.Registrant.Street2,
		RegistrantCity:                           record.Registrant.City,
		RegistrantState:                          record.Registrant.State,
		RegistrantZipCode:                        record.Registrant.ZipCode,
		RegistrantRegion:                         record.Registrant.Region.Description,
		RegistrantCounty:                         record.Registrant.County,
		RegistrantCountry:                        record.Registrant.Country,
		LastActivityDate:                         record.LastActivityDate,
		CertificationIssueDate:                   record.CertificateIssueDate,
		CertificationAirworthinessClassification: record.Certification.AirworthinessClassification.Description,
		CertificationApprovedOperations:          record.Certification.ApprovedOperation.Description,
		AircraftType:                             record.AircraftType.Description,
		EngineType:                               record.EngineType.Description,
		Status:                                   record.StatusCode.Description,
		ModeSCode:                                record.ModeS.Code,
		ModeSCodeHex:                             record.ModeS.CodeHex,
		AirworthinessDate:                        record.AirworthinessDate,
		OwnershipFractional:                      record.Ownership.Fractional.Status,
		OwnershipOtherName1:                      record.Ownership.OtherName1,
		OwnershipOtherName2:                      record.Ownership.OtherName2,
		OwnershipOtherName3:                      record.Ownership.OtherName3,
		OwnershipOtherName4:                      record.Ownership.OtherName4,
		OwnershipOtherName5:                      record.Ownership.OtherName5,
		ExpirationDate:                           record.ExpirationDate,
		UniqueID:                                 record.UniqueID,
		KitManufacturerName:                      record.Kit.ManufacturerName,
		KitModelName:                             record.Kit.ModelName,
	}, nil
}

func bytesToAircraft(bs [][]byte) (*master.Aircraft, error) {
	as := make([]*master.A, len(bs))
	for i, b := range bs {
		a, err := bytesToA(b)
		if err != nil {
			return nil, err
		}
		as[i] = a
	}
	return &master.Aircraft{A: as}, nil
}
