package pkg

import (
	"fmt"
	"strings"
)

func PetuaAndStrings() {
	var fs, ss string
	_, err := fmt.Scan(&fs)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&ss)
	if err != nil {
		return
	}
	fs = strings.ToLower(fs)
	ss = strings.ToLower(ss)
	fmt.Print(strings.Compare(fs, ss))
}
