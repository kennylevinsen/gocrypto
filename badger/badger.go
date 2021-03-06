// Copyright 2015 Kenny Levinsen. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This package implements the Badger message authentication code (MAC).
//
// Specification: http://cryptico.com/images/pages/WP_Badger_Specification.pdf
//
// Review: http://cryptico.com/images/pages/Badger_Security_Report.pdf
//
// Reference: http://cryptico.com/images/pages/Badger_Reference_Code.zip
//
package badger

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/ebfe/estream/rabbit"
)

const (
	Trees  = 4
	Levels = 28
)

type badger struct {
	bitCount    uint64
	bufferIndex uint64
	finalPrng   []byte
	finalKey    [6][Trees]uint32
	levelKey    [Levels][Trees]uint64
	treeBuffer  [Levels][Trees]uint64
	buffer      [16]byte
}

// Hashes a blob of data. Key should be 16 bytes, IV should be 8 bytes.
// Returns 4 * Trees bytes as hash.
func Hash(data, key, iv []byte) (hash []byte, err error) {
	b := badger{}

	err = b.keySetup(key)
	if err != nil {
		return nil, err
	}

	b.process(data)

	hash, err = b.finalize(iv)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

// Sets up the level key and final key. The key parameter should be 16 bytes.
func (b *badger) keySetup(key []byte) error {
	spareFinalKeyIndex := 4
	spareFinalKey := [16]byte{}
	b.finalPrng = key

	rabbit, err := rabbit.NewCipher(key, []byte{})

	if err != nil {
		return errors.New(fmt.Sprintf("Unable to initialize first Rabbit instance: %s", err))
	}

	// We want to extract raw PRNG data, so we're making an empty input
	finalKeyData := make([]byte, Trees*6*4)
	rabbit.XORKeyStream(finalKeyData, finalKeyData)

	// Verify final keys
	for i := 0; i < 6; i++ {
		for j := 0; j < Trees; j++ {
			index := i*16 + j*4
			d := binary.LittleEndian.Uint32(finalKeyData[index : index+4])

			// As long as the key isinvalid...
			for d >= 0xFFFFFFFB {
				if spareFinalKeyIndex == 4 {
					// Fetch some more key material
					rabbit.XORKeyStream(spareFinalKey[:], spareFinalKey[:])
					spareFinalKeyIndex = 0
				}

				// Assemble the new key
				d = binary.BigEndian.Uint32(spareFinalKey[4*spareFinalKeyIndex : 4*spareFinalKeyIndex+4])
				spareFinalKeyIndex++

			}
			b.finalKey[i][j] = d
		}
	}

	// Another case of wanting PRNG
	levelKeyData := make([]byte, 8*Levels*Trees)
	rabbit.XORKeyStream(levelKeyData, levelKeyData)

	// Assign level keys
	for i := 0; i < Levels; i++ {
		for j := 0; j < Trees; j++ {
			index := 8 * (4*i + j)
			b.levelKey[i][j] = binary.LittleEndian.Uint64(levelKeyData[index : index+8])
		}
	}

	return nil
}

// Hashes nodes to the tree buffer
func (b *badger) process(src []byte) {
	length := len(src)
	srcOffset := 0

	// Hash the blocks!
	for length >= 16 {
		localData0 := binary.LittleEndian.Uint64(src[srcOffset : srcOffset+8])
		localData1 := binary.LittleEndian.Uint64(src[srcOffset+8 : srcOffset+16])

		// We do this once per tree
		for i := 0; i < Trees; i++ {
			hashNode(b.levelKey, &b.treeBuffer, b.bitCount>>7, localData0, localData1, i, 0)
		}
		b.bitCount += 0x80
		srcOffset += 16
		length -= 16
	}

	// If we have have an incomplete block left, put it in the buffer
	for length > 0 {
		b.buffer[b.bufferIndex] = src[srcOffset]
		b.bufferIndex++
		srcOffset++
		length--
	}
}

// Finalizes the tree buffer. The iv parameter should be 8 bytes.
func (b *badger) finalize(iv []byte) ([]byte, error) {
	right := [Trees]uint64{}
	bufferMask := b.bitCount >> 7
	counter := 0
	level := 0
	rightFilled := false

	// Here, we deal with the incomplete block, and put it into the temporary
	// result buffer.
	b.bitCount += 8 * b.bufferIndex
	if b.bufferIndex > 0 {
		if b.bufferIndex <= 8 {
			for b.bufferIndex < 8 {
				b.buffer[b.bufferIndex] = 0
				b.bufferIndex++
			}

			for i := 0; i < Trees; i++ {
				right[i] = binary.LittleEndian.Uint64(b.buffer[0:8])
			}
		} else {
			for b.bufferIndex < 16 {
				b.buffer[b.bufferIndex] = 0
				b.bufferIndex++
			}

			bufpart0 := binary.LittleEndian.Uint32(b.buffer[0:4])
			bufpart1 := binary.LittleEndian.Uint32(b.buffer[4:8])
			bufpart2 := binary.LittleEndian.Uint64(b.buffer[8:16])

			for i := 0; i < Trees; i++ {
				lk := b.levelKey[0][i]
				t1 := uint32(lk) + bufpart0
				t2 := uint32(lk>>32) + bufpart1
				right[i] = uint64(t1)*uint64(t2) + bufpart2
			}
		}
		rightFilled = true
	}

	if bufferMask == 0 && b.bufferIndex == 0 {
		// Just fill it with 0's if nothing has been processed
		for i := 0; i < Trees; i++ {
			right[i] = 0
		}
	} else {
		// Find the first full buffer and copy it to the temporary result buffer
		for !rightFilled {
			if bufferMask&1 != 0 {
				for i := 0; i < Trees; i++ {
					right[i] = b.treeBuffer[level][counter+i]
				}
				rightFilled = true
			}
			level++
			bufferMask = bufferMask >> 1
		}

		// Finish the tree
		for bufferMask != 0 {
			if bufferMask&1 != 0 {
				for i := 0; i < Trees; i++ {
					t1 := uint32(b.levelKey[level+1][counter]) + uint32(b.treeBuffer[level][counter])
					t2 := uint32(b.levelKey[level+1][counter]>>32) + uint32(b.treeBuffer[level][counter]>>32)
					right[i] += uint64(t1) * uint64(t2)
					counter++
					if counter >= Trees {
						counter -= Trees
						level++
					}
				}
			} else {
				level++
			}
			bufferMask = bufferMask >> 1
		}
	}

	// The final magic
	mac := make([]byte, Trees*4)
	for i := 0; i < Trees; i++ {
		t := uint64(b.finalKey[0][i]) * (right[i] & 0x07FFFFFF)
		t += uint64(b.finalKey[1][i]) * ((right[i] >> 27) & 0x07FFFFFF)
		t += uint64(b.finalKey[2][i]) * ((right[i] >> 54) | ((b.bitCount & 0x0001FFFF) << 10))
		t += uint64(b.finalKey[3][i]) * ((b.bitCount >> 17) & 0x07FFFFFF)
		t += uint64(b.finalKey[4][i]) * (b.bitCount >> 44)
		t += uint64(b.finalKey[5][i])

		low := uint32(t)
		high := uint32(t >> 32)

		r := low + 5*high

		if r < low || r > 0xFFFFFFFA {
			r -= 0xFFFFFFFB
		}

		binary.LittleEndian.PutUint32(mac[4*i:4*i+4], r)
	}

	// And encrypt the result
	r, err := rabbit.NewCipher(b.finalPrng, iv)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to initialize second Rabbit instance: %s", err))
	}

	r.XORKeyStream(mac, mac)
	return mac, nil
}

// Hashes the two uint64's, left and right, to the hash tree.
func hashNode(key [28][4]uint64, buffer *[28][4]uint64, bufferMask, left, right uint64, counter int, level uint64) {
	for {
		t0 := key[level][counter]
		t1 := uint32(t0) + uint32(left)
		t2 := uint32((t0 >> 32) + (left >> 32))
		right = uint64(t1)*uint64(t2) + right
		if bufferMask&1 == 0 {
			break
		}

		left = buffer[level][counter]
		level++
		bufferMask = bufferMask >> 1
	}
	buffer[level][counter] = right
}
