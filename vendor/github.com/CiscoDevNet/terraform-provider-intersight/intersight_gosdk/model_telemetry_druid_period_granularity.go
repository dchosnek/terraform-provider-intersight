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

// TelemetryDruidPeriodGranularity Period granularities are specified as arbitrary period combinations of years, months, weeks, hours, minutes and seconds (e.g. P2W, P3M, PT1H30M, PT0.750S) in ISO8601 format. They support specifying a time zone which determines where period boundaries start as well as the timezone of the returned timestamps. By default, years start on the first of January, months start on the first of the month and weeks start on Mondays unless an origin is specified. Time zone is optional (defaults to UTC). Origin is optional (defaults to 1970-01-01T00:00:00 in the given time zone).
type TelemetryDruidPeriodGranularity struct {
	// the type of granularity.
	Type string `json:"type"`
	// The period in ISO 8601 format. Examples are P2W, P3M, PT1H30M, PT0.750S.
	Period string `json:"period"`
	// An optional value specifying the time zone. Standard [IANA time zones](http://joda-time.sourceforge.net/timezones.html) are supported.
	TimeZone             *string `json:"timeZone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidPeriodGranularity TelemetryDruidPeriodGranularity

// NewTelemetryDruidPeriodGranularity instantiates a new TelemetryDruidPeriodGranularity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidPeriodGranularity(type_ string, period string) *TelemetryDruidPeriodGranularity {
	this := TelemetryDruidPeriodGranularity{}
	this.Type = type_
	this.Period = period
	return &this
}

// NewTelemetryDruidPeriodGranularityWithDefaults instantiates a new TelemetryDruidPeriodGranularity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidPeriodGranularityWithDefaults() *TelemetryDruidPeriodGranularity {
	this := TelemetryDruidPeriodGranularity{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidPeriodGranularity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidPeriodGranularity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidPeriodGranularity) SetType(v string) {
	o.Type = v
}

// GetPeriod returns the Period field value
func (o *TelemetryDruidPeriodGranularity) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidPeriodGranularity) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *TelemetryDruidPeriodGranularity) SetPeriod(v string) {
	o.Period = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *TelemetryDruidPeriodGranularity) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidPeriodGranularity) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *TelemetryDruidPeriodGranularity) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *TelemetryDruidPeriodGranularity) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o TelemetryDruidPeriodGranularity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["period"] = o.Period
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidPeriodGranularity) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidPeriodGranularity := _TelemetryDruidPeriodGranularity{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidPeriodGranularity); err == nil {
		*o = TelemetryDruidPeriodGranularity(varTelemetryDruidPeriodGranularity)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "period")
		delete(additionalProperties, "timeZone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidPeriodGranularity struct {
	value *TelemetryDruidPeriodGranularity
	isSet bool
}

func (v NullableTelemetryDruidPeriodGranularity) Get() *TelemetryDruidPeriodGranularity {
	return v.value
}

func (v *NullableTelemetryDruidPeriodGranularity) Set(val *TelemetryDruidPeriodGranularity) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidPeriodGranularity) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidPeriodGranularity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidPeriodGranularity(val *TelemetryDruidPeriodGranularity) *NullableTelemetryDruidPeriodGranularity {
	return &NullableTelemetryDruidPeriodGranularity{value: val, isSet: true}
}

func (v NullableTelemetryDruidPeriodGranularity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidPeriodGranularity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}