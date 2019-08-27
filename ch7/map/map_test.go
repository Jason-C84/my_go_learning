package _map

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1:2,2:4,3:9}
	t.Log(m1[2])
	t.Logf("m1.len = %d",len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("m2.len = %d", len(m2))

	m3 := make(map[int]int,10)
	t.Logf("m3.log = %d", len(m3))

}


func TestKeyNotExistingValue(t *testing.T){
	m1 := map[int]int{}
	m1[2] = 1
	if v,ok := m1[3]; ok{
		t.Logf("m1[3] is :%d",v)
	}else{
		t.Log("m1[3] not existing")
	}

}


func TestTravlMap(t *testing.T){
	m1 := map[int]int{1:2,2:4,3:9}
	for k,v := range m1 {
		t.Log(k,v)

	}
}