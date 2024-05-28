package monsters

import "github.com/koenic/rs-odds/pkg"

var TheatreOfBlood = pkg.DefaultMonster{
	Name:   "Theatre of Blood (assuming equal contribution)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "2-man",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Lil' zik",
						Odds: (1.0 / 650),
					},
				},
				{
					{
						Name: "Avernic defender hilt",
						Odds: (1.0 / 2) * (1.0 / 9.1) * (8.0 / 19),
					},
					{
						Name: "Ghrazi rapier",
						Odds: (1.0 / 2) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Sanguinesti staff (uncharged)",
						Odds: (1.0 / 2) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar faceguard",
						Odds: (1.0 / 2) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar chestguard",
						Odds: (1.0 / 2) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar legguards",
						Odds: (1.0 / 2) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Scythe of vitur (uncharged)",
						Odds: (1.0 / 2) * (1.0 / 9.1) * (1.0 / 19),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Lil' zik",
						Amount: 1,
					},
					{
						Name:   "Avernic defender hilt",
						Amount: 1,
					},
					{
						Name:   "Ghrazi rapier",
						Amount: 1,
					},
					{
						Name:   "Sanguinesti staff (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Justiciar faceguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar chestguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar legguards",
						Amount: 1,
					},
					{
						Name:   "Scythe of vitur (uncharged)",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Avernic defender hilt",
						Amount: 1,
					},
					{
						Name:   "Ghrazi rapier",
						Amount: 1,
					},
					{
						Name:   "Sanguinesti staff (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Justiciar faceguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar chestguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar legguards",
						Amount: 1,
					},
					{
						Name:   "Scythe of vitur (uncharged)",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "3-man",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Lil' zik",
						Odds: (1.0 / 650),
					},
				},
				{
					{
						Name: "Avernic defender hilt",
						Odds: (1.0 / 3) * (1.0 / 9.1) * (8.0 / 19),
					},
					{
						Name: "Ghrazi rapier",
						Odds: (1.0 / 3) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Sanguinesti staff (uncharged)",
						Odds: (1.0 / 3) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar faceguard",
						Odds: (1.0 / 3) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar chestguard",
						Odds: (1.0 / 3) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar legguards",
						Odds: (1.0 / 3) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Scythe of vitur (uncharged)",
						Odds: (1.0 / 3) * (1.0 / 9.1) * (1.0 / 19),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Lil' zik",
						Amount: 1,
					},
					{
						Name:   "Avernic defender hilt",
						Amount: 1,
					},
					{
						Name:   "Ghrazi rapier",
						Amount: 1,
					},
					{
						Name:   "Sanguinesti staff (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Justiciar faceguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar chestguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar legguards",
						Amount: 1,
					},
					{
						Name:   "Scythe of vitur (uncharged)",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Avernic defender hilt",
						Amount: 1,
					},
					{
						Name:   "Ghrazi rapier",
						Amount: 1,
					},
					{
						Name:   "Sanguinesti staff (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Justiciar faceguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar chestguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar legguards",
						Amount: 1,
					},
					{
						Name:   "Scythe of vitur (uncharged)",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "4-man",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Lil' zik",
						Odds: (1.0 / 650),
					},
				},
				{
					{
						Name: "Avernic defender hilt",
						Odds: (1.0 / 4) * (1.0 / 9.1) * (8.0 / 19),
					},
					{
						Name: "Ghrazi rapier",
						Odds: (1.0 / 4) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Sanguinesti staff (uncharged)",
						Odds: (1.0 / 4) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar faceguard",
						Odds: (1.0 / 4) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar chestguard",
						Odds: (1.0 / 4) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar legguards",
						Odds: (1.0 / 4) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Scythe of vitur (uncharged)",
						Odds: (1.0 / 4) * (1.0 / 9.1) * (1.0 / 19),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Lil' zik",
						Amount: 1,
					},
					{
						Name:   "Avernic defender hilt",
						Amount: 1,
					},
					{
						Name:   "Ghrazi rapier",
						Amount: 1,
					},
					{
						Name:   "Sanguinesti staff (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Justiciar faceguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar chestguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar legguards",
						Amount: 1,
					},
					{
						Name:   "Scythe of vitur (uncharged)",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Avernic defender hilt",
						Amount: 1,
					},
					{
						Name:   "Ghrazi rapier",
						Amount: 1,
					},
					{
						Name:   "Sanguinesti staff (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Justiciar faceguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar chestguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar legguards",
						Amount: 1,
					},
					{
						Name:   "Scythe of vitur (uncharged)",
						Amount: 1,
					},
				},
			},
		},
		{
			VariationName: "5-man",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Lil' zik",
						Odds: (1.0 / 650),
					},
				},
				{
					{
						Name: "Avernic defender hilt",
						Odds: (1.0 / 5) * (1.0 / 9.1) * (8.0 / 19),
					},
					{
						Name: "Ghrazi rapier",
						Odds: (1.0 / 5) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Sanguinesti staff (uncharged)",
						Odds: (1.0 / 5) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar faceguard",
						Odds: (1.0 / 5) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar chestguard",
						Odds: (1.0 / 5) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Justiciar legguards",
						Odds: (1.0 / 5) * (1.0 / 9.1) * (2.0 / 19),
					},
					{
						Name: "Scythe of vitur (uncharged)",
						Odds: (1.0 / 5) * (1.0 / 9.1) * (1.0 / 19),
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Lil' zik",
						Amount: 1,
					},
					{
						Name:   "Avernic defender hilt",
						Amount: 1,
					},
					{
						Name:   "Ghrazi rapier",
						Amount: 1,
					},
					{
						Name:   "Sanguinesti staff (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Justiciar faceguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar chestguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar legguards",
						Amount: 1,
					},
					{
						Name:   "Scythe of vitur (uncharged)",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Avernic defender hilt",
						Amount: 1,
					},
					{
						Name:   "Ghrazi rapier",
						Amount: 1,
					},
					{
						Name:   "Sanguinesti staff (uncharged)",
						Amount: 1,
					},
					{
						Name:   "Justiciar faceguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar chestguard",
						Amount: 1,
					},
					{
						Name:   "Justiciar legguards",
						Amount: 1,
					},
					{
						Name:   "Scythe of vitur (uncharged)",
						Amount: 1,
					},
				},
			},
		},
	},
}
