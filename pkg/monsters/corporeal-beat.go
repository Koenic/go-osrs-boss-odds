package monsters

import "github.com/koenic/rs-odds/pkg"

var CorporealBeast = pkg.DefaultMonster{
	Name:   "Corporeal Beast",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Corporeal Beast",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet dark core",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Jar of spirits",
						Odds: 1.0 / 1000,
					},
				},
				{
					{
						Name: "Spectral sigil",
						Odds: (1.0 / 585) * (3.0 / 7),
					},
					{
						Name: "Spectral sigil",
						Odds: (1.0 / 585) * (3.0 / 7),
					},
					{
						Name: "Elysian sigil",
						Odds: (1.0 / 585) * (1.0 / 7),
					},
					{
						Name: "Spirit shield",
						Odds: 1.0 / 64,
					},
					{
						Name: "Holy Elixer",
						Odds: 3.0 / 512,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet dark core",
						Amount: 1,
					},
					{
						Name:   "Jar of spirits",
						Amount: 1,
					},
					{
						Name:   "Spectral sigil",
						Amount: 1,
					},
					{
						Name:   "Spectral sigil",
						Amount: 1,
					},
					{
						Name:   "Elysian sigil",
						Amount: 1,
					},
					{
						Name:   "Spirit shield",
						Amount: 1,
					},
					{
						Name:   "Holy Elixer",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Spectral sigil",
						Amount: 1,
					},
					{
						Name:   "Spectral sigil",
						Amount: 1,
					},
					{
						Name:   "Elysian sigil",
						Amount: 1,
					},
					{
						Name:   "Spirit shield",
						Amount: 3,
					},
					{
						Name:   "Holy Elixer",
						Amount: 3,
					},
				},
			},
		},
	},
}
