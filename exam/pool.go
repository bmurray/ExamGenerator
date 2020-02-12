package exam

type Rand interface {
	Shuffle(int, func(int, int))
	Perm(n int) []int
}

type QuestionPool interface {
	AddQuestion(Questioner)
}

type Pool struct {
	QuestionList []Questioner
}

func (p *Pool) AddQuestion(q Questioner) {
	p.QuestionList = append(p.QuestionList, q)
}
func (p Pool) Questions() []Question {
	newQ := make([]Question, 0, len(p.QuestionList))
	for _, q := range p.QuestionList {
		newQ = append(newQ, q.Questions()...)
	}
	return newQ
}

func (p Pool) Randomize(r Rand) []Question {
	newQ := make([]Question, 0, len(p.QuestionList))
	for _, j := range r.Perm(len(p.QuestionList)) {
		x := p.QuestionList[j].Randomize(r)
		// newQ[i] = p.Questions[j].Randomize(r)
		newQ = append(newQ, x...)
	}
	return newQ
}
