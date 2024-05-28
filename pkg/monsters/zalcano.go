package monsters

import "github.com/koenic/rs-odds/pkg"

var Zalcano = pkg.DefaultMonster{
	Name:   "Zalcano (solos)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Zalcano",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Smolcano",
						Odds: 1.0 / 2000,
					},
				},
				{
					{
						Name: "Crystal tool seed",
						Odds: (1.0 / 200) * (39.0 / 40),
					},
					{
						Name: "Uncut onyx",
						Odds: (1.0 / 200) * (1.0 / 40),
					},
				},
				{
					{
						Name: "Zalcano shard",
						Odds: (1.0 / 750),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Smolcano",
						Amount: 1,
					},
					{
						Name:   "Crystal tool seed",
						Amount: 1,
					},
					{
						Name:   "Uncut onyx",
						Amount: 1,
					},
					{
						Name:   "Zalcano shard",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Crystal tool seed",
						Amount: 3,
					},
				},
			},
		},
	},
}
