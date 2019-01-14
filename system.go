package ontap

type System struct{}

type SystemGetVersion struct {
	Base
	Version struct {} `xml:"system-get-version"`
}

func NewSystemGetVersion() *SystemGetVersion {
	sgv := &SystemGetVersion{}
	sgv.Base.Version = apiVersion
	sgv.Base.XMLns = XMLns
	return sgv
}

type SystemGetVersionResult struct {
	Base
	Results struct {
		Status string `xml:"status,attr"`
		Clustered bool `xml:"is-clustered"`
		Version string `xml:"version"`
	} `xml:"results"`
}
