package main

/*
import (
	"io"
	"strconv"
	"strings"
	"golang.org/x/net/html"
)

func All(httpBody io.Reader) []string  {
	links := []string{}
	col := []string{}
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a"{
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
	return links
}

func trimHash(l string) string {
	if strings.Contains(l,"#") {
		var index int
		for n,str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}
//check looks to see if a url exits in the slice
func check(sl []string,s string) bool{
	var check bool
	for _,str := range sl {
		if str == s {
			check =  true
			break
		}
	}
	return check
}

func resolv(sl *[]string,ml []string){
	for _,str := range ml {
		if check(*sl,str) == false {
			*sl = append(*sl,str)
		}
	}
}
*/
