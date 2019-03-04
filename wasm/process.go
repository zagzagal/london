package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doStuff(pool, eval string) bool {
	// Handle the work
	job := deckBuilder(pool)

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
		result = check(eval, job[:handSize])
		if result {
			return result
		}

	}

	return result
}

func process(pool, eval string) {
	// setup
	n := numberOfTries
	rand.Seed(time.Now().UTC().UnixNano())

	deckTemplate := deckBuilder(pool)

	success := 0

	for w := 0; w < n; w++ {
		// Handle the work
		job := make([]byte, len(deckTemplate))
		copy(job, deckTemplate)
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

			result = check(eval, job[:handSize])
			if result {
				break
			}

		}
		if result {
			success++
		}
		updateProg(w, n)
	}

	// output
	fmt.Printf("%d of %d successes [%f%%]\n",
		success,
		n,
		(float64(success)/float64(n))*100,
	)
}
