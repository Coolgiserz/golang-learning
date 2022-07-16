package main

import (
	"fmt"
	"os"
)

//从文件中解析迷宫数据，加载到内存中
func readMaze(fname string) [][]int {

	//打开迷宫数据文件
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var row, col int
	fmt.Fscanf(f, "%d %d", &row, &col) //解析迷宫数组的大小
	fmt.Println(row, col)

	//初始化迷宫数组
	maze := make([][]int, row)
	for i := 0; i < row; i++ {
		maze[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Fscanf(f, "%d", &maze[i][j])
		}
	}
	return maze
}

//打印迷宫数据
func printMaze(maze [][]int) {
	fmt.Println("===Print Slice===")
	row := len(maze)
	for i := 0; i < row; i++ {
		// fmt.Printf("%3d", maze[i])
		fmt.Println(maze[i])

	}
}

//抽象点数据结构
type point struct {
	i, j int
}

func (p point) add(i, j int) point {
	return point{p.i + i, p.j + j}
}

//检查二维数组在某坐标点处的取值
func (p point) at(maze [][]int) (int, bool) {
	row := len(maze)
	if row < 1 {
		panic("Error: wrong maze! row < 1")
	}
	col := len(maze[0])

	//判断是否超出边界，超出则返回false
	if p.i < 0 || p.i > row-1 || p.j < 0 || p.j > col-1 {
		return 0, false
	}
	return maze[p.i][p.j], true
}

var dirs = [4][2]int{
	{0, -1}, {-1, 0}, {0, 1}, {1, 0},
}

//迷宫出路搜索
func walk(maze [][]int, start point, end point) [][]int {
	row, col := len(maze), len(maze[0])
	//------初始化
	steps := make([][]int, row) //存放搜索结果
	for i := 0; i < row; i++ {
		steps[i] = make([]int, col)
	}
	// printMaze(steps)
	// -----核心算法
	//当前点状态：已发现但未探索(steps[i][j]==0)、已发现且已探索(steps[i][j]!=0)、不可发现(maze[i][j]==1、超出边界或为start)
	// 数据结构：队列，存储待探索的点
	var queue = []point{start} //初始化队列
	for len(queue) > 0 {       //队列不为空时进行探索
		cur := queue[0] //当前所要探索的点
		queue = queue[1:]
		if cur == end { //保证找到终点时会退出
			break
		}
		//沿四个方向检测: 对每个方向检查点状态，不可发现、已经发现、或为起点则跳过循环
		for _, dir := range dirs {
			// dir[0][1]
			next := cur.add(dir[0], dir[1])
			if next.i == start.i && next.j == start.j {
				continue
			}
			//超出边界
			tmp, ok := next.at(maze)
			if !ok || tmp == 1 {
				continue
			}

			//已经发现
			tmp, ok = next.at(steps)
			if !ok || tmp != 0 {
				continue
			}

			//之前未发现过，更新steps
			curVal, _ := cur.at(steps)
			steps[next.i][next.j] = curVal + 1

			queue = append(queue, next) //添加到待探索队列
		}
	}

	// -----结果返回
	return steps
}

func constructPath(start point, end point, steps [][]int) []point {
	val, ok := end.at(steps)
	if !ok { //说明不存在路径
		panic("Path not exists!")
	}

	resultPath := make([]point, val+1) //路径长度val+1 （包括start和end）
	resultPath[0] = end
	ind := 1
	// resultPath = append(resultPath, end)
	for val > 0 {
		for _, dir := range dirs {
			cur := end.add(dir[0], dir[1])
			if tmp, _ := cur.at(steps); tmp == val-1 { //找到
				resultPath[ind] = cur
				val--
				ind++
				end = cur

			}
		}
	}
	return resultPath
}

func main() {
	//1. 读取地图数据； 2.广度优先搜索算法（输入迷宫、起点、终点，输出路径）3. 测试：功能测试、边界测试、负用例（迷宫无出路）
	//定义迷宫数据结构、
	maze := readMaze("maze.data")
	printMaze(maze)
	// fmt.Println(dirs[0])

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1}) //假设迷宫数组存在、正确
	// steps := walk(maze, point{0, 0}, point{2, 4}) //假设迷宫数组存在、正确

	// printMaze(steps)

	//打印结果
	for i := range steps {
		for j := 0; j < len(steps[0]); j++ {
			fmt.Printf("%3d", steps[i][j])

		}
		fmt.Println()
	}

	//TODO：从steps中解析从start到end路径
	path := constructPath(point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1}, steps)
	fmt.Println(path)
}
