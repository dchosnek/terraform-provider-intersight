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
)

// RecoveryAbstractBackupConfigAllOf Definition of the list of properties defined in 'recovery.AbstractBackupConfig', excluding properties defined in parent classes.
type RecoveryAbstractBackupConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418.
	FileNamePrefix *string `json:"FileNamePrefix,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Specifies whether the backup will be stored locally or remotely. * `Network Share` - The backup is stored remotely on a separate server. * `Local Storage` - The backup is stored locally on the endpoint.
	LocationType *string `json:"LocationType,omitempty"`
	// Password of Backup server.
	Password *string `json:"Password,omitempty"`
	// The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/.
	Path *string `json:"Path,omitempty"`
	// Protocol for transferring the backup image to the network share location. * `SCP` - Secure Copy Protocol (SCP) to access the file server. * `SFTP` - SSH File Transfer Protocol (SFTP) to access file server. * `FTP` - File Transfer Protocol (FTP) to access file server.
	Protocol *string `json:"Protocol,omitempty"`
	// Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner.
	RetentionCount *int64 `json:"RetentionCount,omitempty"`
	// Username for the backup server.
	UserName             *string `json:"UserName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryAbstractBackupConfigAllOf RecoveryAbstractBackupConfigAllOf

// NewRecoveryAbstractBackupConfigAllOf instantiates a new RecoveryAbstractBackupConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryAbstractBackupConfigAllOf(classId string, objectType string) *RecoveryAbstractBackupConfigAllOf {
	this := RecoveryAbstractBackupConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var locationType string = "Network Share"
	this.LocationType = &locationType
	var protocol string = "SCP"
	this.Protocol = &protocol
	var retentionCount int64 = 10
	this.RetentionCount = &retentionCount
	return &this
}

// NewRecoveryAbstractBackupConfigAllOfWithDefaults instantiates a new RecoveryAbstractBackupConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryAbstractBackupConfigAllOfWithDefaults() *RecoveryAbstractBackupConfigAllOf {
	this := RecoveryAbstractBackupConfigAllOf{}
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	var locationType string = "Network Share"
	this.LocationType = &locationType
	var protocol string = "SCP"
	this.Protocol = &protocol
	var retentionCount int64 = 10
	this.RetentionCount = &retentionCount
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryAbstractBackupConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryAbstractBackupConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryAbstractBackupConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryAbstractBackupConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileNamePrefix returns the FileNamePrefix field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfigAllOf) GetFileNamePrefix() string {
	if o == nil || o.FileNamePrefix == nil {
		var ret string
		return ret
	}
	return *o.FileNamePrefix
}

// GetFileNamePrefixOk returns a tuple with the FileNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetFileNamePrefixOk() (*string, bool) {
	if o == nil || o.FileNamePrefix == nil {
		return nil, false
	}
	return o.FileNamePrefix, true
}

// HasFileNamePrefix returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfigAllOf) HasFileNamePrefix() bool {
	if o != nil && o.FileNamePrefix != nil {
		return true
	}

	return false
}

// SetFileNamePrefix gets a reference to the given string and assigns it to the FileNamePrefix field.
func (o *RecoveryAbstractBackupConfigAllOf) SetFileNamePrefix(v string) {
	o.FileNamePrefix = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfigAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfigAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *RecoveryAbstractBackupConfigAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetLocationType returns the LocationType field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfigAllOf) GetLocationType() string {
	if o == nil || o.LocationType == nil {
		var ret string
		return ret
	}
	return *o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetLocationTypeOk() (*string, bool) {
	if o == nil || o.LocationType == nil {
		return nil, false
	}
	return o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfigAllOf) HasLocationType() bool {
	if o != nil && o.LocationType != nil {
		return true
	}

	return false
}

// SetLocationType gets a reference to the given string and assigns it to the LocationType field.
func (o *RecoveryAbstractBackupConfigAllOf) SetLocationType(v string) {
	o.LocationType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfigAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfigAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *RecoveryAbstractBackupConfigAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfigAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfigAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RecoveryAbstractBackupConfigAllOf) SetPath(v string) {
	o.Path = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfigAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfigAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *RecoveryAbstractBackupConfigAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRetentionCount returns the RetentionCount field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfigAllOf) GetRetentionCount() int64 {
	if o == nil || o.RetentionCount == nil {
		var ret int64
		return ret
	}
	return *o.RetentionCount
}

// GetRetentionCountOk returns a tuple with the RetentionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetRetentionCountOk() (*int64, bool) {
	if o == nil || o.RetentionCount == nil {
		return nil, false
	}
	return o.RetentionCount, true
}

// HasRetentionCount returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfigAllOf) HasRetentionCount() bool {
	if o != nil && o.RetentionCount != nil {
		return true
	}

	return false
}

// SetRetentionCount gets a reference to the given int64 and assigns it to the RetentionCount field.
func (o *RecoveryAbstractBackupConfigAllOf) SetRetentionCount(v int64) {
	o.RetentionCount = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfigAllOf) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfigAllOf) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfigAllOf) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *RecoveryAbstractBackupConfigAllOf) SetUserName(v string) {
	o.UserName = &v
}

func (o RecoveryAbstractBackupConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileNamePrefix != nil {
		toSerialize["FileNamePrefix"] = o.FileNamePrefix
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.LocationType != nil {
		toSerialize["LocationType"] = o.LocationType
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.RetentionCount != nil {
		toSerialize["RetentionCount"] = o.RetentionCount
	}
	if o.UserName != nil {
		toSerialize["UserName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryAbstractBackupConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecoveryAbstractBackupConfigAllOf := _RecoveryAbstractBackupConfigAllOf{}

	if err = json.Unmarshal(bytes, &varRecoveryAbstractBackupConfigAllOf); err == nil {
		*o = RecoveryAbstractBackupConfigAllOf(varRecoveryAbstractBackupConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileNamePrefix")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "LocationType")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "RetentionCount")
		delete(additionalProperties, "UserName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecoveryAbstractBackupConfigAllOf struct {
	value *RecoveryAbstractBackupConfigAllOf
	isSet bool
}

func (v NullableRecoveryAbstractBackupConfigAllOf) Get() *RecoveryAbstractBackupConfigAllOf {
	return v.value
}

func (v *NullableRecoveryAbstractBackupConfigAllOf) Set(val *RecoveryAbstractBackupConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryAbstractBackupConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryAbstractBackupConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryAbstractBackupConfigAllOf(val *RecoveryAbstractBackupConfigAllOf) *NullableRecoveryAbstractBackupConfigAllOf {
	return &NullableRecoveryAbstractBackupConfigAllOf{value: val, isSet: true}
}

func (v NullableRecoveryAbstractBackupConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryAbstractBackupConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}