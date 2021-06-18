package master

import (
	"context"

	pb "github.com/engelsjk/faadb/rpc/master"
)

type Server struct {
	master *Master
}

func NewServer(master *Master) *Server {
	return &Server{master: master}
}

func (s *Server) GetAircraft(ctx context.Context, query *pb.Query) (*pb.Aircraft, error) {
	bs, err := s.master.svc.Get(query.NNumber)
	if err != nil {
		return nil, err
	}
	return bytesToAircraft(bs[0])
}

func (s *Server) GetPossibleAircraft(ctx context.Context, query *pb.Query) (*pb.MultipleAircraft, error) {
	bs, err := s.master.svc.List("nnumber", query.NNumber, "nnumber", false)
	if err != nil {
		return nil, err
	}
	return bytesToMultipleAircraft(bs)
}

func (s *Server) GetMultipleAircraftByRegistrantName(ctx context.Context, query *pb.Query) (*pb.MultipleAircraft, error) {
	bs, err := s.master.svc.List("registrant_name", query.RegistrantName, "registrant.name", true)
	if err != nil {
		return nil, err
	}
	return bytesToMultipleAircraft(bs)
}

func (s *Server) GetMultipleAircraftByRegistrantStreet1(ctx context.Context, query *pb.Query) (*pb.MultipleAircraft, error) {
	bs, err := s.master.svc.List("registrant_street_1", query.RegistrantStreet1, "registrant.street_1", true)
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
