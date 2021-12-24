package ctrl

import "fmt"

// go的for循环类似于c 统一了for和while不再有do while了

//for init; condition post {}
//
//for condition {}
//
//for {}

func test() {

	// 简短声明能让我们更容易在循环中声明下标变量
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	var oldMap = map[int]string{}
	var newMap = map[int]string{}

	// 遍历数组 切片 字符串或者映射 或从信道中读取消息 range可以帮你轻松实现循环
	for key, value := range oldMap {
		newMap[key] = value
	}

	// 若你只需要该遍历中的第一个项 去掉第二个就行了
	// 第一项为索引 第二项为值
	//for key := range m {
	//	if key.expired() {
	//		delete(m, key)
	//	}
	//}

	// 如果遍历第二个项 请使用空白标识符
	sum = 0
	var array = make([]int, 0)
	for _, value := range array {
		sum += value
	}
	print(sum)

	for pos, char := range "日本\\x80語" {
		fmt.Println("character %#U starts at byte position %d\n", char, pos)
	}

	// go没有逗号操作符号 而++和-- 为语句而非表达式
	var a = make([]int, 100)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

}
