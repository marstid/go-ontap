package ontap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"net/http"
	"strconv"
	"strings"
)



type VolumeInfo struct {
	Name string
	Aggr string
	SizeTotal string
	PercentUsed string
	SizeUsed string
	SizeFree string
	State string
}




type Volume struct {
	Client *Client
}

func NewVolume(c *Client) *Volume{

	return &Volume{
		Client: c,
	}
}


type VolumeAttributes struct {
	VolumeIdAttribute VolumeIdAttribute `xml:"volume-id-attributes"`
	VolumeSpaceAttribute VolumeSpaceAttribute `xml:"volume-space-attributes"`
}

type VolumeSpaceAttribute struct {
	Size string  `xml:"size"`
	PercentUsed string `xml:"percentage-size-used"`
	Used string  `xml:"size-used"`
	Free string  `xml:"size-available"`
}

type VolumeIdAttribute struct {
	Name string `xml:"name"`
	AggName string  `xml:"containing-aggregate-name"`
}

type VolumeGetIter struct {
	Base
	Query struct {
		MaxRecords int `xml:"max-records"`
		DesiredAttributes struct {} `xml:"desired-attributes"`
	} `xml:"volume-get-iter"`
}

func NewVolumeGetIter(limit int) *VolumeGetIter {
	vgi := &VolumeGetIter{}
	vgi.Query.MaxRecords = limit
	vgi.Base.Version = apiVersion
	vgi.Base.XMLns = XMLns
	return vgi
}

type VolumeGetIterResult struct {
	Base
	Results struct {
		Status string `xml:"status,attr"`
		AttributesList struct {
			VolumeAttributes []VolumeAttributes `xml:"volume-attributes"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}


func (v *Volume) List(limit int) (vList []string, err error){

	vgi := NewVolumeGetIter(limit)

	output, err := xml.MarshalIndent(vgi, "", "\t")
	payload := bytes.NewReader(output)

	req, err := http.NewRequest("POST", v.Client.Url, payload)
	if err != nil {
		return nil, err
	}

	response, err := v.Client.doRequest(req)
	if err != nil {
		return nil, err
	}

	x := xmlfmt.FormatXML(string(response), "\t", "  ")
	println(x)


	var result NetappVolume
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}



	for k, v := range result.Results.AttributesList.VolumeAttributes {
		fmt.Println(k,v.VolumeIDAttributes.Name, v.VolumeSpaceAttributes.Size)
	}

	fmt.Println()


	return nil, err
}

func (v *Volume) GetVolumeAttributes(limit int) (volumes []VolumeAttributes, err error){
	ixml := &VolumeGetIterShort{}
	ixml.Version = apiVersion
	ixml.Xmlns = XMLns
	ixml.VolumeGetIter.MaxRecords =strconv.Itoa(limit)

	output, err := xml.MarshalIndent(ixml, "", "\t")
	print(xmlfmt.FormatXML(string(output), "\t", "  "))


	payload := bytes.NewReader(output)

	req, err := http.NewRequest("POST", v.Client.Url, payload)
	if err != nil {
		return nil, err
	}

	response, err := v.Client.doRequest(req)
	if err != nil {
		return nil, err
	}

	x := xmlfmt.FormatXML(string(response), "\t", "  ")
	println(x)


	var result NetappVolume
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if strings.Compare(result.Results.Status, "passed") != 1 {
		return nil, fmt.Errorf("%s","foobar")
	}


	return nil, nil
}

