package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `json:"name" ini:"s_name"`
	Score int `json:"score" ini:"s_score"`
}
// 根据反射获取结构体中的方法
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}

	methodObj,_ := t.MethodByName("Sleep")
	fmt.Println(methodObj)
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func main() {
	stu1 := student{
		Name: "白雪公主",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v,kind:%v\n",t.Name(),t.Kind())

	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v ini tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"),field.Tag.Get("ini"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	printMethod(stu1)
}
