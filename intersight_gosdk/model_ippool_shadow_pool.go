/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// IppoolShadowPool Shadow Pool is a tracking object created on behalf of an IP pool, for each VRF.
type IppoolShadowPool struct {
	PoolAbstractPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                   `json:"ObjectType"`
	IpV4Blocks []IppoolIpV4Block        `json:"IpV4Blocks,omitempty"`
	IpV4Config NullableIppoolIpV4Config `json:"IpV4Config,omitempty"`
	IpV6Blocks []IppoolIpV6Block        `json:"IpV6Blocks,omitempty"`
	IpV6Config NullableIppoolIpV6Config `json:"IpV6Config,omitempty"`
	// Number of IPv4 addresses currently in use.
	V4Assigned *int64 `json:"V4Assigned,omitempty"`
	// Number of IPv4 addresses in this pool.
	V4Size *int64 `json:"V4Size,omitempty"`
	// Number of IPv6 addresses currently in use.
	V6Assigned *int64 `json:"V6Assigned,omitempty"`
	// Number of IPv6 addresses in this pool.
	V6Size *int64 `json:"V6Size,omitempty"`
	// An array of relationships to ippoolShadowBlock resources.
	IpBlockHeads         []IppoolShadowBlockRelationship `json:"IpBlockHeads,omitempty"`
	Pool                 *IppoolPoolRelationship         `json:"Pool,omitempty"`
	Vrf                  *VrfVrfRelationship             `json:"Vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolShadowPool IppoolShadowPool

// NewIppoolShadowPool instantiates a new IppoolShadowPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolShadowPool(classId string, objectType string) *IppoolShadowPool {
	this := IppoolShadowPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIppoolShadowPoolWithDefaults instantiates a new IppoolShadowPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolShadowPoolWithDefaults() *IppoolShadowPool {
	this := IppoolShadowPool{}
	var classId string = "ippool.ShadowPool"
	this.ClassId = classId
	var objectType string = "ippool.ShadowPool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolShadowPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolShadowPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolShadowPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolShadowPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpV4Blocks returns the IpV4Blocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpV4Blocks() []IppoolIpV4Block {
	if o == nil {
		var ret []IppoolIpV4Block
		return ret
	}
	return o.IpV4Blocks
}

// GetIpV4BlocksOk returns a tuple with the IpV4Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpV4BlocksOk() (*[]IppoolIpV4Block, bool) {
	if o == nil || o.IpV4Blocks == nil {
		return nil, false
	}
	return &o.IpV4Blocks, true
}

// HasIpV4Blocks returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpV4Blocks() bool {
	if o != nil && o.IpV4Blocks != nil {
		return true
	}

	return false
}

// SetIpV4Blocks gets a reference to the given []IppoolIpV4Block and assigns it to the IpV4Blocks field.
func (o *IppoolShadowPool) SetIpV4Blocks(v []IppoolIpV4Block) {
	o.IpV4Blocks = v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpV4Config() IppoolIpV4Config {
	if o == nil || o.IpV4Config.Get() == nil {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.IpV4Config.Get()
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV4Config.Get(), o.IpV4Config.IsSet()
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpV4Config() bool {
	if o != nil && o.IpV4Config.IsSet() {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given NullableIppoolIpV4Config and assigns it to the IpV4Config field.
func (o *IppoolShadowPool) SetIpV4Config(v IppoolIpV4Config) {
	o.IpV4Config.Set(&v)
}

// SetIpV4ConfigNil sets the value for IpV4Config to be an explicit nil
func (o *IppoolShadowPool) SetIpV4ConfigNil() {
	o.IpV4Config.Set(nil)
}

// UnsetIpV4Config ensures that no value is present for IpV4Config, not even an explicit nil
func (o *IppoolShadowPool) UnsetIpV4Config() {
	o.IpV4Config.Unset()
}

// GetIpV6Blocks returns the IpV6Blocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpV6Blocks() []IppoolIpV6Block {
	if o == nil {
		var ret []IppoolIpV6Block
		return ret
	}
	return o.IpV6Blocks
}

// GetIpV6BlocksOk returns a tuple with the IpV6Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpV6BlocksOk() (*[]IppoolIpV6Block, bool) {
	if o == nil || o.IpV6Blocks == nil {
		return nil, false
	}
	return &o.IpV6Blocks, true
}

// HasIpV6Blocks returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpV6Blocks() bool {
	if o != nil && o.IpV6Blocks != nil {
		return true
	}

	return false
}

// SetIpV6Blocks gets a reference to the given []IppoolIpV6Block and assigns it to the IpV6Blocks field.
func (o *IppoolShadowPool) SetIpV6Blocks(v []IppoolIpV6Block) {
	o.IpV6Blocks = v
}

// GetIpV6Config returns the IpV6Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpV6Config() IppoolIpV6Config {
	if o == nil || o.IpV6Config.Get() == nil {
		var ret IppoolIpV6Config
		return ret
	}
	return *o.IpV6Config.Get()
}

// GetIpV6ConfigOk returns a tuple with the IpV6Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpV6ConfigOk() (*IppoolIpV6Config, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpV6Config.Get(), o.IpV6Config.IsSet()
}

// HasIpV6Config returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpV6Config() bool {
	if o != nil && o.IpV6Config.IsSet() {
		return true
	}

	return false
}

// SetIpV6Config gets a reference to the given NullableIppoolIpV6Config and assigns it to the IpV6Config field.
func (o *IppoolShadowPool) SetIpV6Config(v IppoolIpV6Config) {
	o.IpV6Config.Set(&v)
}

// SetIpV6ConfigNil sets the value for IpV6Config to be an explicit nil
func (o *IppoolShadowPool) SetIpV6ConfigNil() {
	o.IpV6Config.Set(nil)
}

// UnsetIpV6Config ensures that no value is present for IpV6Config, not even an explicit nil
func (o *IppoolShadowPool) UnsetIpV6Config() {
	o.IpV6Config.Unset()
}

// GetV4Assigned returns the V4Assigned field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetV4Assigned() int64 {
	if o == nil || o.V4Assigned == nil {
		var ret int64
		return ret
	}
	return *o.V4Assigned
}

// GetV4AssignedOk returns a tuple with the V4Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetV4AssignedOk() (*int64, bool) {
	if o == nil || o.V4Assigned == nil {
		return nil, false
	}
	return o.V4Assigned, true
}

// HasV4Assigned returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasV4Assigned() bool {
	if o != nil && o.V4Assigned != nil {
		return true
	}

	return false
}

// SetV4Assigned gets a reference to the given int64 and assigns it to the V4Assigned field.
func (o *IppoolShadowPool) SetV4Assigned(v int64) {
	o.V4Assigned = &v
}

// GetV4Size returns the V4Size field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetV4Size() int64 {
	if o == nil || o.V4Size == nil {
		var ret int64
		return ret
	}
	return *o.V4Size
}

// GetV4SizeOk returns a tuple with the V4Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetV4SizeOk() (*int64, bool) {
	if o == nil || o.V4Size == nil {
		return nil, false
	}
	return o.V4Size, true
}

// HasV4Size returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasV4Size() bool {
	if o != nil && o.V4Size != nil {
		return true
	}

	return false
}

// SetV4Size gets a reference to the given int64 and assigns it to the V4Size field.
func (o *IppoolShadowPool) SetV4Size(v int64) {
	o.V4Size = &v
}

// GetV6Assigned returns the V6Assigned field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetV6Assigned() int64 {
	if o == nil || o.V6Assigned == nil {
		var ret int64
		return ret
	}
	return *o.V6Assigned
}

// GetV6AssignedOk returns a tuple with the V6Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetV6AssignedOk() (*int64, bool) {
	if o == nil || o.V6Assigned == nil {
		return nil, false
	}
	return o.V6Assigned, true
}

// HasV6Assigned returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasV6Assigned() bool {
	if o != nil && o.V6Assigned != nil {
		return true
	}

	return false
}

// SetV6Assigned gets a reference to the given int64 and assigns it to the V6Assigned field.
func (o *IppoolShadowPool) SetV6Assigned(v int64) {
	o.V6Assigned = &v
}

// GetV6Size returns the V6Size field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetV6Size() int64 {
	if o == nil || o.V6Size == nil {
		var ret int64
		return ret
	}
	return *o.V6Size
}

// GetV6SizeOk returns a tuple with the V6Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetV6SizeOk() (*int64, bool) {
	if o == nil || o.V6Size == nil {
		return nil, false
	}
	return o.V6Size, true
}

// HasV6Size returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasV6Size() bool {
	if o != nil && o.V6Size != nil {
		return true
	}

	return false
}

// SetV6Size gets a reference to the given int64 and assigns it to the V6Size field.
func (o *IppoolShadowPool) SetV6Size(v int64) {
	o.V6Size = &v
}

// GetIpBlockHeads returns the IpBlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowPool) GetIpBlockHeads() []IppoolShadowBlockRelationship {
	if o == nil {
		var ret []IppoolShadowBlockRelationship
		return ret
	}
	return o.IpBlockHeads
}

// GetIpBlockHeadsOk returns a tuple with the IpBlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowPool) GetIpBlockHeadsOk() (*[]IppoolShadowBlockRelationship, bool) {
	if o == nil || o.IpBlockHeads == nil {
		return nil, false
	}
	return &o.IpBlockHeads, true
}

// HasIpBlockHeads returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasIpBlockHeads() bool {
	if o != nil && o.IpBlockHeads != nil {
		return true
	}

	return false
}

// SetIpBlockHeads gets a reference to the given []IppoolShadowBlockRelationship and assigns it to the IpBlockHeads field.
func (o *IppoolShadowPool) SetIpBlockHeads(v []IppoolShadowBlockRelationship) {
	o.IpBlockHeads = v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetPool() IppoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolShadowPool) SetPool(v IppoolPoolRelationship) {
	o.Pool = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *IppoolShadowPool) GetVrf() VrfVrfRelationship {
	if o == nil || o.Vrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowPool) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolShadowPool) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given VrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolShadowPool) SetVrf(v VrfVrfRelationship) {
	o.Vrf = &v
}

func (o IppoolShadowPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPool, errPoolAbstractPool := json.Marshal(o.PoolAbstractPool)
	if errPoolAbstractPool != nil {
		return []byte{}, errPoolAbstractPool
	}
	errPoolAbstractPool = json.Unmarshal([]byte(serializedPoolAbstractPool), &toSerialize)
	if errPoolAbstractPool != nil {
		return []byte{}, errPoolAbstractPool
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpV4Blocks != nil {
		toSerialize["IpV4Blocks"] = o.IpV4Blocks
	}
	if o.IpV4Config.IsSet() {
		toSerialize["IpV4Config"] = o.IpV4Config.Get()
	}
	if o.IpV6Blocks != nil {
		toSerialize["IpV6Blocks"] = o.IpV6Blocks
	}
	if o.IpV6Config.IsSet() {
		toSerialize["IpV6Config"] = o.IpV6Config.Get()
	}
	if o.V4Assigned != nil {
		toSerialize["V4Assigned"] = o.V4Assigned
	}
	if o.V4Size != nil {
		toSerialize["V4Size"] = o.V4Size
	}
	if o.V6Assigned != nil {
		toSerialize["V6Assigned"] = o.V6Assigned
	}
	if o.V6Size != nil {
		toSerialize["V6Size"] = o.V6Size
	}
	if o.IpBlockHeads != nil {
		toSerialize["IpBlockHeads"] = o.IpBlockHeads
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolShadowPool) UnmarshalJSON(bytes []byte) (err error) {
	type IppoolShadowPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                   `json:"ObjectType"`
		IpV4Blocks []IppoolIpV4Block        `json:"IpV4Blocks,omitempty"`
		IpV4Config NullableIppoolIpV4Config `json:"IpV4Config,omitempty"`
		IpV6Blocks []IppoolIpV6Block        `json:"IpV6Blocks,omitempty"`
		IpV6Config NullableIppoolIpV6Config `json:"IpV6Config,omitempty"`
		// Number of IPv4 addresses currently in use.
		V4Assigned *int64 `json:"V4Assigned,omitempty"`
		// Number of IPv4 addresses in this pool.
		V4Size *int64 `json:"V4Size,omitempty"`
		// Number of IPv6 addresses currently in use.
		V6Assigned *int64 `json:"V6Assigned,omitempty"`
		// Number of IPv6 addresses in this pool.
		V6Size *int64 `json:"V6Size,omitempty"`
		// An array of relationships to ippoolShadowBlock resources.
		IpBlockHeads []IppoolShadowBlockRelationship `json:"IpBlockHeads,omitempty"`
		Pool         *IppoolPoolRelationship         `json:"Pool,omitempty"`
		Vrf          *VrfVrfRelationship             `json:"Vrf,omitempty"`
	}

	varIppoolShadowPoolWithoutEmbeddedStruct := IppoolShadowPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIppoolShadowPoolWithoutEmbeddedStruct)
	if err == nil {
		varIppoolShadowPool := _IppoolShadowPool{}
		varIppoolShadowPool.ClassId = varIppoolShadowPoolWithoutEmbeddedStruct.ClassId
		varIppoolShadowPool.ObjectType = varIppoolShadowPoolWithoutEmbeddedStruct.ObjectType
		varIppoolShadowPool.IpV4Blocks = varIppoolShadowPoolWithoutEmbeddedStruct.IpV4Blocks
		varIppoolShadowPool.IpV4Config = varIppoolShadowPoolWithoutEmbeddedStruct.IpV4Config
		varIppoolShadowPool.IpV6Blocks = varIppoolShadowPoolWithoutEmbeddedStruct.IpV6Blocks
		varIppoolShadowPool.IpV6Config = varIppoolShadowPoolWithoutEmbeddedStruct.IpV6Config
		varIppoolShadowPool.V4Assigned = varIppoolShadowPoolWithoutEmbeddedStruct.V4Assigned
		varIppoolShadowPool.V4Size = varIppoolShadowPoolWithoutEmbeddedStruct.V4Size
		varIppoolShadowPool.V6Assigned = varIppoolShadowPoolWithoutEmbeddedStruct.V6Assigned
		varIppoolShadowPool.V6Size = varIppoolShadowPoolWithoutEmbeddedStruct.V6Size
		varIppoolShadowPool.IpBlockHeads = varIppoolShadowPoolWithoutEmbeddedStruct.IpBlockHeads
		varIppoolShadowPool.Pool = varIppoolShadowPoolWithoutEmbeddedStruct.Pool
		varIppoolShadowPool.Vrf = varIppoolShadowPoolWithoutEmbeddedStruct.Vrf
		*o = IppoolShadowPool(varIppoolShadowPool)
	} else {
		return err
	}

	varIppoolShadowPool := _IppoolShadowPool{}

	err = json.Unmarshal(bytes, &varIppoolShadowPool)
	if err == nil {
		o.PoolAbstractPool = varIppoolShadowPool.PoolAbstractPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpV4Blocks")
		delete(additionalProperties, "IpV4Config")
		delete(additionalProperties, "IpV6Blocks")
		delete(additionalProperties, "IpV6Config")
		delete(additionalProperties, "V4Assigned")
		delete(additionalProperties, "V4Size")
		delete(additionalProperties, "V6Assigned")
		delete(additionalProperties, "V6Size")
		delete(additionalProperties, "IpBlockHeads")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "Vrf")

		// remove fields from embedded structs
		reflectPoolAbstractPool := reflect.ValueOf(o.PoolAbstractPool)
		for i := 0; i < reflectPoolAbstractPool.Type().NumField(); i++ {
			t := reflectPoolAbstractPool.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIppoolShadowPool struct {
	value *IppoolShadowPool
	isSet bool
}

func (v NullableIppoolShadowPool) Get() *IppoolShadowPool {
	return v.value
}

func (v *NullableIppoolShadowPool) Set(val *IppoolShadowPool) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolShadowPool) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolShadowPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolShadowPool(val *IppoolShadowPool) *NullableIppoolShadowPool {
	return &NullableIppoolShadowPool{value: val, isSet: true}
}

func (v NullableIppoolShadowPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolShadowPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}