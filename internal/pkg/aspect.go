package pkg

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

type AspectFragment struct {
	Aspect string    `json:"aspect"`
	Value  string    `json:"value"`
	Time   time.Time `json:"time"`
}

type Aspect struct {
	Name     string        `yaml:"name"`
	Type     string        `yaml:"type"`
	Path     string        `yaml:"path"`
	Pattern  string        `yaml:"pattern"`
	Delay    time.Duration `yaml:"delay"`
	NextTime time.Time
}

func (a *Aspect) FetchData() (fragment *AspectFragment, err error) {

	if a.Type == "parser" {
		fragment, err = a.parser()
		if err != nil {
			fmt.Println("err")
			return
		}
	} else {
		return nil, errors.New("parser type not found")
	}
	a.NextTime = time.Now().Add(a.Delay)
	fmt.Println(a.NextTime)
	return
}

func (a *Aspect) parser() (*AspectFragment, error) {
	b1 := make([]byte, 9999)
	f, _ := os.Open(a.Path)
	f.Read(b1)
	arr := strings.Split(string(b1), "\n")
	for _, line := range arr {
		match, _ := regexp.MatchString(a.Pattern, line)
		if match {
			re := regexp.MustCompile(a.Pattern)
			res := re.ReplaceAllString(line, "$1")
			return &AspectFragment{
				Aspect: a.Name,
				Value:  res,
				Time:   time.Now(),
			}, nil
		}
	}
	return nil, errors.New("not found")
}
