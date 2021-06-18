package reservedserver

type Codes struct {
	ReservationType map[string]string
}

func initCodes() Codes {
	codesReservationType := map[string]string{
		"AA": "Reserved - no fee",
		"A":  "Fee paid, notice for expiration sent",
		"HD": "2 year hold for canceled N-Numbers",
		"FN": "Fee paid, notice for expiration sent",
		"FP": "Fee paid",
		"MF": "Reserved to manufacturer - no fee, no expiration date",
		"MT": "Reserved to manufacturer - no expiration date",
		"NC": "N-Number change is in process",
		"NN": "N-Number change is in process, expiration notice sent",
		"CN": "N-Number change, Expire Notice Sent",
		"CE": "N-Number change Expired",
	}

	return Codes{
		ReservationType: codesReservationType,
	}
}
