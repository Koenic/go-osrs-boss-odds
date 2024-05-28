# Osrs boss completion rates

This is a rewrite of an older project of mine because I wanted to experiment with go.

## Explanation

[Absorbing Markov chains](https://en.wikipedia.org/wiki/Absorbing_Markov_chain) can be used to model states in which one has received a combination of drops. If we model the absorbing to be the state where a player has gotten every drop, we can easily calculate the expected number of steps it takes to be absorbed.

However if we want pretty pictures we can use another probably of transition matrixes. Namely that if we the power n of the matrix the resulting matrix gives the probability of transitioning to a new state in n steps. If we take the resulting matrix and look at the probability of ending up in the absorbing state we know the propbably at that kc.

We can construct the cdf/pdf functions and calculate mean/mode/median based on those.

## Caveats

This only works if we can model the drops as a transition matrix, we can do this in the general case but drops like olmlet and anti dry mechanics need to be calculated individually. That is a bit of a pain so I haven't done so until it becomes relevant.

The number of states in the transition matrix is rougly equal to (2 ^ wanted_drops) since we need to thousands of matrix multiplications per graph the matrix cant be too large. Around 12 drops seems to be the limit of my pc.

Some drop mechanics I'm not 100% sure of, I've gone of the wiki but if I misunderstood it or the mechanics aren't completely understood as of now the results could be slightly of.

## Barrows

Barrows has 24 items which is more than 12, but we can take advantage of the fact that all drops have the same chance of being rolled to make it a 25 x 25 matrix instead of one thats (2 ^ 25) which is slightly faster.
