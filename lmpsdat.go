package lmpsdat

import (
	"bufio"
	"fmt"
	"reflect"

	"github.com/ortiye/lmpsdat/key"
)

// createNames returns a map that links the Names to the corresponding Keys and
// another map that links the Names to the field identifiers of a structure.
func createNames(typ reflect.Type) (map[key.Name]key.Key, map[key.Name]int) {
	names := make([]key.Name, 0)
	namesFields := make(map[key.Name]int, 0)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		v, ok := f.Tag.Lookup("lmpsdat")
		if !ok {
			continue
		}
		n := key.Name(v)
		if key.IsName(n) {
			namesFields[n] = i
			names = append(names, n)
		}
	}
	return key.MakeKeys(names), namesFields
}

// headBody separate the keys. It reproduces what the LAMMPS data parser does.
func headBody(keys map[key.Name]key.Key) (headers, bodies map[key.Name]key.Key) {
	headers = make(map[key.Name]key.Key)
	bodies = make(map[key.Name]key.Key)
	for n, k := range keys {
		if key.IsHeader(k) {
			headers[n] = k
		} else {
			bodies[n] = k
		}
	}
	return
}

func keyDecode(s []byte, keys map[key.Name]key.Key, r *bufio.Scanner) (bool, error) {
	for n, k := range keys {
		if k.Keyword(s) {
			err := k.Decode(s, r)
			if err != nil {
				return true, fmt.Errorf("k.Decode for Key = %s: %w", k.Name(), err)
			}
			delete(keys, n)
			return true, nil
		}
	}
	return false, nil
}
