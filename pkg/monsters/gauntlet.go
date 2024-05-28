package monsters

import "github.com/koenic/rs-odds/pkg"

var Gauntlet = pkg.DefaultMonster{
	Name:   "Gauntlet",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Regular",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Youngllef",
						Odds: 1.0 / 2000,
					},
				},
				{
					{
						Name: "Crystal shard",
						Odds: 1.0,
					},
				},
				{
					{
						Name: "Crystal armour seed",
						Odds: 1.0 / 120,
					},
					{
						Name: "Crystal weapon seed",
						Odds: 1.0 / 120,
					},
					{
						Name: "Enhanced crystal weapon seed",
						Odds: 1.0 / 2000,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Youngllef",
						Amount: 1,
					},
					{
						Name:   "Crystal shard",
						Amount: 1,
					},
					{
						Name:   "Crystal armour seed",
						Amount: 1,
					},
					{
						Name:   "Crystal weapon seed",
						Amount: 1,
					},
					{
						Name:   "Enhanced crystal weapon seed",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "Corrupted",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Youngllef",
						Odds: 1.0 / 800,
					},
				},
				{
					{
						Name: "Crystal shard",
						Odds: 1.0,
					},
				},
				{
					{
						Name: "Crystal armour seed",
						Odds: 1.0 / 50,
					},
					{
						Name: "Crystal weapon seed",
						Odds: 1.0 / 50,
					},
					{
						Name: "Enhanced crystal weapon seed",
						Odds: 1.0 / 400,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Youngllef",
						Amount: 1,
					},
					{
						Name:   "Crystal shard",
						Amount: 1,
					},
					{
						Name:   "Crystal armour seed",
						Amount: 1,
					},
					{
						Name:   "Crystal weapon seed",
						Amount: 1,
					},
					{
						Name:   "Enhanced crystal weapon seed",
						Amount: 1,
					},
				},
			},
		},
	},
}
