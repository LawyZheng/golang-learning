package main

import (
	"fmt"
	"os"
)

type system struct {
}

type student struct {
	studentID string
	name      string
}

func newStudent(studentID string, name string) *student {
	return &student{
		studentID: studentID,
		name:      name,
	}
}

func (s system) showMsg() {
	fmt.Print(`下面是操作菜单:
			1. 查看学生列表
			2. 添加学生
			3. 删除学生
			4. 退出程序
	`)
}

func (s system) getMenu() int {
	fmt.Print("请输入指令:")
	var menu int
	fmt.Scan(&menu)
	return menu
}

func (s system) exitSystem() {
	fmt.Println("退出系统...")
	os.Exit(0)
}

func (s system) showStudents() {
	for k, v := range stuMap {
		fmt.Printf("学号: %s  姓名: %s\n", k, v.name)
	}
}

func (s system) addStudent() {
	var studentID string
	var name string

	fmt.Print("请输入学号:")
	fmt.Scan(&studentID)
	fmt.Print("请输入姓名:")
	fmt.Scan(&name)

	_, ok := stuMap[studentID]
	if ok {
		fmt.Println("该ID已存在！")
		return
	}

	stuMap[studentID] = newStudent(studentID, name)
	fmt.Println("添加成功！")
}

func (s system) deleteStudent() {
	var studentID string

	fmt.Print("请输入学号:")
	fmt.Scan(&studentID)

	_, ok := stuMap[studentID]
	if !ok {
		fmt.Println("该ID不存在！")
		return
	}

	delete(stuMap, studentID)
	fmt.Println("删除成功！")
}

var (
	stuMap map[string]*student
)

func main() {
	stuMap = make(map[string]*student, 10)
	fmt.Println("欢迎来到学生管理系统")
	for {
		var sys system
		sys.showMsg()
		menu := sys.getMenu()
		fmt.Printf("你选择了%d选项\n", menu)

		switch menu {
		case 1:
			sys.showStudents()
		case 2:
			sys.addStudent()
		case 3:
			sys.deleteStudent()
		case 4:
			sys.exitSystem()
		default:
			fmt.Println("没有该选项！")
		}
	}

}
