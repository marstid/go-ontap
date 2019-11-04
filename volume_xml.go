package ontap

import "encoding/xml"

type NetappVolume struct {
	XMLName xml.Name `xml:"netapp"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Results struct {
		Text           string `xml:",chardata"`
		Status         string `xml:"status,attr"`
		AttributesList struct {
			Text             string `xml:",chardata"`
			VolumeAttributes []struct {
				Text                     string `xml:",chardata"`
				Encrypt                  string `xml:"encrypt"`
				VolumeAutosizeAttributes struct {
					Text                   string `xml:",chardata"`
					GrowThresholdPercent   string `xml:"grow-threshold-percent"`
					IsEnabled              string `xml:"is-enabled"`
					MaximumSize            string `xml:"maximum-size"`
					MinimumSize            string `xml:"minimum-size"`
					Mode                   string `xml:"mode"`
					ShrinkThresholdPercent string `xml:"shrink-threshold-percent"`
				} `xml:"volume-autosize-attributes"`
				VolumeCompAggrAttributes struct {
					Text          string `xml:",chardata"`
					TieringPolicy string `xml:"tiering-policy"`
				} `xml:"volume-comp-aggr-attributes"`
				VolumeDirectoryAttributes struct {
					Text       string `xml:",chardata"`
					I2pEnabled string `xml:"i2p-enabled"`
					MaxDirSize string `xml:"max-dir-size"`
					RootDirGen string `xml:"root-dir-gen"`
				} `xml:"volume-directory-attributes"`
				VolumeHybridCacheAttributes struct {
					Text        string `xml:",chardata"`
					Eligibility string `xml:"eligibility"`
				} `xml:"volume-hybrid-cache-attributes"`
				VolumeIDAttributes struct {
					Text     string `xml:",chardata"`
					AggrList struct {
						Text     string `xml:",chardata"`
						AggrName string `xml:"aggr-name"`
					} `xml:"aggr-list"`
					ContainingAggregateName string `xml:"containing-aggregate-name"`
					ContainingAggregateUuid string `xml:"containing-aggregate-uuid"`
					CreationTime            string `xml:"creation-time"`
					ExtentSize              string `xml:"extent-size"`
					Fsid                    string `xml:"fsid"`
					InstanceUuid            string `xml:"instance-uuid"`
					Name                    string `xml:"name"`
					NameOrdinal             string `xml:"name-ordinal"`
					Node                    string `xml:"node"`
					Nodes                   struct {
						Text     string `xml:",chardata"`
						NodeName string `xml:"node-name"`
					} `xml:"nodes"`
					OwningVserverName string `xml:"owning-vserver-name"`
					OwningVserverUuid string `xml:"owning-vserver-uuid"`
					ProvenanceUuid    string `xml:"provenance-uuid"`
					Style             string `xml:"style"`
					StyleExtended     string `xml:"style-extended"`
					Type              string `xml:"type"`
					Uuid              string `xml:"uuid"`
				} `xml:"volume-id-attributes"`
				VolumeInodeAttributes struct {
					Text                     string `xml:",chardata"`
					BlockType                string `xml:"block-type"`
					FilesPrivateUsed         string `xml:"files-private-used"`
					FilesTotal               string `xml:"files-total"`
					FilesUsed                string `xml:"files-used"`
					InodefilePrivateCapacity string `xml:"inodefile-private-capacity"`
					InodefilePublicCapacity  string `xml:"inodefile-public-capacity"`
					InofileVersion           string `xml:"inofile-version"`
				} `xml:"volume-inode-attributes"`
				VolumeLanguageAttributes struct {
					Text                  string `xml:",chardata"`
					IsConvertUcodeEnabled string `xml:"is-convert-ucode-enabled"`
					IsCreateUcodeEnabled  string `xml:"is-create-ucode-enabled"`
					Language              string `xml:"language"`
					NfsCharacterSet       string `xml:"nfs-character-set"`
					OemCharacterSet       string `xml:"oem-character-set"`
				} `xml:"volume-language-attributes"`
				VolumeMirrorAttributes struct {
					Text                     string `xml:",chardata"`
					IsDataProtectionMirror   string `xml:"is-data-protection-mirror"`
					IsLoadSharingMirror      string `xml:"is-load-sharing-mirror"`
					IsMoveMirror             string `xml:"is-move-mirror"`
					IsReplicaVolume          string `xml:"is-replica-volume"`
					IsSnapmirrorSource       string `xml:"is-snapmirror-source"`
					MirrorTransferInProgress string `xml:"mirror-transfer-in-progress"`
					RedirectSnapshotID       string `xml:"redirect-snapshot-id"`
				} `xml:"volume-mirror-attributes"`
				VolumePerformanceAttributes struct {
					Text                      string `xml:",chardata"`
					ExtentEnabled             string `xml:"extent-enabled"`
					FcDelegsEnabled           string `xml:"fc-delegs-enabled"`
					IsAtimeUpdateEnabled      string `xml:"is-atime-update-enabled"`
					MaxWriteAllocBlocks       string `xml:"max-write-alloc-blocks"`
					MinimalReadAhead          string `xml:"minimal-read-ahead"`
					ReadRealloc               string `xml:"read-realloc"`
					SingleInstanceDataLogging string `xml:"single-instance-data-logging"`
				} `xml:"volume-performance-attributes"`
				VolumeSecurityAttributes struct {
					Text                         string `xml:",chardata"`
					VolumeSecurityUnixAttributes struct {
						Text        string `xml:",chardata"`
						Permissions string `xml:"permissions"`
					} `xml:"volume-security-unix-attributes"`
				} `xml:"volume-security-attributes"`
				VolumeSisAttributes struct {
					Text                              string `xml:",chardata"`
					CompressionSpaceSaved             string `xml:"compression-space-saved"`
					DeduplicationSpaceSaved           string `xml:"deduplication-space-saved"`
					DeduplicationSpaceShared          string `xml:"deduplication-space-shared"`
					IsSisLoggingEnabled               string `xml:"is-sis-logging-enabled"`
					IsSisStateEnabled                 string `xml:"is-sis-state-enabled"`
					IsSisVolume                       string `xml:"is-sis-volume"`
					PercentageCompressionSpaceSaved   string `xml:"percentage-compression-space-saved"`
					PercentageDeduplicationSpaceSaved string `xml:"percentage-deduplication-space-saved"`
					PercentageTotalSpaceSaved         string `xml:"percentage-total-space-saved"`
					TotalSpaceSaved                   string `xml:"total-space-saved"`
				} `xml:"volume-sis-attributes"`
				VolumeSnaplockAttributes struct {
					Text         string `xml:",chardata"`
					SnaplockType string `xml:"snaplock-type"`
				} `xml:"volume-snaplock-attributes"`
				VolumeSnapshotAttributes struct {
					Text                           string `xml:",chardata"`
					SnapdirAccessEnabled           string `xml:"snapdir-access-enabled"`
					SnapshotCloneDependencyEnabled string `xml:"snapshot-clone-dependency-enabled"`
					SnapshotCount                  string `xml:"snapshot-count"`
				} `xml:"volume-snapshot-attributes"`
				VolumeSnapshotAutodeleteAttributes struct {
					Text                string `xml:",chardata"`
					Commitment          string `xml:"commitment"`
					DeferDelete         string `xml:"defer-delete"`
					DeleteOrder         string `xml:"delete-order"`
					DestroyList         string `xml:"destroy-list"`
					IsAutodeleteEnabled string `xml:"is-autodelete-enabled"`
					Prefix              string `xml:"prefix"`
					TargetFreeSpace     string `xml:"target-free-space"`
					Trigger             string `xml:"trigger"`
				} `xml:"volume-snapshot-autodelete-attributes"`
				VolumeSpaceAttributes struct {
					Text                            string `xml:",chardata"`
					ExpectedAvailable               string `xml:"expected-available"`
					FilesystemSize                  string `xml:"filesystem-size"`
					IsFilesysSizeFixed              string `xml:"is-filesys-size-fixed"`
					IsSpaceGuaranteeEnabled         string `xml:"is-space-guarantee-enabled"`
					IsSpaceSloEnabled               string `xml:"is-space-slo-enabled"`
					OverProvisioned                 string `xml:"over-provisioned"`
					OverwriteReserve                string `xml:"overwrite-reserve"`
					OverwriteReserveRequired        string `xml:"overwrite-reserve-required"`
					OverwriteReserveUsed            string `xml:"overwrite-reserve-used"`
					OverwriteReserveUsedActual      string `xml:"overwrite-reserve-used-actual"`
					PercentageFractionalReserve     string `xml:"percentage-fractional-reserve"`
					PercentageSizeUsed              string `xml:"percentage-size-used"`
					PercentageSnapshotReserve       string `xml:"percentage-snapshot-reserve"`
					PercentageSnapshotReserveUsed   string `xml:"percentage-snapshot-reserve-used"`
					PhysicalUsed                    string `xml:"physical-used"`
					PhysicalUsedPercent             string `xml:"physical-used-percent"`
					Size                            string `xml:"size"`
					SizeAvailable                   string `xml:"size-available"`
					SizeAvailableForSnapshots       string `xml:"size-available-for-snapshots"`
					SizeTotal                       string `xml:"size-total"`
					SizeUsed                        string `xml:"size-used"`
					SizeUsedBySnapshots             string `xml:"size-used-by-snapshots"`
					SnapshotReserveAvailable        string `xml:"snapshot-reserve-available"`
					SnapshotReserveSize             string `xml:"snapshot-reserve-size"`
					SpaceFullThresholdPercent       string `xml:"space-full-threshold-percent"`
					SpaceGuarantee                  string `xml:"space-guarantee"`
					SpaceMgmtOptionTryFirst         string `xml:"space-mgmt-option-try-first"`
					SpaceNearlyFullThresholdPercent string `xml:"space-nearly-full-threshold-percent"`
					SpaceSlo                        string `xml:"space-slo"`
				} `xml:"volume-space-attributes"`
				VolumeStateAttributes struct {
					Text                      string `xml:",chardata"`
					BecomeNodeRootAfterReboot string `xml:"become-node-root-after-reboot"`
					ForceNvfailOnDr           string `xml:"force-nvfail-on-dr"`
					IgnoreInconsistent        string `xml:"ignore-inconsistent"`
					InNvfailedState           string `xml:"in-nvfailed-state"`
					IsClusterVolume           string `xml:"is-cluster-volume"`
					IsConstituent             string `xml:"is-constituent"`
					IsFlexgroup               string `xml:"is-flexgroup"`
					IsInconsistent            string `xml:"is-inconsistent"`
					IsInvalid                 string `xml:"is-invalid"`
					IsNodeRoot                string `xml:"is-node-root"`
					IsNvfailEnabled           string `xml:"is-nvfail-enabled"`
					IsQuiescedInMemory        string `xml:"is-quiesced-in-memory"`
					IsQuiescedOnDisk          string `xml:"is-quiesced-on-disk"`
					IsUnrecoverable           string `xml:"is-unrecoverable"`
					State                     string `xml:"state"`
				} `xml:"volume-state-attributes"`
				VolumeTransitionAttributes struct {
					Text                  string `xml:",chardata"`
					IsCftPrecommit        string `xml:"is-cft-precommit"`
					IsCopiedForTransition string `xml:"is-copied-for-transition"`
					IsTransitioned        string `xml:"is-transitioned"`
					TransitionBehavior    string `xml:"transition-behavior"`
				} `xml:"volume-transition-attributes"`
			} `xml:"volume-attributes"`
		} `xml:"attributes-list"`
		NextTag struct {
			Text               string `xml:",chardata"`
			VolumeGetIterKeyTd struct {
				Text string `xml:",chardata"`
				Key0 string `xml:"key-0"`
				Key1 string `xml:"key-1"`
			} `xml:"volume-get-iter-key-td"`
		} `xml:"next-tag"`
		NumRecords string `xml:"num-records"`
	} `xml:"results"`
}

type VolumeGetIterShort struct {
	XMLName       xml.Name `xml:"netapp"`
	Text          string   `xml:",chardata"`
	Version       string   `xml:"version,attr"`
	Xmlns         string   `xml:"xmlns,attr"`
	VolumeGetIter struct {
		Text              string `xml:",chardata"`
		MaxRecords        string `xml:"max-records"`
		DesiredAttributes struct {
			VolumeAttributes struct {
				Text                  string `xml:",chardata"`
				VolumeIDAttributes    string `xml:"volume-id-attributes"`
				VolumeSpaceAttributes string `xml:"volume-space-attributes"`
				VolumeStateAttributes string `xml:"volume-state-attributes"`
			} `xml:"volume-attributes,omitempty"`
		} `xml:"desired-attributes"`
	} `xml:"volume-get-iter"`
}

type VolumeGetIterFull struct {
	XMLName       xml.Name `xml:"netapp"`
	Text          string   `xml:",chardata"`
	Version       string   `xml:"version,attr"`
	Xmlns         string   `xml:"xmlns,attr"`
	VolumeGetIter struct {
		Text              string `xml:",chardata"`
		MaxRecords        string `xml:"max-records"`
		DesiredAttributes struct {
			Text string `xml:",chardata"`
		} `xml:"desired-attributes"`
	} `xml:"volume-get-iter"`
}
