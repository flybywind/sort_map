package main

import (
	"fmt"
	"math/rand"
	"sort_map"
	"testing"
)

func makeRandStr(r *rand.Rand, dict []byte, word_len int) string {
	max_rand := len(dict)
	ret := make([]byte, word_len)
	for i := 0; i < word_len; i++ {
		ret[i] = dict[r.Intn(max_rand)]
	}
	return string(ret)
}
func TestStringKey(t *testing.T) {
	src := rand.NewSource(123456)
	r := rand.New(src)
	dict := []byte("abcxyz")
	const (
		key_len = 4
		arr_num = 10
	)
	test_map := map[string]int{}
	for i := 0; i < arr_num; i++ {
		key := makeRandStr(r, dict, key_len)
		test_map[key] = r.Intn(1000)
	}

	by := func(p1, p2 *sort_map.MapPair) bool {
		return !p1.KeyCmp(p2)
	}

	ps := sort_map.NewPairSlice(test_map)
	sort_map.By(by).Sort(ps)

	fmt.Printf("%s\n", ps.String())
}

func TestIntKey(t *testing.T) {
	src := rand.NewSource(123456)
	r := rand.New(src)
	dict := []byte("abcxyz")
	const (
		key_len = 4
		arr_num = 10
	)
	test_map := map[string]int{}
	for i := 0; i < arr_num; i++ {
		key := makeRandStr(r, dict, key_len)
		test_map[key] = r.Intn(1000)
	}

	by := func(p1, p2 *sort_map.MapPair) bool {
		return p1.ValueCmp(p2)
	}

	ps := sort_map.NewPairSlice(test_map)
	sort_map.By(by).Sort(ps)

	fmt.Printf("%s\n", ps.String())
}
