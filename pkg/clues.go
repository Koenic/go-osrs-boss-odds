package pkg

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/james-bowman/sparse"
)

type CluesMonster struct {
	Name             string
	KcName           string
	MonsterVariation []CluesVariation
}

type CluesVariation struct {
	VariationName string
	WantedDrops   [][]DropGoal
	DropTable     []Drop
}

func (monster CluesMonster) GetName() string {
	return monster.Name
}

func (monster CluesMonster) GetKcName() string {
	return monster.KcName
}

func (monster CluesMonster) GetVariations() []DropOdds {
	monsterVariations := []DropOdds{}

	for i := range monster.MonsterVariation {
		p := monster.MonsterVariation[i]
		monsterVariations = append(monsterVariations, &p)
	}
	return monsterVariations
}

func (monsterVariation CluesVariation) GetName() string {
	return monsterVariation.VariationName
}

func (monsterVariation CluesVariation) GetWantedDrops() [][]DropGoal {
	return monsterVariation.WantedDrops
}

func (monster CluesVariation) IsPossible() bool {
	return true
}

func calc_unique_odds(odds float64, gotten int, total_needed int) float64 {
	return (odds * (float64(total_needed-gotten) / float64(total_needed)))
}

func (monsterVariation CluesVariation) GetTransitionMatrices() []DropMatrix {

	drop_goal := []DropGoal{}
	drop_table := []Drop{}
	drop_table_matrices := []DropMatrix{}

	if !monsterVariation.IsPossible() {
		log.Printf("WARN: dropGoal is not possible for monster %s skipping", monsterVariation.VariationName)
		panic("PANIC")
	}

	sort.Slice(monsterVariation.DropTable, func(i, j int) bool { return monsterVariation.DropTable[i].Odds < monsterVariation.DropTable[j].Odds })

	for ind, drop := range monsterVariation.DropTable {
		if ind == 0 || drop.Odds != monsterVariation.DropTable[ind-1].Odds {
			name := strconv.FormatFloat(drop.Odds, 'f', -1, 64)
			drop_table = append(drop_table, Drop{
				Name: name,
				Odds: drop.Odds,
			})
			drop_goal = append(drop_goal, DropGoal{
				Name:   name,
				Amount: 1,
			})
		} else if drop.Odds == monsterVariation.DropTable[ind-1].Odds {
			drop_goal[len(drop_goal)-1].Amount += 1
		}
	}

	for ind, drop := range drop_goal {
		drop_table[ind].Odds *= float64(drop.Amount)
	}

	fmt.Println(drop_goal, drop_table)

	n_rows := calcRows(drop_goal)
	rows := []int{}
	cols := []int{}
	odds := []float64{}

	for row_index := 0; row_index < n_rows; row_index++ {
		row_odds_total := 0.0

		gotten_drops := index_to_gotten_drops(drop_goal, row_index)

		for drop_index, drop := range drop_goal {
			// odds of getting the given drop at index row_index
			if gotten_drops[drop_index] >= drop_goal[drop_index].Amount {
				continue
			}
			drop_table_index := slices.IndexFunc(drop_table, func(tbl_drop Drop) bool {
				return tbl_drop.Name == drop.Name
			})

			// drop is not on this dropTable
			if drop_table_index == -1 {
				continue
			}

			// calculate index after getting the drop
			gotten_drops[drop_index] += 1
			target_index := gotten_drops_to_index(drop_goal, gotten_drops)
			gotten_drops[drop_index] -= 1

			drop_odds := calc_unique_odds(drop_table[drop_table_index].Odds, gotten_drops[drop_index], drop.Amount)
			rows = append(rows, row_index)
			cols = append(cols, target_index)
			odds = append(odds, drop_odds)
			row_odds_total += drop_odds
		}

		// The diagonal of the matrix is equal to the odds of not getting any drop
		rows = append(rows, row_index)
		cols = append(cols, row_index)
		odds = append(odds, 1.0-row_odds_total)
	}

	dropGoalTransitionMatrix := sparse.NewCOO(n_rows, n_rows, rows, cols, odds).ToDense()

	fmt.Println(n_rows)
	// fmt.Println(rows)
	// fmt.Println(cols)
	// fmt.Println(odds)

	drop_table_matrices = append(drop_table_matrices, DropMatrix{
		Matrix:      dropGoalTransitionMatrix,
		Description: strings.Join(getDropNames(monsterVariation.WantedDrops[0]), ", "),
	})

	return drop_table_matrices
}
