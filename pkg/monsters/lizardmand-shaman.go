package monsters

import "github.com/koenic/rs-odds/pkg"

var LizardmanShaman = pkg.DefaultMonster{
	Name:   "Lizardman Shaman",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Shaman",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Dragon warhammer",
						Odds: (1.0 / 5000),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Dragon warhammer",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "Shaman rebalanced",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Dragon warhammer",
						Odds: (1.0 / 3000),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Dragon warhammer",
						Amount: 1,
					},
				},
			},
		},
	},
}
