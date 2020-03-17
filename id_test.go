package trid

import "testing"

func TestNewID(t *testing.T) {
	id := New()
	t.Log(id)
}

func TestIDToString(t *testing.T) {
	id := New()
	s := id.String()
	t.Log(s)
}

func TestFromBytes(t *testing.T) {
	id := New()
	_, err := FromBytes(id.Bytes())
	if err != nil {
		t.Error(err)
	}
}

func TestFromString(t *testing.T) {
	id := New()
	_, err := FromString(id.String())
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkNew(b *testing.B) {
	var id ID
	for n := 0; n < b.N; n++ {
		id = New()
	}

	b.Log(id)
}

func BenchmarkToString(b *testing.B) {
	var s string
	id := New()

	for n := 0; n < b.N; n++ {
		s = id.String()
	}

	b.Log(s)
}

func BenchmarkFromBytes(b *testing.B) {
	var (
		id  ID
		err error
	)

	idBytes := New().Bytes()

	for n := 0; n < b.N; n++ {
		id, err = FromBytes(idBytes)
		if err != nil {
			b.Error(err)
		}
	}

	b.Log(id)
}

func BenchmarkFromString(b *testing.B) {
	var (
		id  ID
		err error
	)

	idStr := New().String()

	for n := 0; n < b.N; n++ {
		id, err = FromString(idStr)
		if err != nil {
			b.Error(err)
		}
	}

	b.Log(id)
}
