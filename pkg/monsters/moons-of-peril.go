package monsters

import "github.com/koenic/rs-odds/pkg"

var MoonsOfPeril = pkg.DefaultMonster{
	Name:   "Moons of Peril (killing all bosses)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Moons of Peril",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Eclipse piece",
						Odds: 4.0 / 224,
					},
				},
				{
					{
						Name: "Blood moon piece",
						Odds: 4.0 / 224,
					},
				},
				{
					{
						Name: "Blue moon piece",
						Odds: 4.0 / 224,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Eclipse piece",
						Amount: 4,
					},
					{
						Name:   "Blood moon piece",
						Amount: 4,
					},
					{
						Name:   "Blue moon piece",
						Amount: 4,
					},
				},
			},
		},
	},
}
