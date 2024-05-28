package monsters

import "github.com/koenic/rs-odds/pkg"

var Zulrah = pkg.DefaultMonster{
	Name:   "Zulrah",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Zulrah",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Jar of swamp",
						Odds: (1.0 / 3000),
					},
					{
						Name: "Pet snakeling",
						Odds: (1.0 / 4000),
					},
				},
				{
					{
						Name: "Tanzanite fang",
						Odds: (1.0 / 256) * (1.0 / 4),
					},
					{
						Name: "Magic fang",
						Odds: (1.0 / 256) * (1.0 / 4),
					},
					{
						Name: "Serpentine visage",
						Odds: (1.0 / 256) * (1.0 / 4),
					},
					{
						Name: "Uncut onyx",
						Odds: (1.0 / 256) * (1.0 / 4),
					},
					{

						Name: "Tanzanite mutagen",
						Odds: (1.0 / 2) * (10.0 / 248) * (10.0 / 2_632),
					},
					{

						Name: "Magma mutagen",
						Odds: (1.0 / 2) * (10.0 / 248) * (10.0 / 2_632),
					},
				},
				{
					{
						Name: "Tanzanite fang",
						Odds: (1.0 / 256) * (1.0 / 4),
					},
					{
						Name: "Magic fang",
						Odds: (1.0 / 256) * (1.0 / 4),
					},
					{
						Name: "Serpentine visage",
						Odds: (1.0 / 256) * (1.0 / 4),
					},
					{
						Name: "Uncut onyx",
						Odds: (1.0 / 256) * (1.0 / 4),
					},
					{
						Name: "Tanzanite mutagen",
						Odds: (1.0 / 2) * (10.0 / 248) * (10.0 / 2_632),
					},
					{
						Name: "Magma mutagen",
						Odds: (1.0 / 2) * (10.0 / 248) * (10.0 / 2_632),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet snakeling",
						Amount: 1,
					},
					{
						Name:   "Jar of swamp",
						Amount: 1,
					},
					{
						Name:   "Tanzanite fang",
						Amount: 1,
					},
					{
						Name:   "Magic fang",
						Amount: 1,
					},
					{
						Name:   "Serpentine visage",
						Amount: 1,
					},
					{
						Name:   "Uncut onyx",
						Amount: 1,
					},
					{

						Name:   "Tanzanite mutagen",
						Amount: 1,
					},
					{
						Name:   "Magma mutagen",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Tanzanite fang",
						Amount: 1,
					},
					{
						Name:   "Magic fang",
						Amount: 1,
					},
					{
						Name:   "Serpentine visage",
						Amount: 1,
					},
				},
			},
		},
	},
}
