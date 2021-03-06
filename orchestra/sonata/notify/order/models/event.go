// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Event Event class is used to describe information structure used for notification.
//
// swagger:discriminator Event eventId
type Event interface {
	runtime.Validatable

	// event
	// Required: true
	Event() *ProductOrderEvent
	SetEvent(*ProductOrderEvent)

	// event Id
	// Required: true
	EventID() string
	SetEventID(string)

	// event time
	// Required: true
	// Format: date-time
	EventTime() *strfmt.DateTime
	SetEventTime(*strfmt.DateTime)

	// event type
	// Required: true
	EventType() ProductOrderEventType
	SetEventType(ProductOrderEventType)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type event struct {
	eventField *ProductOrderEvent

	eventIdField string

	eventTimeField *strfmt.DateTime

	eventTypeField ProductOrderEventType
}

// Event gets the event of this polymorphic type
func (m *event) Event() *ProductOrderEvent {
	return m.eventField
}

// SetEvent sets the event of this polymorphic type
func (m *event) SetEvent(val *ProductOrderEvent) {
	m.eventField = val
}

// EventID gets the event Id of this polymorphic type
func (m *event) EventID() string {
	return "Event"
}

// SetEventID sets the event Id of this polymorphic type
func (m *event) SetEventID(val string) {
}

// EventTime gets the event time of this polymorphic type
func (m *event) EventTime() *strfmt.DateTime {
	return m.eventTimeField
}

// SetEventTime sets the event time of this polymorphic type
func (m *event) SetEventTime(val *strfmt.DateTime) {
	m.eventTimeField = val
}

// EventType gets the event type of this polymorphic type
func (m *event) EventType() ProductOrderEventType {
	return m.eventTypeField
}

// SetEventType sets the event type of this polymorphic type
func (m *event) SetEventType(val ProductOrderEventType) {
	m.eventTypeField = val
}

// UnmarshalEventSlice unmarshals polymorphic slices of Event
func UnmarshalEventSlice(reader io.Reader, consumer runtime.Consumer) ([]Event, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Event
	for _, element := range elements {
		obj, err := unmarshalEvent(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalEvent unmarshals polymorphic Event
func UnmarshalEvent(reader io.Reader, consumer runtime.Consumer) (Event, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalEvent(data, consumer)
}

func unmarshalEvent(data []byte, consumer runtime.Consumer) (Event, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the eventId property.
	var getType struct {
		EventID string `json:"eventId"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("eventId", "body", getType.EventID); err != nil {
		return nil, err
	}

	// The value of eventId is used to determine which type to create and unmarshal the data into
	switch getType.EventID {
	case "Event":
		var result event
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "EventPlus":
		var result EventPlus
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid eventId value: %q", getType.EventID)
}

// Validate validates this event
func (m *event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *event) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event()); err != nil {
		return err
	}

	if m.Event() != nil {
		if err := m.Event().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event")
			}
			return err
		}
	}

	return nil
}

func (m *event) validateEventTime(formats strfmt.Registry) error {

	if err := validate.Required("eventTime", "body", m.EventTime()); err != nil {
		return err
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *event) validateEventType(formats strfmt.Registry) error {

	if err := m.EventType().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eventType")
		}
		return err
	}

	return nil
}
