/*
API Documentation

Source of truth and network automation platform

API version: 1.3.10b1 (1.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// PowerOutletTypeChoices the model 'PowerOutletTypeChoices'
type PowerOutletTypeChoices string

// List of PowerOutletTypeChoices
const (
	IEC_60320_C5  PowerOutletTypeChoices = "iec-60320-c5"
	IEC_60320_C7  PowerOutletTypeChoices = "iec-60320-c7"
	IEC_60320_C13 PowerOutletTypeChoices = "iec-60320-c13"
	IEC_60320_C15 PowerOutletTypeChoices = "iec-60320-c15"
	IEC_60320_C19 PowerOutletTypeChoices = "iec-60320-c19"
	//IEC_60309_P_N_E_4H PowerOutletTypeChoices = "iec-60309-p-n-e-4h"
	//IEC_60309_P_N_E_6H PowerOutletTypeChoices = "iec-60309-p-n-e-6h"
	//IEC_60309_P_N_E_9H PowerOutletTypeChoices = "iec-60309-p-n-e-9h"
	//IEC_60309_2P_E_4H PowerOutletTypeChoices = "iec-60309-2p-e-4h"
	//IEC_60309_2P_E_6H PowerOutletTypeChoices = "iec-60309-2p-e-6h"
	//IEC_60309_2P_E_9H PowerOutletTypeChoices = "iec-60309-2p-e-9h"
	//IEC_60309_3P_E_4H PowerOutletTypeChoices = "iec-60309-3p-e-4h"
	//IEC_60309_3P_E_6H PowerOutletTypeChoices = "iec-60309-3p-e-6h"
	//IEC_60309_3P_E_9H PowerOutletTypeChoices = "iec-60309-3p-e-9h"
	//IEC_60309_3P_N_E_4H PowerOutletTypeChoices = "iec-60309-3p-n-e-4h"
	//IEC_60309_3P_N_E_6H PowerOutletTypeChoices = "iec-60309-3p-n-e-6h"
	//IEC_60309_3P_N_E_9H PowerOutletTypeChoices = "iec-60309-3p-n-e-9h"
	NEMA_1_15R   PowerOutletTypeChoices = "nema-1-15r"
	NEMA_5_15R   PowerOutletTypeChoices = "nema-5-15r"
	NEMA_5_20R   PowerOutletTypeChoices = "nema-5-20r"
	NEMA_5_30R   PowerOutletTypeChoices = "nema-5-30r"
	NEMA_5_50R   PowerOutletTypeChoices = "nema-5-50r"
	NEMA_6_15R   PowerOutletTypeChoices = "nema-6-15r"
	NEMA_6_20R   PowerOutletTypeChoices = "nema-6-20r"
	NEMA_6_30R   PowerOutletTypeChoices = "nema-6-30r"
	NEMA_6_50R   PowerOutletTypeChoices = "nema-6-50r"
	NEMA_10_30R  PowerOutletTypeChoices = "nema-10-30r"
	NEMA_10_50R  PowerOutletTypeChoices = "nema-10-50r"
	NEMA_14_20R  PowerOutletTypeChoices = "nema-14-20r"
	NEMA_14_30R  PowerOutletTypeChoices = "nema-14-30r"
	NEMA_14_50R  PowerOutletTypeChoices = "nema-14-50r"
	NEMA_14_60R  PowerOutletTypeChoices = "nema-14-60r"
	NEMA_15_15R  PowerOutletTypeChoices = "nema-15-15r"
	NEMA_15_20R  PowerOutletTypeChoices = "nema-15-20r"
	NEMA_15_30R  PowerOutletTypeChoices = "nema-15-30r"
	NEMA_15_50R  PowerOutletTypeChoices = "nema-15-50r"
	NEMA_15_60R  PowerOutletTypeChoices = "nema-15-60r"
	NEMA_L1_15R  PowerOutletTypeChoices = "nema-l1-15r"
	NEMA_L5_15R  PowerOutletTypeChoices = "nema-l5-15r"
	NEMA_L5_20R  PowerOutletTypeChoices = "nema-l5-20r"
	NEMA_L5_30R  PowerOutletTypeChoices = "nema-l5-30r"
	NEMA_L5_50R  PowerOutletTypeChoices = "nema-l5-50r"
	NEMA_L6_15R  PowerOutletTypeChoices = "nema-l6-15r"
	NEMA_L6_20R  PowerOutletTypeChoices = "nema-l6-20r"
	NEMA_L6_30R  PowerOutletTypeChoices = "nema-l6-30r"
	NEMA_L6_50R  PowerOutletTypeChoices = "nema-l6-50r"
	NEMA_L10_30R PowerOutletTypeChoices = "nema-l10-30r"
	NEMA_L14_20R PowerOutletTypeChoices = "nema-l14-20r"
	NEMA_L14_30R PowerOutletTypeChoices = "nema-l14-30r"
	NEMA_L14_50R PowerOutletTypeChoices = "nema-l14-50r"
	NEMA_L14_60R PowerOutletTypeChoices = "nema-l14-60r"
	NEMA_L15_20R PowerOutletTypeChoices = "nema-l15-20r"
	NEMA_L15_30R PowerOutletTypeChoices = "nema-l15-30r"
	NEMA_L15_50R PowerOutletTypeChoices = "nema-l15-50r"
	NEMA_L15_60R PowerOutletTypeChoices = "nema-l15-60r"
	NEMA_L21_20R PowerOutletTypeChoices = "nema-l21-20r"
	NEMA_L21_30R PowerOutletTypeChoices = "nema-l21-30r"
	CS6360_C     PowerOutletTypeChoices = "CS6360C"
	CS6364_C     PowerOutletTypeChoices = "CS6364C"
	CS8164_C     PowerOutletTypeChoices = "CS8164C"
	CS8264_C     PowerOutletTypeChoices = "CS8264C"
	CS8364_C     PowerOutletTypeChoices = "CS8364C"
	CS8464_C     PowerOutletTypeChoices = "CS8464C"
	//ITA_E PowerOutletTypeChoices = "ita-e"
	//ITA_F PowerOutletTypeChoices = "ita-f"
	//ITA_G PowerOutletTypeChoices = "ita-g"
	//ITA_H PowerOutletTypeChoices = "ita-h"
	//ITA_I PowerOutletTypeChoices = "ita-i"
	//ITA_J PowerOutletTypeChoices = "ita-j"
	//ITA_K PowerOutletTypeChoices = "ita-k"
	//ITA_L PowerOutletTypeChoices = "ita-l"
	//ITA_M PowerOutletTypeChoices = "ita-m"
	//ITA_N PowerOutletTypeChoices = "ita-n"
	//ITA_O PowerOutletTypeChoices = "ita-o"
	//USB_A PowerOutletTypeChoices = "usb-a"
	//USB_MICRO_B PowerOutletTypeChoices = "usb-micro-b"
	//USB_C PowerOutletTypeChoices = "usb-c"
	HDOT_CX PowerOutletTypeChoices = "hdot-cx"
)

// All allowed values of PowerOutletTypeChoices enum
var AllowedPowerOutletTypeChoicesEnumValues = []PowerOutletTypeChoices{
	"iec-60320-c5",
	"iec-60320-c7",
	"iec-60320-c13",
	"iec-60320-c15",
	"iec-60320-c19",
	"iec-60309-p-n-e-4h",
	"iec-60309-p-n-e-6h",
	"iec-60309-p-n-e-9h",
	"iec-60309-2p-e-4h",
	"iec-60309-2p-e-6h",
	"iec-60309-2p-e-9h",
	"iec-60309-3p-e-4h",
	"iec-60309-3p-e-6h",
	"iec-60309-3p-e-9h",
	"iec-60309-3p-n-e-4h",
	"iec-60309-3p-n-e-6h",
	"iec-60309-3p-n-e-9h",
	"nema-1-15r",
	"nema-5-15r",
	"nema-5-20r",
	"nema-5-30r",
	"nema-5-50r",
	"nema-6-15r",
	"nema-6-20r",
	"nema-6-30r",
	"nema-6-50r",
	"nema-10-30r",
	"nema-10-50r",
	"nema-14-20r",
	"nema-14-30r",
	"nema-14-50r",
	"nema-14-60r",
	"nema-15-15r",
	"nema-15-20r",
	"nema-15-30r",
	"nema-15-50r",
	"nema-15-60r",
	"nema-l1-15r",
	"nema-l5-15r",
	"nema-l5-20r",
	"nema-l5-30r",
	"nema-l5-50r",
	"nema-l6-15r",
	"nema-l6-20r",
	"nema-l6-30r",
	"nema-l6-50r",
	"nema-l10-30r",
	"nema-l14-20r",
	"nema-l14-30r",
	"nema-l14-50r",
	"nema-l14-60r",
	"nema-l15-20r",
	"nema-l15-30r",
	"nema-l15-50r",
	"nema-l15-60r",
	"nema-l21-20r",
	"nema-l21-30r",
	"CS6360C",
	"CS6364C",
	"CS8164C",
	"CS8264C",
	"CS8364C",
	"CS8464C",
	"ita-e",
	"ita-f",
	"ita-g",
	"ita-h",
	"ita-i",
	"ita-j",
	"ita-k",
	"ita-l",
	"ita-m",
	"ita-n",
	"ita-o",
	"usb-a",
	"usb-micro-b",
	"usb-c",
	"hdot-cx",
}

func (v *PowerOutletTypeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerOutletTypeChoices(value)
	for _, existing := range AllowedPowerOutletTypeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerOutletTypeChoices", value)
}

// NewPowerOutletTypeChoicesFromValue returns a pointer to a valid PowerOutletTypeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerOutletTypeChoicesFromValue(v string) (*PowerOutletTypeChoices, error) {
	ev := PowerOutletTypeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerOutletTypeChoices: valid values are %v", v, AllowedPowerOutletTypeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerOutletTypeChoices) IsValid() bool {
	for _, existing := range AllowedPowerOutletTypeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerOutletTypeChoices value
func (v PowerOutletTypeChoices) Ptr() *PowerOutletTypeChoices {
	return &v
}

type NullablePowerOutletTypeChoices struct {
	value *PowerOutletTypeChoices
	isSet bool
}

func (v NullablePowerOutletTypeChoices) Get() *PowerOutletTypeChoices {
	return v.value
}

func (v *NullablePowerOutletTypeChoices) Set(val *PowerOutletTypeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerOutletTypeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerOutletTypeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerOutletTypeChoices(val *PowerOutletTypeChoices) *NullablePowerOutletTypeChoices {
	return &NullablePowerOutletTypeChoices{value: val, isSet: true}
}

func (v NullablePowerOutletTypeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerOutletTypeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
