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
	fmt.Printf("p1.Folded() = %v, p1.Kind() = %v, p1.Narrow() = %v, p1.Wide() = %v\n", p1.Folded(), p1.Kind(), p1.Narrow(), p1.Wide())
	s2 := "❤"
	p2, si2 := width.LookupString(s2)
	fmt.Printf("s2 = |%v|, len = %v, utf8.RuneCountInString = %v, p2 = %v, si2 = %v\n", s2, len(s2), utf8.RuneCountInString(s2), p2.Kind(), si2)
	fmt.Printf("p2.Folded() = %v, p2.Kind() = %v, p2.Narrow() = %v, p2.Wide() = %v\n", p2.Folded(), p2.Kind(), p2.Narrow(), p2.Wide())
	s3 := "┼"
	p3, si3 := width.LookupString(s3)
	fmt.Printf("s3 = |%v|, len = %v, utf8.RuneCountInString = %v, p3 = %v, si3 = %v\n", s3, len(s3), utf8.RuneCountInString(s3), p3.Kind(), si3)
	fmt.Printf("p3.Folded() = %v, p3.Kind() = %v, p3.Narrow() = %v, p3.Wide() = %v\n", p3.Folded(), p3.Kind(), p3.Narrow(), p3.Wide())
}
