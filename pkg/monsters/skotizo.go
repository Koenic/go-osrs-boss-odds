package monsters

import "github.com/koenic/rs-odds/pkg"

var Skotizo = pkg.DefaultMonster{
	Name:   "Skotizo",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Skotizo",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Skotos",
						Odds: 1.0 / 65,
					},
				},
				{
					{
						Name: "Jar of darkness",
						Odds: 1.0 / 200,
					},
				},
				{
					{
						Name: "Dark totem",
						Odds: 1.0 / 200,
					},
				},
				{
					{
						Name: "Dark claw",
						Odds: 1.0 / 200,
					},
				},
				{
					{
						Name: "Ancient shard",
						Odds: 1.0,
					},
				},
				{
					{
						Name: "Uncut onyx",
						Odds: 1.0 / 1000,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Skotos",
						Amount: 1,
					},
					{
						Name:   "Jar of darkness",
						Amount: 1,
					},
					{
						Name:   "Dark totem",
						Amount: 1,
					},
					{
						Name:   "Dark claw",
						Amount: 1,
					},
					{
						Name:   "Ancient shard",
						Amount: 1,
					},
					{
						Name:   "Uncut onyx",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Skotos",
						Amount: 1,
					},
					{
						Name:   "Jar of darkness",
						Amount: 1,
					},
					{
						Name:   "Dark totem",
						Amount: 1,
					},
					{
						Name:   "Dark claw",
						Amount: 1,
					},
					{
						Name:   "Ancient shard",
						Amount: 1,
					},
				},
			},
		},
	},
}
