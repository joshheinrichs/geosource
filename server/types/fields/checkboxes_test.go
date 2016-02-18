package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckboxesUnmarshalValue(t *testing.T) {
	form := CheckboxesForm{}
	data := `[true, false]`
	_, err := form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `[]`
	_, err = form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = ``
	_, err = form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `{}`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)

	data = `{"value": [true, false]}`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)
}

func TestCheckboxesValidate(t *testing.T) {
	form := &CheckboxesForm{}
	assert.Error(t, form.Validate())

	form = nil
	assert.Error(t, form.Validate())

	form = &CheckboxesForm{"foo"}
	assert.NoError(t, form.Validate())

	form = &CheckboxesForm{"foo", "foo"}
	assert.NoError(t, form.Validate())
}

func TestCheckboxesValidateValue(t *testing.T) {
	form := CheckboxesForm{"foo", "bar"}

	value := &CheckboxesValue{true, false}
	err := form.ValidateValue(value)
	assert.NoError(t, err)

	value = &CheckboxesValue{true, false, false}
	err = form.ValidateValue(value)
	assert.Error(t, err)

	value = &CheckboxesValue{}
	err = form.ValidateValue(value)
	assert.Error(t, err)

	err = form.ValidateValue(nil)
	assert.Error(t, err)
}

func TestCheckboxesIsComplete(t *testing.T) {
	value := &CheckboxesValue{}
	assert.True(t, value.IsComplete())
	value = &CheckboxesValue{true, false}
	assert.True(t, value.IsComplete())
	value = nil
	assert.False(t, value.IsComplete())
}
