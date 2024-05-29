package monsters

import "github.com/koenic/rs-odds/pkg"

var DukeSucellus = pkg.DefaultMonster{
	Name:   "Duke Sucellus",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Duke Sucellus",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Baron",
						Odds: 1.0 / 2500,
					},
				},
				{
					{
						Name: "Chromium ingot",
						Odds: (1.0 / 90) * (3.0 / 8),
					},
					{
						Name: "Magus vestige Roll",
						Odds: (1.0 / 90) * (3.0 / 8),
					},
					{
						Name: "Eye of the duke",
						Odds: (1.0 / 90) * (1.0 / 8),
					},
					{
						Name: "Virtus mask",
						Odds: (1.0 / 90) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Virtus robe top",
						Odds: (1.0 / 90) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Virtus robe bottom",
						Odds: (1.0 / 90) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Frozen tablet",
						Odds: 1.0 / 25.8,
					},
					{
						Name: "Awakener's orb",
						Odds: 1.0 / 48.5,
					},
					{
						Name: "Ice quartz",
						Odds: 1.0 / 215.2,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Baron",
						Amount: 1,
					},
					{
						Name:   "Chromium ingot",
						Amount: 1,
					},
					{
						Name:   "Magus vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Eye of the duke",
						Amount: 1,
					},
					{
						Name:   "Virtus mask",
						Amount: 1,
					},
					{
						Name:   "Virtus robe top",
						Amount: 1,
					},
					{
						Name:   "Virtus robe bottom",
						Amount: 1,
					},
					{
						Name:   "Ice quartz",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Chromium ingot",
						Amount: 3,
					},
					{
						Name:   "Magus vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Eye of the duke",
						Amount: 1,
					},
					{
						Name:   "Virtus mask",
						Amount: 1,
					},
					{
						Name:   "Virtus robe top",
						Amount: 1,
					},
					{
						Name:   "Virtus robe bottom",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Chromium ingot",
						Amount: 3,
					},
					{
						Name:   "Magus vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Eye of the duke",
						Amount: 1,
					},
				},
			},
		},
	},
}
