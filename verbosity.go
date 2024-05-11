package bslogger

type Verbosity int

const (
	Minimal Verbosity = iota
	Normal
	All
)

func (v Verbosity) String() string {
	return []string{
		"Minimal", "Normal", "All",
	}[v]
}
