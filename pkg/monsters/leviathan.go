package monsters

import "github.com/koenic/rs-odds/pkg"

var Leviathan = pkg.DefaultMonster{
	Name:   "Leviathan",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Leviathan",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Lil'viathan",
						Odds: 1.0 / 2500,
					},
				},
				{
					{
						Name: "Chromium ingot",
						Odds: (1.0 / 96) * (3.0 / 8),
					},
					{
						Name: "Venator vestige Roll",
						Odds: (1.0 / 96) * (3.0 / 8),
					},
					{
						Name: "Leviathan's lure",
						Odds: (1.0 / 96) * (1.0 / 8),
					},
					{
						Name: "Virtus mask",
						Odds: (1.0 / 96) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Virtus robe top",
						Odds: (1.0 / 96) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Virtus robe bottom",
						Odds: (1.0 / 96) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Scarred tablet",
						Odds: 1.0 / 25.8,
					},
					{
						Name: "Awakener's orb",
						Odds: 1.0 / 48.5,
					},
					{
						Name: "Smoke quartz",
						Odds: 1.0 / 215.2,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Lil'viathan",
						Amount: 1,
					},
					{
						Name:   "Chromium ingot",
						Amount: 1,
					},
					{
						Name:   "Venator vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Leviathan's lure",
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
						Name:   "Smoke quartz",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Chromium ingot",
						Amount: 3,
					},
					{
						Name:   "Venator vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Leviathan's lure",
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
			},
		},
	},
}
