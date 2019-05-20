package ontap

import (
	"bytes"
	"encoding/xml"
	"github.com/go-xmlfmt/xmlfmt"
	"net/http"
)

type ClusterIdInfo struct {
	Contact      string
	Location     string
	Name         string
	SerialNumber string
}

func NewClusterIdGet() *ClusterInfoRequest {
	sgv := &ClusterInfoRequest{}
	sgv.Version = apiVersion
	sgv.Xmlns = XMLns
	return sgv
}

func (c *Client) GetIdClusterInfo() (clusterId IdentityInfo, error error) {

	sgv := NewClusterIdGet()
	output, err := xml.MarshalIndent(sgv, "", "\t")
	payload := bytes.NewReader(output)

	if c.Debug {
		x := xmlfmt.FormatXML(string(output), "\t", "  ")
		println("Request:")
		println(x)
	}

	req, err := http.NewRequest("POST", c.Url, payload)
	if err != nil {
		return clusterId, err
	}

	response, err := c.doRequest(req)
	if err != nil {
		return clusterId, err
	}

	if c.Debug {
		x := xmlfmt.FormatXML(string(response), "\t", "  ")
		println("Response:")
		println(x)
	}

	var result ClusterIdentityInfo
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return clusterId, err
	}

	clusterId = result.Results.Attributes.ClusterIdentityInfo

	return clusterId, nil

}
