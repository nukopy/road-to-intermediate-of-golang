package main

import "fmt"

func main()  {
	// for 文
	cnt := 10
	for i := 0; i < cnt; i++ {
		fmt.Printf("Hello %d\n", i)
	}

	s := []int{10, 20, 30, 40, 50}
	for i, v := range s {
		fmt.Printf("%d: %d\n", i, v)
	}

	// switch 文
	for i, v := range s {
		switch v {
		case 10, 30:
			fmt.Printf("i: %d\n", i)
		case 20:
			fmt.Printf("v: %d\n", v)
		default:
			fmt.Printf("default: %d\n", v)
		}
	}

	// 参照型
	var e = [5]int{1, 2, 3, 4, 5}
	var f = []float64{3.14, 4.25, 5.36}
	var g = map[string]int{
		"Apple": 5,
		"Banana": 6,
	}
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// 構造体
	p := Person{
		name: "John",
		age: 18,
		height: 175.5,
		weight: 69.1,
	}
	fmt.Println(p)
	p.NotDestroyHappyBirthDay()
	fmt.Println(p) // 値レシーバを受け取るメソッドのため、破壊的な変更がない
	p.DestroyHappyBirthDay()
	fmt.Println(p) // ポインタレシーバを受け取るメソッドのため、破壊的な変更がある
}

type Person struct {
	name string
	age int
	height float64
	weight float64
}

func (p Person) NotDestroyHappyBirthDay() {
	// 値レシーバ
	fmt.Printf("Happy birthday, you become %d from %d\n", p.age + 1, p.age)
	p.age++
}

func (p *Person) DestroyHappyBirthDay() {
	// ポインタレシーバ
	fmt.Printf("Happy birthday, you become %d from %d\n", p.age + 1, p.age)
	p.age++
}