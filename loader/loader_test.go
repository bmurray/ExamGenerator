package loader

import (
	"encoding/json"
	"math/rand"
	"strings"
	"testing"

	"github.com/bmurray/ExamGenerator/exam"
)

var data = `
questions:
- question: Test Question A
  correct:
  - A
  answers:
  - B
  - C
  - D
  all: false
  none: false
- question: Test Question B
  correct:
  - D
  answers:
  - X
  - "Y"
  - Z
  all: true
  none: false
- question: "Appendix A"
  group:
  - question: Group Question A
    correct:
    - Y
    answers:
    - Z
    - P
  - question: Group Question B
    correct: 
    - U
    answers:
    - Z
    - P`

func TestLoader(t *testing.T) {
	r := strings.NewReader(data)
	pool := &exam.Pool{}
	err := LoadQuestions(r, pool)
	if err != nil {
		t.Fatal("Cannot load questions", err)
	}
	ra := rand.New(rand.NewSource(42))

	PrintJSON(t, pool.Randomize(ra))

}

func PrintJSON(t testing.TB, i interface{}) {
	data, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Logf("Cannot marshal JSON: %v", err)
		return
	}
	t.Log(string(data))
}
