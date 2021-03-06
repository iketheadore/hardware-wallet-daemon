// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/skycoin/hardware-wallet-daemon/src/models"
)

// PostIntermediatePinMatrixReader is a Reader for the PostIntermediatePinMatrix structure.
type PostIntermediatePinMatrixReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntermediatePinMatrixReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostIntermediatePinMatrixOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostIntermediatePinMatrixDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIntermediatePinMatrixOK creates a PostIntermediatePinMatrixOK with default headers values
func NewPostIntermediatePinMatrixOK() *PostIntermediatePinMatrixOK {
	return &PostIntermediatePinMatrixOK{}
}

/*PostIntermediatePinMatrixOK handles this case with default header values.

success
*/
type PostIntermediatePinMatrixOK struct {
	Payload *models.HttpsuccessResponse
}

func (o *PostIntermediatePinMatrixOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttpsuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntermediatePinMatrixDefault creates a PostIntermediatePinMatrixDefault with default headers values
func NewPostIntermediatePinMatrixDefault(code int) *PostIntermediatePinMatrixDefault {
	return &PostIntermediatePinMatrixDefault{
		_statusCode: code,
	}
}

/*PostIntermediatePinMatrixDefault handles this case with default header values.

error
*/
type PostIntermediatePinMatrixDefault struct {
	_statusCode int

	Payload *models.HTTPErrorResponse
}

// Code gets the status code for the post intermediate pin matrix default response
func (o *PostIntermediatePinMatrixDefault) Code() int {
	return o._statusCode
}

func (o *PostIntermediatePinMatrixDefault) Error() string {
	return o.Payload.Error.Message
}

func (o *PostIntermediatePinMatrixDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
