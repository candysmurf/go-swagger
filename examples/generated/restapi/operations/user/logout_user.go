
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

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// LogoutUserHandlerFunc turns a function with the right signature into a logout user handler
type LogoutUserHandlerFunc func() error

func (fn LogoutUserHandlerFunc) Handle() error {
	return fn()
}

// LogoutUserHandler interface for that can handle valid logout user params
type LogoutUserHandler interface {
	Handle() error
}

// NewLogoutUser creates a new http.Handler for the logout user operation
func NewLogoutUser(ctx *middleware.Context, handler LogoutUserHandler) *LogoutUser {
	return &LogoutUser{Context: ctx, Handler: handler}
}

/*
Logs out current logged in user session
*/
type LogoutUser struct {
	Context *middleware.Context
	Handler LogoutUserHandler
}

func (o *LogoutUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, nil); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	err := o.Handler.Handle() // actually handle the request
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	o.Context.Respond(rw, r, route.Produces, route, nil)

}
