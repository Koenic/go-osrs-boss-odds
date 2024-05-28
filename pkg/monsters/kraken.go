package monsters

import "github.com/koenic/rs-odds/pkg"

var Kraken = pkg.DefaultMonster{
	Name:   "Kraken",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Kraken",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet kraken",
						Odds: 1.0 / 2000,
					},
				},
				{
					{
						Name: "Jar of dirt",
						Odds: 1.0 / 1000,
					},
				},
				{
					{
						Name: "Trident of the seas (full)",
						Odds: 1.0 / 512,
					},
					{
						Name: "Kraken tentacle",
						Odds: 1.0 / 400,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet kraken",
						Amount: 1,
					},
					{
						Name:   "Jar of dirt",
						Amount: 1,
					},
					{
						Name:   "Trident of the seas (full)",
						Amount: 1,
					},
					{
						Name:   "Kraken tentacle",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Kraken tentacle",
						Amount: 10,
					},
				},
			},
		},
	},
}
