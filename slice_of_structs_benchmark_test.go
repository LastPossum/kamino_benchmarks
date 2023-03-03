package kamino_benchmark

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	barkimedes "github.com/barkimedes/go-deepcopy"
	mohae "github.com/mohae/deepcopy"

	"github.com/LastPossum/kamino"
)

type simpleStruct struct {
	Int     int
	Float64 float64
	String  string
}

func genSliceOfStruct(lenStructSlice int) []simpleStruct {
	bigStructSlice := make([]simpleStruct, lenStructSlice)
	for i := 0; i < lenStructSlice; i++ {
		bigStructSlice[i].Int = i
		bigStructSlice[i].Float64 = float64(i)
		bigStructSlice[i].String = strconv.Itoa(i)
	}
	return bigStructSlice
}

func BenchmarkStructsSlice(b *testing.B) {
	for i := 5; i < 10; i++ {
		k := 1 << i
		structs := genSliceOfStruct(k)

		b.Run(fmt.Sprintf("barkimedes for %d structs slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				barkimedes.Anything(structs)
			}
		})

		b.Run(fmt.Sprintf("mohae for %d structs slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mohae.Copy(structs)
			}
		})

		b.Run(fmt.Sprintf("json for %d structs slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneJSON(structs)
			}
		})

		b.Run(fmt.Sprintf("msgpack for %d structs slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cloneMsgPack(structs)
			}
		})

		b.Run(fmt.Sprintf("kamino for %d structs slice", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kamino.Clone(structs)
			}
		})
	}
}

func BenchmarkKaminoTime(b *testing.B) {
	l, _ := time.LoadLocation("Europe/Moscow")
	t := time.Now().In(l)
	for i := 0; i < b.N; i++ {
		kamino.Clone(t)
	}
}

func BenchmarkKaminoFullTime(b *testing.B) {
	l, _ := time.LoadLocation("Europe/Moscow")
	t := time.Now().In(l)
	for i := 0; i < b.N; i++ {
		kamino.Clone(t, kamino.WithForceUnexported())
	}
}
