package obj_pool

import (
	"testing"
	"time"
)

func TestObjPool(t *testing.T){
	 myPool := NewObjPool(10)
	 for i := 0; i < 10; i++ {
	 	if v, err := myPool.GetObjPool(time.Second*1); err != nil {
	 		t.Error(err)
		} else {
			t.Logf("%T \n ", v)
			if err := myPool.ReleaseObjPool(v); err != nil {
				t.Error(err)
			}
		}
	 }
}