// Copyright 2015 Kenny Levinsen. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package badger

import (
	"bytes"
	"fmt"
	"github.com/joushou/gocrypto/badger/testdata"
	"testing"
)

// TODO: Extend keysetup and process tests considerably

func TestKeySetup(t *testing.T) {
	tests := []testdata.KeyTestSet{
		testdata.KeyTestSet1(),
		testdata.KeyTestSet2(),
	}

	for i, test := range tests {
		b := badger{}
		err := b.keySetup(test.Key)
		if err != nil {
			t.Fatal(fmt.Sprintf("Test %d: keySetup failed: %s", i+1, err))
		}
		if b.levelKey != test.LevelKey {
			t.Error(fmt.Sprintf("Test %d: Level key incorrect. Got: %d, Expected: %d", i+1, b.levelKey, test.LevelKey))
		}
		if b.finalKey != test.FinalKey {
			t.Error(fmt.Sprintf("Test %d: Final key incorrect. Got: %d, Expected: %d", i+1, b.finalKey, test.FinalKey))
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []testdata.ProcessTestSet{
		testdata.ProcessTest1(),
		testdata.ProcessTest2(),
	}

	for i, test := range tests {
		b := badger{}
		err := b.keySetup(test.Key)
		if err != nil {
			t.Fatal(fmt.Sprintf("Test %d: keySetup failed: %s", i+1, err))
		}

		b.process(test.Data)
		if test.BitCount != b.bitCount {
			t.Error(fmt.Sprintf("Test %d: Incorrect bit count. Got: %d, Expected: %d", i+1, b.bitCount, test.BitCount))
		}
		if test.BufferIndex != b.bufferIndex {
			t.Error(fmt.Sprintf("Test %d: Incorrect buffer index. Got: %d, Expected: %d", i+1, b.bufferIndex, test.BufferIndex))
		}
		if test.TreeBuffer != b.treeBuffer {
			t.Error(fmt.Sprintf("Test %d: Incorrect tree buffer", i+1))
		}
		if test.Buffer != b.buffer {
			t.Error(fmt.Sprintf("Test %d: Incorrect remaining buffer", i+1))
		}
	}
}

func TestComplete(t *testing.T) {
	test := []testdata.CompleteTestSet{
		testdata.CompleteTest1(),
		testdata.CompleteTest2(),
		testdata.CompleteTest3(),
		testdata.CompleteTest4(),
		testdata.CompleteTest5(),
		testdata.CompleteTest6(),
		testdata.CompleteTest7(),
		testdata.CompleteTest8(),
	}

	for i, test := range test {
		out, err := Hash(test.Data, test.Key, test.Iv)
		if err != nil {
			t.Fatal(fmt.Sprintf("Test %d: Hashing failed: %s", i+1, err))
		}
		if bytes.Compare(out, test.Mac) != 0 {
			t.Error(fmt.Sprintf("Test %d: Mac incorrect.", i+1))
		}
	}
}
