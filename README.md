gclog date split chan
=====================

Split Oracle JDK7's `gc.log` by date.

Oracle JDK7's `gc.log` lotated by file size.

```
2015-03-23T23:48:54.474+0900: 5.700: [GC2015-03-24T17:01:54.474+0900: 5.700: [DefNew: 74464K->8928K(98304K), 0.0110380 secs] 74464K->8928K(491520K), 0.0110990 secs] [Times: user=0.00 sys=0.00, real=0.01 secs]
2015-03-23T23:59:57.826+0900: 9.052: [GC2015-03-24T17:01:57.826+0900: 9.052: [DefNew: 74464K->9964K(98304K), 0.0262860 secs] 74464K->9964K(491520K), 0.0263670 secs] [Times: user=0.01 sys=0.00, real=0.02 secs]
2015-03-24T00:00:01.005+0900: 1868.231: [GC2015-03-24T17:32:57.005+0900: 1868.231: [DefNew: 75500K->12920K(98304K), 0.0456080 secs] 75500K->12920K(491520K), 0.0457090 secs] [Times: user=0.02 sys=0.01, real=0.05 secs]
2015-03-24T00:01:10.643+0900: 2573.869: [GC2015-03-24T17:44:42.643+0900: 2573.869: [DefNew: 78456K->9690K(98304K), 0.0115740 secs] 78456K->9690K(491520K), 0.0117050 secs] [Times: user=0.01 sys=0.00, real=0.01 secs]
```
Above gc.log are contains `2015-03-24`'s log and `2015-03-25`'s log.
But I want to split by date.

For example, above log would be like followings.

```
$ gclogdatesplitchan read /path/to/gc.log
$ ls
gc.log
2015-3-23_gc.log
2015-3-24_gc.log
```

Usage
-----

```
$ git checkout https://github.com/heavenshell/go-gclogdatesplitchan.git
$ go build
$ ./go-gclogdatesplitchan read /path/to/gc.log
```
