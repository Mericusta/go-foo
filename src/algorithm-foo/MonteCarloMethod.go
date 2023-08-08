package algorithmfoo

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// MonteCarloMethod_Estimating_PI 蒙特卡洛法估算圆周率
// @param0                        样本数
func MonteCarloMethod_Estimating_PI(points int) {
	// 已知
	// - 一个边长为 edgeLength 的正方形
	// - 正方形的面积是 edgeLength * edgeLength
	// - 该正方形内存在一个与之相切的圆，半径为 radius = edgeLength / 2
	// - 圆的面积是 pi * radius * radius
	// - 圆和正方形的面积比为 ratio = pi / 4

	// 蒙特卡洛法求解
	// - 在正方形内随机生成 points 个点，坐标为 [x, y]
	// - 计算这些点到圆心的距离从而判断点是否在圆内
	// - 若这些点均匀分布，则会有 inCount = pi / 4 * points 的点落在圆内
	// - 则 pi = inCount / points * 4

	// 以下采用笛卡尔坐标系来计算
	// 所有坐标扩大100倍
	// 正方形左下角为坐标原点
	// 圆心坐标为 [100, 100]
	var (
		edgeLength    float64 = 2
		radius        float64 = edgeLength / 2
		inCount       float64
		circleCenterX float64    = radius * 100
		circleCenterY float64    = radius * 100
		random        *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	)

	for i := 0; i < points; i++ {
		dx := random.Float64()*edgeLength*100 - circleCenterX
		dy := random.Float64()*edgeLength*100 - circleCenterY
		d := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
		if d <= radius*100 {
			inCount++
		}
	}

	pi := inCount / float64(points) * 4
	var delta, v float64
	if pi > math.Pi {
		delta, v = pi/math.Pi-1, 1
	} else if pi < math.Pi {
		delta, v = math.Pi/pi-1, -1
	}
	fmt.Printf("estimating PI = %f\n", pi)
	fmt.Printf("math.Pi = %f\n", math.Pi)
	fmt.Printf("delta = %.2f%%\n", delta*v*10000)
}
