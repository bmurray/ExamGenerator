package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bmurray/ExamGenerator/exam"
	"github.com/bmurray/ExamGenerator/loader"
	"github.com/bmurray/ExamGenerator/writer"
)

func main() {
	var seed int64
	var yaml string
	var output string
	var answers string
	flag.Int64Var(&seed, "seed", 0, "Pick a seed for reproducable results")
	flag.StringVar(&yaml, "yaml", "", "Load questions from yaml file")
	flag.StringVar(&output, "out", "/dev/stdout", "Output exam to file")
	flag.StringVar(&answers, "answers", "/dev/stderr", "Output answers to file")
	flag.Parse()

	if yaml == "" {
		log.Fatal("Requires a file to load questions")
	}
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	f, err := os.Open(yaml)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	out, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Cannot open output file", err)
	}
	ans, err := os.OpenFile(answers, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Cannot open answers file", err)
	}
	r := rand.New(rand.NewSource(seed))
	pool := &exam.Pool{}
	loader.LoadQuestions(f, pool)
	writer.Write(out, ans, pool.Randomize(r))
}
