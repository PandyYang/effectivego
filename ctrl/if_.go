package ctrl

import (
	"log"
	"os"
)

func if_test() int {
	x, y := 1, 2
	if x > 0 {
		return y
	}
	return x
}

// if switch可以接受初始化语句 设置局部变量
func if_test2(file os.File) error {
	if err := file.Chmod(0664); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// if语句如果不会执行到下一条语句 执行体以break continue goto return结束
// 不必要的else将会被省略
func if_test3(file os.File, name string) error {
	_, err := os.Open(name)
	if err != nil {
		return err
	}
	// do something
	return nil
}

func declareAndAllocate(name string) error {
	f, err := os.Open(name)

	// err的重复是合法的 err在第一条语句中被声明 第二条语句中被再次赋值
	a, err := f.Stat()

	print(a)
	return err
}

func declareAndAllocate2(name string) error {
	f, err := os.Open(name)

	// 注意下面的 =（声明语句） 而不是 := （赋值语句）
	// 满足以下条件时 已被声明的变量V可出现在 := 声明中
	// 1. 本次已声明的v处于同一作用域中
	// 2. 在初始化中与其类型相应的值才能赋予v 且
	// 3*. 在此次声明中至少有另有一个新的变量是新声明的
	_, err = f.Stat()

	return err
}
