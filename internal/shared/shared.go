package shared

import "strconv"

func ConvertStringToInt64(str string) (*int64, error) {
	if n, err := strconv.Atoi(str); err == nil {
		n64 := int64(n)
		return &n64, nil
	} else {
		return nil, err
	}
}
