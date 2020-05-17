package ontap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"net/http"
	"strings"
)

type DiskInfo struct {
	Name      string
	DiskType  string
	Model     string
	Online    bool
	Prefailed bool
	Spare     bool
}

func (c *Client) GetDiskInfo() ([]DiskInfo, error) {
	ixml := &DiskInfoRequest{}
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

	var result DiskInfoResponse
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if strings.Compare(result.Results.Status, "passed") != 0 {
		return nil, fmt.Errorf("%s", xmlError)
	}

	var ail []DiskInfo

	for _, v := range result.Results.AttributesList.StorageDiskInfo {
		online := true
		if v.DiskRaidInfo.DiskAggregateInfo.IsOffline == "true" {
			online = false
		}
		prefailed := false
		if v.DiskRaidInfo.DiskAggregateInfo.IsPrefailed == "true" {
			prefailed = true
		}

		spare := false
		if v.DiskRaidInfo.ContainerType == "spare" {
			spare = true
		}

		ai := DiskInfo{
			Name:      v.DiskName,
			DiskType:  v.DiskInventoryInfo.DiskType,
			Model:     v.DiskInventoryInfo.Model,
			Online:    online,
			Prefailed: prefailed,
			Spare:     spare,
		}

		ail = append(ail, ai)
	}

	return ail, nil
}

func (c *Client) GetDiskPerf() ([]PerfCounter, error) {
	// Performance counters of interest
	var counters []string
	counters = append(counters, "base_for_disk_busy")
	counters = append(counters, "disk_busy")
	counters = append(counters, "name")

	var inst []string

	agi, err := c.GetDiskInfo()
	if err != nil {
		return nil, err
	}

	for _, v := range agi {
		inst = append(inst, v.Name)
	}

	pc, err := c.getPerformanceData("disk", counters, inst)
	if err != nil {
		return nil, err
	}

	return pc, nil
}
