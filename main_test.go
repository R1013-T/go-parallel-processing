package main

import (
	"go.uber.org/goleak"
	"testing"
)

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	main()
}
