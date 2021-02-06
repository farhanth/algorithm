/* Nick works at a clothing store. He has a large pile of socks that he must pair by color for
 * sale. Given an array of integers representing the color of each sock, determine how many
 * pairs of socks with matching colors there are. */

/* For example, there are n = 7 socks with colors ar = [1, 2, 1, 2, 1, 3, 2]. There is one pair of
 * color 1 and one of color 2. There are three odd socks left, one of each color. The number of
 * pairs is 2. */

package main

import "fmt"

func main() {
	var jumlahKaosKaki int
	fmt.Print("Masukan Jumlah Kaos Kaki : ")
	fmt.Scan(&jumlahKaosKaki)

	kaosKakiArray := make([]int, 101)

	fmt.Print("Masukan Kaos Kaki (Pisahkan dengan Spasi) : ")
	for i := 0; i < jumlahKaosKaki; i++ {
		var kaosKaki int
		fmt.Scan(&kaosKaki)
		kaosKakiArray[kaosKaki]++
	}

	count := 0

	for _, v := range kaosKakiArray {
		count += v / 2
	}

	fmt.Println(count)
}