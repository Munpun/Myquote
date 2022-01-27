package myquote

import (
	"fmt"

 "rsc.io/quote"
)
func PrintAlleQuotes() {
	fmt.Prinln(MittGlass())
	fmt.Println(MittGo())
	fmt.Println(MittHello())
	fmt.Println(MittOpt())
}

func MittGlass() string {
	return MittGlass()
}
func MittGo() string {
        return quote.Go()
}
func MittHello() string {
        return quote.Hello()
}
func MittOpt() string {
        return quote.Opt()
}
