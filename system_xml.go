package ontap

import "encoding/xml"

type SystemNodeRequest struct {
	XMLName                   xml.Name `xml:"netapp"`
	Text                      string   `xml:",chardata"`
	Version                   string   `xml:"version,attr"`
	Xmlns                     string   `xml:"xmlns,attr"`
	NmsdkVersion              string   `xml:"nmsdk_version,attr"`
	NmsdkPlatform             string   `xml:"nmsdk_platform,attr"`
	NmsdkLanguage             string   `xml:"nmsdk_language,attr"`
	EnvironmentSensorsGetIter string   `xml:"system-node-get-iter"`
}

type SystemNodeResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text           string `xml:",chardata"`
		Status         string `xml:"status,attr"`
		AttributesList struct {
			Text        string            `xml:",chardata"`
			NodeDetails []NodeDetailsInfo `xml:"node-details-info"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}

type NodeDetailsInfo struct {
	Text                        string `xml:",chardata"`
	CpuBusytime                 string `xml:"cpu-busytime"`
	CpuFirmwareRelease          string `xml:"cpu-firmware-release"`
	EnvFailedFanCount           string `xml:"env-failed-fan-count"`
	EnvFailedFanMessage         string `xml:"env-failed-fan-message"`
	EnvFailedPowerSupplyCount   string `xml:"env-failed-power-supply-count"`
	EnvFailedPowerSupplyMessage string `xml:"env-failed-power-supply-message"`
	EnvOverTemperature          string `xml:"env-over-temperature"`
	IsAllFlashOptimized         string `xml:"is-all-flash-optimized"`
	IsDiffSvcs                  string `xml:"is-diff-svcs"`
	IsEpsilonNode               string `xml:"is-epsilon-node"`
	IsNodeClusterEligible       string `xml:"is-node-cluster-eligible"`
	IsNodeHealthy               string `xml:"is-node-healthy"`
	MaximumAggregateSize        string `xml:"maximum-aggregate-size"`
	MaximumNumberOfVolumes      string `xml:"maximum-number-of-volumes"`
	MaximumVolumeSize           string `xml:"maximum-volume-size"`
	Node                        string `xml:"node"`
	NodeLocation                string `xml:"node-location"`
	NodeModel                   string `xml:"node-model"`
	NodeNvramID                 string `xml:"node-nvram-id"`
	NodeOwner                   string `xml:"node-owner"`
	NodeSerialNumber            string `xml:"node-serial-number"`
	NodeStorageConfiguration    string `xml:"node-storage-configuration"`
	NodeSystemID                string `xml:"node-system-id"`
	NodeUptime                  string `xml:"node-uptime"`
	NodeUuid                    string `xml:"node-uuid"`
	NodeVendor                  string `xml:"node-vendor"`
	NvramBatteryStatus          string `xml:"nvram-battery-status"`
	ProductVersion              string `xml:"product-version"`
}
