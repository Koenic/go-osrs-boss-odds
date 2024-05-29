package monsters

import "github.com/koenic/rs-odds/pkg"

var Wintertodt = pkg.DefaultMonster{
	Name:   "Wintertodt (Check wiki to see how many loot rolls you get)",
	KcName: "loot rolls",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Wintertodt",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Dragon axe",
						Odds: 1.0 / 10_000,
					},
					{
						Name: "Phoenix",
						Odds: 1.0 / (5_000 * (9_999.0 / 10_000)),
					},
					{
						Name: "Tome of fire",
						Odds: 1.0 / (150 * (4_999.0 / 5_000) * (9_999.0 / 10_000)),
					},
					{
						Name: "Warm gloves",
						Odds: 1.0 / (150 * (999.0 / 1000) * (4_999.0 / 5_000) * (9_999.0 / 10_000)),
					},
					{
						Name: "Bruma torch",
						Odds: 1.0 / (150 * (149.0 / 150) * (999.0 / 1000) * (4_999.0 / 5_000) * (9_999.0 / 10_000)),
					},
					{
						Name: "Pyromancer outfit piece",
						Odds: 1.0 / (150 * (149.0 / 150) * (149.0 / 150) * (999.0 / 1000) * (4_999.0 / 5_000) * (9_999.0 / 10_000)),
					},
					{
						Name: "Burnt page",
						Odds: 1.0 / (45 * (149.0 / 150) * (149.0 / 150) * (149.0 / 150) * (999.0 / 1000) * (4_999.0 / 5_000) * (9_999.0 / 10_000)),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Dragon axe",
						Amount: 1,
					},
					{
						Name:   "Phoenix",
						Amount: 1,
					},
					{
						Name:   "Tome of fire",
						Amount: 1,
					},
					{
						Name:   "Warm gloves",
						Amount: 1,
					},
					{
						Name:   "Bruma torch",
						Amount: 1,
					},
					{
						Name:   "Pyromancer outfit piece",
						Amount: 4,
					},
					{
						Name:   "Burnt page",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Phoenix",
						Amount: 1,
					},
					{
						Name:   "Tome of fire",
						Amount: 1,
					},
					{
						Name:   "Warm gloves",
						Amount: 1,
					},
					{
						Name:   "Bruma torch",
						Amount: 1,
					},
					{
						Name:   "Pyromancer outfit piece",
						Amount: 4,
					},
					{
						Name:   "Burnt page",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Tome of fire",
						Amount: 1,
					},
					{
						Name:   "Warm gloves",
						Amount: 1,
					},
					{
						Name:   "Bruma torch",
						Amount: 1,
					},
					{
						Name:   "Pyromancer outfit piece",
						Amount: 4,
					},
					{
						Name:   "Burnt page",
						Amount: 1,
					},
				},
			},
		},
	},
}
