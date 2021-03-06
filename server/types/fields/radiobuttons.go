package fields

import (
	"encoding/json"
	"errors"
	"strings"
)

// RadiobuttonsForm is an array of labels for the radiobuttons.
type RadiobuttonsForm []string

func (radiobuttonsForm *RadiobuttonsForm) ValidateForm() error {
	if radiobuttonsForm == nil {
		return errors.New("Missing form.")
	}
	if len(*radiobuttonsForm) <= 0 {
		return errors.New("At least one radiobutton required.")
	}
	labelMap := make(map[string]bool)
	for _, label := range *radiobuttonsForm {
		label = strings.TrimSpace(label)
		if len(label) == 0 {
			return errors.New("All radiobuttons must have labels.")
		}
		_, ok := labelMap[label]
		if ok {
			return errors.New("Duplicate radiobutton label.")
		}
		labelMap[label] = true
	}
	return nil
}

func (radiobuttonsForm *RadiobuttonsForm) ValidateValue(value Value) error {
	radiobuttonsValue, ok := value.(*RadiobuttonsValue)
	if !ok {
		return errors.New("Type mismatch.")
	}
	if radiobuttonsValue == nil {
		return nil
	}
	for _, label := range *radiobuttonsForm {
		if label == string(*radiobuttonsValue) {
			return nil
		}
	}
	return errors.New("No matching radiobutton.")
}

func (radiobuttonsForm *RadiobuttonsForm) UnmarshalValue(blob []byte) (Value, error) {
	if len(blob) <= 0 {
		return nil, nil
	}
	var value RadiobuttonsValue
	err := json.Unmarshal(blob, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// RadiobuttonsValue is a string which corresponds to one of the labels inside
// the form.
type RadiobuttonsValue string

func (radiobuttonsValue *RadiobuttonsValue) IsEmpty() bool {
	return !radiobuttonsValue.IsComplete()
}

func (radiobuttonsValue *RadiobuttonsValue) IsComplete() bool {
	return radiobuttonsValue != nil && len(*radiobuttonsValue) > 0
}
