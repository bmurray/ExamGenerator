package writer

import (
	"fmt"
	"io"
	"strings"

	"github.com/bmurray/ExamGenerator/exam"
)

func Write(out, ans io.Writer, questions []exam.Question) {
	for n, qu := range questions {
		fmt.Fprintf(out, "%d: %s\n", n+1, qu.Question)
		var c []string
		for i, a := range qu.Answers {
			x := string(i%26 + 65)
			fmt.Fprintf(out, "%s. %s\n", x, a.Answer)
			if a.Correct {
				c = append(c, x)
			}
		}
		i := len(qu.Answers)
		if qu.AllOfTheAbove {
			x := string(i%26 + 65)
			fmt.Fprintf(out, "%s. %s\n", x, "All of the above")
			if qu.AllIsCorrect {
				c = append(c, x)
			}
			i += 1
		}
		if qu.NoneOfTheAbove {
			x := string(i%26 + 65)
			fmt.Fprintf(out, "%s. %s\n", x, "None of the above")
			if qu.NoneIsCorrect {
				c = append(c, x)
			}
			i += 1
		}
		fmt.Fprintf(ans, "%d: %s\n", n+1, strings.Join(c, ","))
		// fmt.Fprintln(ans, qu.Question)
	}
}
