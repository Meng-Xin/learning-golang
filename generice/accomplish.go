package main

import "fmt"

type IAnimal interface {
	Barking() string
}

// Accomplish 这里使用了回调函数特性
func Accomplish(i IAnimal) {
	fmt.Println(i.Barking())
}

type Dog struct {
	Name string
}

func (d *Dog) Barking() string {
	var bark string
	d.Name = "小狗:"
	bark = d.Name + "汪汪汪"
	return bark
}

type Cat struct {
	Name string
}

func (c *Cat) Barking() string {
	var bark string
	c.Name = "小猫:"
	bark = c.Name + "喵喵喵"
	return bark
}

type Duck struct {
	Name string
}

func (d *Duck) Barking() string {
	var bark string
	d.Name = "小鸭:"
	bark = d.Name + "嘎嘎嘎"
	return bark
}

func main() {
	var selectAnimal int
	fmt.Scan(&selectAnimal)
	switch selectAnimal {
	case 0:
		dog := new(Dog)
		Accomplish(dog)
	case 1:
		cat := new(Cat)
		Accomplish(cat)
	case 2:
		duck := new(Duck)
		Accomplish(duck)
	}
}
