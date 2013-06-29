package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximum_weighted_subset(weights []int, bound int) int {
	if (len(weights) == 0) {
		return 0
	}

	maxWeight := make([]int, bound + 1)

	for w := 0; w <= bound; w++ {
		if weights[0] <= w {
			maxWeight[w] = weights[0]
		}
	}

	for i := 1; i < len(weights); i++ {
		for w := bound; w >= 0; w-- {
			if weights[i] <= w {
				include := weights[i] + maxWeight[w - weights[i]]
				exclude := maxWeight[w]
				maxWeight[w] = max(include, exclude)
			}   
		}   
	}

	return maxWeight[bound]
}

func assert_equal(x, y int) {
	if x != y {
		panic(fmt.Sprintf("not equal: %d and %d", x, y))
	}
}

func main() {
	assert_equal(maximum_weighted_subset([]int {}, 9), 0)
	assert_equal(maximum_weighted_subset([]int {1, 3, 6}, 9), 9)
	assert_equal(maximum_weighted_subset([]int {1, 3, 6}, 8), 7)
	assert_equal(maximum_weighted_subset([]int {1, 2, 3}, 7), 6)
	assert_equal(maximum_weighted_subset([]int {1, 2, 3}, 6), 6)
	assert_equal(maximum_weighted_subset([]int {1, 2, 3}, 5), 5)
	assert_equal(maximum_weighted_subset([]int {1, 2, 3}, 4), 4)
	assert_equal(maximum_weighted_subset([]int {1, 2, 3}, 3), 3)
	assert_equal(maximum_weighted_subset([]int {1, 2, 3}, 2), 2)
	assert_equal(maximum_weighted_subset([]int {1, 2, 3}, 1), 1)
	assert_equal(maximum_weighted_subset([]int {1, 2, 4}, 7), 7)
	assert_equal(maximum_weighted_subset([]int {3, 5, 7}, 6), 5)
}
