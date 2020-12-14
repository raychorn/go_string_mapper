package mapper

import (
    "unicode"
)

// Interface ...
type Interface interface {
    TransformRune(pos int)
    GetValueAsRuneSlice() []rune
}

// StringSkipper ...
type StringSkipper struct {
    num int
    str string
}

// GetValueAsRuneSlice ...
func (s StringSkipper) GetValueAsRuneSlice() []rune {
    return []rune(s.str)
}

// TransformRune ...
func (s *StringSkipper) TransformRune(pos int) {
    b := []byte(s.str)
    c := unicode.ToLower(rune(s.str[pos]))
    if ((pos+1) % s.num) == 0 {
        c = unicode.ToUpper(rune(c))
        b[pos] = byte(c)
    }
    s.str = string(b)
}

// MapString ...
func MapString(i Interface) {
    for pos := range i.GetValueAsRuneSlice() {
        i.TransformRune(pos)
    }
}

// NewSkipString ...
func NewSkipString(num int, str string) StringSkipper {
    return StringSkipper{num: num, str: str}

}

