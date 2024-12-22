# Learning golang with Advent of Code

https://adventofcode.com/2024/

## Learnings
### Day 1

Gold Star

* Golang doesn't have absolute function for other types included in the standard libs. float64 is the only supported type, which requires creation of a function to turn negative values positive. 
* Assignment with `:=` and `=` is becoming clearer but need to solidify use cases for better memory of which to use.
* Able to load a file and parse it, but this seems more  cumbersome & brittle than it should be. Need to see more patterns of implementation.
* Started testing, struggled with loading packages and package management when there are several subdirectories
* Helper / Utility lib will help going forward
* Testing was very helpful to validate code
* Additional optimization and testing can improve performance
* Need to see more patterns and understand for looping with ranges better.

### Day 2

Gold Star

* More work with modules, beginning to understand utilization more with creation and movement of utils function. Use of `adventofcode.com/2024/utils=../utils` to define module name and local path was helpful to reference. 
* Use of command line tools helpful
* slog is structured logging, doesn't allow formatted lines and prefers variables in k-v pairs. This is helpful for parsing, more to get used to here, but I like it as it manageable logging. 
* Slice append and copy are all pointers - they will update original even through iterations, discovered this with some difficulty.
* Some testing can be done but current code takes less than a second to run even though there is seemingly inefficient looping going on. Current performance doesn't seem worth changing approach given readability and simplicity of functions.
* Added timer to utils to see how long it takes to run.
* Need to look more at `_` in return values - seems to be a black hole, but also you miss out on bubbling errors?
* GitHub learning, need to make main branch and use that for pull requests.
* Need to add more tests `go test utils.go utils_test.go -v`

### Day 3

Gold Star

* understanding ranges better and use of `_`
* regular expressions were straight forward, however wasnt able to leverage `\d` as shorthand for `[0-9]`
* breaking things into functions in a way where the second puzzle doesn't require refactoring is challenging
* CLI makes it easier to switch parameters but harder to debug, testing may resolve this.

### Day 4 