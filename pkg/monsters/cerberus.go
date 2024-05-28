package monsters

import "github.com/koenic/rs-odds/pkg"

var Cerberus = pkg.DefaultMonster{
	Name:   "Cerberus",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Cerberus",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Hellpuppy",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Jar of souls",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Primordial crystals",
						Odds: 1.0 / 512,
					},
					{
						Name: "Pegasian crystal",
						Odds: 1.0 / 512,
					},
					{
						Name: "Eternal crystal",
						Odds: 1.0 / 512,
					},
					{
						Name: "Smouldering stone",
						Odds: 1.0 / 512,
					},
					{
						Name: "Key master teleport",
						Odds: 1.0 / 64,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Hellpuppy",
						Amount: 1,
					},
					{
						Name:   "Jar of souls",
						Amount: 1,
					},
					{
						Name:   "Primordial crystals",
						Amount: 1,
					},
					{
						Name:   "Pegasian crystal",
						Amount: 1,
					},
					{
						Name:   "Eternal crystal",
						Amount: 1,
					},
					{
						Name:   "Smouldering stone",
						Amount: 1,
					},
					{
						Name:   "Key master teleport",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Primordial crystals",
						Amount: 1,
					},
					{
						Name:   "Pegasian crystal",
						Amount: 1,
					},
					{
						Name:   "Eternal crystal",
						Amount: 1,
					},
					{
						Name:   "Smouldering stone",
						Amount: 1,
					},
				},
			},
		},
	},
}
