package exam

type Rand interface {
	Shuffle(int, func(int, int))
	Perm(n int) []int
}

type Pool struct {
	Questions []*Question
}

func (p *Pool) AddQuestion(q *Question) {
	p.Questions = append(p.Questions, q)
}

func (p *Pool) Randomize(r Rand) []*Question {
	newQ := make([]*Question, len(p.Questions))
	for i, j := range r.Perm(len(p.Questions)) {
		newQ[i] = p.Questions[j].Randomize(r)
	}
	return newQ
}
