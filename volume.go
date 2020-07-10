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

type VolumeInfo struct {
	Name                              string
	Aggr                              string
	SizeTotal                         string
	PercentUsed                       string
	SizeUsed                          string
	SizeFree                          string
	State                             string
	SnapPercentUsed                   string
	SnapPercentReserve                string
	CompressionSpaceSaved             string
	CompressionPercentageSpaceSaved   string
	DeduplicationSpaceSaved           string
	DeduplicationPercentageSpaceSaved string
	TotalSpaceSaved                   string
	TotalPercentageSpaceSaved         string
	InodeTotal                        string
	InodeUsed                         string
	InodePercent                      string
}

type Volume struct {
	Client *Client
}

func NewVolume(c *Client) *Volume {

	return &Volume{
		Client: c,
	}
}

type VolumeAttributes struct {
	VolumeIdAttribute    VolumeIdAttribute    `xml:"volume-id-attributes"`
	VolumeSpaceAttribute VolumeSpaceAttribute `xml:"volume-space-attributes"`
}

type VolumeSpaceAttribute struct {
	Size        string `xml:"size"`
	PercentUsed string `xml:"percentage-size-used"`
	Used        string `xml:"size-used"`
	Free        string `xml:"size-available"`
}

type VolumeIdAttribute struct {
	Name    string `xml:"name"`
	AggName string `xml:"containing-aggregate-name"`
}

type VolumeGetIter struct {
	Base
	Query struct {
		MaxRecords        int      `xml:"max-records"`
		DesiredAttributes struct{} `xml:"desired-attributes"`
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
		Status         string `xml:"status,attr"`
		AttributesList struct {
			VolumeAttributes []VolumeAttributes `xml:"volume-attributes"`
		} `xml:"attributes-list"`
	} `xml:"results"`
}

/*

func (v *Volume) List(limit int) (vList []string, err error) {

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
		fmt.Println(k, v.VolumeIDAttributes.Name, v.VolumeSpaceAttributes.Size)
	}

	fmt.Println()

	return nil, err
}

func (v *Volume) GetVolumeAttributes(limit int) (volumes []VolumeAttributes, err error) {
	ixml := &VolumeGetIterShort{}
	ixml.Version = apiVersion
	ixml.Xmlns = XMLns
	ixml.VolumeGetIter.MaxRecords = strconv.Itoa(limit)

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
		return nil, fmt.Errorf("%s", "foobar")
	}

	return nil, nil
}

*/

func (c *Client) GetVolumeToAggrMap() (map[string]string, error) {
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

	//  Exclude internal volumes
	r1, _ := regexp.Compile("^vol0")
	r2, _ := regexp.Compile("_root$")

	for _, v := range result.Results.AttributesList.VolumeAttributes {

		// Skip internal volumes
		if r1.MatchString(v.VolumeIDAttributes.Name) || r2.MatchString(v.VolumeIDAttributes.Name) {
			continue
		}

		// Calculate Inode allocation
		InodeTotal, err := strconv.ParseInt(v.VolumeInodeAttributes.FilesTotal, 10, 64)
		if err != nil {
			InodeTotal = 1
		}

		InodeUsed, err := strconv.ParseInt(v.VolumeInodeAttributes.FilesUsed, 10, 64)
		if err != nil {
			InodeTotal = 1
		}

		InodePercent := float64((InodeUsed / InodeTotal) * 100)

		vol := VolumeInfo{
			Name:                              v.VolumeIDAttributes.Name,
			Aggr:                              v.VolumeIDAttributes.ContainingAggregateName,
			SizeTotal:                         v.VolumeSpaceAttributes.SizeTotal,
			PercentUsed:                       v.VolumeSpaceAttributes.PercentageSizeUsed,
			SizeUsed:                          v.VolumeSpaceAttributes.SizeUsed,
			SizeFree:                          v.VolumeSpaceAttributes.SizeAvailable,
			State:                             v.VolumeStateAttributes.State,
			SnapPercentUsed:                   v.VolumeSpaceAttributes.PercentageSnapshotReserveUsed,
			SnapPercentReserve:                v.VolumeSpaceAttributes.PercentageSnapshotReserve,
			CompressionPercentageSpaceSaved:   v.VolumeSisAttributes.PercentageCompressionSpaceSaved,
			CompressionSpaceSaved:             v.VolumeSisAttributes.CompressionSpaceSaved,
			DeduplicationPercentageSpaceSaved: v.VolumeSisAttributes.PercentageDeduplicationSpaceSaved,
			DeduplicationSpaceSaved:           v.VolumeSisAttributes.DeduplicationSpaceSaved,
			TotalPercentageSpaceSaved:         v.VolumeSisAttributes.PercentageTotalSpaceSaved,
			TotalSpaceSaved:                   v.VolumeSisAttributes.TotalSpaceSaved,
			InodeTotal:                        v.VolumeInodeAttributes.FilesTotal,
			InodeUsed:                         v.VolumeInodeAttributes.FilesUsed,
			InodePercent:                      fmt.Sprintf("%f2", InodePercent),
		}
		volList = append(volList, vol)
	}

	return volList, nil
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

	m, err := c.GetVolumeToAggrMap()
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
