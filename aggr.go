package ontap

type AggrInfo struct {
	Name            string
	SizeTotal       string
	SizeUsed        string
	SizeAvailable   string
	SizeUsedPercent string
	State           string
	Cluster         string
	// Data Compaction
	DataCompactionSpaceSaved        string
	DataCompactionSpaceSavedPercent string
	// Sis
	SisSpaceSaved        string
	SisSpaceSavedPercent string
}
