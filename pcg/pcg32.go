// Copyright 2014, David Minor. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package pcg

import (
  lcglib "github.com/davidminor/gorand/lcg"
)

type Permute64x32 func(uint64) uint32

type Pcg64x32 struct {
  lcg *lcglib.Lcg64
  permute Permute64x32
}

func NewPcg64x32(lcg *lcglib.Lcg64, permuteFunc Permute64x32) *Pcg64x32 {
  // this is how the reference library sets the initial state
  lcg.State += lcg.Stream
  return &Pcg64x32{lcg, permuteFunc}
}

// Get the next random uint32 value
func (pcg *Pcg64x32) Next() uint32 {
  return pcg.permute(pcg.lcg.Next())
}

// Get a random uint32 value evenly distributed in [0,bounds)
func (pcg *Pcg64x32) NextN(bounds uint32) uint32 {
  threshold := (0 - bounds) % bounds
  for {
    result := pcg.Next()
    if result >= threshold {
      return result % bounds;
    }
  }
}

// Set the stream of this PCG
func (rng *Pcg64x32) Stream(stream uint64) {
  // stream must be odd for LCG, so shift left 1 and turn on the 1 bit
  rng.lcg.Stream = stream << 1 | 1
}

// Pcg32 uses the top 37 bits of its 64 bit LCG, XOR'ing the highest half 
// with the lowest, and then randomly rotating the lower 32 of them 
// (which are returned)
type Pcg32 struct {
  *Pcg64x32
}

// Create a new Pcg32 with the given seed
func NewPcg32(seed uint64) Pcg32 {
  lcg := lcglib.NewLcg64(seed)
  return Pcg32{NewPcg64x32(lcg, XshRr)}
}

// Create a new Pcg32 with the given seed and stream
func NewPcg32Stream(seed, stream uint64) Pcg32 {
  lcg := lcglib.NewLcg64Stream(seed, stream)
  return Pcg32{NewPcg64x32(lcg, XshRr)}
}

// Take the highest 37 bits, xor the top half with the bottom,
// then use the top 5 to randomly rotate the next 32 (which we return)
func XshRr(state uint64) uint32 {
  rot := uint32(state >> 59);
  shift := uint32(((state >> 18) ^ state) >> 27)
  return (shift >> rot) | (shift << (32 - rot));
}

