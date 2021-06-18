package dereg

import (
	"context"

	pb "github.com/engelsjk/faadb/rpc/dereg"
)

type Server struct {
	dereg *Dereg
}

func NewServer(dereg *Dereg) *Server {
	return &Server{dereg: dereg}
}

func (s *Server) GetMultipleAircraft(ctx context.Context, query *pb.Query) (*pb.MultipleAircraft, error) {
	bs, err := s.dereg.svc.StartsWith("nnumber_serial", query.NNumber)
	if err != nil {
		return nil, err
	}
	return bytesToMultipleAircraft(bs)
}

func (s *Server) GetMultipleAircraftByRegistrant(ctx context.Context, query *pb.Query) (*pb.MultipleAircraft, error) {
	bs, err := s.dereg.svc.List("registrant_name", query.RegistrantName, "registrant.name", true)
	if err != nil {
		return nil, err
	}
	return bytesToMultipleAircraft(bs)
}

func bytesToAircraft(b []byte) (*pb.Aircraft, error) {
	record := &Record{}
	err := record.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return &pb.Aircraft{
		NNumber:                                  record.NNumber,
		SerialNumber:                             record.SerialNumber,
		ManufacturerAircraftModelCode:            record.Manufacturer.AircraftModelCode,
		ManufacturerEngineModelCode:              record.Manufacturer.EngineModelCode,
		ManufacturerYear:                         record.Manufacturer.Year,
		Status:                                   record.StatusCode.Description,
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
		RegistrantPhysicalAddress:                record.Registrant.PhysicalAddress,
		RegistrantPhysicalAddress2:               record.Registrant.PhysicalAddress2,
		RegistrantPhysicalCity:                   record.Registrant.PhysicalCity,
		RegistrantPhysicalState:                  record.Registrant.PhysicalState,
		RegistrantPhysicalZipCode:                record.Registrant.PhysicalZipCode,
		RegistrantPhysicalCounty:                 record.Registrant.PhysicalCounty,
		RegistrantPhysicalCountry:                record.Registrant.PhysicalCountry,
		CertificationAirworthinessClassification: record.Certification.AirworthinessClassification.Description,
		CertificationApprovedOperations:          record.Certification.ApprovedOperation.Description,
		AirworthinessDate:                        record.AirworthinessDate,
		CancelDate:                               record.CancelDate,
		ExportCountry:                            record.ExportCountry,
		LastActivityDate:                         record.LastActivityDate,
		CertificationIssueDate:                   record.CertificateIssueDate,
		OwnershipOtherName1:                      record.Ownership.OtherName1,
		OwnershipOtherName2:                      record.Ownership.OtherName2,
		OwnershipOtherName3:                      record.Ownership.OtherName3,
		OwnershipOtherName4:                      record.Ownership.OtherName4,
		OwnershipOtherName5:                      record.Ownership.OtherName5,
		KitManufacturerName:                      record.Kit.ManufacturerName,
		KitModelName:                             record.Kit.ModelName,
		ModeSCode:                                record.ModeS.Code,
		ModeSCodeHex:                             record.ModeS.CodeHex,
	}, nil
}

func bytesToMultipleAircraft(bs [][]byte) (*pb.MultipleAircraft, error) {
	aircraft := make([]*pb.Aircraft, len(bs))
	for i, b := range bs {
		a, err := bytesToAircraft(b)
		if err != nil {
			return nil, err
		}
		aircraft[i] = a
	}
	return &pb.MultipleAircraft{Aircraft: aircraft}, nil
}
