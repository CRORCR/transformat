package main

import (
	"fmt"
	"os"
)

//点
type point struct {
	i, j int
}

func main() {
	maze := readMaze("K:/workspace/src/transformat/basic/demo12_maze/readme.txt")
	//验证数据是否正确
	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%d ", val)
	//	}
	//	fmt.Println()
	//}

	//走迷宫  从00-->右下角
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}

//返回第几行  第几列
func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error:%v\n", err)
	}

	//读取行和列
	var row, col int
	//可以按照指定格式读取
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

//定义四个方向  上 左 下 右
var dirs = []point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	//判断当前和下一步是否有效,并且保证不能越界
	//1.判断不能越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}
func walk(maze [][]int, start, end point) [][]int {
	//走过的路,存入
	step := make([][]int, len(maze))
	for i := range step {
		step[i] = make([]int, len(maze[i]))
	}

	//起点加入队列
	Q := []point{start}

	//退出条件  走到终点/队列空
	for len(Q) > 0 {
		//探索队列头
		cur := Q[0]
		Q = Q[1:]

		//如果发现是终点了,就退出
		if cur == end {
			break
		}

		for _, dir := range dirs {
			//新发现节点  == 当前节点加上方向
			next := cur.add(dir)
			//探索下一个节点  下一个节点是0,并且
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			//走过的路线,continue
			val, ok = next.at(step)
			if !ok || val != 0 {
				continue
			}
			//走回原点
			if next == start {
				continue
			}
			//走到这里就可以去探索了
			curStep, _ := cur.at(step)
			step[next.i][next.j] = curStep + 1 //步数填入框中
			//把点加入队列
			Q = append(Q, next)
		}
	}
	return step
}
