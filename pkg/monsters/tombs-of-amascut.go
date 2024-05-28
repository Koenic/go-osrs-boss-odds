package monsters

import "github.com/koenic/rs-odds/pkg"

var TombsOfAmascut = pkg.DefaultMonster{
	Name:   "Tombs of Amascut",
	KcName: "chests",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Solo 300s",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Tumeken's guardian",
						Odds: 0.01 * (19_114 / (350_000.0 - 700*300)),
					},
				},
				{
					{
						Name: "Osmumten's fang",
						Odds: (7.0 / 24) * 0.01 * (19_114 / (10_500.0 - 20*300)),
					},
					{
						Name: "Lightbearer",
						Odds: (7.0 / 24) * 0.01 * (19_114 / (10_500.0 - 20*300)),
					},
					{
						Name: "Elidinis' ward",
						Odds: (3.0 / 24) * 0.01 * (19_114 / (10_500.0 - 20*300)),
					},
					{
						Name: "Masori mask",
						Odds: (2.0 / 24) * 0.01 * (19_114 / (10_500.0 - 20*300)),
					},
					{
						Name: "Masori body",
						Odds: (2.0 / 24) * 0.01 * (19_114 / (10_500.0 - 20*300)),
					},
					{
						Name: "Masori chaps",
						Odds: (2.0 / 24) * 0.01 * (19_114 / (10_500.0 - 20*300)),
					},
					{
						Name: "Tumeken's shadow (uncharged)",
						Odds: (1.0 / 24) * 0.01 * (19_114 / (10_500.0 - 20*300)),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Tumeken's guardian",
						Amount: 1,
					},
					{
						Name:   "Osmumten's fang",
						Amount: 1,
					},
					{
						Name:   "Lightbearer",
						Amount: 1,
					},
					{
						Name:   "Elidinis' ward",
						Amount: 1,
					},
					{
						Name:   "Masori mask",
						Amount: 1,
					},
					{
						Name:   "Masori body",
						Amount: 1,
					},
					{
						Name:   "Masori chaps",
						Amount: 1,
					},
					{
						Name:   "Tumeken's shadow (uncharged)",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Osmumten's fang",
						Amount: 1,
					},
					{
						Name:   "Lightbearer",
						Amount: 1,
					},
					{
						Name:   "Elidinis' ward",
						Amount: 1,
					},
					{
						Name:   "Masori mask",
						Amount: 1,
					},
					{
						Name:   "Masori body",
						Amount: 1,
					},
					{
						Name:   "Masori chaps",
						Amount: 1,
					},
					{
						Name:   "Tumeken's shadow (uncharged)",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "Solo 400s",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Tumeken's guardian",
						Odds: 0.01 * (22_197 / (350_000.0 - 700*400)),
					},
				},
				{
					{
						Name: "Osmumten's fang",
						Odds: (7.0 / 24) * 0.01 * (22_197 / (10_500.0 - 20*400)),
					},
					{
						Name: "Lightbearer",
						Odds: (7.0 / 24) * 0.01 * (22_197 / (10_500.0 - 20*400)),
					},
					{
						Name: "Elidinis' ward",
						Odds: (3.0 / 24) * 0.01 * (22_197 / (10_500.0 - 20*400)),
					},
					{
						Name: "Masori mask",
						Odds: (2.0 / 24) * 0.01 * (22_197 / (10_500.0 - 20*400)),
					},
					{
						Name: "Masori body",
						Odds: (2.0 / 24) * 0.01 * (22_197 / (10_500.0 - 20*400)),
					},
					{
						Name: "Masori chaps",
						Odds: (2.0 / 24) * 0.01 * (22_197 / (10_500.0 - 20*400)),
					},
					{
						Name: "Tumeken's shadow (uncharged)",
						Odds: (1.0 / 24) * 0.01 * (22_197 / (10_500.0 - 20*400)),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Tumeken's guardian",
						Amount: 1,
					},
					{
						Name:   "Osmumten's fang",
						Amount: 1,
					},
					{
						Name:   "Lightbearer",
						Amount: 1,
					},
					{
						Name:   "Elidinis' ward",
						Amount: 1,
					},
					{
						Name:   "Masori mask",
						Amount: 1,
					},
					{
						Name:   "Masori body",
						Amount: 1,
					},
					{
						Name:   "Masori chaps",
						Amount: 1,
					},
					{
						Name:   "Tumeken's shadow (uncharged)",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Osmumten's fang",
						Amount: 1,
					},
					{
						Name:   "Lightbearer",
						Amount: 1,
					},
					{
						Name:   "Elidinis' ward",
						Amount: 1,
					},
					{
						Name:   "Masori mask",
						Amount: 1,
					},
					{
						Name:   "Masori body",
						Amount: 1,
					},
					{
						Name:   "Masori chaps",
						Amount: 1,
					},
					{
						Name:   "Tumeken's shadow (uncharged)",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "Solo 500s",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Tumeken's guardian",
						Odds: 0.01 * (25_779 / (350_000.0 - (700*400 + 100/3))),
					},
				},
				{
					{
						Name: "Osmumten's fang",
						Odds: (7.0 / 24) * 0.01 * (25_779 / (10_500.0 - (20*400 + 100/3))),
					},
					{
						Name: "Lightbearer",
						Odds: (7.0 / 24) * 0.01 * (25_779 / (10_500.0 - (20*400 + 100/3))),
					},
					{
						Name: "Elidinis' ward",
						Odds: (3.0 / 24) * 0.01 * (25_779 / (10_500.0 - (20*400 + 100/3))),
					},
					{
						Name: "Masori mask",
						Odds: (2.0 / 24) * 0.01 * (25_779 / (10_500.0 - (20*400 + 100/3))),
					},
					{
						Name: "Masori body",
						Odds: (2.0 / 24) * 0.01 * (25_779 / (10_500.0 - (20*400 + 100/3))),
					},
					{
						Name: "Masori chaps",
						Odds: (2.0 / 24) * 0.01 * (25_779 / (10_500.0 - (20*400 + 100/3))),
					},
					{
						Name: "Tumeken's shadow (uncharged)",
						Odds: (1.0 / 24) * 0.01 * (25_779 / (10_500.0 - (20*400 + 100/3))),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Tumeken's guardian",
						Amount: 1,
					},
					{
						Name:   "Osmumten's fang",
						Amount: 1,
					},
					{
						Name:   "Lightbearer",
						Amount: 1,
					},
					{
						Name:   "Elidinis' ward",
						Amount: 1,
					},
					{
						Name:   "Masori mask",
						Amount: 1,
					},
					{
						Name:   "Masori body",
						Amount: 1,
					},
					{
						Name:   "Masori chaps",
						Amount: 1,
					},
					{
						Name:   "Tumeken's shadow (uncharged)",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Osmumten's fang",
						Amount: 1,
					},
					{
						Name:   "Lightbearer",
						Amount: 1,
					},
					{
						Name:   "Elidinis' ward",
						Amount: 1,
					},
					{
						Name:   "Masori mask",
						Amount: 1,
					},
					{
						Name:   "Masori body",
						Amount: 1,
					},
					{
						Name:   "Masori chaps",
						Amount: 1,
					},
					{
						Name:   "Tumeken's shadow (uncharged)",
						Amount: 1,
					},
				},
			},
		},
	},
}
