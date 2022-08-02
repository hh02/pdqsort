# 测试

配置：

- goos: windows
- goarch: amd64
- cpu: AMD Ryzen 5 3500U with Radeon Vega Mobile Gfx 

## 1 有序

数组长度：1e7

pdqsort: 1.441s

BenchmarkPdqsortSorted-8   	      31	  33663232 ns/op	      24 B/op	       1 allocs/op

go sort: 2.582s

BenchmarkSortSorted-8   	       2	 810051850 ns/op	      24 B/op	       1 allocs/op

约24倍

## 2 倒序

数组长度：1e7

pdqsort: 2.272s

BenchmarkPdqsortReversed-8   	      32	  35240647 ns/op	      24 B/op	       1 allocs/op

go sort: 2.763s

BenchmarkSortReversed-8   	       2	 876991650 ns/op	      24 B/op	       1 allocs/op

## 3 重复数的数组

数组长度：1e6

pdqsort:

BenchmarkPdqsortSame-8   	     303	   3471135 ns/op	      24 B/op	       1 allocs/op

go sort:

BenchmarkSortSame-8   	      73	  16750745 ns/op	      24 B/op	       1 allocs/op

## 4 短数组


数组长度：12

pdqsort:

BenchmarkPdqsortShort-8   	 9539761	       125.8 ns/op	      24 B/op	       1 allocs/op

go sort:

BenchmarkSortShort-8   	 7667515	       159.1 ns/op	      24 B/op	       1 allocs/op


# 5 随机

数组长度：1e6

pdqsort: 

BenchmarkPdqsortRand-8   	     253	   4362019 ns/op	      24 B/op	       1 allocs/op

go sort:

BenchmarkSortRand-8   	      13	  84741238 ns/op	      24 B/op	       1 allocs/op

