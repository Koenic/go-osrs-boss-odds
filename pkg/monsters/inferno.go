package monsters

import "github.com/koenic/rs-odds/pkg"

var Inferno = pkg.DefaultMonster{
	Name:   "Inferno (gambling capes)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Off task",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Jal-nib-rek",
						Odds: 1.0 / 100,
					},
				},
				{
					{
						Name: "Jal-nib-rek",
						Odds: 1.0 / 100,
					},
				},
				{
					{
						Name: "Infernal cape",
						Odds: 1.0,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Jal-nib-rek",
						Amount: 1,
					},
					{
						Name:   "Infernal cape",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "On task",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Jal-nib-rek",
						Odds: 1.0 / 75,
					},
				},
				{
					{
						Name: "Jal-nib-rek",
						Odds: 1.0 / 100,
					},
				},
				{
					{
						Name: "Infernal cape",
						Odds: 1.0,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Jal-nib-rek",
						Amount: 1,
					},
					{
						Name:   "Infernal cape",
						Amount: 1,
					},
				},
			},
		},
	},
}
