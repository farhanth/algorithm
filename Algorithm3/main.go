/* There is an input number: 1.345.679

 * OUTPUT :

 * 1000000
 * 300000
 * 40000
 * 5000
 * 600
 * 70
 * 9 */

package main

import (
	"fmt"
	"strings"
	"regexp"
)

func main() {
	var builder strings.Builder
	var stringAngka string

	fmt.Print("Masukan Angka : ")
	fmt.Scan(&stringAngka)

	regexp, _ := regexp.Compile("[^0-9]+")
	inputAngka := regexp.ReplaceAllString(stringAngka, "")

	lenAngka := len(inputAngka)

	for i, num := range inputAngka {
		builder.WriteRune(num)
		builder.WriteString(strings.Repeat("0", lenAngka-(i+1)))
		builder.WriteRune('\n')
	}

	fmt.Println(builder.String())

}