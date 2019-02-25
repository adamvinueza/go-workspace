# Balanced Brackets

If you have ever wondered how your code editor knows that your brackets are
incorrectly spelled out, now is your opportunity to find out, by creating and
implementing the algorithm yourself!

Brackets can be any one of three pairs of characters: `{}`, `[]`, or `()`. We
will write our bracket-matching algorithm in two stages. We'll say that a
bracket character is _matched_ to another in a string `s` if the opening
bracket (the first in each of the pairs above occurs before the closing
bracket in the pair (the second character) in `s`.

## Stage 1: Determine if a string has balanced brackets

We'll say that a sequence of brackets in a string `s` is _balanced_ when:

  - `s` contains no unmatched brackets, and
  - every pair of brackets in `s` within matched brackets is also balanced

This is clearly a recursive definition, so maybe we should implement it
recursively? The trouble is that we can't make a recursive call until we have
matching brackets and a string between them. That might be tricky to
implement. Luckily, we can use a definition and a simple theorem about
matching brackets to make our lives easier:

#### Definition:

A pair of brackets `o` and `c` in `s` _directly match_ if there are no brackets
between `o` and `c`.

#### Theorem:
Any string `s` of finite length contains a balanced sequence of brackets if and only if successively eliminating directly matching brackets leaves no brackets in `s`.

#### Proof:
Any pair of nested matching brackets in a balanced sequence of brackets either directly matches or does not. If it does not, then because the sequence containing that pair is balanced, any brackets inside them are either directly matching or not; eventually, because the string is of finite length, there has to be a directly matching pair between the brackets.  Eliminating that directly matching pair must, by the same reasoning, create another directly matching pair. So eventually this process will result in `s` containing no brackets.

What's the significance of this theorem? It means that we can find wither a
string has a balanced sequence of brackets by successively eliminating
directly matching pairs.

### Stacks as matching engines

So if we're checking if the string "[()]" is balanced, we can see that "()" is
directly matching and that the surrounding square brackets are not directly
matching. If we move through the string from left to right, though, we can
as it were put the open square bracket aside, eliminate the "()", then match
the closing square bracket directly with the opening bracket. Surrounding this
string with matching brackets just adds an additional step of indirection:
delay elimination twice, by keeping the most recently encountered open bracket
closest to hand.

There's a data structure that allows us to keep track of things that have been
encountered in a way that the first thing we encounter is the last thing we
deal with: it's called a _stack_, or a _LIFO (last-in-first-out) queue_. Using
a stack to implement the algorithm should be pretty simple:

 - put each encountered open bracket onto the stack
 - when a closing bracket is encountered, look at the top of the stack and see
   if it matches
 - if it's a match, pop the opening bracket off
 - if it's not a match, the sequence is not balanced

## Stage 2: Find the offending brackets

Now that we can determine whether a sequence is balanced, how do we find the
brackets that don't properly match? The answer is simple: by keeping track of
the unbalancing brackets and reporting on them. How do we do that? Well, we
have an empty stack if the sequence is balanced; so what happens if we just
put the offending items onto the stack, and keep going until we reach the end
of the string? then we'll have all the offending items, and all we have to do
to is return whether the stack is empty.

## Stage 3: How to fix unbalanced sequences

To fix an unbalanced sequence, we just keep track of the offending items _and
where they are in the string_, keep track of the balanced sequences _and where
they begin in the string_, and return the collection, perhaps in the context
of a coherent story about what to eliminate and what you'll get after.
