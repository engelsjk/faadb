package engineserver

type Codes struct {
	EngineType map[string]string
}

func initCodes() Codes {

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

	return Codes{
		EngineType: codesEngineType,
	}
}
