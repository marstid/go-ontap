package ontap

import "encoding/xml"

type Base struct {
	XMLName xml.Name `xml:"netapp"`
	Version string   `xml:"version,attr"`
	XMLns string  `xml:"xmlns,attr"`
}
