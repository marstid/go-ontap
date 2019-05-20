package ontap

import "encoding/xml"

type ClusterInfoRequest struct {
	XMLName            xml.Name `xml:"netapp"`
	Text               string   `xml:",chardata"`
	Version            string   `xml:"version,attr"`
	Xmlns              string   `xml:"xmlns,attr"`
	NmsdkVersion       string   `xml:"nmsdk_version,attr"`
	NmsdkPlatform      string   `xml:"nmsdk_platform,attr"`
	NmsdkLanguage      string   `xml:"nmsdk_language,attr"`
	ClusterIdentityGet string   `xml:"cluster-identity-get"`
}

type IdentityInfo struct {
	Text                string `xml:",chardata"`
	ClusterContact      string `xml:"cluster-contact"`
	ClusterLocation     string `xml:"cluster-location"`
	ClusterName         string `xml:"cluster-name"`
	ClusterSerialNumber string `xml:"cluster-serial-number"`
	ClusterUuid         string `xml:"cluster-uuid"`
	RdbUuid             string `xml:"rdb-uuid"`
}

type ClusterIdentityInfo struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text       string `xml:",chardata"`
		Status     string `xml:"status,attr"`
		Attributes struct {
			Text                string       `xml:",chardata"`
			ClusterIdentityInfo IdentityInfo `xml:"cluster-identity-info"`
		} `xml:"attributes"`
	} `xml:"results"`
}
