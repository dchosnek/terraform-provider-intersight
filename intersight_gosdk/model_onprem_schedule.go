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

// OnpremSchedule Schedule is used by Intersight Appliance services to store task scheduling information. For example, appliance backup service uses Schedule to store the backup schedule of the Intersight Appliance. The Upgrade service uses Schedule to store the user-defined schedule for software upgrades of the Intersight Appliance.
type OnpremSchedule struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Schedule a task on a specific day of the month. Valid values are 1 through 31. If monthOfYear is specified, then dayOfMonth value must be valid for that month. DayOfMonth may not be set when dayOfWeek is specfied.
	DayOfMonth *int64 `json:"DayOfMonth,omitempty"`
	// Schedule a task on a specific day of the week. Valid values are 1 through 7, with 1 being Sunday. DayOfWeek may not be specfied when dayOfMonth is specified.
	DayOfWeek *int64 `json:"DayOfWeek,omitempty"`
	// Schedule a task on a specific month of the year. Valid values are 1 through 12, with 1 being January.
	MonthOfYear *int64 `json:"MonthOfYear,omitempty"`
	// Schedule a task to run periodically at an interval. Default unit of the RepeatInterval is in minutes. If the RepeateInterval value is set, then all other properties are ignored by the scheduler. RepeateInterval constraints are enforced by the services that use the schedule. Each service has pre-configured service specific properties for enforcing minimum and maximum values of the RepeatInterval.
	RepeatInterval *int64 `json:"RepeatInterval,omitempty"`
	// Time of the day in seconds. TimeOfDay is required for all schedule configurations, except when the RepeateInterval field is specified.
	TimeOfDay *int64 `json:"TimeOfDay,omitempty"`
	// Timezone to use for the schedule calculation. If a timezone value is not speficied, then the schedule calculation will be based on UTC. * `Pacific/Niue` -  * `Pacific/Pago_Pago` -  * `Pacific/Honolulu` -  * `Pacific/Rarotonga` -  * `Pacific/Tahiti` -  * `Pacific/Marquesas` -  * `America/Anchorage` -  * `Pacific/Gambier` -  * `America/Los_Angeles` -  * `America/Tijuana` -  * `America/Vancouver` -  * `America/Whitehorse` -  * `Pacific/Pitcairn` -  * `America/Dawson_Creek` -  * `America/Denver` -  * `America/Edmonton` -  * `America/Hermosillo` -  * `America/Mazatlan` -  * `America/Phoenix` -  * `America/Yellowknife` -  * `America/Belize` -  * `America/Chicago` -  * `America/Costa_Rica` -  * `America/El_Salvador` -  * `America/Guatemala` -  * `America/Managua` -  * `America/Mexico_City` -  * `America/Regina` -  * `America/Tegucigalpa` -  * `America/Winnipeg` -  * `Pacific/Galapagos` -  * `America/Bogota` -  * `America/Cancun` -  * `America/Cayman` -  * `America/Guayaquil` -  * `America/Havana` -  * `America/Iqaluit` -  * `America/Jamaica` -  * `America/Lima` -  * `America/Nassau` -  * `America/New_York` -  * `America/Panama` -  * `America/Port-au-Prince` -  * `America/Rio_Branco` -  * `America/Toronto` -  * `Pacific/Easter` -  * `America/Caracas` -  * `America/Asuncion` -  * `America/Barbados` -  * `America/Boa_Vista` -  * `America/Campo_Grande` -  * `America/Cuiaba` -  * `America/Curacao` -  * `America/Grand_Turk` -  * `America/Guyana` -  * `America/Halifax` -  * `America/La_Paz` -  * `America/Manaus` -  * `America/Martinique` -  * `America/Port_of_Spain` -  * `America/Porto_Velho` -  * `America/Puerto_Rico` -  * `America/Santo_Domingo` -  * `America/Thule` -  * `Atlantic/Bermuda` -  * `America/St_Johns` -  * `America/Araguaina` -  * `America/Argentina/Buenos_Aires` -  * `America/Bahia` -  * `America/Belem` -  * `America/Cayenne` -  * `America/Fortaleza` -  * `America/Godthab` -  * `America/Maceio` -  * `America/Miquelon` -  * `America/Montevideo` -  * `America/Paramaribo` -  * `America/Recife` -  * `America/Santiago` -  * `America/Sao_Paulo` -  * `Antarctica/Palmer` -  * `Antarctica/Rothera` -  * `Atlantic/Stanley` -  * `America/Noronha` -  * `Atlantic/South_Georgia` -  * `America/Scoresbysund` -  * `Atlantic/Azores` -  * `Atlantic/Cape_Verde` -  * `Africa/Abidjan` -  * `Africa/Accra` -  * `Africa/Bissau` -  * `Africa/Casablanca` -  * `Africa/El_Aaiun` -  * `Africa/Monrovia` -  * `America/Danmarkshavn` -  * `Atlantic/Canary` -  * `Atlantic/Faroe` -  * `Atlantic/Reykjavik` -  * `Etc/GMT` -  * `Europe/Dublin` -  * `Europe/Lisbon` -  * `Europe/London` -  * `Africa/Algiers` -  * `Africa/Ceuta` -  * `Africa/Lagos` -  * `Africa/Ndjamena` -  * `Africa/Tunis` -  * `Africa/Windhoek` -  * `Europe/Amsterdam` -  * `Europe/Andorra` -  * `Europe/Belgrade` -  * `Europe/Berlin` -  * `Europe/Brussels` -  * `Europe/Budapest` -  * `Europe/Copenhagen` -  * `Europe/Gibraltar` -  * `Europe/Luxembourg` -  * `Europe/Madrid` -  * `Europe/Malta` -  * `Europe/Monaco` -  * `Europe/Oslo` -  * `Europe/Paris` -  * `Europe/Prague` -  * `Europe/Rome` -  * `Europe/Stockholm` -  * `Europe/Tirane` -  * `Europe/Vienna` -  * `Europe/Warsaw` -  * `Europe/Zurich` -  * `Africa/Cairo` -  * `Africa/Johannesburg` -  * `Africa/Maputo` -  * `Africa/Tripoli` -  * `Asia/Amman` -  * `Asia/Beirut` -  * `Asia/Damascus` -  * `Asia/Gaza` -  * `Asia/Jerusalem` -  * `Asia/Nicosia` -  * `Europe/Athens` -  * `Europe/Bucharest` -  * `Europe/Chisinau` -  * `Europe/Helsinki` -  * `Europe/Istanbul` -  * `Europe/Kaliningrad` -  * `Europe/Kiev` -  * `Europe/Riga` -  * `Europe/Sofia` -  * `Europe/Tallinn` -  * `Europe/Vilnius` -  * `Africa/Khartoum` -  * `Africa/Nairobi` -  * `Antarctica/Syowa` -  * `Asia/Baghdad` -  * `Asia/Qatar` -  * `Asia/Riyadh` -  * `Europe/Minsk` -  * `Europe/Moscow` -  * `Asia/Tehran` -  * `Asia/Baku` -  * `Asia/Dubai` -  * `Asia/Tbilisi` -  * `Asia/Yerevan` -  * `Europe/Samara` -  * `Indian/Mahe` -  * `Indian/Mauritius` -  * `Indian/Reunion` -  * `Asia/Kabul` -  * `Antarctica/Mawson` -  * `Asia/Aqtau` -  * `Asia/Aqtobe` -  * `Asia/Ashgabat` -  * `Asia/Dushanbe` -  * `Asia/Karachi` -  * `Asia/Tashkent` -  * `Asia/Yekaterinburg` -  * `Indian/Kerguelen` -  * `Indian/Maldives` -  * `Asia/Calcutta` -  * `Asia/Kolkata` -  * `Asia/Colombo` -  * `Asia/Katmandu` -  * `Antarctica/Vostok` -  * `Asia/Almaty` -  * `Asia/Bishkek` -  * `Asia/Dhaka` -  * `Asia/Omsk` -  * `Asia/Thimphu` -  * `Indian/Chagos` -  * `Asia/Rangoon` -  * `Indian/Cocos` -  * `Antarctica/Davis` -  * `Asia/Bangkok` -  * `Asia/Hovd` -  * `Asia/Jakarta` -  * `Asia/Krasnoyarsk` -  * `Asia/Saigon` -  * `Indian/Christmas` -  * `Antarctica/Casey` -  * `Asia/Brunei` -  * `Asia/Choibalsan` -  * `Asia/Hong_Kong` -  * `Asia/Irkutsk` -  * `Asia/Kuala_Lumpur` -  * `Asia/Macau` -  * `Asia/Makassar` -  * `Asia/Manila` -  * `Asia/Shanghai` -  * `Asia/Singapore` -  * `Asia/Taipei` -  * `Asia/Ulaanbaatar` -  * `Australia/Perth` -  * `Asia/Pyongyang` -  * `Asia/Dili` -  * `Asia/Jayapura` -  * `Asia/Seoul` -  * `Asia/Tokyo` -  * `Asia/Yakutsk` -  * `Pacific/Palau` -  * `Australia/Adelaide` -  * `Australia/Darwin` -  * `Antarctica/DumontDUrville` -  * `Asia/Magadan` -  * `Asia/Vladivostok` -  * `Australia/Brisbane` -  * `Australia/Hobart` -  * `Australia/Sydney` -  * `Pacific/Chuuk` -  * `Pacific/Guam` -  * `Pacific/Port_Moresby` -  * `Pacific/Efate` -  * `Pacific/Guadalcanal` -  * `Pacific/Kosrae` -  * `Pacific/Norfolk` -  * `Pacific/Noumea` -  * `Pacific/Pohnpei` -  * `Asia/Kamchatka` -  * `Pacific/Auckland` -  * `Pacific/Fiji` -  * `Pacific/Funafuti` -  * `Pacific/Kwajalein` -  * `Pacific/Majuro` -  * `Pacific/Nauru` -  * `Pacific/Tarawa` -  * `Pacific/Wake` -  * `Pacific/Wallis` -  * `Pacific/Apia` -  * `Pacific/Enderbury` -  * `Pacific/Fakaofo` -  * `Pacific/Tongatapu` -  * `Pacific/Kiritimati` -
	TimeZone *string `json:"TimeZone,omitempty"`
	// Schedule a task on a specific week of the month. Valid values are 1 through 5. Value of 5 means last week of the month. WeekOfMonth may not be set when dayOfMonth is specified.
	WeekOfMonth          *int64 `json:"WeekOfMonth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnpremSchedule OnpremSchedule

// NewOnpremSchedule instantiates a new OnpremSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnpremSchedule(classId string, objectType string) *OnpremSchedule {
	this := OnpremSchedule{}
	this.ClassId = classId
	this.ObjectType = objectType
	var timeZone string = "Pacific/Niue"
	this.TimeZone = &timeZone
	return &this
}

// NewOnpremScheduleWithDefaults instantiates a new OnpremSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnpremScheduleWithDefaults() *OnpremSchedule {
	this := OnpremSchedule{}
	var classId string = "onprem.Schedule"
	this.ClassId = classId
	var objectType string = "onprem.Schedule"
	this.ObjectType = objectType
	var timeZone string = "Pacific/Niue"
	this.TimeZone = &timeZone
	return &this
}

// GetClassId returns the ClassId field value
func (o *OnpremSchedule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OnpremSchedule) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OnpremSchedule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OnpremSchedule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *OnpremSchedule) GetDayOfMonth() int64 {
	if o == nil || o.DayOfMonth == nil {
		var ret int64
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetDayOfMonthOk() (*int64, bool) {
	if o == nil || o.DayOfMonth == nil {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *OnpremSchedule) HasDayOfMonth() bool {
	if o != nil && o.DayOfMonth != nil {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given int64 and assigns it to the DayOfMonth field.
func (o *OnpremSchedule) SetDayOfMonth(v int64) {
	o.DayOfMonth = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *OnpremSchedule) GetDayOfWeek() int64 {
	if o == nil || o.DayOfWeek == nil {
		var ret int64
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetDayOfWeekOk() (*int64, bool) {
	if o == nil || o.DayOfWeek == nil {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *OnpremSchedule) HasDayOfWeek() bool {
	if o != nil && o.DayOfWeek != nil {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given int64 and assigns it to the DayOfWeek field.
func (o *OnpremSchedule) SetDayOfWeek(v int64) {
	o.DayOfWeek = &v
}

// GetMonthOfYear returns the MonthOfYear field value if set, zero value otherwise.
func (o *OnpremSchedule) GetMonthOfYear() int64 {
	if o == nil || o.MonthOfYear == nil {
		var ret int64
		return ret
	}
	return *o.MonthOfYear
}

// GetMonthOfYearOk returns a tuple with the MonthOfYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetMonthOfYearOk() (*int64, bool) {
	if o == nil || o.MonthOfYear == nil {
		return nil, false
	}
	return o.MonthOfYear, true
}

// HasMonthOfYear returns a boolean if a field has been set.
func (o *OnpremSchedule) HasMonthOfYear() bool {
	if o != nil && o.MonthOfYear != nil {
		return true
	}

	return false
}

// SetMonthOfYear gets a reference to the given int64 and assigns it to the MonthOfYear field.
func (o *OnpremSchedule) SetMonthOfYear(v int64) {
	o.MonthOfYear = &v
}

// GetRepeatInterval returns the RepeatInterval field value if set, zero value otherwise.
func (o *OnpremSchedule) GetRepeatInterval() int64 {
	if o == nil || o.RepeatInterval == nil {
		var ret int64
		return ret
	}
	return *o.RepeatInterval
}

// GetRepeatIntervalOk returns a tuple with the RepeatInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetRepeatIntervalOk() (*int64, bool) {
	if o == nil || o.RepeatInterval == nil {
		return nil, false
	}
	return o.RepeatInterval, true
}

// HasRepeatInterval returns a boolean if a field has been set.
func (o *OnpremSchedule) HasRepeatInterval() bool {
	if o != nil && o.RepeatInterval != nil {
		return true
	}

	return false
}

// SetRepeatInterval gets a reference to the given int64 and assigns it to the RepeatInterval field.
func (o *OnpremSchedule) SetRepeatInterval(v int64) {
	o.RepeatInterval = &v
}

// GetTimeOfDay returns the TimeOfDay field value if set, zero value otherwise.
func (o *OnpremSchedule) GetTimeOfDay() int64 {
	if o == nil || o.TimeOfDay == nil {
		var ret int64
		return ret
	}
	return *o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetTimeOfDayOk() (*int64, bool) {
	if o == nil || o.TimeOfDay == nil {
		return nil, false
	}
	return o.TimeOfDay, true
}

// HasTimeOfDay returns a boolean if a field has been set.
func (o *OnpremSchedule) HasTimeOfDay() bool {
	if o != nil && o.TimeOfDay != nil {
		return true
	}

	return false
}

// SetTimeOfDay gets a reference to the given int64 and assigns it to the TimeOfDay field.
func (o *OnpremSchedule) SetTimeOfDay(v int64) {
	o.TimeOfDay = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *OnpremSchedule) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *OnpremSchedule) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *OnpremSchedule) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetWeekOfMonth returns the WeekOfMonth field value if set, zero value otherwise.
func (o *OnpremSchedule) GetWeekOfMonth() int64 {
	if o == nil || o.WeekOfMonth == nil {
		var ret int64
		return ret
	}
	return *o.WeekOfMonth
}

// GetWeekOfMonthOk returns a tuple with the WeekOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremSchedule) GetWeekOfMonthOk() (*int64, bool) {
	if o == nil || o.WeekOfMonth == nil {
		return nil, false
	}
	return o.WeekOfMonth, true
}

// HasWeekOfMonth returns a boolean if a field has been set.
func (o *OnpremSchedule) HasWeekOfMonth() bool {
	if o != nil && o.WeekOfMonth != nil {
		return true
	}

	return false
}

// SetWeekOfMonth gets a reference to the given int64 and assigns it to the WeekOfMonth field.
func (o *OnpremSchedule) SetWeekOfMonth(v int64) {
	o.WeekOfMonth = &v
}

func (o OnpremSchedule) MarshalJSON() ([]byte, error) {
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
	if o.DayOfMonth != nil {
		toSerialize["DayOfMonth"] = o.DayOfMonth
	}
	if o.DayOfWeek != nil {
		toSerialize["DayOfWeek"] = o.DayOfWeek
	}
	if o.MonthOfYear != nil {
		toSerialize["MonthOfYear"] = o.MonthOfYear
	}
	if o.RepeatInterval != nil {
		toSerialize["RepeatInterval"] = o.RepeatInterval
	}
	if o.TimeOfDay != nil {
		toSerialize["TimeOfDay"] = o.TimeOfDay
	}
	if o.TimeZone != nil {
		toSerialize["TimeZone"] = o.TimeZone
	}
	if o.WeekOfMonth != nil {
		toSerialize["WeekOfMonth"] = o.WeekOfMonth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OnpremSchedule) UnmarshalJSON(bytes []byte) (err error) {
	type OnpremScheduleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Schedule a task on a specific day of the month. Valid values are 1 through 31. If monthOfYear is specified, then dayOfMonth value must be valid for that month. DayOfMonth may not be set when dayOfWeek is specfied.
		DayOfMonth *int64 `json:"DayOfMonth,omitempty"`
		// Schedule a task on a specific day of the week. Valid values are 1 through 7, with 1 being Sunday. DayOfWeek may not be specfied when dayOfMonth is specified.
		DayOfWeek *int64 `json:"DayOfWeek,omitempty"`
		// Schedule a task on a specific month of the year. Valid values are 1 through 12, with 1 being January.
		MonthOfYear *int64 `json:"MonthOfYear,omitempty"`
		// Schedule a task to run periodically at an interval. Default unit of the RepeatInterval is in minutes. If the RepeateInterval value is set, then all other properties are ignored by the scheduler. RepeateInterval constraints are enforced by the services that use the schedule. Each service has pre-configured service specific properties for enforcing minimum and maximum values of the RepeatInterval.
		RepeatInterval *int64 `json:"RepeatInterval,omitempty"`
		// Time of the day in seconds. TimeOfDay is required for all schedule configurations, except when the RepeateInterval field is specified.
		TimeOfDay *int64 `json:"TimeOfDay,omitempty"`
		// Timezone to use for the schedule calculation. If a timezone value is not speficied, then the schedule calculation will be based on UTC. * `Pacific/Niue` -  * `Pacific/Pago_Pago` -  * `Pacific/Honolulu` -  * `Pacific/Rarotonga` -  * `Pacific/Tahiti` -  * `Pacific/Marquesas` -  * `America/Anchorage` -  * `Pacific/Gambier` -  * `America/Los_Angeles` -  * `America/Tijuana` -  * `America/Vancouver` -  * `America/Whitehorse` -  * `Pacific/Pitcairn` -  * `America/Dawson_Creek` -  * `America/Denver` -  * `America/Edmonton` -  * `America/Hermosillo` -  * `America/Mazatlan` -  * `America/Phoenix` -  * `America/Yellowknife` -  * `America/Belize` -  * `America/Chicago` -  * `America/Costa_Rica` -  * `America/El_Salvador` -  * `America/Guatemala` -  * `America/Managua` -  * `America/Mexico_City` -  * `America/Regina` -  * `America/Tegucigalpa` -  * `America/Winnipeg` -  * `Pacific/Galapagos` -  * `America/Bogota` -  * `America/Cancun` -  * `America/Cayman` -  * `America/Guayaquil` -  * `America/Havana` -  * `America/Iqaluit` -  * `America/Jamaica` -  * `America/Lima` -  * `America/Nassau` -  * `America/New_York` -  * `America/Panama` -  * `America/Port-au-Prince` -  * `America/Rio_Branco` -  * `America/Toronto` -  * `Pacific/Easter` -  * `America/Caracas` -  * `America/Asuncion` -  * `America/Barbados` -  * `America/Boa_Vista` -  * `America/Campo_Grande` -  * `America/Cuiaba` -  * `America/Curacao` -  * `America/Grand_Turk` -  * `America/Guyana` -  * `America/Halifax` -  * `America/La_Paz` -  * `America/Manaus` -  * `America/Martinique` -  * `America/Port_of_Spain` -  * `America/Porto_Velho` -  * `America/Puerto_Rico` -  * `America/Santo_Domingo` -  * `America/Thule` -  * `Atlantic/Bermuda` -  * `America/St_Johns` -  * `America/Araguaina` -  * `America/Argentina/Buenos_Aires` -  * `America/Bahia` -  * `America/Belem` -  * `America/Cayenne` -  * `America/Fortaleza` -  * `America/Godthab` -  * `America/Maceio` -  * `America/Miquelon` -  * `America/Montevideo` -  * `America/Paramaribo` -  * `America/Recife` -  * `America/Santiago` -  * `America/Sao_Paulo` -  * `Antarctica/Palmer` -  * `Antarctica/Rothera` -  * `Atlantic/Stanley` -  * `America/Noronha` -  * `Atlantic/South_Georgia` -  * `America/Scoresbysund` -  * `Atlantic/Azores` -  * `Atlantic/Cape_Verde` -  * `Africa/Abidjan` -  * `Africa/Accra` -  * `Africa/Bissau` -  * `Africa/Casablanca` -  * `Africa/El_Aaiun` -  * `Africa/Monrovia` -  * `America/Danmarkshavn` -  * `Atlantic/Canary` -  * `Atlantic/Faroe` -  * `Atlantic/Reykjavik` -  * `Etc/GMT` -  * `Europe/Dublin` -  * `Europe/Lisbon` -  * `Europe/London` -  * `Africa/Algiers` -  * `Africa/Ceuta` -  * `Africa/Lagos` -  * `Africa/Ndjamena` -  * `Africa/Tunis` -  * `Africa/Windhoek` -  * `Europe/Amsterdam` -  * `Europe/Andorra` -  * `Europe/Belgrade` -  * `Europe/Berlin` -  * `Europe/Brussels` -  * `Europe/Budapest` -  * `Europe/Copenhagen` -  * `Europe/Gibraltar` -  * `Europe/Luxembourg` -  * `Europe/Madrid` -  * `Europe/Malta` -  * `Europe/Monaco` -  * `Europe/Oslo` -  * `Europe/Paris` -  * `Europe/Prague` -  * `Europe/Rome` -  * `Europe/Stockholm` -  * `Europe/Tirane` -  * `Europe/Vienna` -  * `Europe/Warsaw` -  * `Europe/Zurich` -  * `Africa/Cairo` -  * `Africa/Johannesburg` -  * `Africa/Maputo` -  * `Africa/Tripoli` -  * `Asia/Amman` -  * `Asia/Beirut` -  * `Asia/Damascus` -  * `Asia/Gaza` -  * `Asia/Jerusalem` -  * `Asia/Nicosia` -  * `Europe/Athens` -  * `Europe/Bucharest` -  * `Europe/Chisinau` -  * `Europe/Helsinki` -  * `Europe/Istanbul` -  * `Europe/Kaliningrad` -  * `Europe/Kiev` -  * `Europe/Riga` -  * `Europe/Sofia` -  * `Europe/Tallinn` -  * `Europe/Vilnius` -  * `Africa/Khartoum` -  * `Africa/Nairobi` -  * `Antarctica/Syowa` -  * `Asia/Baghdad` -  * `Asia/Qatar` -  * `Asia/Riyadh` -  * `Europe/Minsk` -  * `Europe/Moscow` -  * `Asia/Tehran` -  * `Asia/Baku` -  * `Asia/Dubai` -  * `Asia/Tbilisi` -  * `Asia/Yerevan` -  * `Europe/Samara` -  * `Indian/Mahe` -  * `Indian/Mauritius` -  * `Indian/Reunion` -  * `Asia/Kabul` -  * `Antarctica/Mawson` -  * `Asia/Aqtau` -  * `Asia/Aqtobe` -  * `Asia/Ashgabat` -  * `Asia/Dushanbe` -  * `Asia/Karachi` -  * `Asia/Tashkent` -  * `Asia/Yekaterinburg` -  * `Indian/Kerguelen` -  * `Indian/Maldives` -  * `Asia/Calcutta` -  * `Asia/Kolkata` -  * `Asia/Colombo` -  * `Asia/Katmandu` -  * `Antarctica/Vostok` -  * `Asia/Almaty` -  * `Asia/Bishkek` -  * `Asia/Dhaka` -  * `Asia/Omsk` -  * `Asia/Thimphu` -  * `Indian/Chagos` -  * `Asia/Rangoon` -  * `Indian/Cocos` -  * `Antarctica/Davis` -  * `Asia/Bangkok` -  * `Asia/Hovd` -  * `Asia/Jakarta` -  * `Asia/Krasnoyarsk` -  * `Asia/Saigon` -  * `Indian/Christmas` -  * `Antarctica/Casey` -  * `Asia/Brunei` -  * `Asia/Choibalsan` -  * `Asia/Hong_Kong` -  * `Asia/Irkutsk` -  * `Asia/Kuala_Lumpur` -  * `Asia/Macau` -  * `Asia/Makassar` -  * `Asia/Manila` -  * `Asia/Shanghai` -  * `Asia/Singapore` -  * `Asia/Taipei` -  * `Asia/Ulaanbaatar` -  * `Australia/Perth` -  * `Asia/Pyongyang` -  * `Asia/Dili` -  * `Asia/Jayapura` -  * `Asia/Seoul` -  * `Asia/Tokyo` -  * `Asia/Yakutsk` -  * `Pacific/Palau` -  * `Australia/Adelaide` -  * `Australia/Darwin` -  * `Antarctica/DumontDUrville` -  * `Asia/Magadan` -  * `Asia/Vladivostok` -  * `Australia/Brisbane` -  * `Australia/Hobart` -  * `Australia/Sydney` -  * `Pacific/Chuuk` -  * `Pacific/Guam` -  * `Pacific/Port_Moresby` -  * `Pacific/Efate` -  * `Pacific/Guadalcanal` -  * `Pacific/Kosrae` -  * `Pacific/Norfolk` -  * `Pacific/Noumea` -  * `Pacific/Pohnpei` -  * `Asia/Kamchatka` -  * `Pacific/Auckland` -  * `Pacific/Fiji` -  * `Pacific/Funafuti` -  * `Pacific/Kwajalein` -  * `Pacific/Majuro` -  * `Pacific/Nauru` -  * `Pacific/Tarawa` -  * `Pacific/Wake` -  * `Pacific/Wallis` -  * `Pacific/Apia` -  * `Pacific/Enderbury` -  * `Pacific/Fakaofo` -  * `Pacific/Tongatapu` -  * `Pacific/Kiritimati` -
		TimeZone *string `json:"TimeZone,omitempty"`
		// Schedule a task on a specific week of the month. Valid values are 1 through 5. Value of 5 means last week of the month. WeekOfMonth may not be set when dayOfMonth is specified.
		WeekOfMonth *int64 `json:"WeekOfMonth,omitempty"`
	}

	varOnpremScheduleWithoutEmbeddedStruct := OnpremScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOnpremScheduleWithoutEmbeddedStruct)
	if err == nil {
		varOnpremSchedule := _OnpremSchedule{}
		varOnpremSchedule.ClassId = varOnpremScheduleWithoutEmbeddedStruct.ClassId
		varOnpremSchedule.ObjectType = varOnpremScheduleWithoutEmbeddedStruct.ObjectType
		varOnpremSchedule.DayOfMonth = varOnpremScheduleWithoutEmbeddedStruct.DayOfMonth
		varOnpremSchedule.DayOfWeek = varOnpremScheduleWithoutEmbeddedStruct.DayOfWeek
		varOnpremSchedule.MonthOfYear = varOnpremScheduleWithoutEmbeddedStruct.MonthOfYear
		varOnpremSchedule.RepeatInterval = varOnpremScheduleWithoutEmbeddedStruct.RepeatInterval
		varOnpremSchedule.TimeOfDay = varOnpremScheduleWithoutEmbeddedStruct.TimeOfDay
		varOnpremSchedule.TimeZone = varOnpremScheduleWithoutEmbeddedStruct.TimeZone
		varOnpremSchedule.WeekOfMonth = varOnpremScheduleWithoutEmbeddedStruct.WeekOfMonth
		*o = OnpremSchedule(varOnpremSchedule)
	} else {
		return err
	}

	varOnpremSchedule := _OnpremSchedule{}

	err = json.Unmarshal(bytes, &varOnpremSchedule)
	if err == nil {
		o.MoBaseComplexType = varOnpremSchedule.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DayOfMonth")
		delete(additionalProperties, "DayOfWeek")
		delete(additionalProperties, "MonthOfYear")
		delete(additionalProperties, "RepeatInterval")
		delete(additionalProperties, "TimeOfDay")
		delete(additionalProperties, "TimeZone")
		delete(additionalProperties, "WeekOfMonth")

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

type NullableOnpremSchedule struct {
	value *OnpremSchedule
	isSet bool
}

func (v NullableOnpremSchedule) Get() *OnpremSchedule {
	return v.value
}

func (v *NullableOnpremSchedule) Set(val *OnpremSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableOnpremSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableOnpremSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnpremSchedule(val *OnpremSchedule) *NullableOnpremSchedule {
	return &NullableOnpremSchedule{value: val, isSet: true}
}

func (v NullableOnpremSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnpremSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
