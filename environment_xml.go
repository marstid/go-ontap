package ontap

import "encoding/xml"

type EnvRequest struct {
	XMLName                   xml.Name `xml:"netapp"`
	Text                      string   `xml:",chardata"`
	Version                   string   `xml:"version,attr"`
	Xmlns                     string   `xml:"xmlns,attr"`
	NmsdkVersion              string   `xml:"nmsdk_version,attr"`
	NmsdkPlatform             string   `xml:"nmsdk_platform,attr"`
	NmsdkLanguage             string   `xml:"nmsdk_language,attr"`
	EnvironmentSensorsGetIter string   `xml:"environment-sensors-get-iter"`
}

type EnvResults struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Results struct {
		Text           string `xml:",chardata"`
		Status         string `xml:"status,attr"`
		AttributesList struct {
			Text               string                   `xml:",chardata"`
			EnvironmentSensors []EnvironmentSensorsInfo `xml:"environment-sensors-info"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}

type EnvironmentSensorsInfo struct {
	Text                  string `xml:",chardata"`
	DiscreteSensorState   string `xml:"discrete-sensor-state"`
	DiscreteSensorValue   string `xml:"discrete-sensor-value"`
	NodeName              string `xml:"node-name"`
	SensorName            string `xml:"sensor-name"`
	SensorType            string `xml:"sensor-type"`
	ThresholdSensorState  string `xml:"threshold-sensor-state"`
	CriticalHighThreshold string `xml:"critical-high-threshold"`
	ThresholdSensorValue  string `xml:"threshold-sensor-value"`
	ValueUnits            string `xml:"value-units"`
	WarningHighThreshold  string `xml:"warning-high-threshold"`
	CriticalLowThreshold  string `xml:"critical-low-threshold"`
	WarningLowThreshold   string `xml:"warning-low-threshold"`
}
