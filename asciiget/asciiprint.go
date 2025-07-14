package asciiartweb

import "fmt"

// AsciiPrint prints the ASCII art representation of the input strings
func AsciiPrint(text, strs []string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		if strs[i] == "" {
			str += "\n"
			continue
		}
		for j := 1; j < 9; j++ {
			for _, c := range strs[i] {
				if c < 32 || c > 126 {
					return ""
				}
				str += (text[(int(c)-32)*9+j])
				fmt.Println(c)
			}
			str += "\n"
		}

	}
	return str
}
