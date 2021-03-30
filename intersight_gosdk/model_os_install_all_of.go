/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// OsInstallAllOf Definition of the list of properties defined in 'os.Install', excluding properties defined in parent classes.
type OsInstallAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the OS install configuration.
	Name                 *string                                                      `json:"Name,omitempty"`
	ConfigurationFile    *OsConfigurationFileRelationship                             `json:"ConfigurationFile,omitempty"`
	Image                *SoftwarerepositoryOperatingSystemFileRelationship           `json:"Image,omitempty"`
	Organization         *OrganizationOrganizationRelationship                        `json:"Organization,omitempty"`
	OsduImage            *FirmwareServerConfigurationUtilityDistributableRelationship `json:"OsduImage,omitempty"`
	Server               *ComputePhysicalRelationship                                 `json:"Server,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship                            `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsInstallAllOf OsInstallAllOf

// NewOsInstallAllOf instantiates a new OsInstallAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsInstallAllOf(classId string, objectType string) *OsInstallAllOf {
	this := OsInstallAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsInstallAllOfWithDefaults instantiates a new OsInstallAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsInstallAllOfWithDefaults() *OsInstallAllOf {
	this := OsInstallAllOf{}
	var classId string = "os.Install"
	this.ClassId = classId
	var objectType string = "os.Install"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsInstallAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsInstallAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsInstallAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsInstallAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsInstallAllOf) SetName(v string) {
	o.Name = &v
}

// GetConfigurationFile returns the ConfigurationFile field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetConfigurationFile() OsConfigurationFileRelationship {
	if o == nil || o.ConfigurationFile == nil {
		var ret OsConfigurationFileRelationship
		return ret
	}
	return *o.ConfigurationFile
}

// GetConfigurationFileOk returns a tuple with the ConfigurationFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetConfigurationFileOk() (*OsConfigurationFileRelationship, bool) {
	if o == nil || o.ConfigurationFile == nil {
		return nil, false
	}
	return o.ConfigurationFile, true
}

// HasConfigurationFile returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasConfigurationFile() bool {
	if o != nil && o.ConfigurationFile != nil {
		return true
	}

	return false
}

// SetConfigurationFile gets a reference to the given OsConfigurationFileRelationship and assigns it to the ConfigurationFile field.
func (o *OsInstallAllOf) SetConfigurationFile(v OsConfigurationFileRelationship) {
	o.ConfigurationFile = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetImage() SoftwarerepositoryOperatingSystemFileRelationship {
	if o == nil || o.Image == nil {
		var ret SoftwarerepositoryOperatingSystemFileRelationship
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetImageOk() (*SoftwarerepositoryOperatingSystemFileRelationship, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given SoftwarerepositoryOperatingSystemFileRelationship and assigns it to the Image field.
func (o *OsInstallAllOf) SetImage(v SoftwarerepositoryOperatingSystemFileRelationship) {
	o.Image = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OsInstallAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetOsduImage returns the OsduImage field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetOsduImage() FirmwareServerConfigurationUtilityDistributableRelationship {
	if o == nil || o.OsduImage == nil {
		var ret FirmwareServerConfigurationUtilityDistributableRelationship
		return ret
	}
	return *o.OsduImage
}

// GetOsduImageOk returns a tuple with the OsduImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetOsduImageOk() (*FirmwareServerConfigurationUtilityDistributableRelationship, bool) {
	if o == nil || o.OsduImage == nil {
		return nil, false
	}
	return o.OsduImage, true
}

// HasOsduImage returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasOsduImage() bool {
	if o != nil && o.OsduImage != nil {
		return true
	}

	return false
}

// SetOsduImage gets a reference to the given FirmwareServerConfigurationUtilityDistributableRelationship and assigns it to the OsduImage field.
func (o *OsInstallAllOf) SetOsduImage(v FirmwareServerConfigurationUtilityDistributableRelationship) {
	o.OsduImage = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *OsInstallAllOf) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *OsInstallAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstallAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *OsInstallAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *OsInstallAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o OsInstallAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ConfigurationFile != nil {
		toSerialize["ConfigurationFile"] = o.ConfigurationFile
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.OsduImage != nil {
		toSerialize["OsduImage"] = o.OsduImage
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsInstallAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsInstallAllOf := _OsInstallAllOf{}

	if err = json.Unmarshal(bytes, &varOsInstallAllOf); err == nil {
		*o = OsInstallAllOf(varOsInstallAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ConfigurationFile")
		delete(additionalProperties, "Image")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OsduImage")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsInstallAllOf struct {
	value *OsInstallAllOf
	isSet bool
}

func (v NullableOsInstallAllOf) Get() *OsInstallAllOf {
	return v.value
}

func (v *NullableOsInstallAllOf) Set(val *OsInstallAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsInstallAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsInstallAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsInstallAllOf(val *OsInstallAllOf) *NullableOsInstallAllOf {
	return &NullableOsInstallAllOf{value: val, isSet: true}
}

func (v NullableOsInstallAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsInstallAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
