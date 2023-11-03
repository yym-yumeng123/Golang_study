package main

import "fmt"

func main() {
	//// map声明没有分配空间
	//var a map[string]string
	//a["a"] = "yym" // panic: assignment to entry in nil map
	//fmt.Println(a)

	// key 是唯一的
	var b map[string]string = make(map[string]string, 10)
	b["name"] = "张三"
	b["name1"] = "李四"
	b["name2"] = "张三"
	fmt.Println(b) // map[name:张三]
	delete(b, "name")
	fmt.Println(b) // map[name:张三]

	// 使用方式
	var city = make(map[string]string)
	city["n1"] = "北京"
	city["n2"] = "上海"
	city["n3"] = "天津"

	var cityMap = map[string]string{
		"h1": "杭州",
		"h2": "阜阳",
	}

	fmt.Println(city, cityMap)

	// 存放多个学生,  每个学生有姓名和sex 信息
	var students = make(map[string]map[string]string)
	students["张三"] = make(map[string]string)
	students["张三"]["sex"] = "男"
	students["张三"]["address"] = "胡同"
	students["李四"] = make(map[string]string)
	students["李四"]["sex"] = "女"
	students["李四"]["address"] = "胡同1"
	fmt.Println(students)

	for k, v := range students {
		fmt.Println(k, v)
		fmt.Println(v["address"])
	}

	// map切片
	// 1.声明一个map切片
	monsters := make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "18"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "小狐狸"
		monsters[1]["age"] = "18"
	}
	fmt.Println(monsters)

	// map是引用类型
	map1 := make(map[int]int, 2)
	map1[10] = 90
	map1[5] = 910
	map1[1] = 900
	modify(map1)
	fmt.Println(map1)

	map2 := make(map[string]Student)
	s1 := Student{"yym", 18, "男"}
	s2 := Student{"yym1", 18, "男"}
	map2["n1"] = s1
	map2["n2"] = s2

	fmt.Println(map2)

	// 练习
	map3 := make(map[string]map[string]string)
	modifyUser(map3, "nickname")
	fmt.Println(map3)
}

type Student struct {
	Name string
	Age  int
	Sex  string
}

func modify(map1 map[int]int) {
	map1[10] = 900
}

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["password"] = "888888"
	} else {
		// 没有这个用户
		users["nickname"] = make(map[string]string, 2)
		users[name]["password"] = "88888"
		users[name]["nickname"] = "昵称" + name
	}
}
