package exam

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
)

func TestPool(t *testing.T) {

	// Produce Predictable Results
	r := rand.New(rand.NewSource(42))

	pool := &Pool{}
	for i := 0; i < 100; i++ {
		var a []Answer
		for j := 0; j < 4; j++ {
			a = append(a, Answer{Correct: true, Answer: fmt.Sprintf("Answer: %d %d", i, j)})
		}

		q := &Question{
			Question: fmt.Sprintf("Test Question %d", i),
			Answers:  a,
		}
		pool.AddQuestion(q)
	}

	t.Log("Sequential Questions")
	PrintJSON(t, pool.Questions)

	random := pool.Randomize(r)
	t.Log("Random Questions")

	PrintJSON(t, random)

}

func PrintJSON(t testing.TB, i interface{}) {
	data, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		t.Logf("Cannot marshal JSON: %v", err)
		return
	}
	t.Log(string(data))
}
