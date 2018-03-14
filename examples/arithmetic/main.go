package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/liuzl/fmr/bnf"
	"github.com/liuzl/goutil"
	"io/ioutil"
	//"os"
)

var inputs = []string{
	//"one", "two", "ten",
	//"minus three minus two",
	"two times two plus three",
	"one add two multiply by two plus three",
	"二加五减三",
	"我的二加五减三",
	/*
		"one plus one",
		"one plus two",
		"one plus three",
		"two plus two",
		"two plus three",
		"three plus one",
		"three plus minus two",
		"two plus two",
		"three minus two",
		"two times two",
		"two times three",
	*/
	//"three plus three minus two",
	//"two times two plus three",
	//"minus four",
}

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile("arithmetic.grammar")
	if err != nil {
		glog.Fatal(err)
	}
	//bnf.Debug = true
	g, err := bnf.CFGrammar(string(b))
	if err != nil {
		glog.Fatal(err)
	}
	//b, _ = goutil.JsonMarshalIndent(g, "", " ")
	//fmt.Println(string(b))
	for _, input := range inputs {
		fmt.Println(input)
		p, err := g.EarleyParse("number", input)
		if err != nil {
			glog.Fatal(err)
		}
		trees := p.GetTrees()
		//fmt.Printf("%+v\n", p)
		fmt.Println("tree number:", len(trees))
		for _, tree := range trees {
			//tree.Print(os.Stdout)
			sem, err := tree.Semantic()
			if err != nil {
				glog.Error(err)
			}
			fmt.Println(sem)
			_, err = goutil.JsonMarshalIndent(tree, "", " ")
			if err != nil {
				glog.Fatal(err)
			}
		}
		fmt.Println()
	}
}