package lesson10

//FallingSquares: 掉落的方块
/*
在无限长的数轴（即 x 轴）上，我们根据给定的顺序放置对应的正方形方块。
第 i 个掉落的方块（positions[i] = (left, side_length)）是正方形，其中 left 表示该方块最左边的点位置(positions[i][0])，side_length 表示该方块的边长(positions[i][1])。
每个方块的底部边缘平行于数轴（即 x 轴），并且从一个比目前所有的落地方块更高的高度掉落而下。在上一个方块结束掉落，并保持静止后，才开始掉落新方块。
方块的底边具有非常大的粘性，并将保持固定在它们所接触的任何长度表面上（无论是数轴还是其他方块）。邻接掉落的边不会过早地粘合在一起，因为只有底边才具有粘性。
返回一个堆叠高度列表 ans 。每一个堆叠高度 ans[i] 表示在通过 positions[0], positions[1], ..., positions[i] 表示的方块掉落结束后，目前所有已经落稳的方块堆叠的最高高度。
*/
/*parameter
positions: 二维整型数组
return: 整型数组
*/
func FallingSquares(positions [][]int) []int {

	tmp := make([]int, len(positions))
	for i := 0; i < len(positions); i++ {

		left1 := positions[i][0]
		size1 := positions[i][1]
		right1 := left1 + size1
		tmp[i] += size1

		for j := i + 1; j < len(positions); j++ {

			left2 := positions[j][0]
			size2 := positions[j][1]
			right2 := left2 + size2

			if left2 < right1 && left1 < right2 {

				tmp[j] = max(tmp[i], tmp[j])
			}
		}
	}

	result := make([]int, 0)
	m := -1
	for i := range tmp {

		m = max(m, tmp[i])
		result = append(result, m)
	}

	return result
}

func max(x, y int) int {

	if x >= y {

		return x
	}

	return y
}
