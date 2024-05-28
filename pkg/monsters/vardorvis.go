package monsters

import "github.com/koenic/rs-odds/pkg"

var Vardorvis = pkg.DefaultMonster{
	Name:   "Vardorvis",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Vardorvis",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Butch",
						Odds: 1.0 / 2500,
					},
				},
				{
					{
						Name: "Chromium ingot",
						Odds: (1.0 / 96) * (3.0 / 8),
					},
					{
						Name: "Ultor vestige Roll",
						Odds: (1.0 / 96) * (3.0 / 8),
					},
					{
						Name: "Executioner's axe head",
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
						Name: "Strangled tablet",
						Odds: 1.0 / 25.8,
					},
					{
						Name: "Awakener's orb",
						Odds: 1.0 / 48.5,
					},
					{
						Name: "Blood quartz",
						Odds: 1.0 / 215.2,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Butch",
						Amount: 1,
					},
					{
						Name:   "Chromium ingot",
						Amount: 1,
					},
					{
						Name:   "Ultor vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Executioner's axe head",
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
						Name:   "Blood quartz",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Chromium ingot",
						Amount: 3,
					},
					{
						Name:   "Ultor vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Executioner's axe head",
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
