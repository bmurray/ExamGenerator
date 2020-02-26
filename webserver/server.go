package webserver

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/bmurray/ExamGenerator/exam"
)

type Server struct {
	mux           *http.ServeMux
	pool          exam.Questioner
	templateFuncs template.FuncMap
}

func (s *Server) AttachServer() {
	s.mux.HandleFunc("/answers", s.getAnswers())
	s.mux.HandleFunc("/questions", s.generateQuestions())
}

func (s *Server) getAnswers() func(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("answers").Funcs(s.templateFuncs).Parse(answer_templ)
	if err != nil {
		log.Fatal("Cannot parse answers template", err)
	}
	return func(w http.ResponseWriter, req *http.Request) {
		seed, questions := s.getQuestions(req)
		err := t.Execute(w, struct {
			Seed      int64
			Questions []exam.Question
		}{
			Seed:      seed,
			Questions: questions,
		})
		if err != nil {
			log.Println("Error executing template", err)
		}
	}
}
func (s *Server) generateQuestions() func(w http.ResponseWriter, req *http.Request) {
	// questions := s.getQuestions(req)
	t, err := template.New("questions").Funcs(s.templateFuncs).Parse(question_templ)
	if err != nil {
		log.Fatal("Cannot parse answers template", err)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		seed, questions := s.getQuestions(req)
		err := t.Execute(w, struct {
			Seed      int64
			Questions []exam.Question
		}{seed, questions})
		if err != nil {
			log.Println("Error executing template", err)
		}
	}
}
func (s Server) getQuestions(req *http.Request) (int64, []exam.Question) {
	seedQ := req.URL.Query().Get("seed")
	seed, _ := strconv.ParseInt(seedQ, 10, 64)
	var questions []exam.Question
	if seed > 0 {
		r := rand.New(rand.NewSource(seed))
		questions = s.pool.Randomize(r)
	} else {
		questions = s.pool.Questions()
	}
	return seed, questions
}

// StartServer starts a web server to print question pools
func StartServer(ctx context.Context, listen string, pool exam.Questioner) error {

	mux := http.NewServeMux()
	svr := &Server{
		mux:  mux,
		pool: pool,
		templateFuncs: template.FuncMap{
			"oneIndex": func(idx int) int {
				return idx + 1
			},
		},
	}
	svr.AttachServer()

	s := &http.Server{
		Addr:           listen,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		<-ctx.Done()
		c, _ := context.WithTimeout(context.Background(), 10*time.Second)
		s.Shutdown(c)
	}()

	return s.ListenAndServe()
}
