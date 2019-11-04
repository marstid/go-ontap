package ontap

import (
	"fmt"
	"os"
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

func TestFC(t *testing.T) {
	client := NewClient(url, user, pw, true)
	client.Debug = true

	data, err := client.GetFCConfig()
	if err != nil {
		t.Error(err)
	}

	if data == nil{
		t.Error(err)
	}


	data2, err := client.GetFCPAdapter()
	if err != nil {
		t.Error(err)
	}

	if data2 == nil{
		t.Error(err)
	}

}
