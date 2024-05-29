package monsters

import "github.com/koenic/rs-odds/pkg"

var ChambersOfXeric = pkg.DefaultMonster{
	Name:   "Chambers Of Xeric",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "40k points",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Dust",
						Odds: (1.0 / 400),
					},
					{
						Name: "Dexterous prayer scroll",
						Odds: ((30_000 / 8676.0) / 100) * (20.0 / 69),
					},
					{
						Name: "Arcane prayer scroll",
						Odds: ((30_000 / 8676.0) / 100) * (20.0 / 69),
					},
					{
						Name: "Twisted buckler",
						Odds: ((30_000 / 8676.0) / 100) * (4.0 / 69),
					},
					{
						Name: "Dragon hunter crossbow",
						Odds: ((30_000 / 8676.0) / 100) * (4.0 / 69),
					},
					{
						Name: "Dinh's bulwark",
						Odds: ((30_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Ancestral hat",
						Odds: ((30_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Ancestral robe top",
						Odds: ((30_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Ancestral robe bottom",
						Odds: ((30_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Dragon claws",
						Odds: ((30_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Elder maul",
						Odds: ((30_000 / 8676.0) / 100) * (2.0 / 69),
					},
					{
						Name: "Kodai insignia",
						Odds: ((30_000 / 8676.0) / 100) * (2.0 / 69),
					},
					{
						Name: "Twisted bow",
						Odds: ((30_000 / 8676.0) / 100) * (2.0 / 69),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Twisted buckler",
						Amount: 1,
					},
					{
						Name:   "Dinh's bulwark",
						Amount: 1,
					},
					{
						Name:   "Ancestral hat",
						Amount: 1,
					},
					{
						Name:   "Dragon claws",
						Amount: 1,
					},
					{
						Name:   "Elder maul",
						Amount: 1,
					},
					{
						Name:   "Kodai insignia",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "60k points",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Dexterous prayer scroll",
						Odds: ((60_000 / 8676.0) / 100) * (20.0 / 69),
					},
					{
						Name: "Arcane prayer scroll",
						Odds: ((60_000 / 8676.0) / 100) * (20.0 / 69),
					},
					{
						Name: "Twisted buckler",
						Odds: ((60_000 / 8676.0) / 100) * (4.0 / 69),
					},
					{
						Name: "Dragon hunter crossbow",
						Odds: ((60_000 / 8676.0) / 100) * (4.0 / 69),
					},
					{
						Name: "Dinh's bulwark",
						Odds: ((60_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Ancestral hat",
						Odds: ((60_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Ancestral robe top",
						Odds: ((60_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Ancestral robe bottom",
						Odds: ((60_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Dragon claws",
						Odds: ((60_000 / 8676.0) / 100) * (3.0 / 69),
					},
					{
						Name: "Elder maul",
						Odds: ((60_000 / 8676.0) / 100) * (2.0 / 69),
					},
					{
						Name: "Kodai insignia",
						Odds: ((60_000 / 8676.0) / 100) * (2.0 / 69),
					},
					{
						Name: "Twisted bow",
						Odds: ((60_000 / 8676.0) / 100) * (2.0 / 69),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Twisted buckler",
						Amount: 1,
					},
					{
						Name:   "Dinh's bulwark",
						Amount: 1,
					},
					{
						Name:   "Ancestral hat",
						Amount: 1,
					},
					{
						Name:   "Dragon claws",
						Amount: 1,
					},
					{
						Name:   "Elder maul",
						Amount: 1,
					},
					{
						Name:   "Kodai insignia",
						Amount: 1,
					},
				},
			},
		},
	},
}
