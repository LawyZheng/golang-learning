package main

import (
	"fmt"
	// "sort"
)

// import "strings"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatch(name string) {
	for _, char := range name {
		if char == 'e' || char == 'E' {
			distribution[name] += 1
			coins -= 1
		} else if char == 'i' || char == 'I' {
			distribution[name] += 2
			coins -= 2
		} else if char == 'o' || char == 'O' {
			distribution[name] += 3
			coins -= 3
		} else if char == 'u' || char == 'U' {
			distribution[name] += 4
			coins -= 4
		}
	}
}

func dispatchCoin() int {
	for _, each_name := range users {
		// fmt.Println(each_name)
		dispatch(each_name)
	}
	fmt.Printf("%+v\n", distribution)
	return coins
}

type student struct {
	name string
	age  int
}

func main() {
	// a1 := 1
	// a2 := 3
	// a3 := 4
	// a4 := 4
	// a5 := 1
	// var b bool = true
	// var s2 = 'h'
	// for i,v:=range l {
	// 	fmt.Printf("%d %d\n", i, v)
	// }
	// fmt.Printf("%T: %v\n", b, b)

	// fmt.Println(a1^a2^a3^a4^a5)

	// fmt.Println(int('9'))
	// fmt.Println(len(s1), len(strings.Split(s1, ".")) )
	// fmt.Printf("%v\n", s2)

	// s := "hello沙河小王子"
	// fmt.Println(len(s))
	// cnt := 0
	// for _, c := range s {
	// 	if ( c > 128){
	// 		cnt ++
	// 		fmt.Printf("%v %c\n", c, c)
	// 	}
	// }

	// fmt.Println(cnt)

	// for i:=1;i<10;i++ {
	// 	for j:=1;j<=i;j++{
	// 		fmt.Printf("%d * %d = %d ", i, j, i*j)
	// 	}
	// 	fmt.Println()
	// }

	// a1 := [...]int{1, 3, 5, 7, 8}
	// s := 0
	// for _, v := range a1 {
	// 	s += v
	// }
	// fmt.Println(s)

	// for i:=0; i<len(a1) ; i++ {
	// 	for j:=i; j<len(a1) ; j++ {
	// 		if a1[i] + a1[j] == 8 {
	// 			fmt.Printf("(%d, %d)\n", i, j)
	// 		}
	// 	}
	// }

	// var a = make([]string, 5, 10)
	// var a []string
	// for i := 0; i < 10; i++ {
	// 	a = append(a, fmt.Sprintf("%v", i))
	// 	fmt.Println(len(a), a)
	// }
	// fmt.Println(a)

	// var a = [...]int{3, 7, 8, 9, 1}
	// s := a[:4]
	// fmt.Printf("s=%v len(s)=%d cap(s)=%d\n", s, len(s), cap(s))

	// s[4] = 5
	// fmt.Printf("s=%v", s)
	// s := a[:]
	// s = append(s, 4)
	// fmt.Printf("%v\n", s)
	// sort.Ints(s)
	// fmt.Printf("%v\n", s)

	// // for i:=0; i<len(s) ; i++{
	// 	// a[i] = s[i]
	// // }
	// fmt.Printf("%v\n", a)

	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1

	// s := "how do you do"
	// word := ""
	// cnt_map := make(map[string]int, 10)

	// for _, c := range s{
	// 	if c != ' '{
	// 		word = fmt.Sprintf("%s%c", word, c)
	// 	}else{
	// 		cnt_map[word]++
	// 		word = ""
	// 	}
	// }
	// cnt_map[word]++
	// word = ""

	// for k, v := range cnt_map{
	// 	fmt.Printf("%s=%d\n", k, v)
	// }

	// 观察下面代码，写出最终的打印结果。
	// type Map map[string][]int
	// m := make(Map)
	// s := []int{1, 2}
	// s = append(s, 3)
	// fmt.Printf("%+v\n", s)
	// m["q1mi"] = s
	// fmt.Printf("Pre Point: %p\n", s)
	// fmt.Printf("map Point: %p\n", m["q1mi"])
	// s = append(s[:1], s[2:]...)
	// fmt.Printf("Pre Point: %p\n", s)
	// fmt.Printf("%+v\n", len(s))
	// fmt.Printf("%+v\n", len(m["q1mi"]))
	// fmt.Printf("map Point: %p\n", m["q1mi"])

	// s := "这是中文字符串"
	// fmt.Println("length:", len(s))
	// fmt.Println("s[0]:", s[0:3])

	// for i, c := range s{
	// 	// fmt.Printf("%T\n", c)
	// 	fmt.Printf("%d: %c\n",i ,c)
	// }

	// for i := range s {
	// 	fmt.Println(i)
	// }

	// left := dispatchCoin()
	// fmt.Println("剩下：", left)

	// m := make(map[string]*student)
	// stus := []student{
	// 	{name: "小王子", age: 18},
	// 	{name: "娜扎", age: 23},
	// 	{name: "大王八", age: 9000},
	// }

	// for _, stu := range stus {
	// 	fmt.Printf("&stu: %p, value: %+v\n", &stu, stu)
	// 	m[stu.name] = &stu
	// }
	// for k, v := range m {
	// 	fmt.Println(k, "=>", v.name)
	// }

}
