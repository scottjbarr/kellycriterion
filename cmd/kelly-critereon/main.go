package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scottjbarr/kellycriteron"
)

func main() {
	winRate := 0.0
	reward := 0.0

	flag.Float64Var(&winRate, "winrate", 0.0, "Win rate e.g. 0.4 = 40%")
	flag.Float64Var(&reward, "reward", 0.0, "The reward ratio. e.g. A return of twice the risk is 2.0")
	flag.Parse()

	rate, err := kellycriteron.Calculate(winRate, reward)
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("%0.3f\n", rate)
}
