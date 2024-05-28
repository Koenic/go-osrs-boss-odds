package monsters

import "github.com/koenic/rs-odds/pkg"

var GrotesqueGuardians = pkg.DefaultMonster{
	Name:   "Grotesque Guardians",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Grotesque Guardians",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Noon",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Jar of stone",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Granite maul",
						Odds: 1.0 / 250,
					},
					{
						Name: "Granite gloves",
						Odds: 1.0 / 500,
					},
					{
						Name: "Granite ring",
						Odds: 1.0 / 500,
					},
					{
						Name: "Granite hammer",
						Odds: 1.0 / 750,
					},
					{
						Name: "Black tourmaline core",
						Odds: 1.0 / 1000,
					},
				},
				{
					{
						Name: "Granite maul",
						Odds: 1.0 / 250,
					},
					{
						Name: "Granite gloves",
						Odds: 1.0 / 500,
					},
					{
						Name: "Granite ring",
						Odds: 1.0 / 500,
					},
					{
						Name: "Granite hammer",
						Odds: 1.0 / 750,
					},
					{
						Name: "Black tourmaline core",
						Odds: 1.0 / 1000,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Noon",
						Amount: 1,
					},
					{
						Name:   "Jar of stone",
						Amount: 1,
					},
					{
						Name:   "Granite maul",
						Amount: 1,
					},
					{
						Name:   "Granite gloves",
						Amount: 1,
					},
					{
						Name:   "Granite ring",
						Amount: 1,
					},
					{
						Name:   "Granite hammer",
						Amount: 1,
					},
					{
						Name:   "Black tourmaline core",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Granite maul",
						Amount: 1,
					},
					{
						Name:   "Granite gloves",
						Amount: 1,
					},
					{
						Name:   "Granite ring",
						Amount: 1,
					},
					{
						Name:   "Granite hammer",
						Amount: 1,
					},
					{
						Name:   "Black tourmaline core",
						Amount: 1,
					},
				},
			},
		},
	},
}
