/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApplicationStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationStats{}

// ApplicationStats struct for ApplicationStats
type ApplicationStats struct {
	// The app's ID
	AppId string `json:"appId"`
	// The app's UID
	AppUid *string `json:"appUid,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	MessageDestinations int32 `json:"messageDestinations"`
}

type _ApplicationStats ApplicationStats

// NewApplicationStats instantiates a new ApplicationStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationStats(appId string, messageDestinations int32) *ApplicationStats {
	this := ApplicationStats{}
	this.AppId = appId
	this.MessageDestinations = messageDestinations
	return &this
}

// NewApplicationStatsWithDefaults instantiates a new ApplicationStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationStatsWithDefaults() *ApplicationStats {
	this := ApplicationStats{}
	return &this
}

// GetAppId returns the AppId field value
func (o *ApplicationStats) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *ApplicationStats) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *ApplicationStats) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise.
func (o *ApplicationStats) GetAppUid() string {
	if o == nil || IsNil(o.AppUid) {
		var ret string
		return ret
	}
	return *o.AppUid
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationStats) GetAppUidOk() (*string, bool) {
	if o == nil || IsNil(o.AppUid) {
		return nil, false
	}
	return o.AppUid, true
}

// HasAppUid returns a boolean if a field has been set.
func (o *ApplicationStats) HasAppUid() bool {
	if o != nil && !IsNil(o.AppUid) {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given string and assigns it to the AppUid field.
func (o *ApplicationStats) SetAppUid(v string) {
	o.AppUid = &v
}

// GetMessageDestinations returns the MessageDestinations field value
func (o *ApplicationStats) GetMessageDestinations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageDestinations
}

// GetMessageDestinationsOk returns a tuple with the MessageDestinations field value
// and a boolean to check if the value has been set.
func (o *ApplicationStats) GetMessageDestinationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageDestinations, true
}

// SetMessageDestinations sets field value
func (o *ApplicationStats) SetMessageDestinations(v int32) {
	o.MessageDestinations = v
}

func (o ApplicationStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	if !IsNil(o.AppUid) {
		toSerialize["appUid"] = o.AppUid
	}
	toSerialize["messageDestinations"] = o.MessageDestinations
	return toSerialize, nil
}

func (o *ApplicationStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appId",
		"messageDestinations",
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

	varApplicationStats := _ApplicationStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationStats)

	if err != nil {
		return err
	}

	*o = ApplicationStats(varApplicationStats)

	return err
}

type NullableApplicationStats struct {
	value *ApplicationStats
	isSet bool
}

func (v NullableApplicationStats) Get() *ApplicationStats {
	return v.value
}

func (v *NullableApplicationStats) Set(val *ApplicationStats) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationStats) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationStats(val *ApplicationStats) *NullableApplicationStats {
	return &NullableApplicationStats{value: val, isSet: true}
}

func (v NullableApplicationStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


