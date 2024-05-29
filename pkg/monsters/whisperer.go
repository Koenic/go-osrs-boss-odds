package monsters

import "github.com/koenic/rs-odds/pkg"

var Whisperer = pkg.DefaultMonster{
	Name:   "The Whisperer",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "The Whisperer",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Wisp",
						Odds: 1.0 / 2500,
					},
				},
				{
					{
						Name: "Chromium ingot",
						Odds: (1.0 / 64) * (3.0 / 8),
					},
					{
						Name: "Bellator vestige Roll",
						Odds: (1.0 / 64) * (3.0 / 8),
					},
					{
						Name: "Siren's staff",
						Odds: (1.0 / 64) * (1.0 / 8),
					},
					{
						Name: "Virtus mask",
						Odds: (1.0 / 64) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Virtus robe top",
						Odds: (1.0 / 64) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Virtus robe bottom",
						Odds: (1.0 / 64) * (1.0 / 8) * (1.0 / 3),
					},
					{
						Name: "Sirenic tablet",
						Odds: 1.0 / 25.8,
					},
					{
						Name: "Awakener's orb",
						Odds: 1.0 / 48.5,
					},
					{
						Name: "Shadow quartz",
						Odds: 1.0 / 215.2,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Wisp",
						Amount: 1,
					},
					{
						Name:   "Chromium ingot",
						Amount: 1,
					},
					{
						Name:   "Bellator vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Siren's staff",
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
						Name:   "Shadow quartz",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Chromium ingot",
						Amount: 3,
					},
					{
						Name:   "Bellator vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Siren's staff",
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
						Name:   "Bellator vestige Roll",
						Amount: 3,
					},
					{
						Name:   "Siren's staff",
						Amount: 1,
					},
				},
			},
		},
	},
}
