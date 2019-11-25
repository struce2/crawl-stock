package models

import "net/url"

type RequestBody struct {
	APP_KEY      string `json:"app_key"`
	APP_VER      string `json:"app_ver"`
	APP_VER_NAME string `json:"app_ver_name"`
	CONN_TYPE    string `json:"conn_type"`
	DEVICE       string `json:"device"`
	DEVICE_MODE  string `json:"device_mode"`
	GSN          string `json:"gsn"`
	HID          string `json:"hid"`
	PID          string `json:"pid"`
	PLATFORM     string `json:"platform"`
	PLATFORM_OS  string `json:"platform_os"`
	RSN          string `json:"rsn"`
	TYPE         string `json:"type"`
	UID          string `json:"uid"`
}

type Response struct {
}

func NewRequestBody(phone string) *RequestBody {
	return &RequestBody{
		APP_KEY:      url.QueryEscape("com.mtk"),
		APP_VER:      "110",
		APP_VER_NAME: url.QueryEscape("8.28.3.127"),
		CONN_TYPE:    "UMTS",
		DEVICE:       "PHONE",
		DEVICE_MODE:  "SM-G9350",
		GSN:          "00000000000000",
		HID:          "860494714667288",
		PID:          "MTK",
		PLATFORM:     "ANDROID",
		PLATFORM_OS:  "19",
		RSN:          "00000000000000",
		TYPE:         "MOB",
		UID:          url.QueryEscape(phone),
	}
}
