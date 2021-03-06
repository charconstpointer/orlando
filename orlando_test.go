package orlando_test

import (
	"fmt"
	"hash/fnv"
	"testing"

	"github.com/charconstpointer/orlando"
)

func TestNewFilter(t *testing.T) {
	filter, err := orlando.NewFilter(100, func(s string) uint32 {
		return 1
	})
	if filter == nil {
		t.Errorf("Expected filter to be created")
	}
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestFilterInsert(t *testing.T) {
	filter, err := orlando.NewFilter(100, func(s string) uint32 {
		return 1
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if err := filter.Insert("test"); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestFilterContains(t *testing.T) {
	hasher := fnv.New32a()
	filter, err := orlando.NewFilter(100, func(s string) uint32 {
		defer hasher.Reset()
		hasher.Write([]byte(s))
		return hasher.Sum32()
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if err := filter.Insert("test"); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !filter.Contains("test") {
		t.Errorf("Expected filter to contain 'test'")
	}
	if filter.Contains("test2") {
		t.Errorf("Expected filter to not contain 'test2'")
	}
}

func BenchmarkFilterInsertContains10(b *testing.B) {
	benchmarkFilterInsertContains(10)
}

func BenchmarkFilterInsertContains100(b *testing.B) {
	benchmarkFilterInsertContains(100)
}

func BenchmarkFilterInsertContains1000(b *testing.B) {
	benchmarkFilterInsertContains(100_0)
}

func BenchmarkFilterInsertContains10000(b *testing.B) {
	benchmarkFilterInsertContains(10_000)
}

func benchmarkFilterInsertContains(i int) {
	hasher := fnv.New32a()
	filter, _ := orlando.NewFilter(100, func(s string) uint32 {
		defer hasher.Reset()
		hasher.Write([]byte(s))
		return hasher.Sum32()
	})
	for j := 0; j < i; j++ {
		filter.Insert(fmt.Sprintf("%d", j))
		filter.Contains(fmt.Sprintf("%d", j))
	}
}
