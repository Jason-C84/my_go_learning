package obj_pool

import (
	"time"
	"errors"
)

type ReusableObj struct {

}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfPool int) *ObjPool{
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfPool)
	for i := 0; i < numOfPool; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool)GetObjPool(timeout time.Duration)(*ReusableObj,error){
	select{
	case ret := <- p.bufChan:
		return ret, nil
	case <- time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *ObjPool)ReleaseObjPool(obj *ReusableObj) error{
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

