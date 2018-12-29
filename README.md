# learning-go

A set of exercises I did to learn Go lang.

# Approach

In order to learn a bit of Go, I decided to do so by solving a few different problems from the [Euler problem](https://projecteuler.net/archives) archive.

To make it more interesting, I also trialed using the [Commit and test or revert approach as explained here](https://medium.com/@tdeniffel/tcr-test-commit-revert-a-test-alternative-to-tdd-6e6b03c22bec) by Kent Beck.

Halfway through it I decided to move off Euler's problems in favour of something more practical. This was mostly due to the problems being repetitively dealing with integer math, and string manipulation.

In order to have a challenge closer to real work, I opted to create:

* [HTTP Server based on TCP sockets.](https://github.com/joseprl89/learning-go/tree/master/server)
* [Ping command line](https://github.com/joseprl89/learning-go/tree/master/ping).

# Challenges

Upon initial inspection of Go, I noticed a few aspects that I expected to prove challenging to my background. Mostly:

* No generics
* Interfaces implicitly implemented
* Go routines and channels
* Separation of code and data (structs and extension functions)

Other challenges found that I wasn't expecting are:

* Go expects your code to be in a very specific structure. This is best explained [here](https://golang.org/doc/code.html).
* Didn't use an IDE while writing the code. Eventually switched to Goland trial version.
* Go is quite strict on its warnings. Thigs like unused variables/imports, or non standard formatting will trigger a build failure. This made me lose progress several times due to the test && commit || revert approach.
* Lack of knowledge of Go's syntax, Go being strict on warnings, and the Commit and test or revrt approach, meant I wasted many iterations due to simple mistakes a syntax highlighter would have solved. Switched to Goland in a trial account.
* Value vs reference type. Despite knowing the asterisk syntax from C/C++ I was confused by the fact you can implement an interface with reference semantics to mutate the object, while value types won't mutate it. See change [here](https://github.com/joseprl89/learning-go/commit/694c3aad68d9260fb4c4f9ce6a2789224f462ccf#diff-d930881d9b72deadbd3d22c7166001f9). Article for reference [here](https://medium.com/@saiyerram/go-interfaces-pointers-4d1d98d5c9c6).
* Uppercase for public, lowercase for private. Initially I thought it was following .Net's uppercasing, but quickly noticed it was not the case when I started using multiple namespaces. Using `package %PACKAGE%_test` for the test files was very helpful to help with encapsulation. See [this article](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c).
