package _039_Combination_Sum

import "sort"

/*
7 6 3 2
7

 6
6 3 2

   3
 3   2
3 2 2


4 3 2 1 8
4,4
4,3,1
4,2,2
4,2,1,1
4,1,1,1,1
3,3,2
3,3,1,1
3,2,2,1
3,2,1,1,1
3,1,1,1,1,1
2,2,2,2
2,2,2,1,1
2,2,1,1,1,1
2,1,1,1,1,1,1
1,1,1,1,1,1,1,1


          4
  4       3        2      1
        3 2 1     2 1     1
				    1     1
					      1

			    3
	 3	        2         1
  3  2  1     2   1       1
        1    2 1  1       1
				  1       1
						  1

			2
 	 	2       1
      2   1     1
    2  1  1     1
       1  1     1
          1     1
                1
         1
         1
		 1
		 1
		 1
		 1
		 1
		 1


*/

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	for index := range candidates {
		backTrack2(&result, []int{candidates[index]}, candidates, index, candidates[index], target)
	}

	return result

}

func backTrack2(result *[][]int, combination []int, candidates []int, root, sum, target int) {
	if sum == target {
		*result = append(*result, append([]int{}, combination[:]...))
		return
	}

	for position := root; position >= 0; position-- {
		if sum+candidates[position] > target {
			continue
		}
		backTrack2(result, append(combination, candidates[position]), candidates, position, sum+candidates[position], target)
	}

}

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	backTracking2(candidates, target, &result, []int{}, 0)
	return result

}

func backTracking2(candidates []int, target int, result *[][]int, subset []int, pos int) {
	if target == 0 {
		*result = append(*result, append([]int{}, subset[:]...))
		return
	}

	for i := pos; i < len(candidates); i++ {
		if target < candidates[i] {
			return
		}
		backTracking2(candidates, target-candidates[i], result, append(subset, candidates[i]), i)

	}
}
