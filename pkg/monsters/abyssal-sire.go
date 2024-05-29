package monsters

import "github.com/koenic/rs-odds/pkg"

var AbyssalSire = pkg.DefaultMonster{
	Name:   "Abyssal Sire",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Abyssal Sire",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Bludgeon piece",
						Odds: (1.0 / 100) * (62.0 / 128),
					},
					{
						Name: "Abyssal dagger",
						Odds: (1.0 / 100) * (26.0 / 128),
					},
					{
						Name: "Abyssal whip",
						Odds: (1.0 / 100) * (12.0 / 128),
					},
					{
						Name: "Jar of miasma",
						Odds: (1.0 / 100) * (13.0 / 128),
					},
					{
						Name: "Abyssal head",
						Odds: (1.0 / 100) * (10.0 / 128),
					},
					{
						Name: "Abyssal orphan",
						Odds: (1.0 / 100) * (5.0 / 128),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Abyssal orphan",
						Amount: 1,
					},
					{
						Name:   "Bludgeon piece",
						Amount: 3,
					},
					{
						Name:   "Abyssal dagger",
						Amount: 1,
					},
					{
						Name:   "Abyssal whip",
						Amount: 1,
					},
					{
						Name:   "Jar of miasma",
						Amount: 1,
					},
					{
						Name:   "Abyssal head",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Bludgeon piece",
						Amount: 3,
					},
					{
						Name:   "Abyssal dagger",
						Amount: 1,
					},
				},
			},
		},
	},
}
