
// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new todos API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for todos API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*AddOne add one API
 */
func (a *Client) AddOne(params AddOneParams, authInfo client.AuthInfoWriter) (*AddOneCreated, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "addOne",
		Params:   &params,
		Reader:   &AddOneReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOneCreated), nil
}

/*DestroyOne destroy one API
 */
func (a *Client) DestroyOne(params DestroyOneParams, authInfo client.AuthInfoWriter) (*DestroyOneNoContent, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "destroyOne",
		Params:   &params,
		Reader:   &DestroyOneReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DestroyOneNoContent), nil
}

/*Find find API
 */
func (a *Client) Find(params FindParams, authInfo client.AuthInfoWriter) (*FindOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "find",
		Params:   &params,
		Reader:   &FindReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindOK), nil
}

/*UpdateOne update one API
 */
func (a *Client) UpdateOne(params UpdateOneParams, authInfo client.AuthInfoWriter) (*UpdateOneOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "updateOne",
		Params:   &params,
		Reader:   &UpdateOneReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOneOK), nil
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
