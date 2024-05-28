package monsters

import "github.com/koenic/rs-odds/pkg"

var CallistoAndArtio = pkg.DefaultMonster{
	Name:   "Callisto And Artio",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Callisto",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Claws of callisto",
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
						Name: "Voidwaker hilt",
						Odds: (1.0 / 360),
					},
					{
						Name: "Tyrannical ring",
						Odds: (1.0 / 512),
					},
				},
				{
					{
						Name: "Callisto cub",
						Odds: (1.0 / 1500),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Callisto cub",
						Amount: 1.0,
					},
					{
						Name:   "Claws of callisto",
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
						Name:   "Voidwaker hilt",
						Amount: 1.0,
					},
					{
						Name:   "Tyrannical ring",
						Amount: 1.0,
					},
				},
				{
					{
						Name:   "Claws of callisto",
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
						Name:   "Voidwaker hilt",
						Amount: 1.0,
					},
					{
						Name:   "Tyrannical ring",
						Amount: 1.0,
					},
				},
			},
		}, {
			VariationName: "Artio",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Claws of callisto",
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
						Name: "Voidwaker hilt",
						Odds: (1.0 / 912),
					},
					{
						Name: "Tyrannical ring",
						Odds: (1.0 / 716),
					},
				},
				{
					{
						Name: "Callisto cub",
						Odds: (1.0 / 2800),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Callisto cub",
						Amount: 1.0,
					},
					{
						Name:   "Claws of callisto",
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
						Name:   "Voidwaker hilt",
						Amount: 1.0,
					},
					{
						Name:   "Tyrannical ring",
						Amount: 1.0,
					},
				},
				{
					{
						Name:   "Claws of callisto",
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
						Name:   "Voidwaker hilt",
						Amount: 1.0,
					},
					{
						Name:   "Tyrannical ring",
						Amount: 1.0,
					},
				},
			},
		},
	},
}
