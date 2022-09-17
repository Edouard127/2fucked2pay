package utils

import (
	"strconv"
	"strings"
)

type Queue struct {
	Position int
}

func (q *Queue) ParseString(str string) {
	s := strings.Split(str, "queue: ")
	if len(s) > 1 {
		s[1] = strings.Replace(s[1], " ", "", -1)
		for i, c := range s[1] {
			if c < '0' || c > '9' {
				s[1] = s[1][0:i]
				break
			}
		}
		n, err := strconv.Atoi(s[1])
		if err == nil {
			q.Update(n)
			return
		}
	}
}

func (q *Queue) Update(p int) {
	q.Position = p
}
