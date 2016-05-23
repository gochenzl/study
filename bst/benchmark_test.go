package bst

import (
	"math/rand"
	"testing"
	"time"
)

func shuffle(a []int) {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < len(a); i++ {
		j := r.Int() % len(a)
		a[i], a[j] = a[j], a[i]
	}
}

//var benchmarkKeys []int
//var benchmarkKeySize int

//func init() {
//	benchmarkKeySize = 1000000000
//	keys := make([]int, benchmarkKeySize)

//	for i := 0; i < benchmarkKeySize; i++ {
//		keys[i] = i
//	}

//	shuffle(keys)
//}

func BenchmarkUnBalancedInsert(b *testing.B) {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	tree := NewUnbalanced(intCompare)
	for i := 0; i < b.N; i++ {
		tree.Insert(r.Int(), 1)
	}
}

func BenchmarkUnBalancedFind(b *testing.B) {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	tree := NewUnbalanced(intCompare)
	keys := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		key := r.Int()
		tree.Insert(key, 1)
		keys = append(keys, key)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := r.Int() % b.N
		tree.Find(keys[index])
	}
}

func BenchmarkAvlInsert(b *testing.B) {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	tree := NewAvl(intCompare)
	for i := 0; i < b.N; i++ {
		tree.Insert(r.Int(), 1)
	}
}

func BenchmarkAvlFind(b *testing.B) {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	tree := NewAvl(intCompare)
	keys := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		key := r.Int()
		tree.Insert(key, 1)
		keys = append(keys, key)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := r.Int() % b.N
		tree.Find(keys[index])
	}
}

func BenchmarkRBInsert(b *testing.B) {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	tree := NewRB(intCompare)
	for i := 0; i < b.N; i++ {
		tree.Insert(r.Int(), 1)
	}
}

func BenchmarkRBFind(b *testing.B) {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	tree := NewRB(intCompare)
	keys := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		key := r.Int()
		tree.Insert(key, 1)
		keys = append(keys, key)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := r.Int() % b.N
		tree.Find(keys[index])
	}
}
