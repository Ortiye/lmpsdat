package key

// IsHeader returns true if the Key is an instance of Header or Box.
func IsHeader(k Key) bool {
	if _, ok := k.(*Header); ok {
		return true
	}
	if _, ok := k.(*Box); ok {
		return true
	}
	return false
}

// IsName returns true if a Name exists and is supported by this package.
func IsName(name Name) bool {
	for _, n := range ListNames {
		if n == name {
			return true
		}
	}
	return false
}

type makeKeys map[Name]Key

// MakeKeys returns the Keys instanced with a list of given Names. It may return
// more Keys than expected: it includes the Keys that are required by other
// Keys.
func MakeKeys(names []Name) map[Name]Key {
	m := make(makeKeys, len(names))
	for _, n := range names {
		m.New(n)
	}
	return m
}

func (m makeKeys) New(name Name) Key {
	if v, ok := m[name]; ok {
		return v
	}

	var v Key
	switch name {
	case NamePairCoeffs:
		v = NewCoeffs(name)
		v.SetKeys(m.New(NameAtomTypes))
	case NameBondCoeffs:
		v = NewCoeffs(name)
		v.SetKeys(m.New(NameBondTypes))
	case NameAngleCoeffs:
		v = NewCoeffs(name)
		v.SetKeys(m.New(NameAngleTypes))
	case NameDihedralCoeffs:
		v = NewCoeffs(name)
		v.SetKeys(m.New(NameDihedralTypes))

	case NameAtomsNbr, NameBondsNbr, NameAnglesNbr, NameDihedralsNbr:
		v = NewHeader(name)
	case NameAtomTypes, NameBondTypes, NameAngleTypes, NameDihedralTypes:
		v = NewHeader(name)
	case NameBoxX, NameBoxY, NameBoxZ:
		v = NewBox(name)

	case NameMasses:
		v = new(Masses)
		v.SetKeys(m.New(NameAtomTypes))

	case NameAtoms:
		v = new(Atoms)
		v.SetKeys(m.New(NameAtomTypes),
			m.New(NameAtomsNbr))

	case NameBonds:
		v = NewLinks(name, 2)
		v.SetKeys(m.New(NameAtomsNbr),
			m.New(NameBondsNbr),
			m.New(NameBondTypes))
	case NameAngles:
		v = NewLinks(name, 3)
		v.SetKeys(m.New(NameAtomsNbr),
			m.New(NameAnglesNbr),
			m.New(NameAngleTypes))
	case NameDihedrals:
		v = NewLinks(name, 4)
		v.SetKeys(m.New(NameAtomsNbr),
			m.New(NameDihedralsNbr),
			m.New(NameDihedralTypes))

	case NameTitle:
		v = new(Title)

	default:
		panic("Name provided is not implemented in this function")
	}

	m[name] = v
	return v
}
