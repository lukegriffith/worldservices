package world

type Status int64

const (
	Pending Status = iota
	Ready
)

func (s Status) String() string {
	switch s {
	case Pending:
		return "pending"
	case Ready:
		return "ready"
	}
	return ""
}
