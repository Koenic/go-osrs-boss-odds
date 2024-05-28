package monsters

import "github.com/koenic/rs-odds/pkg"

var VetIonCalvarIon = pkg.DefaultMonster{
	Name:   "Vet'ion and Calvar'ion",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Vet'ion",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Skull of vet'ion",
						Odds: (1.0 / 196),
					},
					{
						Name: "Dragon 2h sword",
						Odds: (1.0 / 256),
					},
					{
						Name: "Dragon pickaxe",
						Odds: (1.0 / 256),
					},
					{
						Name: "Voidwaker blade",
						Odds: (1.0 / 360),
					},
					{
						Name: "Ring of the gods",
						Odds: (1.0 / 512),
					},
				},
				{
					{
						Name: "Vet'ion jr.",
						Odds: (1.0 / 1500),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Vet'ion jr.",
						Amount: 1.0,
					},
					{
						Name:   "Skull of vet'ion",
						Amount: 1.0,
					},
					{
						Name:   "Dragon 2h sword",
						Amount: 1.0,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1.0,
					},
					{
						Name:   "Voidwaker blade",
						Amount: 1.0,
					},
					{
						Name:   "Ring of the gods",
						Amount: 1.0,
					},
				},
				{
					{
						Name:   "Skull of vet'ion",
						Amount: 1.0,
					},
					{
						Name:   "Dragon 2h sword",
						Amount: 1.0,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1.0,
					},
					{
						Name:   "Voidwaker blade",
						Amount: 1.0,
					},
					{
						Name:   "Ring of the gods",
						Amount: 1.0,
					},
				},
			},
		}, {
			VariationName: "Calvar'ion",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Skull of vet'ion",
						Odds: (1.0 / 618),
					},
					{
						Name: "Dragon 2h sword",
						Odds: (1.0 / 358),
					},
					{
						Name: "Dragon pickaxe",
						Odds: (1.0 / 358),
					},
					{
						Name: "Voidwaker blade",
						Odds: (1.0 / 912),
					},
					{
						Name: "Ring of the gods",
						Odds: (1.0 / 716),
					},
				},
				{
					{
						Name: "Vet'ion jr.",
						Odds: (1.0 / 2800),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Vet'ion jr.",
						Amount: 1.0,
					},
					{
						Name:   "Skull of vet'ion",
						Amount: 1.0,
					},
					{
						Name:   "Dragon 2h sword",
						Amount: 1.0,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1.0,
					},
					{
						Name:   "Voidwaker blade",
						Amount: 1.0,
					},
					{
						Name:   "Ring of the gods",
						Amount: 1.0,
					},
				},
				{
					{
						Name:   "Skull of vet'ion",
						Amount: 1.0,
					},
					{
						Name:   "Dragon 2h sword",
						Amount: 1.0,
					},
					{
						Name:   "Dragon pickaxe",
						Amount: 1.0,
					},
					{
						Name:   "Voidwaker blade",
						Amount: 1.0,
					},
					{
						Name:   "Ring of the gods",
						Amount: 1.0,
					},
				},
			},
		},
	},
}
