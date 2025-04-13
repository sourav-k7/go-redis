package constant

type CommandClass int

const (
	DefaultType CommandClass = iota
	StringType 
)

func (d CommandClass) String() string {
	return [...]string{"Default", "String"}[d]
}

func (d CommandClass) EnumIndex() int {
	return int(d)
}
