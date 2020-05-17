package ontap

import (
	"bytes"
	"encoding/xml"
	"github.com/go-xmlfmt/xmlfmt"
	"net/http"
)

// Get slice of FcConfig
func (c *Client) GetFCConfig() (fcc []FcConfig, error error) {

	egs := &FCConfigRequest{}
	egs.Version = apiVersion
	egs.Xmlns = XMLns

	output, err := xml.MarshalIndent(egs, "", "\t")
	if c.Debug {
		print(xmlfmt.FormatXML(string(output), "\t", "  "))
	}

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

	var result FcConfigReply
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.Results.AttributesList.FcConfigs, nil
}

// Get slice of FcpConfigAdapterInfo
func (c *Client) GetFCPAdapter() (fcc []FcpConfigAdapterInfo, error error) {

	egs := &FCPAdapterRequest{}
	egs.Version = apiVersion
	egs.Xmlns = XMLns

	output, err := xml.MarshalIndent(egs, "", "\t")
	if c.Debug {
		print(xmlfmt.FormatXML(string(output), "\t", "  "))
	}

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

	var result FCPAdapterResponse
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.Results.AttributesList.FcpConfigAdapterInfos, nil
}
