This is a basic golang implementation of pcg64 and pcg32 from http://www.pcg-random.org

PCG is a pseudorandom number generation scheme based on permuting the output
of a linear congruential generator, so that the output doesn't share the
LCG's statistical flaws and visibility into its state.

