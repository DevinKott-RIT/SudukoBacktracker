# SudukoBacktracker
A Suduko Solver (Backtracker) written in Go.

This is my first ever program in Go.

# Usage

**Command Line**: `./SudukoSolver <file>`

**IDE**: Edit your run configuration, add a file name to the program arguments.

# Input Files

Input files must contain 81 numbers only. These numbers may only be 1 through 9 (inclusive).
There must be 9 lines, with 9 numbers each, spaced by whitespace. An example of one line could
be:

`1 2 3 4 5 6 7 8 9`

There can be multiple whitespaces between each number, but there must be 9 numbers total. Empty
lines in the input file are skipped. All of these rules allow for easy game formatting, like so:

![example input][pic1]

[pic1]: https://i.gyazo.com/a00ad2093d46936c5a36507e8f75291c.png

# Output

The output of the program (the solution, if there is one) is printed out to the console. Here
are the outputs of the 3 puzzles included in the repository:

- [easy.txt](https://gist.github.com/DevinKott-RIT/abe7484dbae12111279675fdac64016a)
- [hard.txt](https://gist.github.com/DevinKott-RIT/6d3cb753126d2e4cae912b2ecc1ade6f)
- [input.txt](https://gist.github.com/DevinKott-RIT/50f1c64da8fd1b6d5dbb4a689aa7d735)
- [easy_failure.txt](https://gist.github.com/DevinKott-RIT/98b021dfce1e46ff66bc021521d8e070)

# Future

These are planned features. Who knows if I'll get around to them.

1. Create a GUI to show the progression to the solution.
2. Create different methods to read in the game. Would like it so it just reads in 81 numbers on
a single line, and formats it into the 9 x 9 matrix.
3. Create a Suduko puzzle generator.
