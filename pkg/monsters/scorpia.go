package monsters

import "github.com/koenic/rs-odds/pkg"

var Scorpia = pkg.DefaultMonster{
	Name:   "Scorpia",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Scorpia",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Scorpia's offspring",
						Odds: 1.0 / 2016,
					},
				},
				{
					{
						Name: "Odium shard 3",
						Odds: 1.0 / 256,
					},
					{
						Name: "Malediction shard 3",
						Odds: 1.0 / 256,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Scorpia's offspring",
						Amount: 1,
					},
					{
						Name:   "Odium shard 3",
						Amount: 1,
					},
					{
						Name:   "Malediction shard 3",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Odium shard 3",
						Amount: 1,
					},
					{
						Name:   "Malediction shard 3",
						Amount: 1,
					},
				},
			},
		},
	},
}
