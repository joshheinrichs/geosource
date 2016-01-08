package field

import (
	"errors"
	"log"
)

const (
	TYPE_TEXT         = "text"
	TYPE_NUMBER       = "number"
	TYPE_CHECKBOXES   = "checkboxes"
	TYPE_RADIOBUTTONS = "radiobuttons"
	TYPE_IMAGES       = "images"
)

var ErrMissingType error = errors.New("Missing type.")
var ErrInvalidType error = errors.New("Invalid type.")
var ErrMissingLabel error = errors.New("Missing label.")
var ErrInvalidLabel error = errors.New("Invalid label.")
var ErrMissingValue error = errors.New("Missing value.")
var ErrInvalidValue error = errors.New("Invalid value.")

type Field struct {
	Type         string        `json:"type"`
	Label        string        `json:"label"`
	Required     bool          `json:"required"`
	Text         *Text         `json:"text,omitempty"`
	Number       *Number       `json:"number,omitempty"`
	Checkboxes   *Checkboxes   `json:"checkboxes,omitempty"`
	Radiobuttons *Radiobuttons `json:"radiobuttons,omitempty"`
	Images       *Images       `json:"images,omitempty"`
}

type Value interface {
	Validate() error
	IsEmpty() bool
}

func (field *Field) Validate() error {
	// ensure that only one field is filled out
	if field.Type != TYPE_TEXT && field.Text != nil ||
		field.Type != TYPE_NUMBER && field.Number != nil ||
		field.Type != TYPE_CHECKBOXES && field.Checkboxes != nil ||
		field.Type != TYPE_RADIOBUTTONS && field.Radiobuttons != nil ||
		field.Type != TYPE_IMAGES && field.Images != nil {
		return ErrInvalidValue
	}

	var value Value
	switch field.Type {
	case TYPE_TEXT:
		value = field.Text
	case TYPE_NUMBER:
		value = field.Number
	case TYPE_CHECKBOXES:
		value = field.Checkboxes
	case TYPE_RADIOBUTTONS:
		value = field.Radiobuttons
	case TYPE_IMAGES:
		value = field.Images
		log.Println("Image field")
	default:
		return ErrInvalidType
	}

	err := value.Validate()
	if err != nil {
		return err
	}
	if field.Required && value.IsEmpty() {
		return ErrInvalidValue
	}
	return nil
}

// func (field *Field) Validate() error {
// 	switch field.Type {
// 	// case TYPE_CHECKBOXES:
// 	// 	checkboxes, err := NewCheckboxes(field)
// 	// 	return err
// 	// case TYPE_RADIOBUTTONS:
// 	// 	radiobuttons, err := NewRadiobuttons(field)
// 	// 	return err
// 	case TYPE_TEXT:
// 		_, err := NewText(field)
// 		return err
// 	// case TYPE_IMAGES:
// 	// 	images, err := NewImages(field)
// 	// 	return err
// 	// case TYPE_NUMBER:
// 	// 	number, err := NewNumber(field)
// 	// 	return err
// 	default:
// 		return ErrInvalidType
// 	}
// 	return nil
// }

// func ParseFieldMap(fieldMap map[string]interface{}) {
// 	var field Field
// 	switch field.Type {
// 	case TYPE_CHECKBOX:
// 		fmt.Println("Checkbox")
// 	case TYPE_CHECKBOXES:
// 		fmt.Println("Checkboxes")
// 	case TYPE_RADIOBUTTON:
// 		fmt.Println("Radiobutton")
// 	case TYPE_RADIOBUTTONS:
// 		fmt.Println("Radiobuttons")
// 	default:
// 		fmt.Println("Error")
// 	}
// }