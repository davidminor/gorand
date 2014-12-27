// Copyright 2014, David Minor. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package pcg

import (
  "testing"
  . "github.com/davidminor/uint128"
)

func TestPcg64(t *testing.T) {
  rng := NewPcg64Stream(Uint128{0,uint64(42)}, Uint128{0,uint64(54)})
  for i := 0; i < 5; i++ {
    rng.Uint64()
  }
  
  if next := rng.Uint64(); next != uint64(0x606121f8e3919196) {
    t.Errorf("Got %x, was expecting %x", uint64(0x606121f8e3919196));
  }
}

