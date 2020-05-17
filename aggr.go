package ontap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type AggrInfo struct {
	Name            string
	SizeTotal       string
	SizeUsed        string
	SizeAvailable   string
	SizeUsedPercent string
	State           string
	Cluster         string
	// Data Compaction
	DataCompactionSpaceSaved        string
	DataCompactionSpaceSavedPercent string
	// Sis
	SisSpaceSaved        string
	SisSpaceSavedPercent string
}

func (c *Client) GetAggrInfo(limit int) ([]AggrInfo, error) {
	ixml := &AggrInfoRequest{}
	ixml.Version = apiVersion
	ixml.Xmlns = XMLns
	ixml.AggrGetIter.MaxRecords = strconv.Itoa(limit)

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

	var result AggrInfoResponse
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if strings.Compare(result.Results.Status, "passed") != 0 {
		return nil, fmt.Errorf("%s", xmlError)
	}

	var ail []AggrInfo

	for _, v := range result.Results.AttributesList.AggrAttributes {
		ai := AggrInfo{
			Name:                            v.AggregateName,
			SizeTotal:                       v.AggrSpaceAttributes.SizeTotal,
			SizeUsed:                        v.AggrSpaceAttributes.SizeUsed,
			SizeAvailable:                   v.AggrSpaceAttributes.SizeAvailable,
			SizeUsedPercent:                 v.AggrSpaceAttributes.PercentUsedCapacity,
			Cluster:                         v.AggrOwnershipAttributes.Cluster,
			DataCompactionSpaceSaved:        v.AggrSpaceAttributes.DataCompactionSpaceSaved,
			DataCompactionSpaceSavedPercent: v.AggrSpaceAttributes.DataCompactionSpaceSavedPercent,
			SisSpaceSaved:                   v.AggrSpaceAttributes.SisSpaceSaved,
			SisSpaceSavedPercent:            v.AggrSpaceAttributes.SisSpaceSavedPercent,
		}
		ail = append(ail, ai)
	}

	return ail, nil
}

func (c *Client) GetAggrPerf() ([]PerfCounter, error) {
	// Performance counters of interest
	var counters []string
	counters = append(counters, "user_reads")        // Number of user reads per second to the aggregate. per_sec
	counters = append(counters, "read_data")         // Amount of data read per second from the aggregate. b_per_sec
	counters = append(counters, "user_read_latency") // Average latency per block in microseconds for user read operations. microsec
	counters = append(counters, "user_writes")
	counters = append(counters, "user_write_latency")
	counters = append(counters, "write_data")
	counters = append(counters, "latency")
	counters = append(counters, "aggr_throughput") // Total amount of CP read data, user read data, and user write data per second. b_per_sec

	var inst []string

	agi, err := c.GetAggrInfo(100)
	if err != nil {
		return nil, err
	}

	// Create instance list and exclude internal aggr
	r1, _ := regexp.Compile("^root_")

	for _, v := range agi {
		if !(r1.MatchString(v.Name)) {
			inst = append(inst, v.Name)
		}
	}

	pc, err := c.getPerformanceData("aggregate", counters, inst)
	if err != nil {
		return nil, err
	}

	return pc, nil
}
