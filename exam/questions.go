package exam

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
type Group struct {
	Questions []Question
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
