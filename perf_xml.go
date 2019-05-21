package ontap

import "encoding/xml"

type PerfRequest struct {
	XMLName                xml.Name `xml:"netapp"`
	Version                string   `xml:"version,attr"`
	Xmlns                  string   `xml:"xmlns,attr"`
	PerfObjectGetInstances struct {
		Objectname string `xml:"objectname"`
		Counters   struct {
			Counter []string `xml:"counter"`
		} `xml:"counters"`
		Instances struct {
			Instance []string `xml:"instance"`
		} `xml:"instances"`
	} `xml:"perf-object-get-instances"`
}

type PerfResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text      string `xml:",chardata"`
		Status    string `xml:"status,attr"`
		Instances struct {
			Text         string `xml:",chardata"`
			InstanceData []struct {
				Text     string `xml:",chardata"`
				Counters struct {
					Text        string `xml:",chardata"`
					CounterData []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name"`
						Value string `xml:"value"`
					} `xml:"counter-data"`
				} `xml:"counters"`
				Name   string `xml:"name"`
				SortID string `xml:"sort-id"`
				Uuid   string `xml:"uuid"`
			} `xml:"instance-data"`
		} `xml:"instances"`
		Timestamp string `xml:"timestamp"`
	} `xml:"results"`
}

type PerfCounterRequest struct {
	XMLName                xml.Name `xml:"netapp"`
	Version                string   `xml:"version,attr"`
	Xmlns                  string   `xml:"xmlns,attr"`
	PerfObjectGetInstances struct {
		Objectname string `xml:"objectname"`
	} `xml:"perf-object-counter-list-info"`
}

type PerfCounterResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Text    string   `xml:",chardata"`
	Results struct {
		Text     string `xml:",chardata"`
		Status   string `xml:"status,attr"`
		Counters struct {
			Text        string `xml:",chardata"`
			CounterInfo []struct {
				Text           string `xml:",chardata"`
				Desc           string `xml:"desc"`
				IsDeprecated   string `xml:"is-deprecated"`
				Name           string `xml:"name"`
				PrivilegeLevel string `xml:"privilege-level"`
				Properties     string `xml:"properties"`
				Unit           string `xml:"unit"`
			} `xml:"counter-info"`
		} `xml:"counters"`
	} `xml:"results"`
}

type PerfObjectRequest struct {
	XMLName       xml.Name `xml:"netapp"`
	Text          string   `xml:",chardata"`
	Version       string   `xml:"version,attr"`
	Xmlns         string   `xml:"xmlns,attr"`
	NmsdkVersion  string   `xml:"nmsdk_version,attr"`
	NmsdkPlatform string   `xml:"nmsdk_platform,attr"`
	NmsdkLanguage string   `xml:"nmsdk_language,attr"`
	PerfObject    string   `xml:"perf-object-list-info"`
}

type PerfObjectResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Results struct {
		Text    string `xml:",chardata"`
		Status  string `xml:"status,attr"`
		Objects struct {
			Text           string       `xml:",chardata"`
			ObjectInfoList []ObjectInfo `xml:"object-info"`
		} `xml:"objects"`
	} `xml:"results"`
}

type ObjectInfo struct {
	Text                         string `xml:",chardata"`
	Description                  string `xml:"description"`
	IsDeprecated                 string `xml:"is-deprecated"`
	Name                         string `xml:"name"`
	PrivilegeLevel               string `xml:"privilege-level"`
	GetInstancesPreferredCounter string `xml:"get-instances-preferred-counter"`
}
