///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sun Oct 30 2016 18:40:52]
// Generate from:
//   [github.com/vipally/gx/math/rand/rand.gp]
//   [github.com/vipally/gx/math/rand/rand.gpg] [64]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct 24 2016 20:25:45]
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

//Package rand implements some useful rand object
package rand

import (
	"sync/atomic"
)

var (
	gRand64 = NewRand64S(uint64(RandSeed(0)))
)

//generate a rand number
func Rand64() uint64 {
	return gRand64.Rand()
}

//generate a rand number less than max
func RandMax64(max uint64) uint64 {
	return gRand64.RandMax(max)
}
func RandRange64(min, max uint64) uint64 {
	return gRand64.RandRange(min, max)
}

//rand number generator
//It is thread safe
type Rand64T struct {
	seed uint64
	//lock sync.Mutex
}

//new a initialized rand64 object
func NewRand64S(seed uint64) *Rand64T {
	return &Rand64T{seed: seed}
}

//new a rand64 object initialized by auto-generated seed
func NewRand64() *Rand64T {
	return NewRand64S(gRand64.randBase())
}

//next rand number
func (me *Rand64T) Rand() uint64 {
	var o, n uint64
	for { //mutithread lock-free operation
		o = atomic.LoadUint64(&me.seed)
		n = o*4294955897 + 4094975977
		if atomic.CompareAndSwapUint64(&me.seed, o, n) {
			break
		}
	}
	return n

	//me.seed = me.seed*g_prime_a64 + g_prime_c64
	//return me.seed
}

//new rand seed list
func (me *Rand64T) randBase() uint64 {
	return uint64(RandSeed(uint64(me.Rand())))
}

//generate rand number in range
func (me *Rand64T) RandRange(min, max uint64) uint64 {
	if max < min {
		max, min = min, max
	}
	d := max - min + 1
	r := me.Rand()
	ret := r%d + min

	return ret
}

//generate rand number with max value
func (me *Rand64T) RandMax(max uint64) uint64 {
	return me.RandRange(0, max-1)
}

//get seed
func (me *Rand64T) Seed() uint64 {
	return atomic.LoadUint64(&me.seed)
}

//set seed
func (me *Rand64T) Srand(_seed uint64) uint64 {
	ret := atomic.SwapUint64(&me.seed, _seed) //mutithread lock-free operation
	return ret
}
