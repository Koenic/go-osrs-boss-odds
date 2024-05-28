package monsters

import "github.com/koenic/rs-odds/pkg"

var ChaosElemental = pkg.DefaultMonster{
	Name:   "Chaos Elemental",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Chaos Elemental",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet chaos elemental",
						Odds: 1.0 / 300,
					},
				},
				{
					{
						Name: "Dragon pickaxe",
						Odds: 1.0 / 512,
					},
					{
						Name: "Dragon 2h sword",
						Odds: 1.0 / 64,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet chaos elemental",
						Amount: 1.0,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1.0,
					},
					{
						Name:   "Dragon 2h sword",
						Amount: 1.0,
					},
				},
			},
		},
	},
}
