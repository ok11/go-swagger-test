// Code generated by go-swagger; DO NOT EDIT.

package baz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ok11/go-swagger-test/pkg/gen/models"
)

// RegisterFooCreatedCode is the HTTP code returned for type RegisterFooCreated
const RegisterFooCreatedCode int = 201

/*RegisterFooCreated Created

swagger:response registerFooCreated
*/
type RegisterFooCreated struct {
	/*Fingerprint of the created foo

	 */
	ETag string `json:"ETag"`
	/*URI of the registered foo

	 */
	Location strfmt.URI `json:"Location"`

	/*
	  In: Body
	*/
	Payload *models.Foo `json:"body,omitempty"`
}

// NewRegisterFooCreated creates RegisterFooCreated with default headers values
func NewRegisterFooCreated() *RegisterFooCreated {

	return &RegisterFooCreated{}
}

// WithETag adds the eTag to the register foo created response
func (o *RegisterFooCreated) WithETag(eTag string) *RegisterFooCreated {
	o.ETag = eTag
	return o
}

// SetETag sets the eTag to the register foo created response
func (o *RegisterFooCreated) SetETag(eTag string) {
	o.ETag = eTag
}

// WithLocation adds the location to the register foo created response
func (o *RegisterFooCreated) WithLocation(location strfmt.URI) *RegisterFooCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the register foo created response
func (o *RegisterFooCreated) SetLocation(location strfmt.URI) {
	o.Location = location
}

// WithPayload adds the payload to the register foo created response
func (o *RegisterFooCreated) WithPayload(payload *models.Foo) *RegisterFooCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register foo created response
func (o *RegisterFooCreated) SetPayload(payload *models.Foo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterFooCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header ETag

	eTag := o.ETag
	if eTag != "" {
		rw.Header().Set("ETag", eTag)
	}

	// response header Location

	location := o.Location.String()
	if location != "" {
		rw.Header().Set("Location", location)
	}

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterFooBadRequestCode is the HTTP code returned for type RegisterFooBadRequest
const RegisterFooBadRequestCode int = 400

/*RegisterFooBadRequest Bad Request

swagger:response registerFooBadRequest
*/
type RegisterFooBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterFooBadRequest creates RegisterFooBadRequest with default headers values
func NewRegisterFooBadRequest() *RegisterFooBadRequest {

	return &RegisterFooBadRequest{}
}

// WithPayload adds the payload to the register foo bad request response
func (o *RegisterFooBadRequest) WithPayload(payload *models.Error) *RegisterFooBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register foo bad request response
func (o *RegisterFooBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterFooBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterFooUnauthorizedCode is the HTTP code returned for type RegisterFooUnauthorized
const RegisterFooUnauthorizedCode int = 401

/*RegisterFooUnauthorized Unauthorized

swagger:response registerFooUnauthorized
*/
type RegisterFooUnauthorized struct {
	/*

	 */
	WWWAuthenticate string `json:"WWW_Authenticate"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterFooUnauthorized creates RegisterFooUnauthorized with default headers values
func NewRegisterFooUnauthorized() *RegisterFooUnauthorized {

	return &RegisterFooUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the register foo unauthorized response
func (o *RegisterFooUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *RegisterFooUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the register foo unauthorized response
func (o *RegisterFooUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WithPayload adds the payload to the register foo unauthorized response
func (o *RegisterFooUnauthorized) WithPayload(payload *models.Error) *RegisterFooUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register foo unauthorized response
func (o *RegisterFooUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterFooUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterFooForbiddenCode is the HTTP code returned for type RegisterFooForbidden
const RegisterFooForbiddenCode int = 403

/*RegisterFooForbidden Forbidden

swagger:response registerFooForbidden
*/
type RegisterFooForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterFooForbidden creates RegisterFooForbidden with default headers values
func NewRegisterFooForbidden() *RegisterFooForbidden {

	return &RegisterFooForbidden{}
}

// WithPayload adds the payload to the register foo forbidden response
func (o *RegisterFooForbidden) WithPayload(payload *models.Error) *RegisterFooForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register foo forbidden response
func (o *RegisterFooForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterFooForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterFooInternalServerErrorCode is the HTTP code returned for type RegisterFooInternalServerError
const RegisterFooInternalServerErrorCode int = 500

/*RegisterFooInternalServerError Internal server error

swagger:response registerFooInternalServerError
*/
type RegisterFooInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterFooInternalServerError creates RegisterFooInternalServerError with default headers values
func NewRegisterFooInternalServerError() *RegisterFooInternalServerError {

	return &RegisterFooInternalServerError{}
}

// WithPayload adds the payload to the register foo internal server error response
func (o *RegisterFooInternalServerError) WithPayload(payload *models.Error) *RegisterFooInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register foo internal server error response
func (o *RegisterFooInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterFooInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
