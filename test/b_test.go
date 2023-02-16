package test

import (
	g "dousheng/pkg/global"
	"testing"
	"time"
)

func TestB(t *testing.T) {
	g.A = 2
	time.Sleep(1000 * time.Second)
}
