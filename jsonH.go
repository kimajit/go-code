package Ajit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Salary struct {
	Basic float32
	tax   float64
	total float64
}

type Employee struct {
	FirstName, LastName, Email string
	age                        int
	MonthlySalary              []Salary
}

func EmployeeSalary() (*os.File, bool) {
	set := false
	filename := "demo.json"
	newfile, _ := os.Create(filename)

	data := Employee{
		FirstName: "ajit",
		LastName:  "kumar",
		Email:     "ajithkumar.sinha@srsconsultinginc.com",
		age:       26,
		MonthlySalary: []Salary{
			{
				Basic: 10000,
				tax:   1000,
				total: 11000,
			},
		},
	}
	file, _ := json.MarshalIndent(data, "", "")
	erro := ioutil.WriteFile(newfile.Name(), file, 0644)
	if erro != nil {
		log.Fatal(erro)
	}

	fileinfo, _ := os.Stat(newfile.Name())
	log.Println(fileinfo)
	lenghtOfFile := fileinfo.Size()
	if lenghtOfFile != 0 {
		set = true

	}
	return newfile, set

}
