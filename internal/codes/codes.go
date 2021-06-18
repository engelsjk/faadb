package codes

import (
	"strconv"
	"strings"
)

type Description struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Status struct {
	Code   string `json:"code"`
	Status bool   `json:"status"`
}

func DecodeStatus(id string, m interface{}) bool {
	return decode(id, m).(bool)
}

func DecodeDescription(id string, m interface{}) string {
	return decode(id, m).(string)
}

func ParseInt32(n string) int32 {
	i, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return int32(i)
}

func DecodeDescriptions(runes []rune, codes map[string]string) Description {
	id2s := make([]string, len(runes))
	descriptions := make([]string, len(runes))
	for i, r := range runes {
		id2s[i] = string(r)
		descriptions[i] = DecodeDescription(id2s[i], codes)
	}
	return Description{
		Code:        strings.Join(id2s, ";"),
		Description: strings.Join(descriptions, ";"),
	}
}

func decode(id string, m interface{}) interface{} {
	switch ma := m.(type) {
	case map[string]bool:
		if v, ok := ma[id]; ok {
			return v
		}
		return false
	case map[string]string:
		if v, ok := ma[id]; ok {
			return v
		}
		return "unknown"
	default:
		return nil // bad idea?
	}
}
