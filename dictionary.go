package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-dictionary

	SCALE Dictionary type translates to Go's map type.
*/

import (
	"bytes"
	"sort"
)

type Comparable interface {
	Encodable
	Ordered
}

type Dictionary[K Comparable, V Encodable] map[K]V

func (d Dictionary[K, V]) Encode(buffer *bytes.Buffer) {
	ToCompact(len(d)).Encode(buffer)

	keys := make([]K, 0)
	for k := range d {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, k := range keys {
		v := d[k]
		k.Encode(buffer)
		v.Encode(buffer)
	}
}

func (d Dictionary[K, V]) Bytes() []byte {
	return EncodedBytes(d)
}

func DecodeDictionary[K Comparable, V Encodable](buffer *bytes.Buffer) Dictionary[K, V] {
	result := Dictionary[K, V]{}

	v := DecodeCompact(buffer).ToBigInt()
	size := int(v.Int64())

	for i := 0; i < size; i++ {
		key := decodeByType(*new(K), buffer)
		value := decodeByType(*new(V), buffer)
		result[key.(K)] = value.(V)
	}

	return result
}
