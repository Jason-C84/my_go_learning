package client

import (
	"ch15/series"
	"testing"
	cm "github.com/Jason-C84/concurrent_map"
)

func TestPackage(t *testing.T){
	t.Log(series.GetFiboncciSeries(8))
	t.Log(series.Squarl(10))

	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))

}