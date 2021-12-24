package ctrl

import "fmt"

// go的switch比c的更加通用 其表达式无需为常量或整数
// case语句会自上而下助益进行求值直到匹配为止
// 或switch后面没有表达式 它将匹配true
// 因此可以将一个if-else-if-else链写成一个switch
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// switch不会自动下溯 但case可通过逗号分隔来列举相同的处理条件
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

// break可以使switch提前终止，在go中可以直接设置标签并跳出
func loop() {
	var src = make([]int, 0)
loop:
	for n := 0; n < len(src); n += 1 {
		switch {
		case src[n] < 1:
			break loop
		}
	}
}

func functionOfSomeType() error {
	return nil
}

// switch也可以用来判断接口变量的动态类型
func switch_test() {
	var t interface{}
	t = functionOfSomeType()
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	}
}
