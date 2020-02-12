package loader

import (
	"fmt"
	"io"

	"github.com/bmurray/ExamGenerator/exam"
	"gopkg.in/yaml.v2"
)

type Questions struct {
	Q []Question `yaml:"questions"`
}
type Question struct {
	Question  string     `yaml:"question"`
	Correct   []string   `yaml:"correct"`
	Answers   []string   `yaml:"answers"`
	All       bool       `yaml:"all"`
	None      bool       `yaml:"none"`
	IsCorrect string     `yaml:"iscorrect"`
	Group     []Question `yaml:"group"`
}

func LoadQuestions(r io.Reader, pool exam.QuestionPool) error {
	decoder := yaml.NewDecoder(r)

	q := &Questions{}
	err := decoder.Decode(q)
	if err != nil {
		return fmt.Errorf("Cannot decode input: %w", err)
	}
	for _, question := range q.Q {
		addToPool("", question, pool)
	}
	return nil
}
func addToPool(prefix string, question Question, pool exam.QuestionPool) error {
	if len(question.Group) > 0 {
		return addGroup(prefix, question, pool)
	} else {
		return addQuestion(prefix, question, pool)
	}
}
func addGroup(prefix string, question Question, pool exam.QuestionPool) error {
	p2 := &exam.Pool{}
	pr := question.Question
	if len(prefix) > 0 {
		pr = fmt.Sprintf("%s, %s", prefix, question.Question)
	}
	for _, nq := range question.Group {
		err := addToPool(pr, nq, p2)
		if err != nil {
			return err
		}
	}
	pool.AddQuestion(p2)
	return nil
}
func addQuestion(prefix string, question Question, pool exam.QuestionPool) error {

	if question.Question == "" {
		return fmt.Errorf("Invalid question text")
	}
	var answers []exam.Answer
	for _, c := range question.Correct {
		answers = append(answers, exam.Answer{Correct: true, Answer: c})
	}
	for _, c := range question.Answers {
		answers = append(answers, exam.Answer{Correct: false, Answer: c})
	}
	allIsCorrect := false
	noneIsCorrect := false
	allIsCorrect = question.IsCorrect == "all"
	noneIsCorrect = question.IsCorrect == "all"

	pr := question.Question
	if len(prefix) > 0 {
		pr = fmt.Sprintf("(%s) %s", prefix, question.Question)
	}
	pool.AddQuestion(exam.Question{
		Answers:        answers,
		Question:       pr,
		AllOfTheAbove:  question.All,
		NoneOfTheAbove: question.None,
		AllIsCorrect:   allIsCorrect,
		NoneIsCorrect:  noneIsCorrect,
	})
	return nil
}
