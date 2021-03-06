package fields

import (
	"os"
	"testing"

	"github.com/joshheinrichs/geosource/server/config"
	"github.com/stretchr/testify/assert"
)

func TestImagesUnmarshalValue(t *testing.T) {
	form := ImagesForm{}
	data := `["base64", "base64"]`
	_, err := form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = ``
	_, err = form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `[]`
	_, err = form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `{}`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)
}

func TestImagesValidateForm(t *testing.T) {
	form := &ImagesForm{}
	assert.NoError(t, form.ValidateForm())

	form = nil
	assert.NoError(t, form.ValidateForm())
}

func TestImagesValidateValue(t *testing.T) {
	fieldsConfig = config.New()
	fieldsConfig.Website.Directory = ""
	MediaDir = ""
	ImageDir = ""

	form := ImagesForm{}
	value := &ImagesValue{"data:image/png;base64,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"}
	assert.NoError(t, form.ValidateValue(value))
	assert.NoError(t, os.Remove(fieldsConfig.Website.Directory+MediaDir+ImageDir+(*value)[0]))

	value = &ImagesValue{"foo"}
	assert.Error(t, form.ValidateValue(value))

	value = &ImagesValue{""}
	assert.Error(t, form.ValidateValue(value))

	value = nil
	assert.NoError(t, form.ValidateValue(value))

	assert.Error(t, form.ValidateValue(nil))
}

func TestImagesIsComplete(t *testing.T) {
	value := &ImagesValue{"foo"}
	assert.True(t, value.IsComplete())
	value = &ImagesValue{}
	assert.False(t, value.IsComplete())
	value = nil
	assert.False(t, value.IsComplete())
}
