package ontap

type PerfCounter struct {
	ObjectName string
	Counter    string
	Value      string
}

type PerfCounterInfo struct {
	Name       string
	Desc       string
	Unit       string
	Deprecated string
}
