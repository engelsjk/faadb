package deregserver

import (
	"context"

	"github.com/engelsjk/faadb/rpc/dereg"
)

type Server struct {
	dereg *DeregService
}

func NewServer(dereg *DeregService) *Server {
	return &Server{dereg: dereg}
}

func (s *Server) GetAircraft(ctx context.Context, query *dereg.Query) (*dereg.Aircraft, error) {
	var (
		bs  [][]byte
		err error
	)

	filters := map[string]string{
		"registrant.state":    query.RegistrantState,
		"aircraft_model_code": query.AircraftModelCode,
		"certification.airworthiness_classification.code": query.AirworthinessClassificationCode,
		"certification.approved_operation.code":           query.ApprovedOperationCode,
	}

	if query.NNumber != "" {
		nnumber := query.NNumber
		exact := true
		r := []rune(query.NNumber)
		if string(r[0]) == "*" {
			nnumber = string(r[1:])
			exact = false
		}
		bs, err = s.dereg.svc.List("nnumber", nnumber, "nnumber", exact, filters)
	}
	if query.SerialNumber != "" {
		bs, err = s.dereg.svc.List("serial_number", query.SerialNumber, "serial_number", true, filters)
	}
	if query.ModeSCodeHex != "" {
		bs, err = s.dereg.svc.List("mode_s_code_hex", query.ModeSCodeHex, "mode_s.code_hex", true, filters)
	}
	if query.RegistrantName != "" {
		bs, err = s.dereg.svc.List("registrant_name", query.RegistrantName, "registrant.name", true, filters)
	}
	if query.RegistrantStreet1 != "" {
		bs, err = s.dereg.svc.List("registrant_street_1", query.RegistrantStreet1, "registrant.street_1", true, filters)
	}
	if err != nil {
		return nil, err
	}
	return bytesToAircraft(bs)
}

func bytesToA(b []byte) (*dereg.A, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &dereg.A{
		NNumber:                       record.NNumber,
		SerialNumber:                  record.SerialNumber,
		ManufacturerAircraftModelCode: record.Manufacturer.AircraftModelCode,
		ManufacturerEngineModelCode:   record.Manufacturer.EngineModelCode,
		ManufacturerYear:              record.Manufacturer.Year,
		Status:                        record.StatusCode.Description,
		RegistrantType:                record.Registrant.Type.Description,
		RegistrantName:                record.Registrant.Name,
		RegistrantStreet1:             record.Registrant.Street1,
		RegistrantStreet2:             record.Registrant.Street2,
		RegistrantCity:                record.Registrant.City,
		RegistrantState:               record.Registrant.State,
		RegistrantZipCode:             record.Registrant.ZipCode,
		RegistrantRegion:              record.Registrant.Region.Description,
		RegistrantCounty:              record.Registrant.County,
		RegistrantCountry:             record.Registrant.Country,
		RegistrantPhysicalAddress:     record.Registrant.PhysicalAddress,
		RegistrantPhysicalAddress2:    record.Registrant.PhysicalAddress2,
		RegistrantPhysicalCity:        record.Registrant.PhysicalCity,
		RegistrantPhysicalState:       record.Registrant.PhysicalState,
		RegistrantPhysicalZipCode:     record.Registrant.PhysicalZipCode,
		RegistrantPhysicalCounty:      record.Registrant.PhysicalCounty,
		RegistrantPhysicalCountry:     record.Registrant.PhysicalCountry,
		CertificationAirworthinessClassificationDescription: record.Certification.AirworthinessClassification.Description,
		CertificationAirworthinessClassificationCode:        record.Certification.AirworthinessClassification.Code,
		CertificationApprovedOperationDescription:           record.Certification.ApprovedOperation.Description,
		CertificationApprovedOperationCode:                  record.Certification.ApprovedOperation.Code,
		AirworthinessDate:                                   record.AirworthinessDate,
		CancelDate:                                          record.CancelDate,
		ExportCountry:                                       record.ExportCountry,
		LastActivityDate:                                    record.LastActivityDate,
		CertificationIssueDate:                              record.CertificateIssueDate,
		OwnershipOtherName1:                                 record.Ownership.OtherName1,
		OwnershipOtherName2:                                 record.Ownership.OtherName2,
		OwnershipOtherName3:                                 record.Ownership.OtherName3,
		OwnershipOtherName4:                                 record.Ownership.OtherName4,
		OwnershipOtherName5:                                 record.Ownership.OtherName5,
		KitManufacturerName:                                 record.Kit.ManufacturerName,
		KitModelName:                                        record.Kit.ModelName,
		ModeSCode:                                           record.ModeS.Code,
		ModeSCodeHex:                                        record.ModeS.CodeHex,
	}, nil
}

func bytesToAircraft(bs [][]byte) (*dereg.Aircraft, error) {
	as := make([]*dereg.A, len(bs))
	for i, b := range bs {
		a, err := bytesToA(b)
		if err != nil {
			return nil, err
		}
		as[i] = a
	}
	return &dereg.Aircraft{A: as}, nil
}
