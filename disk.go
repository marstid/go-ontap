package ontap

type DiskInfo struct {
	Name      string
	DiskType  string
	Model     string
	Online    bool
	Prefailed bool
	Spare     bool
}
