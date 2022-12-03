package cppat

func init() {

	fourPermutationsOf6x3x3[1] = []preferences3x3T{
		{
			ID: 1, // Frage 1
			Ppls: [][]int{
				{0, 2, 3},
				{2, 3, 0},
				{3, 0, 2},
			},
		},
		{
			ID: 2, // Frage 2
			Ppls: [][]int{
				{0, 3, 2},
				{2, 2, 1},
				{3, 0, 2},
			},
		},
		{
			ID: 3, // Frage 3
			Ppls: [][]int{
				{2, 0, 3},
				{0, 5, 0},
				{3, 0, 2},
			},
		},
		{
			ID: 4, // Frage 4
			Ppls: [][]int{
				{1, 0, 4},
				{1, 4, 0},
				{3, 1, 1},
			},
		},
		{
			ID: 5, // Frage 5 - rows permutation 2
			Ppls: [][]int{
				{2, 1, 2}, // 2021-05-17 sandro; unchanged
				{2, 0, 3}, // 2021-05-17 sandro; previously {0, 4, 1}
				{1, 4, 0}, // 2021-05-17 sandro; previously {3, 0, 2}
			},
		},
		{
			ID: 6, // Frage 6
			Ppls: [][]int{
				{1, 0, 4},
				{0, 5, 0},
				{4, 0, 1},
			},
		},
	}

}
