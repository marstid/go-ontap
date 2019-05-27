package ontap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"net/http"
	"strings"
)

type System struct{}

type SystemGetVersion struct {
	Base
	Version struct{} `xml:"system-get-version"`
}

func NewSystemGetVersion() *SystemGetVersion {
	sgv := &SystemGetVersion{}
	sgv.Base.Version = apiVersion
	sgv.Base.XMLns = XMLns
	return sgv
}

type SystemGetVersionResult struct {
	Base
	Results struct {
		Status    string `xml:"status,attr"`
		Clustered bool   `xml:"is-clustered"`
		Version   string `xml:"version"`
	} `xml:"results"`
}

func (c *Client) GetSystemNodeInfo() (nodeDetails []NodeDetailsInfo, err error) {
	ixml := &SystemNodeRequest{}
	ixml.Version = apiVersion
	ixml.Xmlns = XMLns

	output, err := xml.MarshalIndent(ixml, "", "\t")

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

	var result SystemNodeResponse
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if strings.Compare(result.Results.Status, "passed") != 0 {
		return nil, fmt.Errorf("%s", xmlError)
	}

	return result.Results.AttributesList.NodeDetails, nil
}

func (c *Client) GetSystemPerf() ([]PerfCounter, error) {
	// Performance counters of interest
	var counters []string
	counters = append(counters, "total_processor_busy")
	counters = append(counters, "total_processor_busy_time")
	counters = append(counters, "cpu_busy")
	counters = append(counters, "cpu_elapsed_time")

	var inst []string

	ni, err := c.GetSystemNodeInfo()
	if err != nil {
		return nil, err
	}
	for _, v := range ni {
		inst = append(inst, v.Node)
	}

	pc, err := c.getPerformanceData("system:node", counters, inst)
	if err != nil {
		return nil, err
	}

	return pc, nil
}
