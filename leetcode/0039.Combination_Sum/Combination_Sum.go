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
		backTrack(&result, []int{candidates[index]}, candidates, index, candidates[index], target)
	}

	return result

}

func backTrack(result *[][]int, combination []int, candidates []int, root, sum, target int) {
	if sum == target {
		*result = append(*result, append([]int{}, combination[:]...))
		return
	}

	for position := root; position >= 0; position-- {
		if sum+candidates[position] > target {
			continue
		}
		backTrack(result, append(combination, candidates[position]), candidates, position, sum+candidates[position], target)
	}

}
