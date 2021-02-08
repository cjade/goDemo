package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	fmt.Println("start")
	// aaa()
	// bbb()
	// ccc()
	// ddd()
	eee()
}

func eee() {
	var graph = make(map[string]map[string]int)
	aa := make(map[string]int)
	aa["aa"] = 12
	aa["bb"] = 12
	aa["cc"] = 12
	graph["aa"] = aa
	fmt.Println(graph["aa"]["cc"])
}

//可以用map实现类似set的功能。程序读取多行输入，但是只打印第一次出现的行
func ddd() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}

	}

	if err := input.Err(); err != nil {
		log.Fatalln(err)
	}
}

// 判断两个map是否相等
func ccc() {
	ages := map[string]int{}
	ages["alice"] = 31
	ages["charlie"] = 34
	agess := map[string]int{}
	// agess["tom"] = 20
	// agess["Jay"] = 34
	agess["alice"] = 31
	agess["charlie"] = 34

	if len(ages) != len(agess) {
		log.Fatalln("不相等")
	}

	for k, av := range ages {
		for ak, ok := agess[k]; !ok || ak != av; {
			log.Fatalln("不相等")
		}
	}
	log.Fatalln("相等")
}

// Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。
// 排序实现
func bbb() {
	ages := map[string]int{}
	ages["alice"] = 31
	ages["charlie"] = 34
	ages["tom"] = 20
	ages["Jay"] = 34
	fmt.Println("排序前")
	for name, age := range ages {
		fmt.Printf("name = %s \t\t age = %d \n", name, age)
	}
	// 因为我们一开始就知道names的最终大小，因此给slice分配一个合适的大小将会更有效
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println(names)
	// 排序
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println("排序后")

	for _, name := range names {
		fmt.Printf("name = %s \t\t age = %d \n", name, ages[name])
	}

	if _, ok := ages["aaa"]; !ok {
		log.Fatal(ages["aaa"] == 0)
	}

}

// map 的几种创建方式
func aaa() {
	// 第一种
	// ages := map[string]int{
	// 	"jade": 18,
	// }

	// 第二种相当于这种写法
	// ages := make(map[string]int)
	// ages["jade"] = 16

	// 第三种
	ages := map[string]int{}
	ages["jade"] = 16

	fmt.Println(ages["jade"])
	//删除元素
	delete(ages, "jade")
}
