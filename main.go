package main

import (
	"github.com/NWBY/go-practice/functions"
	"github.com/NWBY/go-practice/pointers"
	"github.com/NWBY/go-practice/structs"
)

func main() {
	// Functions code
	functions.StandardFunction()
	functions.FunctionWithParams("Sam")
	functions.FunctionWithReturnType()
	functions.VariadicFunction("Wendy", "Courtney")
	functions.RecursiveFunction(15)

	// Pointers code
	pointers.PointersOutter()

	// Struct and methods code
	// Defining a struct
	// Struct literal definition
	b1 := structs.Business{
		Name: "Render Labs Limited",
		Industry: "Software",
		Founded: 2020,
	}

	// Dot notation definition
	var b2 structs.Business
	b2.Name = "Sotent"
	b2.Industry = "Social Media"
	b2.Founded = 2021

	// New function definition
	b3 := new(structs.Business)
	b3.Name = "Hostinger"
	b3.Industry = "Hosting"
	b3.Founded = 2005


	b1.GetName()
	b1.GetIndustry()
	b1.GetYearFounded()

	b2.GetName()
	b2.GetIndustry()
	b2.GetYearFounded()

	b3.GetName()
	b3.GetIndustry()
	b3.GetYearFounded()
}