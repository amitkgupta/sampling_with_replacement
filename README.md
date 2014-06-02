## Question

Given S servers sitting behind a load balancer which spreads request according to a uniform distribution, how many requests R do you need to make to be fairly confident that you've hit each server with a request at least once? ("fairly confident" means the probability of hitting each server at least once is greater than some pre-determined confidence threshhold, e.g. P = 0.95)

## More General Question

Given a sample space of size S, let a "trial" consist of randomly (uniform distribution) picking an element from the sample space, observing it, and replacing it.  What is the probability that within T trials, you will have observed each element in the sample space at least once?

## Output

```
go run main.go
```

This will produce (to STDOUT) a comma-separated table of probability estimates, where the columns correspond to sample size and the rows correspond to the number of trials.  Sample sizes range from 1 to 20, the number of trials range from 1 to 50.  Probability estimates are calculated by repeating the experiment 50,000 times and taking the fraction of successes. 

## Results

![Results](results.png?raw=true)