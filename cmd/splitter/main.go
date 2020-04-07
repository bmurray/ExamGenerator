package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/bmurray/ExamGenerator/aiken"
	"github.com/bmurray/ExamGenerator/exam"
	"github.com/bmurray/ExamGenerator/loader"
)

type outputs []output

func (o outputs) String() string {
	return fmt.Sprintf("OK", len(o))
}
func (o *outputs) Set(s string) error {
	log.Println("Settings", s)
	str := strings.Split(s, ":")

	if len(str) != 2 {
		return fmt.Errorf("Cannot use names, incorrect arguments")
	}
	if len(str[0]) == 0 {
		return fmt.Errorf("Need a name")
	}
	op := output{name: str[0]}

	pos := strings.Split(str[1], ",")

	for _, p := range pos {
		n, err := strconv.Atoi(p)
		if err != nil {
			return err
		}
		op.pattern = append(op.pattern, n)
	}
	*o = append(*o, op)
	log.Println(*o)
	return nil
}

func (o outputs) Files(dir string) ([]outfile, error) {
	max := 0
	pos := make(map[int]*of)
	for _, op := range o {
		f, err := os.OpenFile(filepath.Join(dir, op.name), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0660)
		if err != nil {
			return nil, err
		}
		a := aiken.NewAiken(f)
		k := &of{aiken: a, wr: f}
		for _, pt := range op.pattern {
			if pt > max {
				max = pt
			}
			pos[pt] = k
		}
	}
	if max == 0 {
		return nil, fmt.Errorf("No outputs")
	}
	of := make([]outfile, max+1)
	for i := 0; i <= max; i++ {
		l, ok := pos[i]
		if !ok {
			return nil, fmt.Errorf("Unknown position: %d", i)
		}
		of[i] = l
	}
	return of, nil
}

type of struct {
	aiken *aiken.Aiken
	wr    io.WriteCloser
}

func (o *of) Close() {
	if o.wr == nil {
		return
	}
	o.wr.Close()
	o.wr = nil
}
func (o of) AddQuestion(qu exam.Question) {
	o.aiken.AddQuestion(qu)
}

type outfile interface {
	AddQuestion(exam.Question)
	Close()
}

type output struct {
	name    string
	pattern []int
}

func main() {
	var seed int64
	var yaml string
	var dir string
	var out outputs

	flag.Int64Var(&seed, "seed", 0, "Pick a seed for reproducable results, -1 is random; omit to print the raw questions")
	flag.StringVar(&yaml, "yaml", "", "Load questions from yaml file")
	flag.StringVar(&dir, "dir", "", "Output Directory for Aiken Files")
	flag.Var(&out, "out", "Output names. name:n[,n...], where n is the position")
	flag.Parse()

	// if err := out.Validate(); err != nil {
	// 	log.Fatalf("Cannot validate inputs: %s", err)
	// }
	// total := 0
	// for _, o := range out {
	// 	// total += o.total
	// 	log.Println(o.pattern)
	// }

	if yaml == "" {
		log.Fatal("Requires a file to load questions")
	}
	if seed < 0 {
		seed = time.Now().UnixNano()
	}
	stat, err := os.Stat(dir)
	if err != nil {
		log.Fatal("Cannot stat output directory", err)
	}
	if !stat.IsDir() {
		log.Fatalf("Output dir is not a directory: %s", dir)
	}
	f, err := os.Open(yaml)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	pool := &exam.Pool{}
	err = loader.LoadQuestions(f, pool)
	if err != nil {
		log.Fatal("Cannot load questions", err)
	}
	var questions []exam.Question
	if seed > 0 {
		r := rand.New(rand.NewSource(seed))
		questions = pool.Randomize(r)
	} else {
		questions = pool.Questions()
	}
	_ = questions
	files, err := out.Files(dir)
	if err != nil {
		log.Fatalf("Cannot generate output files: %s", err)
	}
	l := len(files)
	for i, question := range questions {
		files[i%l].AddQuestion(question)
	}
	for _, f := range files {
		f.Close()
	}
}

/*
0 C
1 C
2 B
3 C
4 C
5 B
6 C
7 B

*/
