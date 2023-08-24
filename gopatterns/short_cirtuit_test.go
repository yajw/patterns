package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunShortCircuit(t *testing.T) {
	n := rand.Intn(2) + 2 // 1 - 30

	var fs []fun
	var ret = true

	for i := 0; i < n; i++ {
		v := rand.Intn(3) > 0
		ret = ret && v
		s := rand.Intn(10)
		t.Logf("#%v returns %v, sleep %d s", i, v, s)

		fs = append(fs, func() bool {
			time.Sleep(time.Duration(s) * time.Second)
			return v
		})
	}

	start := time.Now()
	t.Logf("%v start", start)
	v := RunShortCircuit(fs...)
	end := time.Now()
	t.Logf("%v end, cost = %v", end, end.Sub(start))

	t.Log("v =", v)

	if v != ret {
		assert.Equal(t, ret, v)
	}
}
