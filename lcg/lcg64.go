// Copyright 2014, David Minor. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package lcg

var default_multiplier64 = uint64(6364136223846793005)
var default_stream64 = uint64(1442695040888963407)

type Lcg64 struct {
  State, Stream, Multiplier uint64
}

func (lcg *Lcg64) Next() uint64 {
  lcg.State = lcg.State*lcg.Multiplier + lcg.Stream
  return lcg.State
}

func NewLcg64(seed uint64) *Lcg64 {
  return &Lcg64{seed, default_stream64, default_multiplier64}
}

func NewLcg64Stream(seed, stream uint64) *Lcg64 {
  stream = stream << 1 | 1
  return &Lcg64{seed, stream, default_multiplier64}
}

func (lcg *Lcg64) Int63() int64 {
  return int64(lcg.Next() >> 1)
}

func (lcg *Lcg64) Seed(s int64) {
  lcg.State = uint64(s)
}

