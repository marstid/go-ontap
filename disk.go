package ontap

type DiskInfo struct {
	Name            string
	DiskType        string
	SizeUsed        string
	SizeAvailable   string
	SizeUsedPercent string
	State           string
	Cluster         string
}
