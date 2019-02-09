# go-ontap
Golang module to access Netapp ontap API and pull data

Sample usage:
```go
package main

import (
	"fmt"
	"github.com/marstid/go-ontap"
)

func main() {

	client := ontap.NewClient("hostname.tld", "ontap-userid", "password", true)
	client.Debug = false

	// Print ONTAP version
	version, err := client.GetSystemVersion()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(client.Host + " running " + version + "\n\n")
	}

	// Get information about volumes
	volumes, err := client.GetVolumeInfo(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client.Host + " has following Volumes:")
	for _, v := range volumes {
		fmt.Println(v.Name)
	}

	// Get information about disks
	disks, err := client.GetDiskInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println(client.Host + " has following Disks:")
	for _, v := range disks {
		fmt.Println("Name: " + v.Name + ", Type: " + v.DiskType + ", Model: " + v.Model)
	}


}

```