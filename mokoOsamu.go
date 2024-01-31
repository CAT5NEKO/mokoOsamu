package mokoOsamu

import (
	"strings"
)

func ConvertString(input string) string {

	chunks := strings.Fields(input)

	var result string

	for i, chunk := range chunks {
		switch {
		case i == 0:
			result += "あああ"
		case strings.ContainsAny(chunk, "、。"):
			result += "（" + chunk + "）"
		case i == len(chunks)-1:
			result += "もこもこもこ"
		default:
			result += "ああああ"
		}

		if i < len(chunks)-1 {
			result += " "
		}
	}

	return result
}
