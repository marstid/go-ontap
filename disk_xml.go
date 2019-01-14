package ontap

import "encoding/xml"

type DiskInfoRequestName struct {
	XMLName            xml.Name `xml:"netapp"`
	Text               string   `xml:",chardata"`
	Version            string   `xml:"version,attr"`
	Xmlns              string   `xml:"xmlns,attr"`
	StorageDiskGetIter struct {
		Text              string `xml:",chardata"`
		DesiredAttributes struct {
			Text     string `xml:",chardata"`
			DiskName string `xml:"disk-name"`
		} `xml:"desired-attributes"`
	} `xml:"storage-disk-get-iter"`
}

type DiskInfoRequest struct {
	XMLName            xml.Name `xml:"netapp"`
	Text               string   `xml:",chardata"`
	Version            string   `xml:"version,attr"`
	Xmlns              string   `xml:"xmlns,attr"`
	StorageDiskGetIter struct {
		Text string `xml:",chardata"`
	} `xml:"storage-disk-get-iter"`
}

type DiskInfoResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text           string `xml:",chardata"`
		Status         string `xml:"status,attr"`
		AttributesList struct {
			Text            string `xml:",chardata"`
			StorageDiskInfo []struct {
				Text              string `xml:",chardata"`
				DiskInventoryInfo struct {
					Text                           string `xml:",chardata"`
					BytesPerSector                 string `xml:"bytes-per-sector"`
					CapacitySectors                string `xml:"capacity-sectors"`
					ChecksumCompatibility          string `xml:"checksum-compatibility"`
					DiskClass                      string `xml:"disk-class"`
					DiskClusterName                string `xml:"disk-cluster-name"`
					DiskType                       string `xml:"disk-type"`
					DiskUid                        string `xml:"disk-uid"`
					FirmwareRevision               string `xml:"firmware-revision"`
					GrownDefectListCount           string `xml:"grown-defect-list-count"`
					HealthMonitorTimeInterval      string `xml:"health-monitor-time-interval"`
					ImportInProgress               string `xml:"import-in-progress"`
					IsDynamicallyQualified         string `xml:"is-dynamically-qualified"`
					IsMultidiskCarrier             string `xml:"is-multidisk-carrier"`
					IsShared                       string `xml:"is-shared"`
					MediaScrubCount                string `xml:"media-scrub-count"`
					MediaScrubLastDoneTimeInterval string `xml:"media-scrub-last-done-time-interval"`
					Model                          string `xml:"model"`
					ReservationKey                 string `xml:"reservation-key"`
					ReservationType                string `xml:"reservation-type"`
					RightSizeSectors               string `xml:"right-size-sectors"`
					Rpm                            string `xml:"rpm"`
					SerialNumber                   string `xml:"serial-number"`
					Shelf                          string `xml:"shelf"`
					ShelfBay                       string `xml:"shelf-bay"`
					ShelfUid                       string `xml:"shelf-uid"`
					StackID                        string `xml:"stack-id"`
					Vendor                         string `xml:"vendor"`
				} `xml:"disk-inventory-info"`
				DiskMetroclusterInfo struct {
					Text          string `xml:",chardata"`
					IsLocalAttach string `xml:"is-local-attach"`
				} `xml:"disk-metrocluster-info"`
				DiskName          string `xml:"disk-name"`
				DiskOwnershipInfo struct {
					Text             string `xml:",chardata"`
					DiskUid          string `xml:"disk-uid"`
					HomeNodeID       string `xml:"home-node-id"`
					HomeNodeName     string `xml:"home-node-name"`
					IsFailed         string `xml:"is-failed"`
					OwnerNodeID      string `xml:"owner-node-id"`
					OwnerNodeName    string `xml:"owner-node-name"`
					Pool             string `xml:"pool"`
					ReservedByNodeID string `xml:"reserved-by-node-id"`
				} `xml:"disk-ownership-info"`
				DiskPaths struct {
					Text         string `xml:",chardata"`
					DiskPathInfo []struct {
						Text                    string `xml:",chardata"`
						ArrayName               string `xml:"array-name"`
						DiskName                string `xml:"disk-name"`
						DiskPort                string `xml:"disk-port"`
						DiskPortName            string `xml:"disk-port-name"`
						DiskUid                 string `xml:"disk-uid"`
						InitiatorIoKbps         string `xml:"initiator-io-kbps"`
						InitiatorIops           string `xml:"initiator-iops"`
						InitiatorLunInUseCount  string `xml:"initiator-lun-in-use-count"`
						InitiatorPort           string `xml:"initiator-port"`
						InitiatorPortSpeed      string `xml:"initiator-port-speed"`
						InitiatorSideSwitchPort string `xml:"initiator-side-switch-port"`
						LunIoKbps               string `xml:"lun-io-kbps"`
						LunIops                 string `xml:"lun-iops"`
						LunNumber               string `xml:"lun-number"`
						LunPathUseState         string `xml:"lun-path-use-state"`
						Node                    string `xml:"node"`
						PathIoKbps              string `xml:"path-io-kbps"`
						PathIops                string `xml:"path-iops"`
						PathLinkErrors          string `xml:"path-link-errors"`
						PathLunInUseCount       string `xml:"path-lun-in-use-count"`
						PathQuality             string `xml:"path-quality"`
						PreferredTargetPort     string `xml:"preferred-target-port"`
						TargetIoKbps            string `xml:"target-io-kbps"`
						TargetIops              string `xml:"target-iops"`
						TargetIqn               string `xml:"target-iqn"`
						TargetLunInUseCount     string `xml:"target-lun-in-use-count"`
						TargetPortAccessState   string `xml:"target-port-access-state"`
						TargetSideSwitchPort    string `xml:"target-side-switch-port"`
						TargetWwnn              string `xml:"target-wwnn"`
						TargetWwpn              string `xml:"target-wwpn"`
						Tpgn                    string `xml:"tpgn"`
					} `xml:"disk-path-info"`
				} `xml:"disk-paths"`
				DiskRaidInfo struct {
					Text              string `xml:",chardata"`
					ActiveNodeName    string `xml:"active-node-name"`
					ContainerType     string `xml:"container-type"`
					DiskAggregateInfo struct {
						Text             string `xml:",chardata"`
						AggregateName    string `xml:"aggregate-name"`
						ChecksumType     string `xml:"checksum-type"`
						IsMediaScrubbing string `xml:"is-media-scrubbing"`
						IsOffline        string `xml:"is-offline"`
						IsPrefailed      string `xml:"is-prefailed"`
						IsReconstructing string `xml:"is-reconstructing"`
						IsReplacing      string `xml:"is-replacing"`
						IsZeroed         string `xml:"is-zeroed"`
						IsZeroing        string `xml:"is-zeroing"`
						PlexName         string `xml:"plex-name"`
						RaidGroupName    string `xml:"raid-group-name"`
					} `xml:"disk-aggregate-info"`
					DiskSharedInfo struct {
						Text             string `xml:",chardata"`
						ChecksumType     string `xml:"checksum-type"`
						IsMediaScrubbing string `xml:"is-media-scrubbing"`
						IsOffline        string `xml:"is-offline"`
						IsPrefailed      string `xml:"is-prefailed"`
						IsReconstructing string `xml:"is-reconstructing"`
						IsReplacing      string `xml:"is-replacing"`
						IsZeroed         string `xml:"is-zeroed"`
						IsZeroing        string `xml:"is-zeroing"`
					} `xml:"disk-shared-info"`
					DiskUid           string `xml:"disk-uid"`
					EffectiveDiskType string `xml:"effective-disk-type"`
					EffectiveRpm      string `xml:"effective-rpm"`
					PhysicalBlocks    string `xml:"physical-blocks"`
					Position          string `xml:"position"`
					SparePool         string `xml:"spare-pool"`
					StandardDiskType  string `xml:"standard-disk-type"`
					UsedBlocks        string `xml:"used-blocks"`
				} `xml:"disk-raid-info"`
				DiskStatsInfo struct {
					Text                string `xml:",chardata"`
					AverageLatency      string `xml:"average-latency"`
					BytesPerSector      string `xml:"bytes-per-sector"`
					DiskIoKbps          string `xml:"disk-io-kbps"`
					DiskIops            string `xml:"disk-iops"`
					DiskUid             string `xml:"disk-uid"`
					PathErrorCount      string `xml:"path-error-count"`
					PowerOnTimeInterval string `xml:"power-on-time-interval"`
					SectorsRead         string `xml:"sectors-read"`
					SectorsWritten      string `xml:"sectors-written"`
				} `xml:"disk-stats-info"`
				DiskUid string `xml:"disk-uid"`
			} `xml:"storage-disk-info"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}
