package key

import (
	"reflect"
	"testing"
)

// TestKeys checks if the Keys provided by this package implements the Key
// interface. Moreover, it checks each individual methods of the Key interface:
// they must be not equal to nil.
func TestKeys(t *testing.T) {
	var i interface{}

	i = NewBox(NameBoxX)
	if _, ok := i.(Key); !ok {
		t.Error("Box (instanced with NewBox) does not implement Key")
	}

	i = NewCoeffs(NameAngleCoeffs)
	if _, ok := i.(Key); !ok {
		t.Error("Coeffs (instanced with NewCoeffs) does not implement Key")
	}

	i = NewHeader(NameAtomTypes)
	if _, ok := i.(Key); !ok {
		t.Error("Header (instanced with NewHeader) does not implement Key")
	}

	i = NewLinks(NameBonds, 2)
	if _, ok := i.(Key); !ok {
		t.Error("Links (instanced with NewLinks) does not implement Key")
	}

	i = new(Atoms)
	if _, ok := i.(Key); !ok {
		t.Error("Atoms (instanced by `hand`) does not implement Key")
	}

	i = new(Masses)
	if _, ok := i.(Key); !ok {
		t.Error("Masses (instanced by `hand`) does not implement Key")
	}

	i = new(Title)
	if _, ok := i.(Title); !ok {
		t.Error("Title (instanced by `hand`) does not implement Key")
	}
}

func TestNames(t *testing.T) {
	b := NewBox(NameBoxX)
	if b.Name() != NameBoxX {
		t.Errorf("incorrect name. Expected %q", NameBoxX)
	}

	c := NewCoeffs(NameAngleCoeffs)
	if c.Name() != NameAngleCoeffs {
		t.Errorf("incorrect name. Expected %q", NameAngleCoeffs)
	}

	h := NewHeader(NameAtomTypes)
	if h.Name() != NameAtomTypes {
		t.Errorf("incorrect name. Expected %q", NameAtomTypes)
	}

	l := NewLinks(NameBonds, 2)
	if l.Name() != NameBonds {
		t.Errorf("incorrect name. Expected %q", NameBonds)
	}

	a := new(Atoms)
	if a.Name() != NameAtoms {
		t.Errorf("incorrect name. Expected %q", NameAtoms)
	}

	m := new(Masses)
	if m.Name() != NameMasses {
		t.Errorf("incorrect name. Expected %q", NameMasses)
	}

	tt := new(Title)
	if tt.Name() != NameTitle {
		t.Errorf("incorrect name. Expected %q", NameTitle)
	}
}

func TestKeywords(t *testing.T) {
	k := []interface{}{
		NewBox(NameBoxX),
		NewCoeffs(NameAngleCoeffs),
		NewHeader(NameAtomTypes),
		NewLinks(NameBonds, 2),
		new(Atoms),
		new(Masses),
		new(Title),
	}

	qTrue := [][][]byte{
		{
			[]byte("150.0 150.0 xlo xhi"),
			[]byte("   150.0\t\v0 xlo\t xhi"),
		},
		{
			[]byte("Angle Coeffs"),
			[]byte("\t\v Angle Coeffs"),
		},
		{
			[]byte("5000 atom types"),
			[]byte("0 atom\t types"),
		},
		{
			[]byte("\t Bonds"),
			[]byte("Bonds"),
		},
		{
			[]byte("\t Atoms"),
			[]byte("Atoms"),
		},
		{
			[]byte("\t Masses"),
			[]byte("Masses"),
		},
		{},
	}

	qFalse := [][][]byte{
		{
			[]byte("150.0 xlo xhi"),
			[]byte("xlo\t xhi"),
			[]byte(""),
		},
		{
			[]byte("Angle\t Coeffs"),
			[]byte("Angle"),
			[]byte(""),
		},
		{
			[]byte("atom types"),
			[]byte(""),
			[]byte(" atom types"),
		},
		{
			[]byte(""),
			[]byte("0 Bonds"),
		},
		{
			[]byte(""),
			[]byte("0 Atoms"),
		},
		{
			[]byte(""),
			[]byte("0 Masses"),
		},
		{
			[]byte("title"),
		},
	}

	for i, key := range k {
		for _, qT := range qTrue[i] {
			keyword := key.(Key).Keyword(qT)
			if !keyword {
				t.Errorf("error Key %q, got %v instead of %v for %q", key.(Key).Name(), keyword, true, qT)
			}
		}

		for _, qF := range qFalse[i] {
			keyword := key.(Key).Keyword(qF)
			if keyword {
				t.Errorf("error Key %q, got %v instead of %v for %q", key.(Key).Name(), keyword, true, qF)
			}
		}
	}
}

func TestSetGets(t *testing.T) {
	k := []interface{}{
		NewBox(NameBoxX),
		NewCoeffs(NameAngleCoeffs),
		NewHeader(NameAtomTypes),
		NewLinks(NameBonds, 2),
		new(Atoms),
		new(Masses),
		new(Title),
	}

	q := []interface{}{
		[2]float64{},
		make(map[int][]float64),
		50,
		make(map[int]*Link),
		make(map[int]*Atom),
		make(map[int]float64),
		"title",
	}

	for i, key := range k {
		if err := key.(Key).Set(q[i]); err != nil {
			t.Error(err)
		}
		get := key.(Key).Get()
		kindSet := reflect.TypeOf(q[i]).Kind()
		kindGet := reflect.TypeOf(get).Kind()
		if kindSet != kindGet {
			t.Errorf("Get not the same kind as Set (%v vs %v)", kindGet, kindSet)
		}
		if err := key.(Key).Set(nil); err == nil {
			t.Error("got an nil error")
		}
	}
}
