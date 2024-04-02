package scanner

type Scanner struct {
	Start   int
	Current int
	Line    int
}

func NewScanner() *Scanner {
	return &Scanner{
		Line: 1,
	}
}
