package algorithmfoo

import (
	"fmt"
	"math/rand"
	"time"
)

// 基于 线性回归 实现的梯度下降算法

func GradientDescentFoo() {
	random, x, y := generateRandomData(100) // 生成 100 条随机数据
	biasedX := appendBiasItem(x)            // 添加偏置项

	// 批量梯度下降
	batchGradientDescentLearning := 0.01
	batchGradientDescentLearningIterations := 1000
	batchTheta, batchLoss := batchGradientDescent(random, biasedX, y, batchGradientDescentLearning, batchGradientDescentLearningIterations)
	_, _ = batchTheta, batchLoss
	fmt.Printf("batchTheta %v\n", toLine(batchTheta))
}

func toLine(s [][]float64) string {
	l := make([]float64, len(s))
	for row := 0; row < len(s); row++ {
		l[row] = s[row][0]
	}
	return fmt.Sprintf("%v", l)
}

// 批量梯度下降
// - learningRate : 学习率
// - iterations   : 迭代次数
func batchGradientDescent(random *rand.Rand, x [][]float64, y [][]float64, learningRate float64, iterations int) ([][]float64, [][]float64) {
	// m: 样本数量, n: 特征数量
	m, n := len(x), len(x[0])

	// 初始化参数向量 theta
	theta := make([][]float64, n)
	for i := 0; i < n; i++ {
		theta[i] = []float64{random.Float64()} // 随机初始化参数
	}

	// 记录每次迭代的损失值
	lossHistory := make([][]float64, 0, iterations)
	for t := 0; t < iterations; t++ {
		// 批量梯度下降的"灵魂"，一次性计算所有样本的平均梯度
		prediction := dot(x, theta) // 计算预测值
		diff := sub(prediction, y)  // 预测值与真实值的差异
		xT := transpose(x)          // x 的转置矩阵
		gradient := dot(xT, diff)   // 计算梯度，未归一化
		normalizeGradient := make([][]float64, len(gradient))
		for row := 0; row < len(gradient); row++ {
			normalizeGradient[row] = make([]float64, len(gradient[0]))
			for col := 0; col < len(gradient[0]); col++ {
				normalizeGradient[row][col] = gradient[row][col] / float64(m) // 归一化梯度
			}
		}
		// 更新参数 theta
		for row := 0; row < len(theta); row++ {
			gradient := normalizeGradient[row][0]
			theta[row][0] -= learningRate * gradient
		}
		// 计算当前损失值
		loss := computeLoss(x, y, theta)
		// 记录损失值
		lossHistory = append(lossHistory, loss)
	}

	// 返回经过迭代的参数 和 所有迭代过程中的损失值历史
	return theta, lossHistory
}

// 生成原始随机数据
func generateRandomData(size int) (*rand.Rand, [][]float64, [][]float64) {
	// 随机种子
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// x 数据，shape (size, 1)
	x := make([][]float64, size)
	for i := 0; i < size; i++ {
		x[i] = []float64{random.Float64() * 10} // 生成 0-100 之间的随机数
	}

	// y 数据，假设真实关系为 y = 2x + 3 + 噪声
	y := make([][]float64, size)
	for row, cols := range x {
		noise := random.NormFloat64() // 正态分布噪声
		y[row] = []float64{2*cols[0] + 3 + noise}
	}

	return random, x, y
}

// 添加 偏置项
func appendBiasItem(square [][]float64) [][]float64 {
	biasedX := make([][]float64, len(square))
	for i, val := range square {
		biasedX[i] = []float64{1, val[0]} // 添加偏置项 1
	}
	return biasedX
}

// 定义损失函数
func computeLoss(x, y, theta [][]float64) []float64 {
	// 计算样本数量
	l := len(y)

	// 计算预测值
	// 设计矩阵 X (第0列是偏置项1，第1列是特征值)
	// x = np.array([
	//     [1, 2],   # 样本1: x₀=1, x₁=2
	//     [1, 3],   # 样本2: x₀=1, x₁=3
	//     [1, 5]    # 样本3: x₀=1, x₁=5
	// ])
	// 参数向量 theta
	// theta = np.array([
	//     [0.1],    # 截距项 θ₀
	//     [0.5]     # 特征权重 θ₁
	// ])
	predictions := dot(x, theta)

	// 计算预测值与真实值的差异
	diff := sub(predictions, y)
	// 计算差异的平方
	// - 消除负值（避免正负误差抵消）
	// - 放大较大误差（对异常值更敏感）
	diffSquared := make([][]float64, len(diff))
	for row, cols := range diff {
		diffSquared[row] = make([]float64, len(cols))
		for col := 0; col < len(cols); col++ {
			diffSquared[row][col] = cols[col] * cols[col]
		}
	}

	diffSquaredSum := make([]float64, len(diffSquared))
	for row, cols := range diffSquared {
		diffSquaredSum[row] = 0.0
		for col := 0; col < len(cols); col++ {
			diffSquaredSum[row] += cols[col]
		}
	}

	// 计算均方误差（MSE）
	average := make([]float64, len(diffSquaredSum))
	for row := range diffSquaredSum {
		average[row] = diffSquaredSum[row] / float64(l) // sum/l：求平均（真正的 MSE）
	}
	magicValue := float64(1) / float64(2.0) // 1/2：方便求导（后续梯度下降时，导数会消去这个 2）
	mse := make([]float64, len(average))
	for row := range average {
		mse[row] = magicValue * average[row]
	}

	return mse
}

// 计算矩阵的点积
// 矩阵 (3, 2)
// [[1, 2],
// [1, 4],
// [1, 6]]
// 矩阵 (2, 1)
// [[0.5],
// [1.0]]
// 计算过程
// 1.0*0.5 + 2.0*1.0 = 2.5
// 1.0*0.5 + 4.0*1.0 = 4.5
// 1.0*0.5 + 6.0*1.0 = 6.5
// 结果
// [[2.5],
// [4.5],
// [6.5]]
func dot(l, r [][]float64) [][]float64 {
	// 1. 维度检查
	if len(l) == 0 || len(r) == 0 {
		return nil
	}

	lRows, lCols := len(l), len(l[0])
	rRows, rCols := len(r), len(r[0])

	// 关键：左矩阵的列数必须等于右矩阵的行数
	if lCols != rRows {
		panic(fmt.Sprintf("square not match: l[%d,%d] * r[%d,%d]", lRows, lCols, rRows, rCols))
	}

	// 2. 创建结果矩阵 (m × p)
	result := make([][]float64, lRows)
	for i := range result {
		result[i] = make([]float64, rCols)
	}

	// 3. 矩阵乘法核心
	for i := 0; i < lRows; i++ { // 遍历左矩阵的行
		for j := 0; j < rCols; j++ { // 遍历右矩阵的列
			sum := 0.0
			for k := 0; k < lCols; k++ { // 累加: l[i][k] * r[k][j]
				sum += l[i][k] * r[k][j]
			}
			result[i][j] = sum
		}
	}

	return result
}

// 矩阵减法
func sub(l, r [][]float64) [][]float64 {
	result := make([][]float64, len(l))
	for row := 0; row < len(l); row++ {
		result[row] = make([]float64, len(l[row]))
		for col := 0; col < len(l[0]); col++ {
			result[row][col] = l[row][col] - r[row][col]
		}
	}
	return result
}

// 矩阵转置
func transpose(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 {
		return [][]float64{}
	}
	rows := len(matrix)
	cols := len(matrix[0])
	transposed := make([][]float64, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = make([]float64, rows)
		for j := 0; j < rows; j++ {
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}
