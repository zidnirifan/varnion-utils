package tools

import (
	"encoding/json"
	"strconv"

	"github.com/google/uuid"
)

// String - UUID
func ConvertStringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

// String - Int
func ConvertStringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// Int - String
func ConvertIntToString(i int) string {
	return strconv.Itoa(i)
}

// String - Float
func ConvertStringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Float - String
func ConvertFloatToString(f float64, presisi int) string {
	return strconv.FormatFloat(f, 'f', presisi, 64)
}

// String - JSON
func ConvertStringToJSON(s string, data any) (any, error) {
	err := json.Unmarshal([]byte(s), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// JSON - String
func ConvertJSONToString(data any) ([]byte, error) {
	return json.Marshal(data)
}
