package monsters

import "github.com/koenic/rs-odds/pkg"

var ThermonuclearSmokeSevil = pkg.DefaultMonster{
	Name:   "Thermonuclear smoke devil",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Thermonuclear smoke devil",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet smoke devil",
						Odds: 1.0 / 4,
					},
				},
				{
					{
						Name: "Jar of smoke",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Occult necklace",
						Odds: 1.0 / 350,
					},
					{
						Name: "Smoke battlestaff",
						Odds: 1.0 / 512,
					},
					{
						Name: "Dragon chainbody",
						Odds: 1.0 / 2000,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet smoke devil",
						Amount: 1,
					},
					{
						Name:   "Jar of smoke",
						Amount: 1,
					},
					{
						Name:   "Occult necklace",
						Amount: 1,
					},
					{
						Name:   "Smoke battlestaff",
						Amount: 1,
					},
					{
						Name:   "Dragon chainbody",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Occult necklace",
						Amount: 1,
					},
					{
						Name:   "Smoke battlestaff",
						Amount: 1,
					},
					{
						Name:   "Dragon chainbody",
						Amount: 1,
					},
				},
			},
		},
	},
}
