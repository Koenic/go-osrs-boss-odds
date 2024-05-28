package monsters

import "github.com/koenic/rs-odds/pkg"

var PhantomMuspah = pkg.DefaultMonster{
	Name:   "Phantom Muspah",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Phantom Muspah",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Muphin",
						Odds: (1.0 / 2500),
					},
				},
				{
					{
						Name: "Ancient essence",
						Odds: (60.0 / 100),
					},
					{
						Name: "Ancient essence",
						Odds: (23.0 / 100),
					},
					{
						Name: "Ancient essence",
						Odds: (10.0 / 100),
					},
					{
						Name: "Frozen cache",
						Odds: (4.0 / 100),
					},
					{
						Name: "Ancient icon",
						Odds: (2.0 / 100),
					},
					{
						Name: "Venator shard",
						Odds: (1.0 / 100),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Muphin",
						Amount: 1,
					},
					{
						Name:   "Ancient essence",
						Amount: 1,
					},
					{
						Name:   "Frozen cache",
						Amount: 1,
					},
					{
						Name:   "Ancient icon",
						Amount: 1,
					},
					{
						Name:   "Venator shard",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Ancient essence",
						Amount: 1,
					},
					{
						Name:   "Frozen cache",
						Amount: 1,
					},
					{
						Name:   "Ancient icon",
						Amount: 1,
					},
					{
						Name:   "Venator shard",
						Amount: 5,
					},
				},
			},
		},
	},
}
