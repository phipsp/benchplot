# benchplot

benchplot can be used to plot the output of a benchmark run on varying configuration. The configuration parameter which is subject to change needs to be part of the benchmark title separated by an underline. This parameter will then be used as the value on the x-axis.

An example would be:

```
BenchmarkTest/configParam_1-8                     1        3908798545 ns/op
BenchmarkTest/configParam_2-8                     1        44077476157 ns/op
BenchmarkTest/configParam_3-8                     1        55245503168 ns/op
BenchmarkTest/configParam_4-8                     1        71775742126 ns/op
BenchmarkTest/configParam_5-8                     1        76281380343 ns/op
BenchmarkTest/configParam_6-8                     1        85576719571 ns/op
BenchmarkTest/configParam_7-8                     1        110055220182 ns/op
BenchmarkTest/configParam_8-8                     1        87608523965 ns/op
```

Running benchplot on this input creates the following output png:

![benchplot](/example/bench.png)
