// Copyright 2014, David Minor. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package pcg

import (
  . "github.com/davidminor/uint128"
)

var multiplier = Uint128{2549297995355413924,4865540595714422341}
var default_stream = Uint128{6364136223846793005,1442695040888963407}

type Pcg64 struct {
  state, stream Uint128
}

// create a new PCG with the given seed and the default stream
func NewPcg64(seed Uint128) *Pcg64 {
  rng := &Pcg64{}
  rng.Seed(seed)
  return rng
}

// create a new PCG with the given seed and stream
func NewPcg64Stream(state, stream Uint128) *Pcg64 {
  rng := &Pcg64{}
  rng.Stream(stream)
  rng.Seed(state)
  return rng
}

// seed the PCG's initial state
func (rng *Pcg64) Seed(seed Uint128) {
  if rng.stream.L == 0 && rng.stream.H == 0 {
    rng.Stream(default_stream)
  }
  rng.state = seed.Add(rng.stream)
  rng.bump()
}

// set the stream of this PCG
func (rng *Pcg64) Stream(stream Uint128) {
  // stream must be odd for LCG, so shift left 1 and turn on the 1 bit
  rng.stream = stream.ShiftLeft(1)
  rng.stream.L |= uint64(1)
}

// advance the LCG one step
func (rng *Pcg64) bump() {
  rng.state = rng.state.Mult(multiplier).Add(rng.stream)
}

// xor the top bits with the bottom and randomly rotate them
func output(internal Uint128) uint64 {
  rot := internal.H >> 58;
  shift := internal.L ^ internal.H;
  return (shift >> rot) | (shift << (64 - rot));
}

// returns the next uint64 from the generator
func (rng *Pcg64) Uint64() uint64 {
  rng.bump()
  return output(rng.state)
}

// returns an uint64 uniformly distributed in [0,bounds)
func (rng *Pcg64) Uint64n(bounds uint64) uint64 {
  threshold := (uint64(0) - bounds) % bounds
  for {
    result := rng.Uint64()
    if result >= threshold {
      return result % bounds;
    }
  }
}

