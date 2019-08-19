// 输入二位平面极坐标上的点（半径和角度（度））。计算对应的笛卡尔坐标系的点的 x 和 y 并输出。
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type polar struct {
	radius float64
	ø      float64
}

type cartesian struct {
	x float64
	y float64
}

// 用于打印结果
const result = "Polar: radius=%.02f angle=%.02f degrees -- Cartesian: x=%.02f y=%.02f\n"

// 输入提醒
var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "window" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

// 通过极坐标点计算笛卡尔坐标 x, y
func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions // 等待 questions
			ø := polarCoord.ø * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(ø)
			y := polarCoord.radius * math.Sin(ø)
			answers <- cartesian{x, y} // 给出结果
		}
	}()

	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n') // 读取 stdin 直到 \n
		if err != nil {
			break
		}
		line = line[:len(line)-1] // 去除末尾 \n
		if numbers := strings.Fields(line); len(numbers) == 2 {
			polars, err := floatsForStrins(numbers)
			if err != nil {
				fmt.Fprintln(os.Stderr, "invalid number")
				continue
			}
			questions <- polar{polars[0], polars[1]} // 生成问题
			coord := <-answers                       // 等待解答
			fmt.Printf(result, polars[0], polars[1], coord.x, coord.y)
		} else {
			fmt.Fprintln(os.Stderr, "invalid input")
		}
	}
}

// 转换 string 数组为 float64 数组
func floatsForStrins(numbers []string) ([]float64, error) {
	var floats []float64
	for _, number := range numbers {
		x, err := strconv.ParseFloat(number, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, x)
	}

	return floats, nil
}
