package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	default:
		fmt.Printf("unknown type", t)
	}
}

//func TestBaseType(t *testing.T) {
//	var f float64 = 12
//	CheckType(f)
//	CheckType(&f)
//}

type Employee struct {
	Id   int
	Name string `format:"normal""`
	Age  int
}

func (e *Employee) UpdateAge(newage int) {
	e.Age = newage
}

func (e *Employee) UpdateName(newname string) {
	e.Name = newname
}

func (e *Employee) UpdateId(newid int) {
	e.Id = newid
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{1, "Mike", 30}
	f := []Employee{{2, "Neek", 31}, {3, "Cook", 32}}
	t.Logf("Value: id: %d, name: %s, age: %d \n", reflect.ValueOf(*e).FieldByName("Id"), reflect.ValueOf(*e).FieldByName("Name"), reflect.ValueOf(*e).FieldByName("Age"))
	t.Log(reflect.TypeOf(*e).FieldByName("Id"))
	t.Log(reflect.TypeOf(*e).FieldByName("Name"))
	t.Log(reflect.TypeOf(*e).FieldByName("Age"))

	if namevalue, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("get field name value failed")
	} else {
		t.Log("Tag:format", namevalue.Tag.Get("format"))
		//t.Log("Tag:format", namevalue)
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(111)})
	t.Log(e)

	//字符串
	reflect.ValueOf(e).MethodByName("UpdateName").Call([]reflect.Value{reflect.ValueOf("world")})
	reflect.ValueOf(e).MethodByName("UpdateId").Call([]reflect.Value{reflect.ValueOf(12344)})
	t.Log(e)

	//数组
	reflect.ValueOf(&f[1]).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(131)})
	t.Log(f[0])
	t.Log(f[1])
	t.Log(f)
}
