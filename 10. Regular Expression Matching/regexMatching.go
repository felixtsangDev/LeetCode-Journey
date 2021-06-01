package main

type patternType struct {
	IsZeroOrAll bool
	Value       string
}

func extractPattern(p string) (pattern []patternType) {
	temp := patternType{}
	for i, l := range p {
		if l == '*' {
			continue
		}
		if p[i+1] != '*' {
			temp.Value += string(l)
		} else {
			if temp.Value != "" {
				temp.IsZeroOrAll = false
				pattern = append(pattern, temp)
			}
			temp.IsZeroOrAll = true
			temp.Value = string(l)
			pattern = append(pattern, temp)
			temp.Value = ""
			if i+1 >= len(p) {
				break
			}
		}
	}
	return
}

func isMatch(s string, p string) bool {
	// oneOrManyIndex := []int{}
	pattern := extractPattern(p)
	for _, pt := range pattern {
		for _, a := range s {
			if pt.IsZeroOrAll {

			}
		}
	}
	return true
}

func main() {
	isMatch("aa", "a*baaaaaaaab*.*")
}
