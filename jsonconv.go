package commonservicemodels

import "encoding/json"

// ToJSON generates json out of the passed type
func ToJSON(v interface{}) (string, error) {
	byteData, err := json.Marshal(v)
	stringData := string(byteData)
	return stringData, err
}

//FromJSON get struct type from Json string
func FromJSON(jsonString string, v interface{}) error {

	err := json.Unmarshal([]byte(jsonString), v)
	return err
}
