package convertapi

import (
	"fmt"
	"regexp"
)

func ExtractTitle(s string) (string, error) {
	// Regular expression to match the title before the file extension(s)
	
	r := regexp.MustCompile(`^(.*?)(\.\w+)*$`)

	// Find the matched string based on the regular expression
	matches := r.FindStringSubmatch(s)

	// Check if there are matches
	if len(matches) > 1 {

		return matches[1], nil
	}
	return "", fmt.Errorf("no title found")
}

// func Title(t string, ext string) string {
// 	str := strings.Split(t, fmt.Sprintf(".%v", ext))
// 	fmt.Println(str)
// 	return ""
// }