package monsters

import "github.com/koenic/rs-odds/pkg"

var Obor = pkg.DefaultMonster{
	Name:   "Obor",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Obor",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Hill giant club",
						Odds: (1.0 / 118),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Hill giant club",
						Amount: 1,
					},
				},
			},
		},
	},
}
