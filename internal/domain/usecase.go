package domain

import (
	"encoding/json"
	"fmt"
)

func (i *Indicator) JSONString() string {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}`, err)
	}
	return string(jsonBytes)
}

func (s *Scale) JSONString() string {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}`, err)
	}
	return string(jsonBytes)
}
