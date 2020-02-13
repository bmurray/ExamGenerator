package exam

import "fmt"

type Questioner interface {
	Randomize(Rand) []Question
	Questions() []Question
}
type Question struct {
	Question       string
	Answers        []Answer
	AllOfTheAbove  bool
	NoneOfTheAbove bool
	AllIsCorrect   bool
	NoneIsCorrect  bool
}
type Answer struct {
	Answer  string
	Correct bool
}

func (q Question) Randomize(r Rand) []Question {

	nq := Question{
		Question:       q.Question,
		Answers:        make([]Answer, len(q.Answers)),
		AllOfTheAbove:  q.AllOfTheAbove,
		NoneOfTheAbove: q.NoneOfTheAbove,
		AllIsCorrect:   q.AllIsCorrect,
		NoneIsCorrect:  q.NoneIsCorrect,
	}
	for i, j := range r.Perm(len(q.Answers)) {
		nq.Answers[i] = q.Answers[j]
	}

	return []Question{nq}
}
func (q Question) Questions() []Question {

	return []Question{q}
}

func (qu Question) AllAnswers() []string {
	var c []string
	for i, a := range qu.Answers {
		x := string(i%26 + 65)
		c = append(c, fmt.Sprintf("%s. %s", x, a.Answer))
		if a.Correct {
		}
	}
	i := len(qu.Answers)
	if qu.AllOfTheAbove {
		x := string(i%26 + 65)
		c = append(c, fmt.Sprintf("%s. %s", x, "All of the above"))
		i++
	}
	if qu.NoneOfTheAbove {
		x := string(i%26 + 65)
		c = append(c, fmt.Sprintf("%s. %s", x, "None of the above"))
		i++
	}

	return c

}
func (qu Question) CorrectAnswers() []string {
	var c []string
	for i, a := range qu.Answers {
		x := string(i%26 + 65)
		if a.Correct {
			c = append(c, x)
		}
	}
	i := len(qu.Answers)
	if qu.AllOfTheAbove {
		x := string(i%26 + 65)
		if qu.AllIsCorrect {
			c = append(c, x)
		}
		i++
	}
	if qu.NoneOfTheAbove {
		x := string(i%26 + 65)
		if qu.NoneIsCorrect {
			c = append(c, x)
		}
		i++
	}
	return c
}
