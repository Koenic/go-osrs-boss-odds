package monsters

import "github.com/koenic/rs-odds/pkg"

var GiantMole = pkg.DefaultMonster{
	Name:   "Giant Mole",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Giant Mole",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Baby Mole",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Mole Skin",
						Odds: 1.0,
					},
				},
				{
					{
						Name: "Mole Claw",
						Odds: 1.0,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{{
				{
					Name:   "Baby Mole",
					Amount: 1,
				},
				{
					Name:   "Mole Skin",
					Amount: 1,
				},
				{
					Name:   "Mole Claw",
					Amount: 1,
				},
			}},
		},
	},
}
