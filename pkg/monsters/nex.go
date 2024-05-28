package monsters

import "github.com/koenic/rs-odds/pkg"

var Nex = pkg.DefaultMonster{
	Name:   "Nex (assuming equal contribution)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "2-man",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Nexling",
						Odds: 1.0 / 172,
					},
				},
				{
					{
						Name: "Nihil shard",
						Odds: 8.0 / 209,
					},
					{
						Name: "Nihil shard",
						Odds: 5.0 / 82,
					},
				},
				{
					{
						Name: "Zaryte vambraces",
						Odds: (1.0 / 43) * (3.0 / 12) * 1.05,
					},
					{
						Name: "Nihil horn",
						Odds: (1.0 / 43) * (2.0 / 12) * 1.05,
					},
					{
						Name: "Torva full helm (damaged)",
						Odds: (1.0 / 43) * (2.0 / 12) * 1.05,
					},
					{
						Name: "Torva platebody (damaged)",
						Odds: (1.0 / 43) * (2.0 / 12) * 1.05,
					},
					{
						Name: "Torva platelegs (damaged)",
						Odds: (1.0 / 43) * (2.0 / 12) * 1.05,
					},
					{
						Name: "Ancient hilt",
						Odds: (1.0 / 43) * (1.0 / 12) * 1.05,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Nexling",
						Amount: 1,
					},
					{
						Name:   "Nihil shard",
						Amount: 1,
					},
					{
						Name:   "Zaryte vambraces",
						Amount: 1,
					},
					{
						Name:   "Nihil horn",
						Amount: 1,
					},
					{
						Name:   "Torva full helm (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platebody (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platelegs (damaged)",
						Amount: 1,
					},
					{
						Name:   "Ancient hilt",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Zaryte vambraces",
						Amount: 1,
					},
					{
						Name:   "Nihil horn",
						Amount: 1,
					},
					{
						Name:   "Torva full helm (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platebody (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platelegs (damaged)",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "3-man",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Nexling",
						Odds: 1.0 / 172,
					},
				},
				{
					{
						Name: "Nihil shard",
						Odds: 8.0 / 209,
					},
					{
						Name: "Nihil shard",
						Odds: 5.0 / 82,
					},
				},
				{
					{
						Name: "Zaryte vambraces",
						Odds: (1.0 / 86) * (3.0 / 12) * (1 + .1/3),
					},
					{
						Name: "Nihil horn",
						Odds: (1.0 / 86) * (2.0 / 12) * (1 + .1/3),
					},
					{
						Name: "Torva full helm (damaged)",
						Odds: (1.0 / 86) * (2.0 / 12) * (1 + .1/3),
					},
					{
						Name: "Torva platebody (damaged)",
						Odds: (1.0 / 86) * (2.0 / 12) * (1 + .1/3),
					},
					{
						Name: "Torva platelegs (damaged)",
						Odds: (1.0 / 86) * (2.0 / 12) * (1 + .1/3),
					},
					{
						Name: "Ancient hilt",
						Odds: (1.0 / 86) * (1.0 / 12) * (1 + .1/3),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Nexling",
						Amount: 1,
					},
					{
						Name:   "Nihil shard",
						Amount: 1,
					},
					{
						Name:   "Zaryte vambraces",
						Amount: 1,
					},
					{
						Name:   "Nihil horn",
						Amount: 1,
					},
					{
						Name:   "Torva full helm (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platebody (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platelegs (damaged)",
						Amount: 1,
					},
					{
						Name:   "Ancient hilt",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Zaryte vambraces",
						Amount: 1,
					},
					{
						Name:   "Nihil horn",
						Amount: 1,
					},
					{
						Name:   "Torva full helm (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platebody (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platelegs (damaged)",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "4-man",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Nexling",
						Odds: 1.0 / 172,
					},
				},
				{
					{
						Name: "Nihil shard",
						Odds: 8.0 / 209,
					},
					{
						Name: "Nihil shard",
						Odds: 5.0 / 82,
					},
				},
				{
					{
						Name: "Zaryte vambraces",
						Odds: (1.0 / 129) * (3.0 / 12) * (1 + .1/4),
					},
					{
						Name: "Nihil horn",
						Odds: (1.0 / 129) * (2.0 / 12) * (1 + .1/4),
					},
					{
						Name: "Torva full helm (damaged)",
						Odds: (1.0 / 129) * (2.0 / 12) * (1 + .1/4),
					},
					{
						Name: "Torva platebody (damaged)",
						Odds: (1.0 / 129) * (2.0 / 12) * (1 + .1/4),
					},
					{
						Name: "Torva platelegs (damaged)",
						Odds: (1.0 / 129) * (2.0 / 12) * (1 + .1/4),
					},
					{
						Name: "Ancient hilt",
						Odds: (1.0 / 129) * (1.0 / 12) * (1 + .1/4),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Nexling",
						Amount: 1,
					},
					{
						Name:   "Nihil shard",
						Amount: 1,
					},
					{
						Name:   "Zaryte vambraces",
						Amount: 1,
					},
					{
						Name:   "Nihil horn",
						Amount: 1,
					},
					{
						Name:   "Torva full helm (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platebody (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platelegs (damaged)",
						Amount: 1,
					},
					{
						Name:   "Ancient hilt",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Zaryte vambraces",
						Amount: 1,
					},
					{
						Name:   "Nihil horn",
						Amount: 1,
					},
					{
						Name:   "Torva full helm (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platebody (damaged)",
						Amount: 1,
					},
					{
						Name:   "Torva platelegs (damaged)",
						Amount: 1,
					},
				},
			},
		},
	},
}
