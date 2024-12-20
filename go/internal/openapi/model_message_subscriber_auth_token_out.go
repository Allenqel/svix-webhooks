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

// checks if the MessageSubscriberAuthTokenOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageSubscriberAuthTokenOut{}

// MessageSubscriberAuthTokenOut struct for MessageSubscriberAuthTokenOut
type MessageSubscriberAuthTokenOut struct {
	BridgeToken string `json:"bridgeToken"`
	Token string `json:"token"`
}

type _MessageSubscriberAuthTokenOut MessageSubscriberAuthTokenOut

// NewMessageSubscriberAuthTokenOut instantiates a new MessageSubscriberAuthTokenOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageSubscriberAuthTokenOut(bridgeToken string, token string) *MessageSubscriberAuthTokenOut {
	this := MessageSubscriberAuthTokenOut{}
	this.BridgeToken = bridgeToken
	this.Token = token
	return &this
}

// NewMessageSubscriberAuthTokenOutWithDefaults instantiates a new MessageSubscriberAuthTokenOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageSubscriberAuthTokenOutWithDefaults() *MessageSubscriberAuthTokenOut {
	this := MessageSubscriberAuthTokenOut{}
	return &this
}

// GetBridgeToken returns the BridgeToken field value
func (o *MessageSubscriberAuthTokenOut) GetBridgeToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BridgeToken
}

// GetBridgeTokenOk returns a tuple with the BridgeToken field value
// and a boolean to check if the value has been set.
func (o *MessageSubscriberAuthTokenOut) GetBridgeTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BridgeToken, true
}

// SetBridgeToken sets field value
func (o *MessageSubscriberAuthTokenOut) SetBridgeToken(v string) {
	o.BridgeToken = v
}

// GetToken returns the Token field value
func (o *MessageSubscriberAuthTokenOut) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *MessageSubscriberAuthTokenOut) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *MessageSubscriberAuthTokenOut) SetToken(v string) {
	o.Token = v
}

func (o MessageSubscriberAuthTokenOut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageSubscriberAuthTokenOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bridgeToken"] = o.BridgeToken
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *MessageSubscriberAuthTokenOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bridgeToken",
		"token",
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

	varMessageSubscriberAuthTokenOut := _MessageSubscriberAuthTokenOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageSubscriberAuthTokenOut)

	if err != nil {
		return err
	}

	*o = MessageSubscriberAuthTokenOut(varMessageSubscriberAuthTokenOut)

	return err
}

type NullableMessageSubscriberAuthTokenOut struct {
	value *MessageSubscriberAuthTokenOut
	isSet bool
}

func (v NullableMessageSubscriberAuthTokenOut) Get() *MessageSubscriberAuthTokenOut {
	return v.value
}

func (v *NullableMessageSubscriberAuthTokenOut) Set(val *MessageSubscriberAuthTokenOut) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageSubscriberAuthTokenOut) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageSubscriberAuthTokenOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageSubscriberAuthTokenOut(val *MessageSubscriberAuthTokenOut) *NullableMessageSubscriberAuthTokenOut {
	return &NullableMessageSubscriberAuthTokenOut{value: val, isSet: true}
}

func (v NullableMessageSubscriberAuthTokenOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageSubscriberAuthTokenOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


