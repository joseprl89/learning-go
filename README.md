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
* Lack of knowledge of Go's syntax, Go being strict on warnings, and the Commit and test or revrt approach, meant I wasted many iterations due to simple mistakes a syntax highlighter would have solved. Switched to Goland in a trial account.
* Value vs reference type. Despite knowing the asterisk syntax from C/C++ I was confused by the fact you can implement an interface with reference semantics to mutate the object, while value types won't mutate it. See change [here](https://github.com/joseprl89/learning-go/commit/694c3aad68d9260fb4c4f9ce6a2789224f462ccf#diff-d930881d9b72deadbd3d22c7166001f9). Article for reference [here](https://medium.com/@saiyerram/go-interfaces-pointers-4d1d98d5c9c6)

# Expected work

I am moving away from the Euler problems as they are heavily focused on problems that are easily solved via procedural exercises.

I'd like to start focusing on exercises that will be more challenging for the design of the code, as Go has a few characteristics that makes it differ from other languages known to me:

* No generics
* Interfaces implicitly implemented
* Go routines and channels
* Only structs and no classes

To do so, I will attmpt to do, if time allows:

* [HTTP Server based on TCP sockets.](https://github.com/joseprl89/learning-go/tree/master/server)
* [Ping command line](https://github.com/joseprl89/learning-go/tree/master/ping).
* [Exercises on classes](https://github.com/karan/Projects#classes).


