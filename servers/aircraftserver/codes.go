package aircraftserver

type Codes struct {
	AircraftType             map[string]string
	EngineType               map[string]string
	AircraftCategoryCode     map[string]string
	BuilderCertificationCode map[string]string
	AircraftWeight           map[string]string // ???
}

func initCodes() Codes {
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

	codesAircraftCategoryCode := map[string]string{
		"1": "Land",
		"2": "Sea",
		"3": "Amphibian",
	}

	codesBuilderCertificationCode := map[string]string{
		"0": "Type Certificated",
		"1": "Not Type Certificated",
		"2": "Light Sport",
	}

	codesAircraftWeight := map[string]string{
		"CLASS 1": "Up to 12,499",
		"CLASS 2": "12,500 - 19,199",
		"CLASS 3": "20,000 and over",
		"CLASS 4": "UAV up to 55",
	}

	return Codes{
		AircraftType:             codesAircraftType,
		EngineType:               codesEngineType,
		AircraftCategoryCode:     codesAircraftCategoryCode,
		BuilderCertificationCode: codesBuilderCertificationCode,
		AircraftWeight:           codesAircraftWeight,
	}
}
