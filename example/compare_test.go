package example

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

// 1. 基础类型变量的比较
func TestBasic(t *testing.T)  {
	// 类型完全不同，不能比较
	//fmt.Println(1 == "1")

	// 类型自定义关系，不能比较，但可以强制转换后比较
	var a int = 1
	type newInt int
	var b newInt = 1
	//fmt.Println(a == b)
	fmt.Println(a == int(b))

	// 类型别名关系，可以比较
	type intAlias = int
	var c intAlias = 1
	fmt.Println(a == c)
}

// 2. 复合类型变量的比较
// 复合类型array和struct是逐个元素进行比较，且元素也需是可比较的
func TestComp(t *testing.T)  {
	var arr1 [2]int = [2]int{1, 2}
	var arr2 [3]int = [3]int{1, 2}
	fmt.Println(arr1, arr2)
	//fmt.Println(arr1 == arr2)

	type Stu struct {
		name string
		age  int
		Addr []string
	}
	s1 := Stu{"tom", 20,[]string{"China", "Beijing"}}
	s2 := Stu{"jerry", 30, []string{"China", "Shanghai"}}
	fmt.Println(s1, s2)
	//fmt.Println(s1 == s2)
}

// 3. 引用类型变量的比较
func TestRef(t *testing.T)  {
	// 普通引用类型，判断两个引用类型存储的是否为同一个变量
	type Stu struct {
		name string
		age  int
	}
	s1 := &Stu{"tom", 20}
	s2 := &Stu{"tom", 20}
	fmt.Println(s1 == s2)	// false

	// channel 同普通引用类型的比较规则
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)
	fmt.Println(ch1 == ch2)	// false

	// slice 不可比较（除nil外，可以自己实现比较函数，逐个元素比较）
	a := []string{}
	b := []string{}
	fmt.Println(a, b)
	//fmt.Println(a == b)

	c := []byte{'f', 'o', 'o'}
	d := []byte{'f', 'o', 'o'}
	fmt.Println(bytes.Equal(c, d))	// true

	// map 同slice比较规则
	// reflect.DeepEqual()比较，只要变量的类型和值相同，结果就为true
	m1 := map[string]int{"foo": 1, "bar": 2}
	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(reflect.DeepEqual(m1, m2))	// true



	// 函数类型也不可比较
	fun1 := func(int) int {return 1}
	fun2 := func(int) int {return 1}
	fmt.Println(&fun1, &fun2)
	//fmt.Println(fun1 == fun2)
	fmt.Println(reflect.DeepEqual(fun1, fun2))	// false
}

// 4. 接口类型interface
// 接口类型变量包含接口变量存储的值和值的类型两部分，即接口的动态类型和动态值
// 只有动态类型和动态值都相同，且动态类型可比较时，两个接口变量才相同
func TestInterface(t *testing.T)  {
	d1 := Dog{"tom"}
	d2 := Dog{"tom"}
	c1 := Cat{"tom"}

	fmt.Println(compare(d1, d2))	// true
	fmt.Println(compare(d1, c1))	// false
}

// 总结
// 基础类型完全相同，可比较
// 复合类型只有每个元素可比较，且类型和值都相同，才相等
// 普通引用类型（&var）和channel可比较，只有指向同一个变量才相等
// slice、map、func不可比较，但可以使用reflect或第三方的cmp包来进行比较
// 自定义类型不可比较
// 类型别名可比较


