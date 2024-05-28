package monsters

import "github.com/koenic/rs-odds/pkg"

var FightCaves = pkg.DefaultMonster{
	Name:   "Fight Caves (gambling capes)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Off task",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "TzRek-jad",
						Odds: 1.0 / 200,
					},
				},
				{
					{
						Name: "TzRek-jad",
						Odds: 1.0 / 200,
					},
				},
				{
					{
						Name: "Fire Cape",
						Odds: 1.0,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "TzRek-jad",
						Amount: 1,
					},
					{
						Name:   "Fire Cape",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "On task",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "TzRek-jad",
						Odds: 1.0 / 100,
					},
				},
				{
					{
						Name: "TzRek-jad",
						Odds: 1.0 / 200,
					},
				},
				{
					{
						Name: "Fire Cape",
						Odds: 1.0,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "TzRek-jad",
						Amount: 1,
					},
					{
						Name:   "Fire Cape",
						Amount: 1,
					},
				},
			},
		},
	},
}
