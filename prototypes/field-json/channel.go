package main

import (
	"./fields"
	"encoding/json"
	"errors"
)

type Channel struct {
	Name   string          `json:"name"`
	Fields []*fields.Field `json:"fields"`
}

// Unmarshals the given JSON blob, returning a Channel on success, or an error
// if unsuccessful. All fields must be valid and empty for parsing to succeed.
func UnmarshalChannel(blob []byte) (*Channel, error) {

	unmarshalChannel := struct {
		Channel
		JsonFields []json.RawMessage `json:"fields"`
	}{}

	json.Unmarshal(blob, &unmarshalChannel)

	channelFields := make([]*fields.Field, len(unmarshalChannel.JsonFields))

	for i, jsonField := range unmarshalChannel.JsonFields {
		field, err := fields.UnmarshalField(jsonField)
		if err != nil {
			return nil, err
		}
		if !field.IsEmpty() {
			return nil, errors.New("Non-empty field submitted")
		}
		channelFields[i] = field
	}
	unmarshalChannel.Channel.Fields = channelFields

	return &unmarshalChannel.Channel, nil
}

// Unmarshals the given JSON blob into a Submission, and attempts to validate
// and return it in Post form. If the submission is invalid due to either an
// unmarshalling error or a form mismatch, an error is returned.
func (channel *Channel) UnmarshalSubmission(blob []byte) (*Post, error) {

	unmarshalSubmission := struct {
		Submission
		JsonValues []json.RawMessage `json:"values"`
	}{}

	json.Unmarshal(blob, &unmarshalSubmission)
	if len(unmarshalSubmission.JsonValues) != len(channel.Fields) {
		return nil, errors.New("An invalid number of values were provided.")
	}

	post := Post{
		Title:   unmarshalSubmission.Title,
		Channel: unmarshalSubmission.Channel,
		Fields:  make([]*fields.Field, len(channel.Fields)),
	}

	for i, field := range channel.Fields {
		value, err := field.Form.UnmarshalValue(unmarshalSubmission.JsonValues[i])
		if err != nil {
			return nil, err
		}
		post.Fields[i] = &fields.Field{
			Label:    field.Label,
			Type:     field.Type,
			Required: field.Required,
			Form:     field.Form,
			Value:    value,
		}
	}

	return &post, nil
}
