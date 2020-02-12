package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/bmurray/ExamGenerator/exam"
	"github.com/bmurray/ExamGenerator/loader"
	"github.com/bmurray/ExamGenerator/webserver"
	"github.com/bmurray/ExamGenerator/writer"
)

func main() {
	var seed int64
	var yaml string
	var output string
	var answers string
	var listen string
	flag.Int64Var(&seed, "seed", 0, "Pick a seed for reproducable results, -1 is random; omit to print the raw questions")
	flag.StringVar(&yaml, "yaml", "", "Load questions from yaml file")
	flag.StringVar(&output, "out", "/dev/stdout", "Output exam to file")
	flag.StringVar(&answers, "answers", "/dev/stderr", "Output answers to file")
	flag.StringVar(&listen, "listen", "", "Start Webserver on address (eg, :8080)")
	flag.Parse()

	if yaml == "" {
		log.Fatal("Requires a file to load questions")
	}
	if seed < 0 {
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
	pool := &exam.Pool{}
	loader.LoadQuestions(f, pool)
	if len(listen) > 0 {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		ctx, cancel := context.WithCancel(context.Background())
		go func(func()) {
			<-c
			cancel()
		}(cancel)
		err := webserver.StartServer(ctx, listen, pool)
		if err != nil {
			log.Fatalf("Webserver failed: %v", err)
		}
	} else {
		var questions []exam.Question
		if seed > 0 {
			r := rand.New(rand.NewSource(seed))
			questions = pool.Randomize(r)
		} else {
			questions = pool.Questions()
		}
		writer.Write(out, ans, questions)
	}
}
