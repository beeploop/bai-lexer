package scanner

func (S *Scanner) Advance(source string) byte {
	S.Current++
	return source[S.Current-1]
}
