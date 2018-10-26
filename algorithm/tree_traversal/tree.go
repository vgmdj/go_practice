package tree_traversal

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

//TestTree
/*

			1
         /    \
       2       3
     /  \
    4    5
    \    /
    6    7




*/

func TestTree() *Tree {
	return &Tree{
		Val: 1,
		Left: &Tree{
			Val: 2,
			Left: &Tree{
				Val: 4,
				Right: &Tree{
					Val: 6,
				},
			},
			Right: &Tree{
				Val: 5,
				Left: &Tree{
					Val: 7,
				},
			},
		},
		Right: &Tree{
			Val: 3,
		},
	}

}
