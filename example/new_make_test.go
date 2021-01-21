package example

import (
	"fmt"
	"testing"
)

// 值类型：int float bool string array struct 变量直接存储值，内存通常在栈中分配
// 引用类型 pointer slice map channel interface function 变量存储的是一个地址，这个地址对应的空间才是真正存储的值，内存通常在堆中分配，通过GC回收

// 变量声明
// 值类型变量在声明后会默认分配好内存，如果不指定默认值，则使用零值
// 引用类型变量在声明后，要为它们分配内存，否知直接使用会panic

// new()和make()都可以用来创建分配类型内存
// new()在分配内存后返回指向该类型的指针，同时把分配的内存置为零值
// make()只能用于slice/map/chan的内存创建，并且返回该类型本身

func TestSlice(t *testing.T)  {
	s1 := new([]int)
	fmt.Printf("%T\n", s1)
	fmt.Println(*s1)

	s2 := make([]int, 1)
	fmt.Printf("%T\n", s2)
	fmt.Println(s2)
}

