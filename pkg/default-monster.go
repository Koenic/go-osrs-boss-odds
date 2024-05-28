package pkg

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/james-bowman/sparse"
	"gonum.org/v1/gonum/mat"
)

type Drop struct {
	Name string
	Odds float64
}

type DropGoal struct {
	Name   string
	Amount int
}

type Monster interface {
	GetName() string
	GetKcName() string
	GetVariations() []DropOdds
}

type DefaultMonster struct {
	Name             string
	KcName           string
	MonsterVariation []MonsterVariation
}

type MonsterVariation struct {
	VariationName string
	DropTables    [][]Drop
	WantedDrops   [][]DropGoal
}

func calcRows(goal []DropGoal) int {
	total := 1
	for _, item := range goal {
		total *= (item.Amount) + 1
	}
	return total
}

func index_to_gotten_drops(goal []DropGoal, index int) []int {
	index_bounds := []int{}
	array_size := 1
	gotten_drops := []int{}
	for _, drop := range goal {
		index_bounds = append(index_bounds, drop.Amount+1)
		array_size *= drop.Amount + 1
	}

	for i := range goal {
		gotten_drops = append(gotten_drops, index%index_bounds[i])
		index /= index_bounds[i]
	}
	return gotten_drops
}

func gotten_drops_to_index(goal []DropGoal, gotten_drops []int) int {
	index := 0
	size := 1
	for i, drop := range goal {
		dimSize := drop.Amount + 1
		index += gotten_drops[i] * size
		size *= dimSize
	}
	return index
}

func eye(n int) *mat.Dense {
	d := make([]float64, n*n)
	for i := 0; i < n*n; i += n + 1 {
		d[i] = 1
	}
	return mat.NewDense(n, n, d)
}

func getDropNames(drop_goal []DropGoal) []string {

	drop_names := make([]string, len(drop_goal))
	for ind, drop := range drop_goal {
		drop_names[ind] = fmt.Sprintf("%dx %s", drop.Amount, drop.Name)
	}

	return drop_names
}

func (monster DefaultMonster) GetName() string {
	return monster.Name
}

func (monster DefaultMonster) GetKcName() string {
	return monster.KcName
}

func (monster DefaultMonster) GetVariations() []DropOdds {
	monsterVariations := []DropOdds{}

	for i := range monster.MonsterVariation {
		p := monster.MonsterVariation[i]
		monsterVariations = append(monsterVariations, &p)
	}
	return monsterVariations
}

func (monsterVariation MonsterVariation) GetName() string {
	return monsterVariation.VariationName
}

func (monsterVariation MonsterVariation) GetWantedDrops() [][]DropGoal {
	return monsterVariation.WantedDrops
}

func (monster MonsterVariation) IsPossible() bool {
	possible_drops := []string{}
	for _, drop_table := range monster.DropTables {
		for _, drop := range drop_table {
			if slices.Index(possible_drops, drop.Name) == -1 {
				possible_drops = append(possible_drops, drop.Name)
			}
		}
	}

	for _, dropGoal := range monster.WantedDrops {
		for _, drop := range dropGoal {
			if slices.Index(possible_drops, drop.Name) == -1 {
				return false
			}
		}
	}

	return true
}

func (monsterVariation MonsterVariation) GetTransitionMatrices() []DropMatrix {

	drop_table_matrices := []DropMatrix{}

	for ind, dropGoal := range monsterVariation.WantedDrops {
		if !monsterVariation.IsPossible() {
			log.Printf("WARN: dropGoal %d is not possible for monster %s skipping", ind, monsterVariation.VariationName)
			continue
		}

		n_rows := calcRows(dropGoal)

		dropGoalTransitionMatrix := eye(n_rows)

		for _, dropTable := range monsterVariation.DropTables {
			rows := []int{}
			cols := []int{}
			odds := []float64{}

			for row_index := 0; row_index < n_rows; row_index++ {
				row_odds_total := 0.0

				gotten_drops := index_to_gotten_drops(dropGoal, row_index)

				for drop_index, drop := range dropGoal {
					// odds of getting the given drop at index row_index
					if gotten_drops[drop_index] >= dropGoal[drop_index].Amount {
						continue
					}
					drop_table_index := slices.IndexFunc(dropTable, func(tbl_drop Drop) bool {
						return tbl_drop.Name == drop.Name
					})

					// drop is not on this dropTable
					if drop_table_index == -1 {
						continue
					}

					// calculate index after getting the drop
					gotten_drops[drop_index] += 1
					target_index := gotten_drops_to_index(dropGoal, gotten_drops)
					gotten_drops[drop_index] -= 1

					rows = append(rows, row_index)
					cols = append(cols, target_index)
					odds = append(odds, dropTable[drop_table_index].Odds)
					row_odds_total += dropTable[drop_table_index].Odds
				}

				// The diagonal of the matrix is equal to the odds of not getting any drop
				rows = append(rows, row_index)
				cols = append(cols, row_index)
				odds = append(odds, 1.0-row_odds_total)
			}

			dense := sparse.NewCOO(n_rows, n_rows, rows, cols, odds).ToDense()
			dropGoalTransitionMatrix.Mul(dropGoalTransitionMatrix, dense)
		}

		drop_table_matrices = append(drop_table_matrices, DropMatrix{
			Matrix:      dropGoalTransitionMatrix,
			Description: strings.Join(getDropNames(dropGoal), ", "),
		})
	}

	return drop_table_matrices
}
