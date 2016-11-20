// Copyright 2014, David Minor. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package lcg

import (
  . "github.com/davidminor/uint128"
)

var default_multiplier128 = Uint128{2549297995355413924,4865540595714422341}
var default_stream128 = Uint128{6364136223846793005,1442695040888963407}

type Lcg128 struct {
  State, Stream, Multiplier Uint128
}

func (lcg *Lcg128) Next() Uint128 {
  lcg.State = lcg.State.Mult(lcg.Multiplier).Add(lcg.Stream)
  return lcg.State
}

func NewLcg128(seed Uint128) *Lcg128 {
  return &Lcg128{seed, default_stream128, default_multiplier128}
}

func NewLcg128Stream(seed, stream Uint128) *Lcg128 {
  stream = stream.ShiftLeft(1)
  stream.L |= 1
  return &Lcg128{seed, stream, default_multiplier128}
}

func (lcg *Lcg128) Int63() int64 {
  n := lcg.Next()
  return int64((n.H ^ n.L) >> 1)
}

func (lcg *Lcg128) Seed(s int64) {
  lcg64 := NewLcg64(uint64(s))
  lcg.State.H, lcg.State.L = lcg64.Next(), lcg64.Next()
}

