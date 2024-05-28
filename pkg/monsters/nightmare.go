package monsters

import "github.com/koenic/rs-odds/pkg"

var Nightmare = pkg.DefaultMonster{
	Name:   "Nightmare (assuming solos)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Nightmare",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Little nightmare",
						Odds: 1.0 / 800,
					},
				},
				{
					{
						Name: "Jar of dreams",
						Odds: 1.0 / 1900,
					},
				},
				{
					{
						Name: "Nightmare staff",
						Odds: (1.0 / 120) * (3.0 / 10),
					},
					{
						Name: "Inquisitor's great helm",
						Odds: (1.0 / 120) * (2.0 / 10),
					},
					{
						Name: "Inquisitor's hauberk",
						Odds: (1.0 / 120) * (2.0 / 10),
					},
					{
						Name: "Inquisitor's plateskirt",
						Odds: (1.0 / 120) * (2.0 / 10),
					},
					{
						Name: "Inquisitor's mace",
						Odds: (1.0 / 120) * (1.0 / 10),
					},
				},
				{
					{
						Name: "Eldritch orb",
						Odds: (1.0 / 600) * (1.0 / 3),
					},
					{
						Name: "Harmonised orb",
						Odds: (1.0 / 600) * (1.0 / 3),
					},
					{
						Name: "Volatile orb",
						Odds: (1.0 / 600) * (1.0 / 3),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Little nightmare",
						Amount: 1,
					},
					{
						Name:   "Jar of dreams",
						Amount: 1,
					},
					{
						Name:   "Nightmare staff",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's great helm",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's hauberk",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's plateskirt",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's mace",
						Amount: 1,
					},
					{
						Name:   "Eldritch orb",
						Amount: 1,
					},
					{
						Name:   "Harmonised orb",
						Amount: 1,
					},
					{
						Name:   "Volatile orb",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Nightmare staff",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's great helm",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's hauberk",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's plateskirt",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's mace",
						Amount: 1,
					},
					{
						Name:   "Eldritch orb",
						Amount: 1,
					},
					{
						Name:   "Harmonised orb",
						Amount: 1,
					},
					{
						Name:   "Volatile orb",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "Nightmare rebalanced",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Little nightmare",
						Odds: 1.0 / 800,
					},
				},
				{
					{
						Name: "Jar of dreams",
						Odds: 1.0 / 1900,
					},
				},
				{
					{
						Name: "Nightmare staff",
						Odds: (1.0 / 120) * (3.0 / 10) / 0.8,
					},
					{
						Name: "Inquisitor's great helm",
						Odds: (1.0 / 120) * (2.0 / 10) / 0.7,
					},
					{
						Name: "Inquisitor's hauberk",
						Odds: (1.0 / 120) * (2.0 / 10) / 0.7,
					},
					{
						Name: "Inquisitor's plateskirt",
						Odds: (1.0 / 120) * (2.0 / 10) / 0.7,
					},
					{
						Name: "Inquisitor's mace",
						Odds: (1.0 / 120) * (1.0 / 10) / 0.625,
					},
				},
				{
					{
						Name: "Eldritch orb",
						Odds: (1.0 / 600) * (1.0 / 3) / 0.533,
					},
					{
						Name: "Harmonised orb",
						Odds: (1.0 / 600) * (1.0 / 3) / 0.533,
					},
					{
						Name: "Volatile orb",
						Odds: (1.0 / 600) * (1.0 / 3) / 0.533,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Little nightmare",
						Amount: 1,
					},
					{
						Name:   "Jar of dreams",
						Amount: 1,
					},
					{
						Name:   "Nightmare staff",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's great helm",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's hauberk",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's plateskirt",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's mace",
						Amount: 1,
					},
					{
						Name:   "Eldritch orb",
						Amount: 1,
					},
					{
						Name:   "Harmonised orb",
						Amount: 1,
					},
					{
						Name:   "Volatile orb",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Nightmare staff",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's great helm",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's hauberk",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's plateskirt",
						Amount: 1,
					},
					{
						Name:   "Inquisitor's mace",
						Amount: 1,
					},
					{
						Name:   "Eldritch orb",
						Amount: 1,
					},
					{
						Name:   "Harmonised orb",
						Amount: 1,
					},
					{
						Name:   "Volatile orb",
						Amount: 1,
					},
				},
			},
		},
	},
}
