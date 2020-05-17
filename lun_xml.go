package ontap

import "encoding/xml"

type LunListResult struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text           string `xml:",chardata"`
		Status         string `xml:"status,attr"`
		AttributesList struct {
			Text    string        `xml:",chardata"`
			LunInfo []LunInfoType `xml:"lun-info"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}

type LunInfoType struct {
	Text                      string `xml:",chardata"`
	Alignment                 string `xml:"alignment"`
	Application               string `xml:"application"`
	ApplicationUuid           string `xml:"application-uuid"`
	BlockSize                 int    `xml:"block-size"`
	Class                     string `xml:"class"`
	Comment                   string `xml:"comment"`
	CreationTimestamp         string `xml:"creation-timestamp"`
	IsClone                   bool   `xml:"is-clone"`
	IsCloneAutodeleteEnabled  bool   `xml:"is-clone-autodelete-enabled"`
	IsInconsistentImport      bool   `xml:"is-inconsistent-import"`
	IsRestoreInaccessible     bool   `xml:"is-restore-inaccessible"`
	IsSpaceAllocEnabled       bool   `xml:"is-space-alloc-enabled"`
	IsSpaceReservationEnabled bool   `xml:"is-space-reservation-enabled"`
	Mapped                    bool   `xml:"mapped"`
	MultiprotocolType         string `xml:"multiprotocol-type"`
	Node                      string `xml:"node"`
	Online                    bool   `xml:"online"`
	Path                      string `xml:"path"`
	PrefixSize                string `xml:"prefix-size"`
	Qtree                     string `xml:"qtree"`
	ReadOnly                  bool   `xml:"read-only"`
	Serial7Mode               string `xml:"serial-7-mode"`
	SerialNumber              string `xml:"serial-number"`
	ShareState                string `xml:"share-state"`
	Size                      int    `xml:"size"`
	SizeUsed                  int    `xml:"size-used"`
	Staging                   string `xml:"staging"`
	State                     string `xml:"state"`
	SuffixSize                string `xml:"suffix-size"`
	Uuid                      string `xml:"uuid"`
	Volume                    string `xml:"volume"`
	Vserver                   string `xml:"vserver"`
}
