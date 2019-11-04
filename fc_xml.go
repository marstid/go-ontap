package ontap

import "encoding/xml"

type FCConfigRequest struct {
	XMLName          xml.Name `xml:"netapp"`
	Text             string   `xml:",chardata"`
	Version          string   `xml:"version,attr"`
	Xmlns            string   `xml:"xmlns,attr"`
	NmsdkVersion     string   `xml:"nmsdk_version,attr"`
	NmsdkPlatform    string   `xml:"nmsdk_platform,attr"`
	NmsdkLanguage    string   `xml:"nmsdk_language,attr"`
	FCConfigListInfo string   `xml:"fc-config-list-info"`
}

type FcConfig struct {
	Text          string `xml:",chardata"`
	AdapterName   string `xml:"adapter-name"`
	AdapterState  string `xml:"adapter-state"`
	AdapterStatus string `xml:"adapter-status"`
	AdapterType   string `xml:"adapter-type"`
	NodeName      string `xml:"node-name"`
}

type FcConfigReply struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text           string `xml:",chardata"`
		Status         string `xml:"status,attr"`
		AttributesList struct {
			Text      string     `xml:",chardata"`
			FcConfigs []FcConfig `xml:"fc-config"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}

type FCPAdapterRequest struct {
	XMLName          xml.Name `xml:"netapp"`
	Text             string   `xml:",chardata"`
	Version          string   `xml:"version,attr"`
	Xmlns            string   `xml:"xmlns,attr"`
	NmsdkVersion     string   `xml:"nmsdk_version,attr"`
	NmsdkPlatform    string   `xml:"nmsdk_platform,attr"`
	NmsdkLanguage    string   `xml:"nmsdk_language,attr"`
	FCConfigListInfo string   `xml:"fcp-adapter-get-iter"`
}

type FcpConfigAdapterInfo struct {
	Text                   string `xml:",chardata"`
	Adapter                string `xml:"adapter"`
	ConnectionEstablished  string `xml:"connection-established"`
	DataLinkRate           string `xml:"data-link-rate"`
	DataProtocolsSupported struct {
		Text                   string   `xml:",chardata"`
		FcpAdapterDataProtocol []string `xml:"fcp-adapter-data-protocol"`
	} `xml:"data-protocols-supported"`
	FabricEstablished                    string `xml:"fabric-established"`
	FirmwareRev                          string `xml:"firmware-rev"`
	HardwareRev                          string `xml:"hardware-rev"`
	InfoName                             string `xml:"info-name"`
	IsSfpDiagnosticsInternallyCalibrated string `xml:"is-sfp-diagnostics-internally-calibrated"`
	IsSfpOpticalTransceiverValid         string `xml:"is-sfp-optical-transceiver-valid"`
	IsSfpRxPowerInRange                  string `xml:"is-sfp-rx-power-in-range"`
	IsSfpTxPowerInRange                  string `xml:"is-sfp-tx-power-in-range"`
	MaxSpeed                             string `xml:"max-speed"`
	MediaType                            string `xml:"media-type"`
	Node                                 string `xml:"node"`
	NodeName                             string `xml:"node-name"`
	PhysicalProtocol                     string `xml:"physical-protocol"`
	PortAddress                          string `xml:"port-address"`
	PortName                             string `xml:"port-name"`
	SfpConnector                         string `xml:"sfp-connector"`
	SfpDateCode                          string `xml:"sfp-date-code"`
	SfpEncoding                          string `xml:"sfp-encoding"`
	SfpFcSpeedcapabilities               string `xml:"sfp-fc-speedcapabilities"`
	SfpFormfactor                        string `xml:"sfp-formfactor"`
	SfpPartNumber                        string `xml:"sfp-part-number"`
	SfpRev                               string `xml:"sfp-rev"`
	SfpRxPower                           string `xml:"sfp-rx-power"`
	SfpSerialNumber                      string `xml:"sfp-serial-number"`
	SfpTxPower                           string `xml:"sfp-tx-power"`
	SfpVendorName                        string `xml:"sfp-vendor-name"`
	SfpVendorOui                         string `xml:"sfp-vendor-oui"`
	SfpWavelength                        string `xml:"sfp-wavelength"`
	Speed                                string `xml:"speed"`
	State                                string `xml:"state"`
	StatusAdmin                          string `xml:"status-admin"`
	StatusExtended                       string `xml:"status-extended"`
	SwitchPort                           string `xml:"switch-port"`
	FabricName                           string `xml:"fabric-name"`
}

type FCPAdapterResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text           string `xml:",chardata"`
		Status         string `xml:"status,attr"`
		AttributesList struct {
			Text                  string                 `xml:",chardata"`
			FcpConfigAdapterInfos []FcpConfigAdapterInfo `xml:"fcp-config-adapter-info"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}
