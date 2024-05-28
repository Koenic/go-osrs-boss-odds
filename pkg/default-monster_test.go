package pkg

import (
	"testing"
)

func TestDefaultMonster(t *testing.T) {
	twoDiceTest := DefaultMonster{
		Name: "Dice",
		MonsterVariation: []MonsterVariation{
			{
				DropTables: [][]Drop{
					{
						{
							Name: "six",
							Odds: 1.0 / 6,
						},
					},
					{
						{
							Name: "six",
							Odds: 1.0 / 6,
						},
					},
				},
				WantedDrops: [][]DropGoal{{
					{
						Name:   "six",
						Amount: 2,
					},
				}},
			},
		},
	}

	matrices := twoDiceTest.MonsterVariation[0].GetTransitionMatrices()

	if matrices[0].Matrix.At(0, 2) != 1.0/36 {
		t.Fatalf(`TransitionMatrixes[0,2] = %f, want 1.0/36`, matrices[0].Matrix.At(0, 2))
	}
}

func TestPossibleDrops(t *testing.T) {
	possible := DefaultMonster{
		Name: "Dice",
		MonsterVariation: []MonsterVariation{
			{
				DropTables: [][]Drop{
					{
						{
							Name: "five",
							Odds: 1.0 / 6,
						},
					},
					{
						{
							Name: "six",
							Odds: 1.0 / 6,
						},
					},
				},
				WantedDrops: [][]DropGoal{{
					{
						Name:   "six",
						Amount: 2,
					},
				}},
			},
		},
	}

	impossible := DefaultMonster{
		Name: "Dice",
		MonsterVariation: []MonsterVariation{
			{DropTables: [][]Drop{
				{
					{
						Name: "five",
						Odds: 1.0 / 6,
					},
				},
				{
					{
						Name: "six",
						Odds: 1.0 / 6,
					},
				},
			},
				WantedDrops: [][]DropGoal{{
					{
						Name:   "four",
						Amount: 2,
					},
				}},
			},
		},
	}

	if !possible.MonsterVariation[0].IsPossible() || impossible.MonsterVariation[0].IsPossible() {
		t.Fail()
	}
}
