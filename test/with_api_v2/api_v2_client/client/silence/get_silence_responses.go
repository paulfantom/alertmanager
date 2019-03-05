// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package silence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
)

// GetSilenceReader is a Reader for the GetSilence structure.
type GetSilenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSilenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSilenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSilenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSilenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSilenceOK creates a GetSilenceOK with default headers values
func NewGetSilenceOK() *GetSilenceOK {
	return &GetSilenceOK{}
}

/*GetSilenceOK handles this case with default header values.

Get silence response
*/
type GetSilenceOK struct {
	Payload *models.GettableSilence
}

func (o *GetSilenceOK) Error() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceOK  %+v", 200, o.Payload)
}

func (o *GetSilenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GettableSilence)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSilenceNotFound creates a GetSilenceNotFound with default headers values
func NewGetSilenceNotFound() *GetSilenceNotFound {
	return &GetSilenceNotFound{}
}

/*GetSilenceNotFound handles this case with default header values.

A silence with the specified ID was not found
*/
type GetSilenceNotFound struct {
}

func (o *GetSilenceNotFound) Error() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceNotFound ", 404)
}

func (o *GetSilenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSilenceInternalServerError creates a GetSilenceInternalServerError with default headers values
func NewGetSilenceInternalServerError() *GetSilenceInternalServerError {
	return &GetSilenceInternalServerError{}
}

/*GetSilenceInternalServerError handles this case with default header values.

Internal server error
*/
type GetSilenceInternalServerError struct {
	Payload string
}

func (o *GetSilenceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /silence/{silenceID}][%d] getSilenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSilenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
