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
	Question  string   `yaml:"question"`
	Correct   []string `yaml:"correct"`
	Answers   []string `yaml:"answers"`
	All       bool     `yaml:"all"`
	None      bool     `yaml:"none"`
	IsCorrect string   `yaml:"iscorrect"`
}

func LoadQuestions(r io.Reader, pool *exam.Pool) error {
	decoder := yaml.NewDecoder(r)

	q := &Questions{}
	err := decoder.Decode(q)
	if err != nil {
		return fmt.Errorf("Cannot decode input: %w", err)
	}
	for _, question := range q.Q {
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

		pool.AddQuestion(&exam.Question{
			Answers:        answers,
			Question:       question.Question,
			AllOfTheAbove:  question.All,
			NoneOfTheAbove: question.None,
			AllIsCorrect:   allIsCorrect,
			NoneIsCorrect:  noneIsCorrect,
		})

	}

	return nil
}
