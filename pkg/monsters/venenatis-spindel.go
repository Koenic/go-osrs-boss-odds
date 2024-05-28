package monsters

import "github.com/koenic/rs-odds/pkg"

var VenenatisSpindel = pkg.DefaultMonster{
	Name:   "Venenatis and Spindel",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Venenatis",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Fangs of venenatis",
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
						Name: "Voidwaker gem",
						Odds: (1.0 / 360),
					},
					{
						Name: "Treasonous ring",
						Odds: (1.0 / 512),
					},
				},
				{
					{
						Name: "Venenatis spiderling",
						Odds: (1.0 / 1500),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Venenatis spiderling",
						Amount: 1.0,
					},
					{
						Name:   "Fangs of venenatis",
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
						Name:   "Voidwaker gem",
						Amount: 1.0,
					},
					{
						Name:   "Treasonous ring",
						Amount: 1.0,
					},
				},
				{
					{
						Name:   "Fangs of venenatis",
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
						Name:   "Voidwaker gem",
						Amount: 1.0,
					},
					{
						Name:   "Treasonous ring",
						Amount: 1.0,
					},
				},
			},
		}, {
			VariationName: "Spindel",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Fangs of venenatis",
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
						Name: "Voidwaker gem",
						Odds: (1.0 / 912),
					},
					{
						Name: "Treasonous ring",
						Odds: (1.0 / 716),
					},
				},
				{
					{
						Name: "Venenatis spiderling",
						Odds: (1.0 / 2800),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Venenatis spiderling",
						Amount: 1.0,
					},
					{
						Name:   "Fangs of venenatis",
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
						Name:   "Voidwaker gem",
						Amount: 1.0,
					},
					{
						Name:   "Treasonous ring",
						Amount: 1.0,
					},
				},
				{
					{
						Name:   "Fangs of venenatis",
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
						Name:   "Voidwaker gem",
						Amount: 1.0,
					},
					{
						Name:   "Treasonous ring",
						Amount: 1.0,
					},
				},
			},
		},
	},
}
