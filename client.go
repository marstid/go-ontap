package ontap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
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

// Get slice of EnvironmentSensorInfo
func (c *Client) GetEnvironment() (sensorInfo []EnvironmentSensorsInfo, error error) {

	egs := &EnvRequest{}
	egs.Version = apiVersion
	egs.Xmlns = XMLns

	output, err := xml.MarshalIndent(egs, "", "\t")
	payload := bytes.NewReader(output)

	req, err := http.NewRequest("POST", c.Url, payload)
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var result EnvResults
	err = xml.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}

	return result.Results.AttributesList.EnvironmentSensors, nil
}

func (c *Client) GetNodePerf() ([]PerfCounter, error) {
	// Performance counters of interest
	var counters []string
	counters = append(counters, "cpu_busy")
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
