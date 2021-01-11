# ADVENT OF CODE 2020

I'm doing the [Advent of Code 2020](https://adventofcode.com/2020/). This repo
contains my solutions.

My goals for this project are to:
- Gain fluency in Golang.
- Learn to better write Unit Tests.
- Get the correct solutions.

My goals for this project are not:
- Be the first to complete any problems.
- Write 'the best' solution, whatever that means.

## Day 10

### Part 1
The basic approach works here:
- Sort the input.
- Count the Jumps of Ones and Threes.
- Don't forget the additional Three at the end (device adaptor) or to insert a
zero at the start (power outlet). Because I did.

### Part 2

- Looking at the sorted input, I can see that there are no jumps of 2, only
jumps of 1 or 3.
- Since a jump of 3 is the maximum allowable, the possible paths 'converge' at
these points. Therefore I can split the sequence into smaller, monotonically
increasing sequences. I can more easily calculate the total paths in these, and
multiply them together.
- It turns out these follow a sequence, where the number of paths in a sequence is
the sum of the previous three entries in the sequence:
1,1,2,4,7,13 ...

## Day 12

### Part 2
 - Floating Point Math is expensive, therefore use integer rules.
 - A rotation of 180&deg; is simply a sign change such that `x,y -> -x,-y`
 - A rotation of +270&deg; = -90&deg; and so on.
 - For clockwise turns, `x,y(+90&deg;) = y,-x(0&deg;)`.
 - For widdershins turns, `x,y(-90&deg;) = -y,x(0&deg;)`.
 - This is all a lot easier if you assume the origin for the waypoint is the boat.

