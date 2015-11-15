
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

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/examples/generated/models"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// UpdatePetParams contains all the bound params for the update pet operation
// typically these are obtained from a http.Request
type UpdatePetParams struct {
	// Pet object that needs to be added to the store
	Body *models.Pet
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdatePetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	if err := route.Consumer.Consume(r.Body, o.Body); err != nil {
		res = append(res, errors.NewParseError("body", "body", "", err))
	} else {
		if err := o.Body.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
