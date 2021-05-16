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

func (i *Indicator) Validate() error {
	if err := i.checkParamsConsistency(); err != nil {
		return err
	}

	return nil
}

func (i *Indicator) checkParamsConsistency() error {
	if i.ValueParams == nil {
		return nil
	}

	if !(i.ValueParams.Min+i.ValueParams.Step < i.ValueParams.Max) {
		return fmt.Errorf(
			"min + step must be less than max, got !(%f + %f < %f)",
			i.ValueParams.Min, i.ValueParams.Step, i.ValueParams.Max,
		)
	}

	if i.ValueParams.Min > i.ValueParams.Default {
		return fmt.Errorf("default mustn't be less than min")
	}

	if i.ValueParams.Max < i.ValueParams.Default {
		return fmt.Errorf("default mustn't be greater than max")
	}

	return nil
}

func (s *Scale) JSONString() string {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}`, err)
	}
	return string(jsonBytes)
}

func (d *Dictionary) JSONString() string {
	jsonBytes, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}`, err)
	}
	return string(jsonBytes)
}

func (e *DictionaryEntry) JSONString() string {
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}`, err)
	}
	return string(jsonBytes)
}
