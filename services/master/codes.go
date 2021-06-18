package master

import (
	"strings"

	"github.com/engelsjk/faadb/internal/codes"
)

type Codes struct {
	RegistrantType      map[string]string
	RegistrantRegion    map[string]string
	Certification       CertificationCodes
	AircraftType        map[string]string
	EngineType          map[string]string
	StatusCode          map[string]string
	FractionalOwnership map[string]bool
}

type CertificationCodes struct {
	AirworthinessClassification map[string]string
	ApprovedOperation           ApprovedOperationCodes
}

type ApprovedOperationCodesMultiple struct {
	Part1 map[string]string
	Part2 map[string]string
}

type ApprovedOperationCodes struct {
	Standard            map[string]string
	Limited             map[string]string
	Restricted          map[string]string
	Experimental        map[string]string
	Provisional         map[string]string
	Multiple            ApprovedOperationCodesMultiple
	Primary             map[string]string
	SpecialFlightPermit map[string]string
	LightSport          map[string]string
}

func initCodes() Codes {
	codesRegistrantType := map[string]string{
		"1": "Individual",
		"2": "Partnership",
		"3": "Corporation",
		"4": "Co-Owned",
		"5": "Government",
		"7": "LLC",
		"8": "Non Citizen Corporation",
		"9": "Non Citizen Co-Owned",
	}

	codesRegistrationRegion := map[string]string{
		"1": "Eastern",
		"2": "SouthWestern",
		"3": "Central",
		"4": "Western-Pacific",
		"5": "Alaskan",
		"7": "Southern",
		"8": "European",
		"C": "Great Lakes",
		"E": "New England",
		"S": "Northwest Mountain",
	}

	codesAirworthinessClassification := map[string]string{
		"1": "Standard",
		"2": "Limited",
		"3": "Restricted",
		"4": "Experimental",
		"5": "Provisional",
		"6": "Multiple",
		"7": "Primary",
		"8": "Special Flight Permit",
		"9": "Light Sport",
	}

	codesApprovedOpsStandard := map[string]string{
		"":  "",
		"N": "Normal",
		"U": "Utility",
		"A": "Acrobatic",
		"T": "Transport",
		"G": "Glider",
		"B": "Balloon",
		"C": "Commuter",
		"O": "Other",
	}

	codesApprovedOpsLimited := map[string]string{"": ""}

	codesApprovedOpsRestricted := map[string]string{
		"0": "Other",
		"1": "Agriculture and Pest Control",
		"2": "Aerial Surveying",
		"3": "Aerial Advertising",
		"4": "Forest",
		"5": "Patrolling",
		"6": "Weather Control",
		"7": "Carriage of Cargo",
	}

	codesApprovedOpsExperimental := map[string]string{
		"0":  "To show compliance with FAR",
		"1":  "Research and Development",
		"2":  "Amateur Built",
		"3":  "Exhibition",
		"4":  "Racing",
		"5":  "Crew Training",
		"6":  "Market Survey",
		"7":  "Operating Kit Built Aircraft",
		"8A": "Reg. Prior to 01/31/08",
		"8B": "Operating Light-Sport Kit-Built",
		"8C": "Operating Light-Sport Previously issued cert under 21.190",
		"9A": "Unmanned Aircraft - Research and Development",
		"9B": "Unmanned Aircraft - Market Survey",
		"9C": "Unmanned Aircraft - Crew Training",
		"9D": "Unmanned Aircraft – Exhibition",
		"9E": "Unmanned Aircraft – Compliance With CFR",
	}

	codesApprovedOpsProvisional := map[string]string{
		"0": "Class I",
		"1": "Class II",
	}

	codesApprovedOpsMultiple1 := map[string]string{
		"1": "Standard",
		"2": "Limited",
		"3": "Restricted",
	}

	codesApprovedOpsMultiple2 := map[string]string{
		"0": "Other",
		"1": "Agriculture and Pest Control",
		"2": "Aerial Surveying",
		"3": "erial Advertising",
		"4": "Forest",
		"5": "Patrolling",
		"6": "Weather Control",
		"7": "Carriage of Cargo",
	}

	codesApprovedOpsMultiple := ApprovedOperationCodesMultiple{
		Part1: codesApprovedOpsMultiple1,
		Part2: codesApprovedOpsMultiple2,
	}

	codesApprovedOpsPrimary := map[string]string{"": ""}

	codesApprovedOpsSpecialFlightPermit := map[string]string{
		"1": "Ferry flight for repairs, alterations, maintenance or storage",
		"2": "Evacuate from area of impending danger",
		"3": "Operation in excess of maximum certificated",
		"4": "Delivery or export",
		"5": "Production flight testing",
		"6": "Customer Demo",
	}

	codesApprovedOpsLightSport := map[string]string{
		"A": "Airplane",
		"G": "Glider",
		"L": "Lighter than Air",
		"P": "Power-Parachute",
		"W": "Weight-Shift-Control",
	}

	codesCertification := CertificationCodes{
		AirworthinessClassification: codesAirworthinessClassification,
		ApprovedOperation: ApprovedOperationCodes{
			Standard:            codesApprovedOpsStandard,
			Limited:             codesApprovedOpsLimited,
			Restricted:          codesApprovedOpsRestricted,
			Experimental:        codesApprovedOpsExperimental,
			Provisional:         codesApprovedOpsProvisional,
			Multiple:            codesApprovedOpsMultiple,
			Primary:             codesApprovedOpsPrimary,
			SpecialFlightPermit: codesApprovedOpsSpecialFlightPermit,
			LightSport:          codesApprovedOpsLightSport,
		},
	}

	codesAircraftType := map[string]string{
		"1": "Glider",
		"2": "Balloon",
		"3": "Blimp/Dirigible",
		"4": "Fixed wing single engine",
		"5": "Fixed wing multi engine",
		"6": "Rotorcraft",
		"7": "Weight-shift-control",
		"8": "Powered Parachute",
		"9": "Gyroplane",
		"H": "Hybrid Lift",
		"O": "Other",
	}

	codesEngineType := map[string]string{
		"0":  "None",
		"1":  "Reciprocating",
		"2":  "Turbo-prop",
		"3":  "Turbo-shaft",
		"4":  "Turbo-jet",
		"5":  "Turbo-fan",
		"6":  "Ramjet",
		"7":  "2 Cycle",
		"8":  "4 Cycle",
		"9":  "Unknown",
		"10": "Electric",
		"11": "Rotary",
	}

	codesStatusCode := map[string]string{
		"A":  "The Triennial Aircraft Registration form was mailed and has not been returned by the Post Office",
		"D":  "Expired Dealer",
		"E":  "The Certificate of Aircraft Registration was revoked by enforcement action",
		"M":  "Aircraft registered to the manufacturer under their Dealer Certificate",
		"N":  "Non-citizen Corporations which have not returned their flight hour reports",
		"R":  "Registration pending",
		"S":  "Second Triennial Aircraft Registration Form has been mailed and has not been returned by the Post Office",
		"T":  "Valid Registration from a Trainee",
		"V":  "Valid Registration",
		"W":  "Certificate of Registration has been deemed Ineffective or Invalid",
		"X":  "Enforcement Letter",
		"Z":  "Permanent Reserved",
		"1":  "Triennial Aircraft Registration form was returned by the Post Office as undeliverable",
		"2":  "N-Number Assigned –but has notyet been registered",
		"3":  "N-Number assigned as a Non Type   Certificated aircraft -but has not yet been registered",
		"4":  "N-Number assigned as import -but has not yet been registered",
		"5":  "Reserved N-Number",
		"6":  "Administratively canceled",
		"7":  "Sale reported",
		"8":  "A second attempt has been made at mailing a Triennial Aircraft Registration form to the owner with no response",
		"9":  "Certificate of Registration has been revoked",
		"10": "N-Number assigned, has not been registered and is pending cancellation",
		"11": "N-Number assigned as a Non Type Certificated (Amateur) but has not been registered that is pending cancellation",
		"12": "N-Number assigned as import but has not been registered that is pending cancellation",
		"13": "Registration Expired",
		"14": "First Notice for Re-Registration/Renewal",
		"15": "Second Notice for Re-Registration/Renewal",
		"16": "Registration Expired – Pending Cancellation",
		"17": "Sale Reported – Pending Cancellation",
		"18": "Sale Reported – Canceled",
		"19": "Registration Pending – Pending Cancellation",
		"20": "Registration Pending – Canceled",
		"21": "Revoked – Pending Cancellation",
		"22": "Revoked – Canceled",
		"23": "Expired Dealer (Pending Cancellation)",
		"24": "Third Notice for Re-Registration/Renewal",
		"25": "First Notice for Registration Renewal",
		"26": "Second Notice for Registration Renewal",
		"27": "Registration Expired",
		"28": "Third Notice for Registration Renewal",
		"29": "Registration Expired – Pending Cancellation",
	}

	return Codes{
		RegistrantType:   codesRegistrantType,
		RegistrantRegion: codesRegistrationRegion,
		Certification:    codesCertification,
		AircraftType:     codesAircraftType,
		EngineType:       codesEngineType,
		StatusCode:       codesStatusCode,
	}
}

func decodeCertification(id string, certificationCodes CertificationCodes) Certification {

	certification := Certification{}

	runes := []rune(id)

	if len(runes) == 0 {
		return certification
	}

	id1 := string(runes[0])

	certification.AirworthinessClassification = codes.Description{
		Code:        id1,
		Description: codes.DecodeDescription(id1, certificationCodes.AirworthinessClassification),
	}

	if len(runes) == 1 {
		return certification
	}

	switch id1 {
	case "1": // Standard
		certification.ApprovedOperation = codes.DecodeDescriptions(
			runes[1:],
			certificationCodes.ApprovedOperation.Standard,
		)
	case "2": // Limited
		certification.ApprovedOperation = codes.Description{Code: "", Description: ""}
	case "3": // Restricted
		certification.ApprovedOperation = codes.DecodeDescriptions(
			runes[1:],
			certificationCodes.ApprovedOperation.Restricted,
		)
	case "4": // Experimental
		certification.ApprovedOperation = codes.DecodeDescriptions(
			runes[1:],
			certificationCodes.ApprovedOperation.Experimental,
		)
	case "5": // Provisional
		id2 := string(runes[1])
		certification.ApprovedOperation = codes.Description{
			Code:        string(id2),
			Description: codes.DecodeDescription(id2, certificationCodes.ApprovedOperation.Provisional),
		}
	case "6": // Multiple
		part1 := codes.DecodeDescriptions(
			runes[1:],
			certificationCodes.ApprovedOperation.Multiple.Part1,
		)
		part2 := codes.DecodeDescriptions(
			runes[3:],
			certificationCodes.ApprovedOperation.Multiple.Part2,
		)
		certification.ApprovedOperation = codes.Description{
			Code:        strings.Join([]string{part1.Code, part2.Code}, ";"),
			Description: strings.Join([]string{part1.Description, part2.Description}, ";"),
		}
	case "7": // Primary
		certification.ApprovedOperation = codes.Description{Code: "", Description: ""}
	case "8": // Special Flight Permit
		certification.ApprovedOperation = codes.DecodeDescriptions(
			runes[1:],
			certificationCodes.ApprovedOperation.SpecialFlightPermit,
		)
	case "9": // Light Sport
		certification.ApprovedOperation = codes.DecodeDescriptions(
			runes[1:],
			certificationCodes.ApprovedOperation.LightSport,
		)
	default:
		// ???
	}

	return certification
}
