package monsters

import "github.com/koenic/rs-odds/pkg"

var Vorkath = pkg.DefaultMonster{
	Name:   "Vorkath",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Vorkath",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Vorki",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Draconic visage",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Skeletal visage",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Jar of decay",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Dragonbone necklace",
						Odds: 1.0 / 1000,
					},
				},
				{
					{
						Name: "Vorkath's head",
						Odds: 1.0 / 50,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Vorki",
						Amount: 1,
					},
					{
						Name:   "Draconic visage",
						Amount: 1,
					},
					{
						Name:   "Skeletal visage",
						Amount: 1,
					},
					{
						Name:   "Jar of decay",
						Amount: 1,
					},
					{
						Name:   "Dragonbone necklace",
						Amount: 1,
					},
					{
						Name:   "Vorkath's head",
						Amount: 1,
					},
				},
			},
		},
	},
}
