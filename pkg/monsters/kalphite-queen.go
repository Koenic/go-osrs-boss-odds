package monsters

import "github.com/koenic/rs-odds/pkg"

var KalphiteQueen = pkg.DefaultMonster{
	Name:   "Kalphite Queen",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Kalphite Queen",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Kalphite princess",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Jar of sand",
						Odds: 1.0 / 2000,
					},
				},
				{
					{
						Name: "Kq head",
						Odds: 1.0 / 128,
					},
				},
				{
					{
						Name: "Dragon chainbody",
						Odds: 1.0 / 128,
					},
				},
				{
					{
						Name: "Dragon 2h sword",
						Odds: 1.0 / 256,
					},
				},
				{
					{
						Name: "Dragon pickaxe",
						Odds: 1.0 / 400,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Kalphite princess",
						Amount: 1,
					},
					{
						Name:   "Jar of sand",
						Amount: 1,
					},
					{
						Name:   "Kq head",
						Amount: 1,
					},
					{
						Name:   "Dragon chainbody",
						Amount: 1,
					},
					{
						Name:   "Dragon 2h sword",
						Amount: 1,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Kq head",
						Amount: 1,
					},
					{
						Name:   "Dragon chainbody",
						Amount: 1,
					},
					{
						Name:   "Dragon 2h sword",
						Amount: 1,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1,
					},
				},
			},
		},
	},
}
