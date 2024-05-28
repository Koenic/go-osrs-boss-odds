package monsters

import "github.com/koenic/rs-odds/pkg"

var Sarachnis = pkg.DefaultMonster{
	Name:   "Sarachnis",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Sarachnis",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Sraracha",
						Odds: (1.0 / 3000),
					},
				},
				{
					{
						Name: "Jar of eyes",
						Odds: (1.0 / 2000),
					},
				},
				{
					{
						Name: "Sarachnis cudgel",
						Odds: (1.0 / 384),
					},
				},
				{
					{
						Name: "Giant egg sac(full)",
						Odds: (1.0 / 20),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Sraracha",
						Amount: 1,
					},
					{
						Name:   "Jar of eyes",
						Amount: 1,
					},
					{
						Name:   "Sarachnis cudgel",
						Amount: 1,
					},
					{
						Name:   "Giant egg sac(full)",
						Amount: 1,
					},
				},
			},
		},
	},
}
