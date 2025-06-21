package main

import (
	"fmt"

	"theparadance.com/quan-lang/env"
	lang "theparadance.com/quan-lang/quan-lang"
)

func main() {
	program := `
		fn calculateInterest(principal, rate, time) {
			return principal * rate * time / 100;
		}
		fn calculateInterest2(principal, rate, time) {
			return principal * rate * time / 100;
		}
		interest = calculateInterest(loanAmount, 3, 1);
	`

	env, _ := lang.Execuate(program, &env.Env{
		Vars: map[string]interface{}{"loanAmount": 100000},
	})

	// fmt.Println("interest =", env.Vars["interest"]) // Should print the calculated interest
	fmt.Printf("format: %f\n", env.Vars["interest"])
}
