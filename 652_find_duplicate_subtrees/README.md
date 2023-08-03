20230724: Worth solving again!

About the first approach's Complexity Analysis in the solution:
Space complexity is O(n^2). It looks right to me.
However, according to the master theorem, isn't time complexity theta(n log n)?
=> T(n) = 2T(n/2) + theta(n)
=> To say that theta(n log n) = O(n^2) is fine though..
