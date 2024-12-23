package main

import (
	"fmt"
	"log"

	"github.com/expr-lang/expr"
)

func main() {
	fmt.Println("main start")

	env := map[string]any{
		"creditScore": 810,
	}

	expressions := []string{
		"creditScore >= 700",
		"creditScore >= 680",
		"creditScore >= 1000",
	}

	var programs []any
	//var programs []*vm.Program <-- unable to do this since vm.Program is not accessible.

	for _, exprStr := range expressions {
		program, err := expr.Compile(exprStr)
		if err != nil {
			log.Fatalf("failed to compile expression: ", exprStr)
		}
		programs = append(programs, program)
	}

	for _, program := range programs {
		//compiledProgram, ok := program.(*expr.Program) <-- unable to do this. type not accessible
		compiledProgram, ok := program.(any)
		if !ok {
			log.Printf("program has an invalid type")
		}

		result, err := expr.Run(compiledProgram, env) // causes error: need to assert type, but not able to.
		if err != nil {
			fmt.Errorf("error evaluating...")
		}
		fmt.Println("the result is: ", result)
	}
}
