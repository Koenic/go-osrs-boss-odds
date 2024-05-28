package monsters

import "github.com/koenic/rs-odds/pkg"

var Tempoross = pkg.DefaultMonster{
	Name:   "Tempoross",
	KcName: "loot rolls",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Tempoross",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Tiny tempor",
						Odds: 1.0 / 8000,
					},
					{
						Name: "Spirit flakes",
						Odds: 1.0 / 4,
					},
					{
						Name: "Soaked page",
						Odds: 149.0 / 8000,
					},
					{
						Name: "Fish barrel",
						Odds: 1.0 / 400,
					},
					{
						Name: "Tackle box",
						Odds: 1.0 / 400,
					},
					{
						Name: "Big harpoonfish",
						Odds: 1.0 / 1600,
					},
					{
						Name: "Tome of water (empty)",
						Odds: 1.0 / 1600,
					},
					{
						Name: "Dragon harpoon",
						Odds: 1.0 / 8000,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Tiny tempor",
						Amount: 1,
					},
					{
						Name:   "Spirit flakes",
						Amount: 1,
					},
					{
						Name:   "Soaked page",
						Amount: 1,
					},
					{
						Name:   "Fish barrel",
						Amount: 1,
					},
					{
						Name:   "Tackle box",
						Amount: 1,
					},
					{
						Name:   "Big harpoonfish",
						Amount: 1,
					},
					{
						Name:   "Tome of water (empty)",
						Amount: 1,
					},
					{
						Name:   "Dragon harpoon",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Spirit flakes",
						Amount: 1,
					},
					{
						Name:   "Soaked page",
						Amount: 1,
					},
					{
						Name:   "Fish barrel",
						Amount: 1,
					},
					{
						Name:   "Tackle box",
						Amount: 1,
					},
					{
						Name:   "Tome of water (empty)",
						Amount: 1,
					},
				},
			},
		},
	},
}
