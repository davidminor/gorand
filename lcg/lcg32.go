// Copyright 2014, David Minor. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package lcg

var default_multiplier32 = uint32(747796405)
var default_stream32 = uint32(2891336453)

type Lcg32 struct {
  State, Stream, Multiplier uint32
}

func (lcg *Lcg32) Next() uint32 {
  lcg.State = lcg.State*lcg.Multiplier + lcg.Stream
  return lcg.State
}

func NewLcg32(seed uint32) *Lcg32 {
  return &Lcg32{seed, default_stream32, default_multiplier32}
}

func NewLcg32Stream(seed, stream uint32) *Lcg32 {
  stream = stream << 1 | 1
  return &Lcg32{seed, stream, default_multiplier32}
}

func (lcg *Lcg32) Int63() int64 {
  return int64(lcg.Next() >> 1) | int64(lcg.Next()) << 30
}

func (lcg *Lcg32) Seed(s int64) {
  lcg.State = uint32(s)
}

