package main

import (
	"fmt"
)

type School struct {
	StudentManage []*Student
}

func NewSchool() *School {
	this := new(School)
	this.StudentManage = make([]*Student, 0)
	return this
}


type Student struct {
	Id   int
	Name string
	Age  int
}

func NewStudent(name string, age int) *Student {
	this := new(Student)
	this.Name = name
	this.Age = age
	return this
}
func main() {
	school := NewSchool()
	students := []*Student{{Name: "小黄", Age: 15}, {Name: "小李", Age: 12}, {Name: "小红", Age: 14}}
	// 把上边数组数据来三轮添加 -----------分配id 错误示例-------------
	// for i := 0; i < 3; i++ {
	// 	school.StudentManage = append(school.StudentManage, students...)
	// }
	for i := 0; i < 3; i++ {
		for i := 0; i < len(students); i++ {
			school.StudentManage = append(school.StudentManage, NewStudent(students[i].Name, students[i].Id))
		}

	}

	// 给 9 个学生设置id

	for i := 0; i < len(school.StudentManage); i++ {
		school.StudentManage[i].Id = i
	}
	fmt.Println("------------------------------初始化完成------------------------------")
	// 显示玩家信息
	for i := 0; i < len(school.StudentManage); i++ {
		fmt.Println("学生Id：", school.StudentManage[i].Id)
	}

}
