package monsters

import "github.com/koenic/rs-odds/pkg"

var Bryophyta = pkg.DefaultMonster{
	Name:   "Bryophyta",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Bryophyta",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Bryophyta's essence",
						Odds: (1.0 / 118),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Bryophyta's essence",
						Amount: 1,
					},
				},
			},
		},
	},
}
