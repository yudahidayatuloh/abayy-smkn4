package utils

import "strings"

func GetWarna(jurusan string) string {
	jurusan = strings.ToUpper(jurusan)

	switch {
	case strings.Contains(jurusan, "PPLG"):
		return "bg-green-800"
	case strings.Contains(jurusan, "TJKT"):
		return "bg-blue-800"
	case strings.Contains(jurusan, "TBSM"):
		return "bg-red-800"
	case strings.Contains(jurusan, "DKV"):
		return "bg-yellow-600"
	case strings.Contains(jurusan, "TOI"):
		return "bg-black"
	default:
		return "bg-gray-800"
	}
}
