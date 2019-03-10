package ontap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/go-xmlfmt/xmlfmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const basePath string = "/servlets/netapp.servlets.admin.XMLrequest_filer"
const agent string = "go-ontap"
const apiVersion string = "1.3"
const XMLns = "http://www.netapp.com/filer/admin"
const xmlError = "ONTAP API Call failed"

type Client struct {
	UserName  string
	Password  string
	Host      string
	VerifySSL bool
	SSL       bool
	TimeOut   time.Duration
	Url       string
	Debug     bool
}

// Instantiate new client
func NewClient(host, username, password string, ssl bool) *Client {

	url := "https://" + host + basePath
	if !ssl {
		url = "http://" + host + basePath
	}

	return &Client{
		UserName:  username,
		Password:  password,
		Host:      host,
		VerifySSL: ssl,
		SSL:       true,
		TimeOut:   10,
		Url:       url,
		Debug:     false,
	}
}

func (c *Client) GetSystemVersion() (version string, error error) {

	sgv := NewSystemGetVersion()
	output, err := xml.MarshalIndent(sgv, "", "\t")
	payload := bytes.NewReader(output)

	req, err := http.NewRequest("POST", c.Url, payload)
	if err != nil {
		return "", err
	}

	response, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	var result SystemGetVersionResult
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	s := strings.Split(result.Results.Version, ":")

	return s[0], nil

}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.UserName, c.Password)
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("UserAgent", agent)

	httpClient := &http.Client{
		Timeout: time.Second * c.TimeOut,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.VerifySSL,
			},
		},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO check context implications
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil

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
			if c.Debug {
				println(v.DiskRaidInfo.DiskAggregateInfo.IsOffline)
			}
		}
		prefailed := false
		if v.DiskRaidInfo.DiskAggregateInfo.IsPrefailed == "true" {
			prefailed = true
			if c.Debug {
				println(v.DiskRaidInfo.DiskAggregateInfo.IsPrefailed)
			}
		}

		ai := DiskInfo{
			Name:      v.DiskName,
			DiskType:  v.DiskInventoryInfo.DiskType,
			Model:     v.DiskInventoryInfo.Model,
			Online:    online,
			Prefailed: prefailed,
		}

		ail = append(ail, ai)
	}

	return ail, nil
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
			Name:            v.AggregateName,
			SizeTotal:       v.AggrSpaceAttributes.SizeTotal,
			SizeUsed:        v.AggrSpaceAttributes.SizeUsed,
			SizeAvailable:   v.AggrSpaceAttributes.SizeAvailable,
			SizeUsedPercent: v.AggrSpaceAttributes.PercentUsedCapacity,
			Cluster:         v.AggrOwnershipAttributes.Cluster,
		}

		ail = append(ail, ai)
	}

	return ail, nil
}

func (c *Client) GetAggrPerf() ([]PerfCounter, error) {
	// Performance counters of interest
	var counters []string
	counters = append(counters, "user_reads")
	counters = append(counters, "read_data")
	counters = append(counters, "user_read_latency")
	counters = append(counters, "user_writes")
	counters = append(counters, "user_write_latency")
	counters = append(counters, "write_data")
	counters = append(counters, "latency")
	counters = append(counters, "aggr_throughput")
	counters = append(counters, "latency")

	var inst []string

	agi, err := c.GetAggrInfo(10)
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

func (c *Client) GetVolumePerf() ([]PerfCounter, error) {
	// Performance counters of interest
	var counter []string
	counter = append(counter, "total_ops")
	counter = append(counter, "avg_latency")
	counter = append(counter, "read_ops")
	counter = append(counter, "read_latency")
	counter = append(counter, "read_data")
	counter = append(counter, "write_ops")
	counter = append(counter, "write_latency")
	counter = append(counter, "write_data")

	var inst []string

	m, err := c.getVolumeToAggrMap()
	if err != nil {
		return nil, err
	}

	// Create instance list and exclude internal volumes
	r1, _ := regexp.Compile("^vol0")
	r2, _ := regexp.Compile("_root$")

	for k := range m {
		if !(r1.MatchString(k) || r2.MatchString(k)) {
			inst = append(inst, k)
		}
	}

	pc, err := c.getPerformanceData("volume", counter, inst)
	if err != nil {
		return nil, err
	}

	return pc, nil
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

func (c *Client) getVolumeToAggrMap() (map[string]string, error) {
	m := make(map[string]string)
	list, err := c.GetVolumeInfo(1000)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		m[v.Name] = v.Aggr
	}

	return m, nil
}

// Returns custom struct of Volumes Info
func (c *Client) GetVolumeInfo(limit int) (volumes []VolumeInfo, err error) {
	ixml := &VolumeGetIterShort{}
	ixml.Version = apiVersion
	ixml.Xmlns = XMLns
	ixml.VolumeGetIter.MaxRecords = strconv.Itoa(limit)

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

	var result NetappVolume
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	if strings.Compare(result.Results.Status, "passed") != 0 {
		return nil, fmt.Errorf("%s", xmlError)
	}

	// Convert to custom struct
	var volList []VolumeInfo
	for _, v := range result.Results.AttributesList.VolumeAttributes {
		vol := VolumeInfo{
			Name:        v.VolumeIDAttributes.Name,
			Aggr:        v.VolumeIDAttributes.AggrList.AggrName,
			SizeTotal:   v.VolumeSpaceAttributes.SizeTotal,
			PercentUsed: v.VolumeSpaceAttributes.PercentageSizeUsed,
			SizeUsed:    v.VolumeSpaceAttributes.SizeUsed,
			SizeFree:    v.VolumeSpaceAttributes.SizeAvailable,
			State:       v.VolumeStateAttributes.State}
		volList = append(volList, vol)
	}

	return volList, nil
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
