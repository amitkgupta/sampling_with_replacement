package main

import "fmt"
import "math/rand"

const MAX_SAMPLE_SIZE = 20
const MAX_TRIALS = 50
const EXPERIMENTS_FOR_ESTIMATE = 50000
const EXPERIMENTS_FOR_ESTIMATE_AS_FLOAT = 50000.0

func main() {
  for trials := 1; trials < MAX_TRIALS; trials++ {
    for sample_size := 1; sample_size < MAX_SAMPLE_SIZE; sample_size++ {
      fmt.Printf("%.4f,", est_prob_of_sampling_all(sample_size, trials))
    }
    fmt.Printf("\n")
  }
}

func est_prob_of_sampling_all(sample_size int, trials int) float64 {
  failures := 0 

  for experiment := 0; experiment < EXPERIMENTS_FOR_ESTIMATE; experiment++ {
    seen := make([]bool, sample_size)

    for trial := 0; trial < trials; trial++ {
      seen[rand.Intn(sample_size)] = true
    }

    for element := 0; element < sample_size; element++ {
      if !seen[element] {
        failures++
        break
      }
    }
  }

  return float64(EXPERIMENTS_FOR_ESTIMATE - failures) / EXPERIMENTS_FOR_ESTIMATE_AS_FLOAT
}
