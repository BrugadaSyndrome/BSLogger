package bslogger

type verbosity int

const (
	Minimal verbosity = iota
	Normal
	All
)

func (v verbosity) String() string {
	return []string{
		"Minimal", "Normal", "All",
	}[v]
}
