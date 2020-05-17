package ontap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"net/http"
	"strings"
)

type PerfCounter struct {
	ObjectName string
	Counter    string
	Value      string
}

type PerfCounterInfo struct {
	Name       string
	Desc       string
	Unit       string
	Deprecated string
}

func (c *Client) GetPerfCounters(object string) (volumes []PerfCounterInfo, err error) {
	ixml := &PerfCounterRequest{}
	ixml.Version = apiVersion
	ixml.Xmlns = XMLns
	ixml.PerfObjectGetInstances.Objectname = object

	output, err := xml.MarshalIndent(ixml, "", "\t")
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

	var result PerfCounterResponse
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if strings.Compare(result.Results.Status, "passed") != 0 {
		return nil, fmt.Errorf("%s", xmlError)
	}

	var pci []PerfCounterInfo
	for _, v := range result.Results.Counters.CounterInfo {
		pc := PerfCounterInfo{
			Name:       v.Name,
			Desc:       v.Desc,
			Unit:       v.Unit,
			Deprecated: v.IsDeprecated,
		}
		pci = append(pci, pc)
	}

	return pci, nil
}

func (c *Client) GetPerfObject() (oi []ObjectInfo, err error) {

	ixml := &PerfObjectRequest{}
	ixml.Version = apiVersion
	ixml.Xmlns = XMLns

	output, err := xml.MarshalIndent(ixml, "", "\t")
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

	var result PerfObjectResponse
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if strings.Compare(result.Results.Status, "passed") != 0 {
		return nil, fmt.Errorf("%s", xmlError)
	}

	return result.Results.Objects.ObjectInfoList, nil
}

func (c *Client) getPerformanceData(object string, counters []string, instances []string) ([]PerfCounter, error) {
	ixml := &PerfRequest{}
	ixml.Version = apiVersion
	ixml.Xmlns = XMLns
	ixml.PerfObjectGetInstances.Objectname = object
	if counters != nil {
		ixml.PerfObjectGetInstances.Counters.Counter = counters
	}
	if instances != nil {
		ixml.PerfObjectGetInstances.Instances.Instance = instances
	}

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

	var result PerfResponse
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if strings.Compare(result.Results.Status, "passed") != 0 {
		return nil, fmt.Errorf("%s", xmlError)
	}

	// Convert to custom struct
	var perfC []PerfCounter
	for _, v := range result.Results.Instances.InstanceData {
		for _, p := range v.Counters.CounterData {
			prf := PerfCounter{ObjectName: v.Name, Counter: p.Name, Value: p.Value}
			perfC = append(perfC, prf)
		}
	}

	return perfC, nil
}
