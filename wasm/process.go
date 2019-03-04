package main

import (
	"fmt"
	"math/rand"
)

func inputVal(poo, eval string) bool {
	job, err := deckBuilder(pool)
	if err != nil {
		return false
	}

	_, err = handCheck(eval, job[:handSize])
	if err != nil {
		return false
	}
	return true
}

func deckCheck(pool, eval string) bool {
	// Handle the work
	job, err := deckBuilder(pool)
	if err != nil {
		return false
	}

	result := false
	for try := handSize; try != maxMul-1; try-- {
		rand.Shuffle(len(job), func(i, j int) {
			job[i], job[j] = job[j], job[i]
		})
		if debug {
			fmt.Printf("try [%d/%d]\n",
				handSize-try+1,
				handSize-maxMul+1,
			)
		}
		result, _ = handCheck(eval, job[:handSize])
		if result {
			return result
		}

	}

	return result
}
