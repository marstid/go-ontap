package ontap

import "encoding/xml"

type AggrInfoRequest struct {
	XMLName     xml.Name `xml:"netapp"`
	Version     string   `xml:"version,attr"`
	Xmlns       string   `xml:"xmlns,attr"`
	AggrGetIter struct {
		MaxRecords string `xml:"max-records"`
	} `xml:"aggr-get-iter"`
}

type AggrInfoResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text           string `xml:",chardata"`
		Status         string `xml:"status,attr"`
		AttributesList struct {
			Text           string `xml:",chardata"`
			AggrAttributes []struct {
				Text             string `xml:",chardata"`
				AggrFsAttributes struct {
					Text      string `xml:",chardata"`
					BlockType string `xml:"block-type"`
					Fsid      string `xml:"fsid"`
					Type      string `xml:"type"`
				} `xml:"aggr-fs-attributes"`
				AggrInodeAttributes struct {
					Text                     string `xml:",chardata"`
					FilesPrivateUsed         string `xml:"files-private-used"`
					FilesTotal               string `xml:"files-total"`
					FilesUsed                string `xml:"files-used"`
					InodefilePrivateCapacity string `xml:"inodefile-private-capacity"`
					InodefilePublicCapacity  string `xml:"inodefile-public-capacity"`
					InofileVersion           string `xml:"inofile-version"`
					MaxfilesAvailable        string `xml:"maxfiles-available"`
					MaxfilesPossible         string `xml:"maxfiles-possible"`
					MaxfilesUsed             string `xml:"maxfiles-used"`
					PercentInodeUsedCapacity string `xml:"percent-inode-used-capacity"`
				} `xml:"aggr-inode-attributes"`
				AggrOwnershipAttributes struct {
					Text      string `xml:",chardata"`
					Cluster   string `xml:"cluster"`
					HomeID    string `xml:"home-id"`
					HomeName  string `xml:"home-name"`
					OwnerID   string `xml:"owner-id"`
					OwnerName string `xml:"owner-name"`
				} `xml:"aggr-ownership-attributes"`
				AggrPerformanceAttributes struct {
					Text                      string `xml:",chardata"`
					FreeSpaceRealloc          string `xml:"free-space-realloc"`
					MaxWriteAllocBlocks       string `xml:"max-write-alloc-blocks"`
					SingleInstanceDataLogging string `xml:"single-instance-data-logging"`
				} `xml:"aggr-performance-attributes"`
				AggrRaidAttributes struct {
					Text              string `xml:",chardata"`
					AggregateType     string `xml:"aggregate-type"`
					ChecksumStatus    string `xml:"checksum-status"`
					ChecksumStyle     string `xml:"checksum-style"`
					DiskCount         string `xml:"disk-count"`
					HaPolicy          string `xml:"ha-policy"`
					HasLocalRoot      string `xml:"has-local-root"`
					HasPartnerRoot    string `xml:"has-partner-root"`
					IsChecksumEnabled string `xml:"is-checksum-enabled"`
					IsComposite       string `xml:"is-composite"`
					IsEncrypted       string `xml:"is-encrypted"`
					IsHybrid          string `xml:"is-hybrid"`
					IsHybridEnabled   string `xml:"is-hybrid-enabled"`
					IsInconsistent    string `xml:"is-inconsistent"`
					IsMirrored        string `xml:"is-mirrored"`
					IsRootAggregate   string `xml:"is-root-aggregate"`
					MirrorStatus      string `xml:"mirror-status"`
					MountState        string `xml:"mount-state"`
					PlexCount         string `xml:"plex-count"`
					Plexes            struct {
						Text           string `xml:",chardata"`
						PlexAttributes struct {
							Text        string `xml:",chardata"`
							IsOnline    string `xml:"is-online"`
							IsResyncing string `xml:"is-resyncing"`
							PlexName    string `xml:"plex-name"`
							PlexStatus  string `xml:"plex-status"`
							Pool        string `xml:"pool"`
							Raidgroups  struct {
								Text                string `xml:",chardata"`
								RaidgroupAttributes struct {
									Text                        string `xml:",chardata"`
									ChecksumStyle               string `xml:"checksum-style"`
									IsCacheTier                 string `xml:"is-cache-tier"`
									IsRecomputingParity         string `xml:"is-recomputing-parity"`
									IsReconstructing            string `xml:"is-reconstructing"`
									RaidgroupName               string `xml:"raidgroup-name"`
									RecomputingParityPercentage string `xml:"recomputing-parity-percentage"`
									ReconstructionPercentage    string `xml:"reconstruction-percentage"`
								} `xml:"raidgroup-attributes"`
							} `xml:"raidgroups"`
							ResyncingPercentage string `xml:"resyncing-percentage"`
						} `xml:"plex-attributes"`
					} `xml:"plexes"`
					RaidLostWriteState string `xml:"raid-lost-write-state"`
					RaidSize           string `xml:"raid-size"`
					RaidStatus         string `xml:"raid-status"`
					RaidType           string `xml:"raid-type"`
					State              string `xml:"state"`
					UsesSharedDisks    string `xml:"uses-shared-disks"`
				} `xml:"aggr-raid-attributes"`
				AggrSnaplockAttributes struct {
					Text         string `xml:",chardata"`
					IsSnaplock   string `xml:"is-snaplock"`
					SnaplockType string `xml:"snaplock-type"`
				} `xml:"aggr-snaplock-attributes"`
				AggrSnapshotAttributes struct {
					Text                        string `xml:",chardata"`
					FilesTotal                  string `xml:"files-total"`
					FilesUsed                   string `xml:"files-used"`
					IsSnapshotAutoCreateEnabled string `xml:"is-snapshot-auto-create-enabled"`
					IsSnapshotAutoDeleteEnabled string `xml:"is-snapshot-auto-delete-enabled"`
					MaxfilesAvailable           string `xml:"maxfiles-available"`
					MaxfilesPossible            string `xml:"maxfiles-possible"`
					MaxfilesUsed                string `xml:"maxfiles-used"`
					PercentInodeUsedCapacity    string `xml:"percent-inode-used-capacity"`
					PercentUsedCapacity         string `xml:"percent-used-capacity"`
					SizeAvailable               string `xml:"size-available"`
					SizeTotal                   string `xml:"size-total"`
					SizeUsed                    string `xml:"size-used"`
					SnapshotReservePercent      string `xml:"snapshot-reserve-percent"`
				} `xml:"aggr-snapshot-attributes"`
				AggrSpaceAttributes struct {
					Text                            string `xml:",chardata"`
					CapacityTierUsed                string `xml:"capacity-tier-used"`
					DataCompactedCount              string `xml:"data-compacted-count"`
					DataCompactionSpaceSaved        string `xml:"data-compaction-space-saved"`
					DataCompactionSpaceSavedPercent string `xml:"data-compaction-space-saved-percent"`
					HybridCacheSizeTotal            string `xml:"hybrid-cache-size-total"`
					PercentUsedCapacity             string `xml:"percent-used-capacity"`
					PhysicalUsed                    string `xml:"physical-used"`
					PhysicalUsedPercent             string `xml:"physical-used-percent"`
					SisSharedCount                  string `xml:"sis-shared-count"`
					SisSpaceSaved                   string `xml:"sis-space-saved"`
					SisSpaceSavedPercent            string `xml:"sis-space-saved-percent"`
					SizeAvailable                   string `xml:"size-available"`
					SizeTotal                       string `xml:"size-total"`
					SizeUsed                        string `xml:"size-used"`
					TotalReservedSpace              string `xml:"total-reserved-space"`
				} `xml:"aggr-space-attributes"`
				AggrVolumeCountAttributes struct {
					Text         string `xml:",chardata"`
					FlexvolCount string `xml:"flexvol-count"`
				} `xml:"aggr-volume-count-attributes"`
				AggregateName                         string `xml:"aggregate-name"`
				AggregateUuid                         string `xml:"aggregate-uuid"`
				AutobalanceAvailableThresholdPercent  string `xml:"autobalance-available-threshold-percent"`
				AutobalanceState                      string `xml:"autobalance-state"`
				AutobalanceStateChangeCounter         string `xml:"autobalance-state-change-counter"`
				AutobalanceUnbalancedThresholdPercent string `xml:"autobalance-unbalanced-threshold-percent"`
				IsAutobalanceEligible                 string `xml:"is-autobalance-eligible"`
				IsCftPrecommit                        string `xml:"is-cft-precommit"`
				IsObjectStoreAttachEligible           string `xml:"is-object-store-attach-eligible"`
				IsTransitionOutOfSpace                string `xml:"is-transition-out-of-space"`
				Nodes                                 struct {
					Text     string `xml:",chardata"`
					NodeName string `xml:"node-name"`
				} `xml:"nodes"`
			} `xml:"aggr-attributes"`
		} `xml:"attributes-list"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}
