package sort_map

import (
	"fmt"
	"reflect"
	"sort"
)

type MapPair struct {
	Key, Value reflect.Value
}

func (m *MapPair) KeyCmp(m2 *MapPair) bool {
	switch m.Key.Interface().(type) {
	case int, int8, int16, int32, int64:
		v1 := m.Key.Int()
		v2 := m2.Key.Int()
		return v1 < v2
	case float32, float64:
		v1 := m.Key.Float()
		v2 := m2.Key.Float()
		return v1 < v2
	case uint, uint8, uint16, uint32, uint64:
		v1 := m.Key.Uint()
		v2 := m2.Key.Uint()
		return v1 < v2
	case string:
		v1 := m.Key.String()
		v2 := m2.Key.String()
		return v1 < v2
	default:
		panic("not Comparable")
	}
}
func (m *MapPair) ValueCmp(m2 *MapPair) bool {
	switch m.Value.Interface().(type) {
	case int, int8, int16, int32, int64:
		v1 := m.Value.Int()
		v2 := m2.Value.Int()
		return v1 < v2
	case float32, float64:
		v1 := m.Value.Float()
		v2 := m2.Value.Float()
		return v1 < v2
	case uint, uint8, uint16, uint32, uint64:
		v1 := m.Value.Uint()
		v2 := m2.Value.Uint()
		return v1 < v2
	case string:
		v1 := m.Value.String()
		v2 := m2.Value.String()
		return v1 < v2
	default:
		panic("not Comparable")
	}
}
func (m *MapPair) String() string {
	return fmt.Sprintf("%v -> %v", m.Key, m.Value)
}

type PairSlice []MapPair

func (ps *PairSlice) String() string {
	ret := "["
	for _, e := range *ps {
		ret += fmt.Sprintf("%s, ", e.String())
	}
	return ret + "]"
}

func NewPairSlice(I interface{}) PairSlice {
	V := reflect.ValueOf(I)
	keys := V.MapKeys()
	ps := PairSlice{}
	for _, k := range keys {
		m := MapPair{
			Key:   reflect.Value(k),
			Value: V.MapIndex(k),
		}
		ps = append(ps, m)
	}
	return ps
}

type By func(p1, p2 *MapPair) bool

func (b By) Sort(ps PairSlice) {
	wrapper := &pairSliceWrapper{
		pair_slice: ps,
		by:         b,
	}
	sort.Sort(wrapper)
}

type pairSliceWrapper struct {
	pair_slice PairSlice
	by         By
}

func (ps *pairSliceWrapper) Len() int {
	return len(ps.pair_slice)
}

func (ps *pairSliceWrapper) Swap(i, j int) {
	ps.pair_slice[i], ps.pair_slice[j] = ps.pair_slice[j], ps.pair_slice[i]
}

func (ps *pairSliceWrapper) Less(i, j int) bool {
	return ps.by(&ps.pair_slice[i], &ps.pair_slice[j])
}
