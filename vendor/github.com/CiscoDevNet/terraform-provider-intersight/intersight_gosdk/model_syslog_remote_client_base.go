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

// SyslogRemoteClientBase Lists properties that are common to all remote logging clients. It serves as a base class for all remote logging clients.
type SyslogRemoteClientBase struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Enables/disables remote logging for the endpoint If enabled, log messages will be sent to the syslog server mentioned in the Hostname/IP Address field.
	Enabled *bool `json:"Enabled,omitempty"`
	// Hostname or IP Address of the syslog server where log should be stored.
	Hostname *string `json:"Hostname,omitempty"`
	// Lowest level of messages to be included in the remote log. * `warning` - Use logging level warning for logs classified as warning. * `emergency` - Use logging level emergency for logs classified as emergency. * `alert` - Use logging level alert for logs classified as alert. * `critical` - Use logging level critical for logs classified as critical. * `error` - Use logging level error for logs classified as error. * `notice` - Use logging level notice for logs classified as notice. * `informational` - Use logging level informational for logs classified as informational. * `debug` - Use logging level debug for logs classified as debug.
	MinSeverity *string `json:"MinSeverity,omitempty"`
	// Port number used for logging on syslog server.
	Port *int64 `json:"Port,omitempty"`
	// Transport layer protocol for transmission of log messages to syslog server. * `udp` - Use User Datagram Protocol (UDP) for syslog remote server connection. * `tcp` - Use Transmission Control Protocol (TCP) for syslog remote server connection.
	Protocol             *string `json:"Protocol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyslogRemoteClientBase SyslogRemoteClientBase

// NewSyslogRemoteClientBase instantiates a new SyslogRemoteClientBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogRemoteClientBase(classId string, objectType string) *SyslogRemoteClientBase {
	this := SyslogRemoteClientBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var minSeverity string = "warning"
	this.MinSeverity = &minSeverity
	var port int64 = 514
	this.Port = &port
	var protocol string = "udp"
	this.Protocol = &protocol
	return &this
}

// NewSyslogRemoteClientBaseWithDefaults instantiates a new SyslogRemoteClientBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogRemoteClientBaseWithDefaults() *SyslogRemoteClientBase {
	this := SyslogRemoteClientBase{}
	var classId string = "syslog.RemoteLoggingClient"
	this.ClassId = classId
	var objectType string = "syslog.RemoteLoggingClient"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var minSeverity string = "warning"
	this.MinSeverity = &minSeverity
	var port int64 = 514
	this.Port = &port
	var protocol string = "udp"
	this.Protocol = &protocol
	return &this
}

// GetClassId returns the ClassId field value
func (o *SyslogRemoteClientBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SyslogRemoteClientBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SyslogRemoteClientBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SyslogRemoteClientBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SyslogRemoteClientBase) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBase) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SyslogRemoteClientBase) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SyslogRemoteClientBase) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SyslogRemoteClientBase) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBase) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SyslogRemoteClientBase) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SyslogRemoteClientBase) SetHostname(v string) {
	o.Hostname = &v
}

// GetMinSeverity returns the MinSeverity field value if set, zero value otherwise.
func (o *SyslogRemoteClientBase) GetMinSeverity() string {
	if o == nil || o.MinSeverity == nil {
		var ret string
		return ret
	}
	return *o.MinSeverity
}

// GetMinSeverityOk returns a tuple with the MinSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBase) GetMinSeverityOk() (*string, bool) {
	if o == nil || o.MinSeverity == nil {
		return nil, false
	}
	return o.MinSeverity, true
}

// HasMinSeverity returns a boolean if a field has been set.
func (o *SyslogRemoteClientBase) HasMinSeverity() bool {
	if o != nil && o.MinSeverity != nil {
		return true
	}

	return false
}

// SetMinSeverity gets a reference to the given string and assigns it to the MinSeverity field.
func (o *SyslogRemoteClientBase) SetMinSeverity(v string) {
	o.MinSeverity = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SyslogRemoteClientBase) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBase) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SyslogRemoteClientBase) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SyslogRemoteClientBase) SetPort(v int64) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SyslogRemoteClientBase) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogRemoteClientBase) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SyslogRemoteClientBase) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SyslogRemoteClientBase) SetProtocol(v string) {
	o.Protocol = &v
}

func (o SyslogRemoteClientBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.MinSeverity != nil {
		toSerialize["MinSeverity"] = o.MinSeverity
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SyslogRemoteClientBase) UnmarshalJSON(bytes []byte) (err error) {
	type SyslogRemoteClientBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Enables/disables remote logging for the endpoint If enabled, log messages will be sent to the syslog server mentioned in the Hostname/IP Address field.
		Enabled *bool `json:"Enabled,omitempty"`
		// Hostname or IP Address of the syslog server where log should be stored.
		Hostname *string `json:"Hostname,omitempty"`
		// Lowest level of messages to be included in the remote log. * `warning` - Use logging level warning for logs classified as warning. * `emergency` - Use logging level emergency for logs classified as emergency. * `alert` - Use logging level alert for logs classified as alert. * `critical` - Use logging level critical for logs classified as critical. * `error` - Use logging level error for logs classified as error. * `notice` - Use logging level notice for logs classified as notice. * `informational` - Use logging level informational for logs classified as informational. * `debug` - Use logging level debug for logs classified as debug.
		MinSeverity *string `json:"MinSeverity,omitempty"`
		// Port number used for logging on syslog server.
		Port *int64 `json:"Port,omitempty"`
		// Transport layer protocol for transmission of log messages to syslog server. * `udp` - Use User Datagram Protocol (UDP) for syslog remote server connection. * `tcp` - Use Transmission Control Protocol (TCP) for syslog remote server connection.
		Protocol *string `json:"Protocol,omitempty"`
	}

	varSyslogRemoteClientBaseWithoutEmbeddedStruct := SyslogRemoteClientBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSyslogRemoteClientBaseWithoutEmbeddedStruct)
	if err == nil {
		varSyslogRemoteClientBase := _SyslogRemoteClientBase{}
		varSyslogRemoteClientBase.ClassId = varSyslogRemoteClientBaseWithoutEmbeddedStruct.ClassId
		varSyslogRemoteClientBase.ObjectType = varSyslogRemoteClientBaseWithoutEmbeddedStruct.ObjectType
		varSyslogRemoteClientBase.Enabled = varSyslogRemoteClientBaseWithoutEmbeddedStruct.Enabled
		varSyslogRemoteClientBase.Hostname = varSyslogRemoteClientBaseWithoutEmbeddedStruct.Hostname
		varSyslogRemoteClientBase.MinSeverity = varSyslogRemoteClientBaseWithoutEmbeddedStruct.MinSeverity
		varSyslogRemoteClientBase.Port = varSyslogRemoteClientBaseWithoutEmbeddedStruct.Port
		varSyslogRemoteClientBase.Protocol = varSyslogRemoteClientBaseWithoutEmbeddedStruct.Protocol
		*o = SyslogRemoteClientBase(varSyslogRemoteClientBase)
	} else {
		return err
	}

	varSyslogRemoteClientBase := _SyslogRemoteClientBase{}

	err = json.Unmarshal(bytes, &varSyslogRemoteClientBase)
	if err == nil {
		o.MoBaseComplexType = varSyslogRemoteClientBase.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "MinSeverity")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Protocol")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableSyslogRemoteClientBase struct {
	value *SyslogRemoteClientBase
	isSet bool
}

func (v NullableSyslogRemoteClientBase) Get() *SyslogRemoteClientBase {
	return v.value
}

func (v *NullableSyslogRemoteClientBase) Set(val *SyslogRemoteClientBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogRemoteClientBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogRemoteClientBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogRemoteClientBase(val *SyslogRemoteClientBase) *NullableSyslogRemoteClientBase {
	return &NullableSyslogRemoteClientBase{value: val, isSet: true}
}

func (v NullableSyslogRemoteClientBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogRemoteClientBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}