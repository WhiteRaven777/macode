package macode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlatMap(t *testing.T) {
	m1 := map[string]int{
		"three": 3,
		"one":   1,
		"two":   2,
	}

	p1 := New(m1)
	if m1e := p1.Encode(m1); fmt.Sprint(m1e) != "[{one 1} {three 3} {two 2}]" {
		t.Errorf("Encode: %+x", m1e)
	} else if m1d := p1.Decode(m1e); !reflect.DeepEqual(m1d, m1) {
		t.Errorf("Decode: %+x", m1d)
	}

	// ---

	m2 := map[rune]int{
		'c': 3,
		'a': 1,
		'b': 2,
	}

	p2 := New(m2)
	if m2e := p2.Encode(m2); fmt.Sprint(m2e) != "[{97 1} {98 2} {99 3}]" {
		t.Errorf("Encode: %+x", m2e)
	} else if m2d := p2.Decode(m2e); !reflect.DeepEqual(m2d, m2) {
		t.Errorf("Decode: %+x", m2d)
	}

	// ---

	m3 := map[string]int{
		string([]byte{3}): 3,
		string([]byte{1}): 1,
		string([]byte{2}): 2,
	}

	p3 := New(m3)
	if m3e := p3.Encode(m3); fmt.Sprint(m3e) != "[{\u0001 1} {\u0002 2} {\u0003 3}]" {
		t.Errorf("Encode: %+x", m3e)
	} else if m3d := p3.Decode(m3e); !reflect.DeepEqual(m3d, m3) {
		t.Errorf("Decode: %+x", m3d)
	}
}
