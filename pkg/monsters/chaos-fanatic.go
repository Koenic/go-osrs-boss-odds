package monsters

import "github.com/koenic/rs-odds/pkg"

var ChaosFanatic = pkg.DefaultMonster{
	Name:   "Chaos Fanatic",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Chaos Fanatic",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet chaos elemental",
						Odds: 1.0 / 1000,
					},
				},
				{
					{
						Name: "Odium shard 1",
						Odds: 1.0 / 256,
					},
					{
						Name: "Malediction shard 1",
						Odds: 1.0 / 256,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet chaos elemental",
						Amount: 1.0,
					},
					{
						Name:   "Odium shard 1",
						Amount: 1.0,
					},
					{
						Name:   "Malediction shard 1",
						Amount: 1.0,
					},
				},
				{
					{
						Name:   "Odium shard 1",
						Amount: 1.0,
					},
					{
						Name:   "Malediction shard 1",
						Amount: 1.0,
					},
				},
			},
		},
	},
}
