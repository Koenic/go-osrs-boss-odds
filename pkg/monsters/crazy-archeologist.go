package monsters

import "github.com/koenic/rs-odds/pkg"

var CrazyArchaeologist = pkg.DefaultMonster{
	Name:   "Crazy Archaeologist",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Crazy Archaeologist",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Fedora",
						Odds: 1.0 / 128,
					},
					{
						Name: "Odium shard 2",
						Odds: 1.0 / 256,
					},
					{
						Name: "Malediction shard 2",
						Odds: 1.0 / 256,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Fedora",
						Amount: 1.0,
					},
					{
						Name:   "Odium shard 2",
						Amount: 1.0,
					},
					{
						Name:   "Malediction shard 2",
						Amount: 1.0,
					},
				},
				{
					{
						Name:   "Odium shard 2",
						Amount: 1.0,
					},
					{
						Name:   "Malediction shard 2",
						Amount: 1.0,
					},
				},
			},
		},
	},
}
