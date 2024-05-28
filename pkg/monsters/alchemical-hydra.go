package monsters

import "github.com/koenic/rs-odds/pkg"

var AlchemicalHydra = pkg.DefaultMonster{
	Name:   "Alchemical Hydra",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Alchemical Hydra",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Brimstone ring piece",
						Odds: 1.0 / (180.0 / ((1999.0 / 2000) * (1999.0 / 2000) * (999.0 / 1000) * (511.0 / 512) * (511.0 / 512))),
					},
					{
						Name: "Hydra leather",
						Odds: 1.0 / (512.0 / ((1999.0 / 2000) * (1999.0 / 2000) * (999.0 / 1000) * (511.0 / 512))),
					},
					{
						Name: "Hydra tail",
						Odds: 1.0 / (512.0 / ((1999.0 / 2000) * (1999.0 / 2000) * (999.0 / 1000))),
					},
					{
						Name: "Hydra's claw",
						Odds: 1.0 / (1000.0 / ((1999.0 / 2000) * (1999.0 / 2000))),
					},
					{
						Name: "Dragon knife",
						Odds: 1.0 / (2000.0 / (1999.0 / 2000)),
					},
					{
						Name: "Dragon thrownaxe",
						Odds: 1.0 / 2000.0,
					},
				},
				{
					{
						Name: "Alchemical hydra heads",
						Odds: 1.0 / 256.0,
					},
				},
				{
					{
						Name: "Jar of chemicals",
						Odds: 1.0 / 2000.0,
					},
				},
				{
					{
						Name: "Ikkle hydra",
						Odds: 1.0 / 3000.0,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Ikkle hydra",
						Amount: 1,
					},
					{
						Name:   "Brimstone ring piece",
						Amount: 3,
					},
					{
						Name:   "Hydra leather",
						Amount: 1,
					},
					{
						Name:   "Hydra tail",
						Amount: 1,
					},
					{
						Name:   "Hydra's claw",
						Amount: 1,
					},
					{
						Name:   "Dragon knife",
						Amount: 1,
					},
					{
						Name:   "Dragon thrownaxe",
						Amount: 1,
					},
					{
						Name:   "Alchemical hydra heads",
						Amount: 1,
					},
					{
						Name:   "Jar of chemicals",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Brimstone ring piece",
						Amount: 3,
					},
					{
						Name:   "Hydra leather",
						Amount: 1,
					},
					{
						Name:   "Hydra tail",
						Amount: 1,
					},
					{
						Name:   "Hydra's claw",
						Amount: 1,
					},
				},
			},
		},
	},
}
