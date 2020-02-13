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
		for _, a := range qu.AllAnswers() {
			fmt.Fprintf(out, "%s\n", a)
		}
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(ans, "%d: %s\n", n+1, strings.Join(qu.CorrectAnswers(), ","))
		// fmt.Fprintln(ans, qu.Question)
	}
}
