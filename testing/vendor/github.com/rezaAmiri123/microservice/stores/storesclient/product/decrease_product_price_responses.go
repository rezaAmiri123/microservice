// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rezaAmiri123/microservice/stores/storesclient/models"
)

// DecreaseProductPriceReader is a Reader for the DecreaseProductPrice structure.
type DecreaseProductPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DecreaseProductPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDecreaseProductPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDecreaseProductPriceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDecreaseProductPriceOK creates a DecreaseProductPriceOK with default headers values
func NewDecreaseProductPriceOK() *DecreaseProductPriceOK {
	return &DecreaseProductPriceOK{}
}

/*
DecreaseProductPriceOK describes a response with status code 200, with default header values.

A successful response.
*/
type DecreaseProductPriceOK struct {
	Payload models.StorespbDecreaseProductPriceResponse
}

// IsSuccess returns true when this decrease product price o k response has a 2xx status code
func (o *DecreaseProductPriceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this decrease product price o k response has a 3xx status code
func (o *DecreaseProductPriceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this decrease product price o k response has a 4xx status code
func (o *DecreaseProductPriceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this decrease product price o k response has a 5xx status code
func (o *DecreaseProductPriceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this decrease product price o k response a status code equal to that given
func (o *DecreaseProductPriceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the decrease product price o k response
func (o *DecreaseProductPriceOK) Code() int {
	return 200
}

func (o *DecreaseProductPriceOK) Error() string {
	return fmt.Sprintf("[POST /v1/api/stores/decrease_product_price][%d] decreaseProductPriceOK  %+v", 200, o.Payload)
}

func (o *DecreaseProductPriceOK) String() string {
	return fmt.Sprintf("[POST /v1/api/stores/decrease_product_price][%d] decreaseProductPriceOK  %+v", 200, o.Payload)
}

func (o *DecreaseProductPriceOK) GetPayload() models.StorespbDecreaseProductPriceResponse {
	return o.Payload
}

func (o *DecreaseProductPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDecreaseProductPriceDefault creates a DecreaseProductPriceDefault with default headers values
func NewDecreaseProductPriceDefault(code int) *DecreaseProductPriceDefault {
	return &DecreaseProductPriceDefault{
		_statusCode: code,
	}
}

/*
DecreaseProductPriceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DecreaseProductPriceDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this decrease product price default response has a 2xx status code
func (o *DecreaseProductPriceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this decrease product price default response has a 3xx status code
func (o *DecreaseProductPriceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this decrease product price default response has a 4xx status code
func (o *DecreaseProductPriceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this decrease product price default response has a 5xx status code
func (o *DecreaseProductPriceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this decrease product price default response a status code equal to that given
func (o *DecreaseProductPriceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the decrease product price default response
func (o *DecreaseProductPriceDefault) Code() int {
	return o._statusCode
}

func (o *DecreaseProductPriceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/api/stores/decrease_product_price][%d] decreaseProductPrice default  %+v", o._statusCode, o.Payload)
}

func (o *DecreaseProductPriceDefault) String() string {
	return fmt.Sprintf("[POST /v1/api/stores/decrease_product_price][%d] decreaseProductPrice default  %+v", o._statusCode, o.Payload)
}

func (o *DecreaseProductPriceDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DecreaseProductPriceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
