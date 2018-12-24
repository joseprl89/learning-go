# learning-go

A set of exercises I did to learn Go lang.

# Approach

In order to learn a bit of Go, I decided to do so by solving a few different problems from the [Euler problem](https://projecteuler.net/archives) archive.

To make it more interesting, I also trialed using the [Commit and test or revert approach as explained here](https://medium.com/@tdeniffel/tcr-test-commit-revert-a-test-alternative-to-tdd-6e6b03c22bec) by Kent Beck.

# Challenges

The main challenges encountered through it is:

* Go expects your code to be in a very specific structure. This is best explained [here](https://golang.org/doc/code.html).
* Didn't use an IDE while writing the code.
* Go is quite strict on its warnings. Unused variables trigger a build failure. 
* Lack of knowledge of Go's syntax, Go being strict on warnings, and the Commit and test or revrt approach, meant I wasted many iterations due to simple mistakes a syntax highlighter would have solved.

# Expected work

I am moving away from the Euler problems as they are heavily focused on problems that are easily solved via procedural exercises.

I'd like to start focusing on exercises that will be more challenging for the design of the code, as Go has a few characteristics that makes it differ from other languages known to me:

* No generics
* Interfaces implicitly implemented
* Go routines and channels

To do so, I will attmpt to do:

* Ping command line.
* HTTP Server based on TCP sockets.
* [Exercises on classes](https://github.com/karan/Projects#classes).


