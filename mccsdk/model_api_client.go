/*
MACH composer Cloud (MCC) Public API

# Introduction  MACH composer Cloud is a platform and API to facilitate and coordinate work across teams that build composable architectures using MACH technology.  All operations available in MACH composer cloud are available through this API. For more information about using it in your MACH architecture, have a look at the [documentation website](https://docs.machcomposer.io/cloud/index.html).

API version: 0.1.0
Contact: mach@labdigital.nl
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mccsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the ApiClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiClient{}

// ApiClient struct for ApiClient
type ApiClient struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// description about the api client
	Description *string `json:"description,omitempty"`
	// the client id
	ClientId string `json:"client_id"`
	// the client id
	ClientSecret string     `json:"client_secret"`
	LastUsedAt   *time.Time `json:"last_used_at,omitempty"`
	// Scope
	Scope []string `json:"scope"`
}

type _ApiClient ApiClient

// NewApiClient instantiates a new ApiClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClient(id string, createdAt time.Time, clientId string, clientSecret string, scope []string) *ApiClient {
	this := ApiClient{}
	this.Id = id
	this.CreatedAt = createdAt
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Scope = scope
	return &this
}

// NewApiClientWithDefaults instantiates a new ApiClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClientWithDefaults() *ApiClient {
	this := ApiClient{}
	return &this
}

// GetId returns the Id field value
func (o *ApiClient) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiClient) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiClient) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ApiClient) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ApiClient) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ApiClient) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiClient) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClient) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiClient) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiClient) SetDescription(v string) {
	o.Description = &v
}

// GetClientId returns the ClientId field value
func (o *ApiClient) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ApiClient) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ApiClient) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *ApiClient) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *ApiClient) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *ApiClient) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *ApiClient) GetLastUsedAt() time.Time {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClient) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ApiClient) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *ApiClient) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetScope returns the Scope field value
func (o *ApiClient) GetScope() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *ApiClient) GetScopeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope, true
}

// SetScope sets field value
func (o *ApiClient) SetScope(v []string) {
	o.Scope = v
}

func (o ApiClient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["client_secret"] = o.ClientSecret
	if !IsNil(o.LastUsedAt) {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	toSerialize["scope"] = o.Scope
	return toSerialize, nil
}

func (o *ApiClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"client_id",
		"client_secret",
		"scope",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiClient := _ApiClient{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiClient)

	if err != nil {
		return err
	}

	*o = ApiClient(varApiClient)

	return err
}

type NullableApiClient struct {
	value *ApiClient
	isSet bool
}

func (v NullableApiClient) Get() *ApiClient {
	return v.value
}

func (v *NullableApiClient) Set(val *ApiClient) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClient) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClient(val *ApiClient) *NullableApiClient {
	return &NullableApiClient{value: val, isSet: true}
}

func (v NullableApiClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
