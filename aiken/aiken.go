package aiken

import (
	"github.com/bmurray/ExamGenerator/exam"
	"io"
	"fmt"
	"strings"
)

type Aiken struct {
	out io.Writer
}

func NewAiken(o io.Writer) *Aiken {
	return &Aiken{out: o}
}

func (a Aiken) AddQuestion(qu exam.Question) error {
	fmt.Fprintf(a.out, "%s\n", qu.Question)
	for _, aa := range qu.AllAnswers() {
		fmt.Fprintf(a.out, "%s\n", aa)
	}
	// fmt.Fprintf(a.out, "\n")
	fmt.Fprintf(a.out, "ANSWER: %s\n\n", strings.Join(qu.CorrectAnswers(), ","))
	return nil
}