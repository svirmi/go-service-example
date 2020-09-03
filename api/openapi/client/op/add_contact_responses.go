// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/powerman/go-service-goswagger-clean-example/api/openapi/model"
)

// AddContactReader is a Reader for the AddContact structure.
type AddContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddContactCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddContactDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddContactCreated creates a AddContactCreated with default headers values
func NewAddContactCreated() *AddContactCreated {
	return &AddContactCreated{}
}

/*AddContactCreated handles this case with default header values.

Contact added.
*/
type AddContactCreated struct {
	Payload *model.Contact
}

func (o *AddContactCreated) Error() string {
	return fmt.Sprintf("[POST /contacts][%d] addContactCreated  %+v", 201, o.Payload)
}

func (o *AddContactCreated) GetPayload() *model.Contact {
	return o.Payload
}

func (o *AddContactCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddContactDefault creates a AddContactDefault with default headers values
func NewAddContactDefault(code int) *AddContactDefault {
	return &AddContactDefault{
		_statusCode: code,
	}
}

/*AddContactDefault handles this case with default header values.

- 409.1000: contact already exists

*/
type AddContactDefault struct {
	_statusCode int

	Payload *model.Error
}

// Code gets the status code for the add contact default response
func (o *AddContactDefault) Code() int {
	return o._statusCode
}

func (o *AddContactDefault) Error() string {
	return fmt.Sprintf("[POST /contacts][%d] addContact default  %+v", o._statusCode, o.Payload)
}

func (o *AddContactDefault) GetPayload() *model.Error {
	return o.Payload
}

func (o *AddContactDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
