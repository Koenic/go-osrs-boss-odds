package monsters

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/james-bowman/sparse"
	"github.com/koenic/rs-odds/pkg"
)

type BarrowsMonster struct {
	Name             string
	KcName           string
	MonsterVariation []BarrowsVariation
}

type BarrowsVariation struct {
	VariationName string
	WantedDrops   [][]pkg.DropGoal
}

var Barrows = BarrowsMonster{
	Name:   "Barrows (killing all 6)",
	KcName: "Chests",
	MonsterVariation: []BarrowsVariation{
		{
			VariationName: "Barrows",
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Ahrim's hood",
						Amount: 1,
					},
					{
						Name:   "Ahrim's robetop",
						Amount: 1,
					},
					{
						Name:   "Ahrim's robeskirt",
						Amount: 1,
					},
					{
						Name:   "Ahrim's staff",
						Amount: 1,
					},
					{
						Name:   "Dharok's helm",
						Amount: 1,
					},
					{
						Name:   "Dharok's platebody",
						Amount: 1,
					},
					{
						Name:   "Dharok's platelegs",
						Amount: 1,
					},
					{
						Name:   "Dharok's greataxe",
						Amount: 1,
					},
					{
						Name:   "Guthan's helm",
						Amount: 1,
					},
					{
						Name:   "Guthan's platebody",
						Amount: 1,
					},
					{
						Name:   "Guthan's chainskirt",
						Amount: 1,
					},
					{
						Name:   "Guthan's warspear",
						Amount: 1,
					},
					{
						Name:   "Karil's coif",
						Amount: 1,
					},
					{
						Name:   "Karil's leathertop",
						Amount: 1,
					},
					{
						Name:   "Karil's leatherskirt",
						Amount: 1,
					},
					{
						Name:   "Karil's crossbow",
						Amount: 1,
					},
					{
						Name:   "Torag's helm",
						Amount: 1,
					},
					{
						Name:   "Torag's platebody",
						Amount: 1,
					},
					{
						Name:   "Torag's platelegs",
						Amount: 1,
					},
					{
						Name:   "Torag's hammers",
						Amount: 1,
					},
					{
						Name:   "Verac's helm",
						Amount: 1,
					},
					{
						Name:   "Verac's brassard",
						Amount: 1,
					},
					{
						Name:   "Verac's plateskirt",
						Amount: 1,
					},
					{
						Name:   "Verac's flail",
						Amount: 1,
					},
				},
			},
		},
	},
}

func getDropNames(drop_goal []pkg.DropGoal) []string {

	drop_names := make([]string, len(drop_goal))
	for ind, drop := range drop_goal {
		drop_names[ind] = fmt.Sprintf("%dx %s", drop.Amount, drop.Name)
	}

	return drop_names
}

func (monster BarrowsMonster) GetName() string {
	return monster.Name
}

func (monster BarrowsMonster) GetKcName() string {
	return monster.KcName
}

func (monster BarrowsMonster) GetVariations() []pkg.DropOdds {
	monsterVariations := []pkg.DropOdds{}

	for i := range monster.MonsterVariation {
		p := monster.MonsterVariation[i]
		monsterVariations = append(monsterVariations, &p)
	}
	return monsterVariations
}

func (monsterVariation BarrowsVariation) GetName() string {
	return monsterVariation.VariationName
}

func (monsterVariation BarrowsVariation) GetWantedDrops() [][]pkg.DropGoal {
	return monsterVariation.WantedDrops
}

func (monster BarrowsVariation) IsPossible() bool {
	return true
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func BinCoeff(n int, k int) float64 {
	return float64(Factorial(n) / (Factorial(k) * Factorial(n-k)))
}

func BarrowsUniqueItemOdds(gotten int, new_items int) float64 {
	odds := 0.0

	for k := new_items; k < 8; k++ {
		oddsOfKItems := BinCoeff(7, k) * math.Pow(1.0/102, float64(k)) * math.Pow(101.0/102, float64(7-k)) * BinCoeff(k, new_items)

		noDupesOdds := 1.0
		for l := 0; l < new_items; l++ {
			noDupesOdds *= float64(24-gotten-l) / float64(24-l)
		}

		noDupesInChestOdds := 1.0
		for l := 0; l < k-new_items; l++ {
			noDupesInChestOdds *= 1.0 - (float64(24-new_items-gotten-l) / float64(24-new_items-l))
		}

		odds += oddsOfKItems * noDupesOdds * noDupesInChestOdds
	}
	return odds
}

func (monsterVariation BarrowsVariation) GetTransitionMatrices() []pkg.DropMatrix {

	drop_table_matrices := []pkg.DropMatrix{}

	for ind, dropGoal := range monsterVariation.WantedDrops {
		if !monsterVariation.IsPossible() {
			log.Printf("WARN: dropGoal %d is not possible for monster %s skipping", ind, monsterVariation.VariationName)
			continue
		}

		n_rows := 25

		rows := []int{}
		cols := []int{}
		odds := []float64{}

		for row_index := 0; row_index < n_rows; row_index++ {
			row_odds_total := 0.0

			col_max := 8
			if n_rows-row_index < col_max {
				col_max = n_rows - row_index
			}

			for n_items := 1; n_items < col_max; n_items++ {
				value := BarrowsUniqueItemOdds(row_index, n_items)
				odds = append(odds, value)
				row_odds_total += value

				fmt.Println(row_index, n_items, value)
				rows = append(rows, row_index)
				cols = append(cols, row_index+n_items)
			}

			// The diagonal of the matrix is equal to the odds of not getting any drop
			rows = append(rows, row_index)

			cols = append(cols, row_index)
			odds = append(odds, 1.0-row_odds_total)
		}

		fmt.Println(n_rows)
		fmt.Println(rows)
		fmt.Println(cols)
		fmt.Println(odds)
		dropGoalTransitionMatrix := sparse.NewCOO(n_rows, n_rows, rows, cols, odds).ToDense()

		drop_table_matrices = append(drop_table_matrices, pkg.DropMatrix{
			Matrix:      dropGoalTransitionMatrix,
			Description: strings.Join(getDropNames(dropGoal), ", "),
		})
	}

	return drop_table_matrices
}
