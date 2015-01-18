// Copyright 2014, David Minor. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package pcg

import (
  . "github.com/davidminor/uint128"
  lcglib "github.com/davidminor/gorand/lcg"
)

type Permute128x64 func(Uint128) uint64

type Pcg128x64 struct {
  lcg *lcglib.Lcg128
  permute Permute128x64
}

func NewPcg128x64(lcg *lcglib.Lcg128, permuteFunc Permute128x64) *Pcg128x64 {
  // this is how the reference library sets the initial state
  lcg.State = lcg.State.Add(lcg.Stream)
  // reference library uses "post-advance" state of LCG, so we do an 
  // initial advance to match its output
  lcg.Next()
  return &Pcg128x64{lcg, permuteFunc}
}

// Get the next random uint64 value
func (pcg *Pcg128x64) Next() uint64 {
  return pcg.permute(pcg.lcg.Next())
}

// Get a random uint64 value evenly distributed in [0,bounds)
func (pcg *Pcg128x64) NextN(bounds uint64) uint64 {
  threshold := (0 - bounds) % bounds
  for {
    result := pcg.Next()
    if result >= threshold {
      return result % bounds;
    }
  }
}

// Set the stream of this PCG
func (rng *Pcg128x64) Stream(streamH, streamL uint64) {
  // stream must be odd for LCG, so shift left 1 and turn on the 1 bit
  rng.lcg.Stream = Uint128{streamH,streamL}.ShiftLeft(1)
  rng.lcg.Stream.L |= 1
}

// Pcg64 uses XOR of high and low bits combined with random shift
type Pcg64 struct {
  *Pcg128x64
}

// Create a new Pcg64 with the given high and low bits of seed
func NewPcg64(seedH, seedL uint64) Pcg64 {
  lcg := lcglib.NewLcg128(Uint128{seedH, seedL})
  return Pcg64{NewPcg128x64(lcg, XslRr)}
}

// Create a new Pcg64 with the given high and low bits of seed and stream
func NewPcg64Stream(seedH, seedL, streamH, streamL uint64) Pcg64 {
  lcg := lcglib.NewLcg128Stream(Uint128{seedH, seedL}, 
    Uint128{streamH, streamL})
  return Pcg64{NewPcg128x64(lcg, XslRr)}
}

// Permute functions

// Xor the state's top bits with the bottom and randomly rotate them
// based on the highest 6 bits.
func XslRr(state Uint128) uint64 {
  h, l := state.H, state.L
  shift := l ^ h;
  rot := h >> 58;
  return (shift >> rot) | (shift << (64 - rot));
}

