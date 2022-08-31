package go_merkle

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

var defaultHash = func(data ...[]byte) []byte {
	d := []byte{}
	for _, arr := range data {
		d = append(d, arr...)
	}

	return sha256.New().Sum(d)
}

func TestNewTree(t *testing.T) {
	var table = []struct {
		testCaseId   int
		hashF        HashF
		contents     []Content
		expectedRoot []byte
	}{
		{
			testCaseId: 0,
			hashF:      defaultHash,
			contents: []Content{
				&DefaultContent{data: []byte{1, 2, 3, 4, 5}},
			},
			expectedRoot: defaultHash([]byte{1, 2, 3, 4, 5}),
		},
		{
			testCaseId: 1,
			hashF:      defaultHash,
			contents: []Content{
				&DefaultContent{data: []byte{1, 2, 3, 4, 5}},
				&DefaultContent{data: []byte{5, 4, 3, 2, 1}},
			},
			expectedRoot: defaultHash(defaultHash([]byte{5, 4, 3, 2, 1}), defaultHash([]byte{1, 2, 3, 4, 5})),
		},
		{
			testCaseId: 2,
			hashF:      defaultHash,
			contents: []Content{
				&DefaultContent{data: []byte{1}},
				&DefaultContent{data: []byte{2}},
				&DefaultContent{data: []byte{3}},
			},
			expectedRoot: defaultHash(defaultHash([]byte{3}), defaultHash(defaultHash([]byte{2}), defaultHash([]byte{1}))),
		},
	}

	for _, test := range table {
		if !bytes.Equal(NewTree(defaultHash, test.contents...).Root(), test.expectedRoot) {
			t.Fatalf("content hashes mismatch, test: %d", test.testCaseId)
		}
	}
}

func TestPath(t *testing.T) {
	var table = []struct {
		testCaseId      int
		hashF           HashF
		contents        []Content
		contentToSearch Content
		expectedPath    [][]byte
	}{
		{
			testCaseId:      0,
			hashF:           defaultHash,
			contentToSearch: &DefaultContent{data: []byte{1}},
			contents: []Content{
				&DefaultContent{data: []byte{1}},
			},
			expectedPath: [][]byte{},
		},
		{
			testCaseId:      1,
			hashF:           defaultHash,
			contentToSearch: &DefaultContent{data: []byte{1}},
			contents: []Content{
				&DefaultContent{data: []byte{1}},
				&DefaultContent{data: []byte{2}},
			},
			expectedPath: [][]byte{defaultHash([]byte{2})},
		},
		{
			testCaseId:      2,
			hashF:           defaultHash,
			contentToSearch: &DefaultContent{data: []byte{1}},
			contents: []Content{
				&DefaultContent{data: []byte{1}},
				&DefaultContent{data: []byte{2}},
				&DefaultContent{data: []byte{3}},
			},
			expectedPath: [][]byte{defaultHash([]byte{2}), defaultHash([]byte{3})},
		},
		{
			testCaseId:      3,
			hashF:           defaultHash,
			contentToSearch: &DefaultContent{data: []byte{2}},
			contents: []Content{
				&DefaultContent{data: []byte{1}},
				&DefaultContent{data: []byte{2}},
				&DefaultContent{data: []byte{3}},
			},
			expectedPath: [][]byte{defaultHash([]byte{1}), defaultHash([]byte{3})},
		},
		{
			testCaseId:      4,
			hashF:           defaultHash,
			contentToSearch: &DefaultContent{data: []byte{3}},
			contents: []Content{
				&DefaultContent{data: []byte{1}},
				&DefaultContent{data: []byte{2}},
				&DefaultContent{data: []byte{3}},
			},
			expectedPath: [][]byte{defaultHash(defaultHash([]byte{2}), defaultHash([]byte{1}))},
		},
	}

	for _, test := range table {
		tree := NewTree(defaultHash, test.contents...)
		path, ok := tree.Path(test.contentToSearch)
		if !ok {
			t.Fatalf("content was not found, test: %d", test.testCaseId)
		}

		if len(path) != len(test.expectedPath) {
			t.Fatalf("wrong path size, test: %d", test.testCaseId)
		}

		for i := range path {
			if !bytes.Equal(path[i], test.expectedPath[i]) {
				t.Fatalf("path hashes mismatch, test: %d", test.testCaseId)
			}
		}
	}
}
