package monsters

import "github.com/koenic/rs-odds/pkg"

var GeneralGraardor = pkg.DefaultMonster{
	Name:   "General Graardor (including minions)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "General Graardor",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet general graardor",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Bandos chestplate",
						Odds: 1.0 / 381,
					},
					{
						Name: "Bandos tassets",
						Odds: 1.0 / 381,
					},
					{
						Name: "Bandos boots",
						Odds: 1.0 / 381,
					},
					{
						Name: "Bandos hilt",
						Odds: 1.0 / 508,
					},
					{
						Name: "Godsword shard 1",
						Odds: 1.0 / 762,
					},
					{
						Name: "Godsword shard 2",
						Odds: 1.0 / 762,
					},
					{
						Name: "Godsword shard 3",
						Odds: 1.0 / 762,
					},
				},
				{
					{
						Name: "Bandos chestplate",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Bandos tassets",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Bandos boots",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Godsword shard 1",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 2",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 3",
						Odds: 1.0 / 1_524,
					},
				},
				{
					{
						Name: "Bandos chestplate",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Bandos tassets",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Bandos boots",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Godsword shard 1",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 2",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 3",
						Odds: 1.0 / 1_524,
					},
				},
				{
					{
						Name: "Bandos chestplate",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Bandos tassets",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Bandos boots",
						Odds: 1.0 / 16_256,
					},
					{
						Name: "Godsword shard 1",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 2",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 3",
						Odds: 1.0 / 1_524,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet general graardor",
						Amount: 1,
					},
					{
						Name:   "Bandos chestplate",
						Amount: 1,
					},
					{
						Name:   "Bandos tassets",
						Amount: 1,
					},
					{
						Name:   "Bandos boots",
						Amount: 1,
					},
					{
						Name:   "Bandos hilt",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 1",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 2",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 3",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Bandos chestplate",
						Amount: 1,
					},
					{
						Name:   "Bandos tassets",
						Amount: 1,
					},
					{
						Name:   "Bandos boots",
						Amount: 1,
					},
					{
						Name:   "Bandos hilt",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 1",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 2",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 3",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Bandos chestplate",
						Amount: 1,
					},
					{
						Name:   "Bandos tassets",
						Amount: 1,
					},
					{
						Name:   "Bandos boots",
						Amount: 1,
					},
					{
						Name:   "Bandos hilt",
						Amount: 1,
					},
				},
			},
		},
	},
}
