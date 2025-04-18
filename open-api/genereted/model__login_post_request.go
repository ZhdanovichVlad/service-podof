/*
backend service

Сервис для управления ПВЗ и приемкой товаров

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LoginPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginPostRequest{}

// LoginPostRequest struct for LoginPostRequest
type LoginPostRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type _LoginPostRequest LoginPostRequest

// NewLoginPostRequest instantiates a new LoginPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginPostRequest(email string, password string) *LoginPostRequest {
	this := LoginPostRequest{}
	this.Email = email
	this.Password = password
	return &this
}

// NewLoginPostRequestWithDefaults instantiates a new LoginPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginPostRequestWithDefaults() *LoginPostRequest {
	this := LoginPostRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *LoginPostRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *LoginPostRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *LoginPostRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *LoginPostRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LoginPostRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LoginPostRequest) SetPassword(v string) {
	o.Password = v
}

func (o LoginPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *LoginPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLoginPostRequest := _LoginPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLoginPostRequest)

	if err != nil {
		return err
	}

	*o = LoginPostRequest(varLoginPostRequest)

	return err
}

type NullableLoginPostRequest struct {
	value *LoginPostRequest
	isSet bool
}

func (v NullableLoginPostRequest) Get() *LoginPostRequest {
	return v.value
}

func (v *NullableLoginPostRequest) Set(val *LoginPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginPostRequest(val *LoginPostRequest) *NullableLoginPostRequest {
	return &NullableLoginPostRequest{value: val, isSet: true}
}

func (v NullableLoginPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


