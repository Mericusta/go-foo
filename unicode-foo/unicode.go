package unicodefoo

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/width"
)

func UnicodeLengthFoo() {
	s1 := "*"
	p1, si1 := width.LookupString(s1)
	fmt.Printf("s1 = |%v|, len = %v, utf8.RuneCountInString = %v, p1 = %v, si1 = %v\n", s1, len(s1), utf8.RuneCountInString(s1), p1.Kind(), si1)
	s2 := "❤"
	p2, si2 := width.LookupString(s2)
	fmt.Printf("s2 = |%v|, len = %v, utf8.RuneCountInString = %v, p2 = %v, si2 = %v\n", s2, len(s2), utf8.RuneCountInString(s2), p2.Kind(), si2)
	s3 := "┼"
	p3, si3 := width.LookupString(s3)
	fmt.Printf("s3 = |%v|, len = %v, utf8.RuneCountInString = %v, p3 = %v, si3 = %v\n", s3, len(s3), utf8.RuneCountInString(s3), p3.Kind(), si3)
}
