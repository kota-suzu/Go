package main

import (
	"fmt"
	"math"
)

// Vertexは2次元空間における点を表し、XとYの座標を持ちます。
type Vertex struct {
	X, Y float64
}

// Absは頂点の絶対値を計算して返します。
// ピタゴラスの定理を用いて、原点(0,0)からの頂点までの距離を見つけます。
// この関数はVertex vを入力として受け取り、距離をfloat64として返します。
func Abs(v Vertex) float64 {
	// math.Sqrtは平方根を計算します。
	// v.X*v.X + v.Y*v.Y は頂点の座標の二乗の和を計算します。
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// X=3、Y=4でVertexインスタンスを作成します。
	v := Vertex{3, 4}
	// 頂点の絶対値をコンソールに出力します。
	// これは原点から点(3,4)までの距離を計算する方法を示しています。
	fmt.Println(Abs(v))

}
