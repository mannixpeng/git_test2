package main

import (
	"testing"
)

type P struct {
	Age int
}

func TestUnmarshalAndPrint(t *testing.T) {
	statistics := make([]*P, 0)
	for i := 0; i < 100; i++ {
		statistics = append(statistics, &P{Age: i})
	}

	t.Run("testing unmarshalAndPrint()", func(t *testing.T) {
		//err := unmarshalAndPrint(strings.NewReader(`[{"name": "Dubi Gal", "age": 900}]`))
		//assert.Nil(t, err)
		m := make(map[int]int)
		for i, e := range statistics {
			t.Logf("%p", e)
			m[i] = e.Age
		}
		t.Logf("%+v", m)
	})
}
