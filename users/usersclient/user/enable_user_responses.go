// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rezaAmiri123/microservice/users/usersclient/models"
)

// EnableUserReader is a Reader for the EnableUser structure.
type EnableUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEnableUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableUserOK creates a EnableUserOK with default headers values
func NewEnableUserOK() *EnableUserOK {
	return &EnableUserOK{}
}

/*
EnableUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type EnableUserOK struct {
	Payload models.UserspbEnableUserResponse
}

// IsSuccess returns true when this enable user o k response has a 2xx status code
func (o *EnableUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable user o k response has a 3xx status code
func (o *EnableUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable user o k response has a 4xx status code
func (o *EnableUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable user o k response has a 5xx status code
func (o *EnableUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enable user o k response a status code equal to that given
func (o *EnableUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enable user o k response
func (o *EnableUserOK) Code() int {
	return 200
}

func (o *EnableUserOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/enable_user][%d] enableUserOK  %+v", 200, o.Payload)
}

func (o *EnableUserOK) String() string {
	return fmt.Sprintf("[PATCH /v1/enable_user][%d] enableUserOK  %+v", 200, o.Payload)
}

func (o *EnableUserOK) GetPayload() models.UserspbEnableUserResponse {
	return o.Payload
}

func (o *EnableUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableUserDefault creates a EnableUserDefault with default headers values
func NewEnableUserDefault(code int) *EnableUserDefault {
	return &EnableUserDefault{
		_statusCode: code,
	}
}

/*
EnableUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EnableUserDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this enable user default response has a 2xx status code
func (o *EnableUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enable user default response has a 3xx status code
func (o *EnableUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enable user default response has a 4xx status code
func (o *EnableUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enable user default response has a 5xx status code
func (o *EnableUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enable user default response a status code equal to that given
func (o *EnableUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enable user default response
func (o *EnableUserDefault) Code() int {
	return o._statusCode
}

func (o *EnableUserDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/enable_user][%d] enableUser default  %+v", o._statusCode, o.Payload)
}

func (o *EnableUserDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/enable_user][%d] enableUser default  %+v", o._statusCode, o.Payload)
}

func (o *EnableUserDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *EnableUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}