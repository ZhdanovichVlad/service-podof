/*
backend service

Сервис для управления ПВЗ и приемкой товаров

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PvzGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PvzGet200ResponseInner{}

// PvzGet200ResponseInner struct for PvzGet200ResponseInner
type PvzGet200ResponseInner struct {
	Pvz *PVZ `json:"pvz,omitempty"`
	Receptions []PvzGet200ResponseInnerReceptionsInner `json:"receptions,omitempty"`
}

// NewPvzGet200ResponseInner instantiates a new PvzGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPvzGet200ResponseInner() *PvzGet200ResponseInner {
	this := PvzGet200ResponseInner{}
	return &this
}

// NewPvzGet200ResponseInnerWithDefaults instantiates a new PvzGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPvzGet200ResponseInnerWithDefaults() *PvzGet200ResponseInner {
	this := PvzGet200ResponseInner{}
	return &this
}

// GetPvz returns the Pvz field value if set, zero value otherwise.
func (o *PvzGet200ResponseInner) GetPvz() PVZ {
	if o == nil || IsNil(o.Pvz) {
		var ret PVZ
		return ret
	}
	return *o.Pvz
}

// GetPvzOk returns a tuple with the Pvz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PvzGet200ResponseInner) GetPvzOk() (*PVZ, bool) {
	if o == nil || IsNil(o.Pvz) {
		return nil, false
	}
	return o.Pvz, true
}

// HasPvz returns a boolean if a field has been set.
func (o *PvzGet200ResponseInner) HasPvz() bool {
	if o != nil && !IsNil(o.Pvz) {
		return true
	}

	return false
}

// SetPvz gets a reference to the given PVZ and assigns it to the Pvz field.
func (o *PvzGet200ResponseInner) SetPvz(v PVZ) {
	o.Pvz = &v
}

// GetReceptions returns the Receptions field value if set, zero value otherwise.
func (o *PvzGet200ResponseInner) GetReceptions() []PvzGet200ResponseInnerReceptionsInner {
	if o == nil || IsNil(o.Receptions) {
		var ret []PvzGet200ResponseInnerReceptionsInner
		return ret
	}
	return o.Receptions
}

// GetReceptionsOk returns a tuple with the Receptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PvzGet200ResponseInner) GetReceptionsOk() ([]PvzGet200ResponseInnerReceptionsInner, bool) {
	if o == nil || IsNil(o.Receptions) {
		return nil, false
	}
	return o.Receptions, true
}

// HasReceptions returns a boolean if a field has been set.
func (o *PvzGet200ResponseInner) HasReceptions() bool {
	if o != nil && !IsNil(o.Receptions) {
		return true
	}

	return false
}

// SetReceptions gets a reference to the given []PvzGet200ResponseInnerReceptionsInner and assigns it to the Receptions field.
func (o *PvzGet200ResponseInner) SetReceptions(v []PvzGet200ResponseInnerReceptionsInner) {
	o.Receptions = v
}

func (o PvzGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PvzGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pvz) {
		toSerialize["pvz"] = o.Pvz
	}
	if !IsNil(o.Receptions) {
		toSerialize["receptions"] = o.Receptions
	}
	return toSerialize, nil
}

type NullablePvzGet200ResponseInner struct {
	value *PvzGet200ResponseInner
	isSet bool
}

func (v NullablePvzGet200ResponseInner) Get() *PvzGet200ResponseInner {
	return v.value
}

func (v *NullablePvzGet200ResponseInner) Set(val *PvzGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePvzGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePvzGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePvzGet200ResponseInner(val *PvzGet200ResponseInner) *NullablePvzGet200ResponseInner {
	return &NullablePvzGet200ResponseInner{value: val, isSet: true}
}

func (v NullablePvzGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePvzGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


