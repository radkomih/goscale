package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeString(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       string
		expectation []byte
	}{
		{
			label:       "(abc)",
			input:       "abc",
			expectation: []byte{0x0c, 0x61, 0x62, 0x63},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			SliceU8ToSequenceU8(StringToSliceU8(testExample.input)).Encode(buffer)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeString(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation string
	}{
		{
			label:       "(0x0c616263)",
			input:       []byte{0x0c, 0x61, 0x62, 0x63},
			expectation: "abc",
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			SliceU8ToSequenceU8(StringToSliceU8(string(testExample.expectation))).Encode(buffer)

			result := SliceU8ToString(DecodeSequenceU8(buffer).Values)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeU8Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[U8]
		expectation []byte
	}{
		{
			label:       "(0x616263)",
			input:       Sequence[U8]{Values: []U8{0x61, 0x62, 0x63}},
			expectation: []byte{0x0c, 0x61, 0x62, 0x63},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeU8Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation Sequence[U8]
	}{
		{
			label:       "(0x616263)",
			input:       []byte{0x0c, 0x61, 0x62, 0x63},
			expectation: Sequence[U8]{Values: []U8{0x61, 0x62, 0x63}},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.expectation.Encode(buffer)

			result := DecodeSequenceU8(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeBoolSequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[Bool]
		expectation []byte
	}{
		{
			label:       "([false,true,true])",
			input:       Sequence[Bool]{Values: []Bool{false, true, true}},
			expectation: []byte{0x0c, 0x00, 0x01, 0x01},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeCompactSequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[Compact]
		expectation []byte
	}{
		{
			label: "()",
			input: Sequence[Compact]{Values: []Compact{42, 63, 64, 65535, 1073741823}},
			expectation: []byte{
				0x14,
				0xa8,
				0xfc,
				0x01, 0x01,
				0xfe, 0xff, 0x03, 0x00,
				0xfe, 0xff, 0xff, 0xff,
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeI8Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[I8]
		expectation []byte
	}{
		{
			label:       "(0x616263)",
			input:       Sequence[I8]{Values: []I8{0x61, 0x62, 0x63}},
			expectation: []byte{0x0c, 0x61, 0x62, 0x63},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeI16Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[I16]
		expectation []byte
	}{
		{
			label:       "[]int16{0,1,-1,2,-2,3,-3}",
			input:       Sequence[I16]{Values: []I16{0, 1, -1, 2, -2, 3, -3}},
			expectation: []byte{0x1c, 0x00, 0x00, 0x01, 0x00, 0xff, 0xff, 0x02, 0x00, 0xfe, 0xff, 0x03, 0x00, 0xfd, 0xff},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeU16Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[U16]
		expectation []byte
	}{
		{
			label:       "([4,8,15,16,23,42])",
			input:       Sequence[U16]{Values: []U16{4, 8, 15, 16, 23, 42}},
			expectation: []byte{0x18, 0x04, 0x00, 0x08, 0x00, 0x0f, 0x00, 0x10, 0x00, 0x17, 0x00, 0x2a, 0x00},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeNestedSequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[Sequence[U8]]
		expectation []byte
	}{
		{
			label: "()",
			input: Sequence[Sequence[U8]]{Values: []Sequence[U8]{
				{Values: []U8{0x33, 0x55}},
				{Values: []U8{0x77, 0x99}},
			}},
			expectation: []byte{0x08, 0x08, 0x33, 0x55, 0x08, 0x77, 0x99},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

// TODO check for better string SCALE representation
// func Test_EncodeStringSequence(t *testing.T) {
// 	var testExamples = []struct {
// 		label       string
// 		input       Sequence[string]
// 		expectation []byte
// 	}{
// 		{
// 			label: "([\"a1\",\"b2\",\"c3\"])",
// 			input: Sequence[string]{Values: []string{"a1", "b2", "c3"}},
// 			expectation: []byte{
// 				0x0c,
// 				0x08, 0x61, 0x31,
// 				0x08, 0x62, 0x32,
// 				0x08, 0x63, 0x33,
// 			},
// 		},
// 		{
// 			label: "([\"Hamlet\",\"Война и мир\",\"三国演义\",\"أَلْف لَيْلَة وَلَيْلَة\u200e\"])",
// 			input: []string{"Hamlet", "Война и мир", "三国演义", "أَلْف لَيْلَة وَلَيْلَة\u200e"},
// 			expectation: []byte{
// 				0x10,
// 				0x18, 0x48, 0x61, 0x6d, 0x6c, 0x65, 0x74,
// 				0x50, 0xd0, 0x92, 0xd0, 0xbe, 0xd0, 0xb9, 0xd0, 0xbd, 0xd0, 0xb0, 0x20, 0xd0, 0xb8, 0x20, 0xd0, 0xbc, 0xd0, 0xb8, 0xd1, 0x80,
// 				0x30, 0xe4, 0xb8, 0x89, 0xe5, 0x9b, 0xbd, 0xe6, 0xbc, 0x94, 0xe4, 0xb9, 0x89,
// 				0xbc, 0xd8, 0xa3, 0xd9, 0x8e, 0xd9, 0x84, 0xd9, 0x92, 0xd9, 0x81, 0x20, 0xd9, 0x84, 0xd9, 0x8e, 0xd9, 0x8a, 0xd9, 0x92, 0xd9, 0x84, 0xd9, 0x8e, 0xd8, 0xa9, 0x20, 0xd9, 0x88, 0xd9, 0x8e, 0xd9, 0x84, 0xd9, 0x8e, 0xd9, 0x8a, 0xd9, 0x92, 0xd9, 0x84, 0xd9, 0x8e, 0xd8, 0xa9, 0xe2, 0x80, 0x8e,
// 			},
// 		},
// 	}

// 	for _, testExample := range testExamples {
// 		t.Run(testExample.label, func(t *testing.T) {
// 			buffer := &bytes.Buffer{}

// 			testExample.input.Encode(buffer)

// 			assertEqual(t, buffer.Bytes(), testExample.expectation)
// 		})
// 	}
// }