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

// PostTransactionSignReader is a Reader for the PostTransactionSign structure.
type PostTransactionSignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTransactionSignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostTransactionSignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostTransactionSignDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostTransactionSignOK creates a PostTransactionSignOK with default headers values
func NewPostTransactionSignOK() *PostTransactionSignOK {
	return &PostTransactionSignOK{}
}

/*PostTransactionSignOK handles this case with default header values.

success
*/
type PostTransactionSignOK struct {
	Payload *models.HttpsuccessResponse
}

func (o *PostTransactionSignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttpsuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionSignDefault creates a PostTransactionSignDefault with default headers values
func NewPostTransactionSignDefault(code int) *PostTransactionSignDefault {
	return &PostTransactionSignDefault{
		_statusCode: code,
	}
}

/*PostTransactionSignDefault handles this case with default header values.

error
*/
type PostTransactionSignDefault struct {
	_statusCode int

	Payload *models.HTTPErrorResponse
}

// Code gets the status code for the post transaction sign default response
func (o *PostTransactionSignDefault) Code() int {
	return o._statusCode
}

func (o *PostTransactionSignDefault) Error() string {
	return o.Payload.Error.Message
}

func (o *PostTransactionSignDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
