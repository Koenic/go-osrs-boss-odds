package monsters

import "github.com/koenic/rs-odds/pkg"

var FortisColloseum = pkg.DefaultMonster{
	Name:   "Fortis Colosseum	(allways completing all waves)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Fortis Colosseum (gambling quivers)",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Sunfire splinters",
						Odds: 1.0,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 344.4,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 3_100,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 620,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 305.6,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 2_750,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 550,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 266.7,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 2_400,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 480,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 243,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 2_187,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 437.3,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 243,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 2_187,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 437.3,
					},
					{
						Name: "Tonalztics of ralos (uncharged)",
						Odds: 1.0 / 1_312,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 201.5,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 1_813,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 362.7,
					},
					{
						Name: "Tonalztics of ralos (uncharged)",
						Odds: 1.0 / 1_088,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 160,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 1_440,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 288,
					},
					{
						Name: "Tonalztics of ralos (uncharged)",
						Odds: 1.0 / 864,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 118.5,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 1_067,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 213.3,
					},
					{
						Name: "Tonalztics of ralos (uncharged)",
						Odds: 1.0 / 640,
					},
				},
				{
					{
						Name: "Echo crystal",
						Odds: 1.0 / 77.04,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 693.3,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 138.7,
					},
					{
						Name: "Tonalztics of ralos (uncharged)",
						Odds: 1.0 / 416,
					},
				},
				{
					{
						Name: "Dizana's quiver (uncharged)",
						Odds: 1.0,
					},
				},
				{
					{
						Name: "Smol heredit",
						Odds: 1.0 / 200,
					},
				},
				{
					{
						Name: "Smol heredit",
						Odds: 1.0 / 200,
					},
					{
						Name: "Echo crystal",
						Odds: 1.0 / 320,
					},
					{
						Name: "Sunfire fanatic piece",
						Odds: 3.0 / 64,
					},
					{
						Name: "Tonalztics of ralos (uncharged)",
						Odds: 1.0 / 192,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Smol heredit",
						Amount: 1,
					},
					{
						Name:   "Echo crystal",
						Amount: 1,
					},
					{
						Name:   "Sunfire fanatic piece",
						Amount: 3,
					},
					{
						Name:   "Tonalztics of ralos (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Dizana's quiver (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Sunfire splinters",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Echo crystal",
						Amount: 1,
					},
					{
						Name:   "Sunfire fanatic piece",
						Amount: 3,
					},
					{
						Name:   "Tonalztics of ralos (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Dizana's quiver (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Sunfire splinters",
						Amount: 1,
					},
				},
			},
		},
	},
}
