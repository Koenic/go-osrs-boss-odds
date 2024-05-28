package monsters

import "github.com/koenic/rs-odds/pkg"

var Scurrius = pkg.DefaultMonster{
	Name:   "Scurrius (MVP solo)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Scurrius",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Scurry",
						Odds: 1.0 / 3000,
					},
				},
				{
					{
						Name: "Scurrius' spine",
						Odds: 1.0 / 33,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Scurry",
						Amount: 1,
					},
					{
						Name:   "Scurrius' spine",
						Amount: 1,
					},
				},
			},
		},
	},
}
