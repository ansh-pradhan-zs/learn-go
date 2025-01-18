package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for key, val := range words {
		 ss = append(ss, kv{key,val})
	}

	sort.Slice(ss, func (i,j int) bool{
		return ss[i].val > ss[j].val
	})

	fmt.Println("Top three common words are: ")

	for _, s := range ss[:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
}