package monsters

import "github.com/koenic/rs-odds/pkg"

var Hespori = pkg.DefaultMonster{
	Name:   "Hespori (99 farming)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Hespori",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Attas seed",
						Odds: 1.0 / 3,
					},
					{
						Name: "Iasor seed",
						Odds: 1.0 / 3,
					},
					{
						Name: "Kronos seed",
						Odds: 1.0 / 3,
					},
				},
				{
					{
						Name: "Bottomless compost bucket",
						Odds: 1.0 / 35,
					},
				},
				{
					{
						Name: "Tangleroot",
						Odds: 1.0 / 4525,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Tangleroot",
						Amount: 1,
					},
					{
						Name:   "Attas seed",
						Amount: 1,
					},
					{
						Name:   "Iasor seed",
						Amount: 1,
					},
					{
						Name:   "Kronos seed",
						Amount: 1,
					},
					{
						Name:   "Bottomless compost bucket",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Attas seed",
						Amount: 1,
					},
					{
						Name:   "Iasor seed",
						Amount: 1,
					},
					{
						Name:   "Kronos seed",
						Amount: 1,
					},
					{
						Name:   "Bottomless compost bucket",
						Amount: 1,
					},
				},
			},
		},
	},
}
