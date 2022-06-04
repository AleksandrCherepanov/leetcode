package arrays_and_string

import "strings"

func addBinary(a, b string) string {
	additional := false
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}

	result := ""
	for i := len(a) - 1; i >= 0; i-- {
		sum := ""
		if string(a[i]) == "0" && string(b[i]) == "0" {
			sum = "0"
			if additional {
				sum = "1"
				additional = false
			}
			result = sum + result
			continue
		}

		if string(a[i]) == "1" && string(b[i]) == "1" {
			sum = "0"
			if additional {
				sum = "1"
			}
			additional = true
			result = sum + result
			continue
		}

		if additional {
			sum = "0"
		} else {
			sum = "1"
		}

		result = sum + result
	}

	if additional {
		result = "1" + result
	}

	return result
}
