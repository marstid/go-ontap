package ontap

import (
	"bytes"
	"encoding/xml"
	"github.com/go-xmlfmt/xmlfmt"
	"net/http"
)

type LunGetIter struct {
	Base
	Query struct {
		MaxRecords        int      `xml:"max-records"`
		DesiredAttributes struct{} `xml:"desired-attributes"`
	} `xml:"lun-get-iter"`
}

func NewLunGetIter(limit int) *LunGetIter {
	lgi := &LunGetIter{}
	lgi.Query.MaxRecords = limit
	lgi.Base.Version = apiVersion
	lgi.Base.XMLns = XMLns
	return lgi
}

func (c *Client) GetLunInfo(limit int) (lunInfo []LunInfoType, err error) {

	gi := NewLunGetIter(limit)

	output, err := xml.MarshalIndent(gi, "", "\t")
	payload := bytes.NewReader(output)

	req, err := http.NewRequest("POST", c.Url, payload)
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		x := xmlfmt.FormatXML(string(response), "\t", "  ")
		println(x)
	}

	var result LunListResult
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.Results.AttributesList.LunInfo, nil
}
