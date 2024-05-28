package monsters

import "github.com/koenic/rs-odds/pkg"

var KingBlackDragon = pkg.DefaultMonster{
	Name:   "King Black Dragon",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "King Black Dragon",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Prince black dragon",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Kbd heads",
						Odds: 1.0 / 128,
					},
				},
				{
					{
						Name: "Dragon pickaxe",
						Odds: 1.0 / 1000,
					},
				},
				{
					{
						Name: "Draconic visage",
						Odds: 1.0 / 5000,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Prince black dragon",
						Amount: 1,
					},
					{
						Name:   "Kbd heads",
						Amount: 1,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1,
					},
					{
						Name:   "Draconic visage",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Kbd heads",
						Amount: 1,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1,
					},
					{
						Name:   "Draconic visage",
						Amount: 1,
					},
				},
			},
		},
	},
}
