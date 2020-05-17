package ontap

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

var url string
var user string
var pw string

func init() {
	user = os.Getenv("NETAPP_UID")
	pw = os.Getenv("NETAPP_PW")
	url = os.Getenv("NETAPP_URL")

	if user == "" || pw == "" || url == "" {
		fmt.Println("Env must be set")
		os.Exit(1)
	}
}

func TestClient(t *testing.T) {
	client := NewClient(url, user, pw, true)
	client.Debug = false

	ver, err := client.GetSystemVersion()
	if err != nil {
		t.Error(err)
	}
	if ver == "" {
		t.Error("GetSystemVersion returned empty string")
	}
}

func TestLun(t *testing.T) {
	client := NewClient(url, user, pw, true)
	client.Debug = false

	_, err := client.GetLunInfo(10)
	if err != nil {
		t.Error(err)
	}
	/*
		for _, re := range res {
			fmt.Printf("Name: %s, Space: %t\n", re.Volume, re.IsSpaceAllocEnabled)
		}
	*/
}

/*
func TestFC(t *testing.T) {
	client := NewClient(url, user, pw, true)
	client.Debug = true

	data, err := client.GetFCConfig()
	if err != nil {
		t.Error(err)
	}

	if data == nil {
		t.Error(err)
	}

	data2, err := client.GetFCPAdapter()
	if err != nil {
		t.Error(err)
	}

	if data2 == nil {
		t.Error(err)
	}

	fmt.Println(data2[0].StatusAdmin)

}
*/

/*
func TestPFD(t *testing.T){
	client := NewClient(url, user, pw, true)

	pf, err := client.GetPerfObject()
	if err != nil {
		t.Error(err)
	}
	for _, value := range pf {
		if strings.Contains(value.Name, "aggr"){
			fmt.Printf("%s\n%s\n\n",value.Name, value.Description)
		}

	}

}
*/

/*
func TestPFD(t *testing.T) {
	client := NewClient(url, user, pw, true)

	pf, err := client.GetPerfCounters("aggregate")
	if err != nil {
		t.Error(err)
	}
	for _, value := range pf {

		fmt.Printf("%s\n%s\n%s\n\n", value.Name, value.Desc, value.Unit)

	}

}
*/

func TestVolInf(t *testing.T) {
	client := NewClient(url, user, pw, true)
	client.Debug = false
	vl, err := client.GetVolumeInfo(4)
	if err != nil {
		t.Error(err)
	}

	for _, value := range vl {

		fmt.Printf("Name: %s\n", value.Name)
		fmt.Printf("Aggregate: %s\n", value.Aggr)
		fmt.Printf("Dedup: %s\n", value.DeduplicationPercentageSpaceSaved)
		fmt.Printf("Comp: %s\n", value.CompressionPercentageSpaceSaved)
		fmt.Printf("P space saved: %s\n", value.TotalPercentageSpaceSaved)
		i, _ := strconv.Atoi(value.TotalSpaceSaved)
		fmt.Printf("Space saved: %d GB\n", i/1024/1024/1024)
		fmt.Println()
	}
}

func TestAggrInf(t *testing.T) {
	client := NewClient(url, user, pw, true)
	client.Debug = false
	vl, err := client.GetAggrInfo(4)
	if err != nil {
		t.Error(err)
	}

	for _, value := range vl {

		fmt.Printf("Name: %s\n", value.Name)
		fmt.Printf("P space saved: %s\n", value.SisSpaceSavedPercent)
		i, _ := strconv.Atoi(value.SisSpaceSaved)
		fmt.Printf("Space saved: %d GB\n", i/1024/1024/1024)
		fmt.Printf("CP space saved: %s\n", value.DataCompactionSpaceSavedPercent)
		j, _ := strconv.Atoi(value.DataCompactionSpaceSaved)
		fmt.Printf("Compacting Space saved: %d GB\n", j/1024/1024/1024)
		fmt.Println()
	}
}
