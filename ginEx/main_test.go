/**
 * File: main_test
 * DESC:
 * Author: jtzhang
 * Email: jtzhang@yoozoo.com
 * DATE:  2022/1/7 4:19 PM
 **/

package main

import "testing"

// Search for a wildcard segment and check the name for invalid characters.
// Returns -1 as index, if no wildcard was found.
func findWildcard(path string) (wildcard string, i int, valid bool) {
	// Find start
	for start, c := range []byte(path) {
		// A wildcard starts with ':' (param) or '*' (catch-all)
		if c != ':' && c != '*' {
			continue
		}

		// Find end and check for invalid characters
		valid = true
		for end, c := range []byte(path[start+1:]) {
			switch c {
			case '/':
				return path[start : start+1+end], start, valid
			case ':', '*':
				valid = false
			}
		}
		return path[start:], start, valid
	}
	return "", -1, false
}

func TestSTT(t *testing.T) {
	w, i, v := findWildcard("/add")
	t.Logf("%s=====%d---%v", w, i, v)

	w, i, v = findWildcard("/add/test/m1")
	t.Logf("%s=====%d---%v", w, i, v)

	w, i, v = findWildcard("/add/test/m1:1")
	t.Logf("%s====%d---%v", w, i, v)

	w, i, v = findWildcard("/add/test/m1:1*3")
	t.Logf("%s====%d---%v", w, i, v)

}
