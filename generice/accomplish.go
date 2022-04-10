package main

import "fmt"

type IAnimal interface {
	Barking() string
}

func Accomplish(i IAnimal) {
	fmt.Println(i.Barking)
}

type Dog struct {
	Name string
}

func (d *Dog) Barking() string {
	d.Name = "小狗:"
	return d.Name + "汪汪汪"
}

type Cat struct {
	Name string
}

func (c *Cat) Barking() string {
	c.Name = "小猫:"
	return c.Name + "喵喵喵"
}

type Duck struct {
	Name string
}

func (d *Duck) Barking() string {
	d.Name = "小鸭:"
	return d.Name + "嘎嘎嘎"
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
