package algorithmfoo

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// 基于 线性回归分析 实现的梯度下降算法

// 线性回归分析的目标函数 -> y = 10x + 5
// 梯度下降算法的目标是通过不断调整参数，最小化损失函数，从而找到最佳拟合线。
// 换句话说就是，梯度下降就是为了通过大量数据找到这个函数中的 10 和 5 这两个参数。
func targetFunc(x float64) float64 { return 10*x + 5 }

func GradientDescentFoo() {
	random, x, y := generateRandomData(targetFunc, false, 100) // 生成随机数据
	biasedX := appendBiasItem(x)                               // 添加偏置项

	// 批量梯度下降
	// 不同的学习率，迭代次数，都会影响结果
	// 从目前来看
	// 学习率在 0.05，迭代 1000次，综合效果还不错（时间成本）
	// 学习率在 0.05，迭代 10000次，效果最佳，最接近真实值
	// 学习率在 0.01，迭代 10000次，效果次之，与真实值误差率不超过 0.01%
	// 学习率在 0.001，迭代 10000次，效果不太好，收敛太慢了
	// 学习率在 0.1，迭代 1000次，发散严重，theta 趋向于无穷（float64 越界）
	batchGradientDescentLearning := 0.05
	batchGradientDescentLearningIterations := 1000 // 在线性回归分析中，迭代部署越多，模型效果越好
	batchTheta, batchLoss := batchGradientDescent(random, biasedX, y, batchGradientDescentLearning, batchGradientDescentLearningIterations)
	for _, cols := range batchTheta {
		if math.IsNaN(cols[0]) {
			panic("batchGradientDescent got NaN")
		}
	}
	// 返回结果 batchTheta 是列向量 (5, 10)
	// 为什么是 (5, 10)，因为是偏置项+权重
	// 我们构造的偏置矩阵 biasedX 的第一列就是为了解决偏置项的（也就是常数项）
	fmt.Printf("batchTheta result %v\n", toLine(batchTheta)) // 返回结果 batchTheta 是列向量

	plotX := make([]float64, len(batchLoss))
	plotY := make([]float64, len(batchLoss))
	for i := 0; i < len(batchLoss); i++ {
		plotX[i] = float64(i)
		plotY[i] = batchLoss[i]
	}
	toPlot(
		"batchGradientDescent Loss over Iterations",
		"Iterations", "Loss", plotX, plotY,
		"batch_gradient_descent_loss.png",
	)
}

func toPlot(title, xLabel, yLabel string, plotX, plotY []float64, saveAs string) {
	// 创建图表
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	// 添加数据
	points := make(plotter.XYs, len(plotX))
	for i := range plotX {
		points[i].X = plotX[i]
		points[i].Y = plotY[i]
	}

	// 创建折线
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	line.Width = vg.Points(2)

	// 添加到图表并保存
	p.Add(line)
	err = p.Save(4*vg.Inch, 3*vg.Inch, saveAs)
	if err != nil {
		panic(err)
	}
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
func batchGradientDescent(random *rand.Rand, x [][]float64, y [][]float64, learningRate float64, iterations int) ([][]float64, []float64) {
	// rows: 样本数量, cols: 特征数量
	rows, cols := len(x), len(x[0])

	// 初始化参数向量 theta
	theta := make([][]float64, cols)
	for i := 0; i < cols; i++ {
		theta[i] = []float64{random.Float64()} // 随机初始化参数
	}

	// 记录每次迭代的损失值
	lossHistory := make([]float64, 0, iterations)
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
				normalizeGradient[row][col] = gradient[row][col] / float64(rows) // 归一化梯度
			}
		}
		// 更新参数 theta
		for row := 0; row < len(theta); row++ {
			gradient := normalizeGradient[row][0]
			theta[row][0] -= learningRate * gradient
			if theta[row][0] > 1e10 || math.IsNaN(theta[row][0]) {
				panic("theta diverged or got NaN")
			}
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
func generateRandomData(realFunc func(float64) float64, withNoise bool, size int) (*rand.Rand, [][]float64, [][]float64) {
	// 随机种子
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// x 数据，shape (size, 1)
	x := make([][]float64, size)
	for i := 0; i < size; i++ {
		x[i] = []float64{random.Float64() * 10} // 生成 0-100 之间的随机数
	}

	// y 数据，假设真实关系为 y = realFunc(x) + 噪声
	y := make([][]float64, size)
	for row, cols := range x {
		if withNoise {
			noise := random.NormFloat64() // 正态分布噪声
			y[row] = []float64{realFunc(cols[0]) + noise}
		} else {
			y[row] = []float64{realFunc(cols[0])}
		}
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
func computeLoss(x, y, theta [][]float64) float64 {
	// 计算样本数量
	l := len(y)

	// 计算预测值
	// 设计矩阵 X (第0列是偏置项1，第1列是特征值)
	// x = np.array([
	//     [1, x1],   # 样本1: x₀=1, x₁=2
	//     [1, x2],   # 样本2: x₀=1, x₁=3
	//     [1, x3]    # 样本3: x₀=1, x₁=5
	// ])
	// 参数向量 theta
	// theta = np.array([
	//     [t1],    # 截距项 θ₀ -> 什么是截距项？
	//     [t2]     # 特征权重 θ₁
	// ])
	// 在线性回归分析里面，
	// 截距项（intercept term）是指回归方程中的常数项，
	// 表示当所有自变量（特征）都为零时，因变量（目标变量）的预期值。
	// 它是回归线与纵轴（y轴）的交点。在数学表达式中，截距项通常用 θ₀ 表示。
	// 特征权重（feature weight）是指每个自变量（特征）对因变量（目标变量）的影响程度。
	// 它表示当某个特征增加一个单位时，因变量的预期变化量。
	// 在数学表达式中，特征权重通常用 θ₁, θ₂, ... θn 表示。
	// 用高等数学的语言来说，就是斜率。
	// 结果
	// predictions = np.array([
	//     [1*t1 + x1*t2],    # 预测值1
	//     [1*t1 + x2*t2],    # 预测值2
	//     [1*t1 + x3*t2]     # 预测值3
	// ])
	predictions := dot(x, theta)

	// 计算预测值与真实值的差异
	// 预测值
	// predictions = np.array([
	//     [1*t1 + x1*t2],    # 预测值1
	//     [1*t1 + x2*t2],    # 预测值2
	//     [1*t1 + x3*t2]     # 预测值3
	// ])
	// 真实值
	// y = np.array([
	//     [y1],    # 真实值1
	//     [y2],    # 真实值2
	//     [y3]     # 真实值3
	// ])
	// 差异
	// diff = np.array([
	//     [(1*t1 + x1*t2) - y1],    # 差异1
	//     [(1*t1 + x2*t2) - y2],    # 差异2
	//     [(1*t1 + x3*t2) - y3]     # 差异3
	// ])
	diff := sub(predictions, y)

	// 计算差异的平方
	// - 消除负值（避免正负误差抵消）
	// - 放大较大误差（对异常值更敏感）
	// 逐个元素放大，形状不变
	// diffSquaredSum = np.array([
	//     [((1*t1 + x1*t2) - y1)^2],
	//     [((1*t1 + x2*t2) - y2)^2],
	//     [((1*t1 + x3*t2) - y3)^2]
	// ])
	diffSquared := make([][]float64, len(diff))
	for row, cols := range diff {
		diffSquared[row] = make([]float64, len(cols))
		for col := 0; col < len(cols); col++ {
			diffSquared[row][col] = cols[col] * cols[col]
		}
	}

	// 对所有元素求和，压缩成一个标量 -> 为什么要压缩成一个标量？
	diffSquaredSum := 0.0
	for _, cols := range diffSquared {
		for _, v := range cols {
			diffSquaredSum += v
		}
	}

	// 计算均方误差（MSE）
	average := diffSquaredSum / float64(l) // sum/l：求平均（真正的 MSE）
	magicValue := float64(1) / float64(2)  // 1/2：方便求导（后续梯度下降时，导数会消去这个 2）
	mse := average * magicValue

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
