package key

import (
	"bytes"
	"testing"
)

func TestDelComments(t *testing.T) {
	q := [][]byte{
		[]byte("hello # this is a comment"),
		[]byte("hello # this # is a comment"),
		[]byte("# hello # this is a comment"),
		[]byte("hello ##### this is a comment"),
	}

	a := [][]byte{
		[]byte("hello "),
		[]byte("hello "),
		[]byte(""),
		[]byte("hello "),
	}

	for i, s := range q {
		d := delComments(s)
		if !bytes.Equal(d, a[i]) {
			t.Errorf("%q is not equal to %q", d, a[i])
		}
	}
}

func TestSortIntsMap(t *testing.T) {
	m := make(map[int]bool)
	m[5] = true
	m[9] = true
	m[1] = true

	keys := sortIntsMap(m)
	if keys[0] != 1 || keys[1] != 5 || keys[2] != 9 {
		t.Errorf("incorrect sorting, got %v", keys)
	}
}

func TestKeyword(t *testing.T) {
	q := [][]byte{
		[]byte("hello foo bar"),
		[]byte(" hello foo bar"),
		[]byte("\r hello foo bar"),
		[]byte("\n hello foo bar"),
		[]byte("\f hello foo bar"),
		[]byte("\t hello foo bar"),
		[]byte("\v hello foo bar"),
		[]byte("\u00A0 hello foo bar"),
		[]byte("\u0085 hello foo bar"),
		[]byte("\n\t\v\f\r\u00A0\u0085 hello foo bar"),
	}

	for _, s := range q {
		if !keyword(s, []byte("hello")) {
			t.Errorf("%q does not have `hello` prefix", s)
		}
	}

	q = [][]byte{
		[]byte("not hello foo bar"),
		[]byte(" not hello foo bar"),
		[]byte(""),
		[]byte(" "),
	}

	for _, s := range q {
		if keyword(s, []byte("hello")) {
			t.Errorf("%q does have `hello` prefix", s)
		}
	}
}

func TestKeywordHeader(t *testing.T) {
	p := [][]byte{
		[]byte("hello"),
		[]byte("foo"),
	}

	q := [][]byte{
		[]byte("hello foo "),
		[]byte(" hello\t foo"),
		[]byte("\n\t\v\f\r\u00A0\u0085 hello\n\t\v\f\r\u00A0\u0085 foo"),
	}

	for _, s := range q {
		if !keywordHeader(s, p) {
			t.Errorf("%q does not have prefix", s)
		}
	}

	q = [][]byte{
		[]byte("\t hello # foo "),
		[]byte("# hellofoo"),
		[]byte("# hello#foo"),
		[]byte("hello"),
	}

	for _, s := range q {
		if keywordHeader(s, p) {
			t.Errorf("%q does have prefix", s)
		}
	}

	if keywordHeader(q[0], nil) {
		t.Errorf("must be false (without prefix)")
	}
}
