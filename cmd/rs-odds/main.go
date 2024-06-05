// cmd/yourapp/main.go
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"

	"gonum.org/v1/gonum/mat"

	"github.com/koenic/rs-odds/pkg/monsters"
)

type variationResult struct {
	VariationName string
	Cdf           []float64
	X_vals        []int64
}

type exportData struct {
	Variation    []variationResult
	Monster_name string
	Kc_name      string
	Description  string
}

var goal_completion = 0.9999
var data_points = 5000

func filesExist(monsterName string) bool {

	f, err := os.Open("./data")
	if err != nil {
		fmt.Println(err)
		panic("couldnt oppen file")
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		panic("couldnt read dir")
	}

	for _, v := range files {
		if strings.Contains(v.Name(), monsterName) {
			return true
		}
	}

	return false
}

func main() {

	testRun := false
	argsWithoutProg := ""

	testRun = true
	// if len(os.Args) >= 2 {
	// 	argsWithoutProg = os.Args[1]
	// 	testRun = true
	// 	f, err := os.Create("myprogram.prof")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	pprof.StartCPUProfile(f)
	// 	defer pprof.StopCPUProfile()
	// }

	for _, monster := range monsters.All_Monsters {
		monster_name_formatted := strings.Join(strings.Split(monster.GetName(), " "), "-")

		if filesExist(monster_name_formatted) {
			continue
		}
		fmt.Println("Starting", monster.GetName())

		data := make([]exportData, len(monster.GetVariations()[0].GetWantedDrops()))

		for variation_index, monsterVariation := range monster.GetVariations() {
			fmt.Println("Starting making matrices for", monsterVariation.GetName())
			matrices := monsterVariation.GetTransitionMatrices()
			fmt.Println("Done making matrices for", monsterVariation.GetName())
			for matrix_index, matrix := range matrices {
				i, j := matrix.Matrix.Dims()
				if strings.Compare(argsWithoutProg, "test") == 0 && i > 5 {
					goal_completion = 0.999
				}

				matrix_base := mat.DenseCopyOf(matrix.Matrix)

				cdf := []float64{}
				x_vals := []int64{}

				iterations := 1
				step_size := 1

				if testRun {
					// Determine increased step size for test runs
					for matrix_base.At(0, j-1) < goal_completion {
						matrix_base.Mul(matrix_base, matrix_base)
						iterations <<= 1
					}

					scaled_datapoints := data_points / int(math.Log2(float64(i)))
					if scaled_datapoints == 0 {
						scaled_datapoints = 1
					}
					step_size = ((iterations) / scaled_datapoints)

					if step_size == 0 {
						step_size = 1
					}
					fmt.Println("step size", step_size)
				}

				cdf = append(cdf, 0)
				x_vals = append(x_vals, 0)
				cdf = append(cdf, matrix.Matrix.At(0, j-1))
				x_vals = append(x_vals, 1)
				matrix_base = mat.DenseCopyOf(matrix.Matrix)
				matrix_mult := mat.DenseCopyOf(matrix.Matrix)

				for i = 1; i < step_size-1; i <<= 1 {
					matrix_mult.Mul(matrix_mult, matrix_mult)
				}
				step_size = i

				fmt.Println("step size", step_size)
				for i := int64(1); matrix_base.At(0, j-1) < goal_completion; i++ {
					matrix_base.Mul(matrix_base, matrix_mult)
					cdf = append(cdf, matrix_base.At(0, j-1))
					x_vals = append(x_vals, 1+i*int64(step_size))
					if i%1000 == 0 {
						fmt.Println("step", 1+i*int64(step_size), matrix_base.At(0, j-1))
					}
					fmt.Println("step", 1+i*int64(step_size), matrix_base.At(0, j-1))
				}

				if data[matrix_index].Monster_name == "" {
					data[matrix_index] = exportData{
						Variation:    make([]variationResult, len(monster.GetVariations())),
						Monster_name: monster.GetName(),
						Description:  matrix.Description,
						Kc_name:      monster.GetKcName(),
					}
				}

				data[matrix_index].Variation[variation_index] = variationResult{
					X_vals:        x_vals,
					Cdf:           cdf,
					VariationName: monsterVariation.GetName(),
				}
			}
		}

		for ind, dat := range data {

			file, _ := os.OpenFile(fmt.Sprintf("data/%s_%d.json", monster_name_formatted, ind), os.O_CREATE|os.O_WRONLY, os.ModePerm)
			file.Truncate(0)
			file.Seek(0, 0)
			defer file.Close()
			encoder := json.NewEncoder(file)
			err := encoder.Encode(dat)

			if err != nil {
				panic(err.Error())
			}
		}

	}
}
