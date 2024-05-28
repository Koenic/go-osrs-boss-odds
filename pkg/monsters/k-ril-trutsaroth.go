package monsters

import "github.com/koenic/rs-odds/pkg"

var KRilTsutsaroth = pkg.DefaultMonster{
	Name:   "K'ril Tsutsaroth (including minions)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "K'ril Tsutsaroth",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet k'ril tsutsaroth",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Steam battlestaff",
						Odds: 1.0 / 127,
					},
					{
						Name: "Zamorakian spear",
						Odds: 1.0 / 127,
					},
					{
						Name: "Staff of the dead",
						Odds: 1.0 / 508,
					},
					{
						Name: "Zamorak hilt",
						Odds: 1.0 / 508,
					},
					{
						Name: "Godsword shard 1",
						Odds: 1.0 / 762,
					},
					{
						Name: "Godsword shard 2",
						Odds: 1.0 / 762,
					},
					{
						Name: "Godsword shard 3",
						Odds: 1.0 / 762,
					},
				},
				{
					{
						Name: "Zamorakian spear",
						Odds: 1.0 / 5_376,
					},
					{
						Name: "Godsword shard 1",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 2",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 3",
						Odds: 1.0 / 1_524,
					},
				},
				{
					{
						Name: "Zamorakian spear",
						Odds: 1.0 / 5_376,
					},
					{
						Name: "Godsword shard 1",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 2",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 3",
						Odds: 1.0 / 1_524,
					},
				},
				{
					{
						Name: "Zamorakian spear",
						Odds: 1.0 / 5_376,
					},
					{
						Name: "Godsword shard 1",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 2",
						Odds: 1.0 / 1_524,
					},
					{
						Name: "Godsword shard 3",
						Odds: 1.0 / 1_524,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet k'ril tsutsaroth",
						Amount: 1,
					},
					{
						Name:   "Steam battlestaff",
						Amount: 1,
					},
					{
						Name:   "Zamorakian spear",
						Amount: 1,
					},
					{
						Name:   "Staff of the dead",
						Amount: 1,
					},
					{
						Name:   "Zamorak hilt",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 1",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 2",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 3",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Steam battlestaff",
						Amount: 1,
					},
					{
						Name:   "Zamorakian spear",
						Amount: 1,
					},
					{
						Name:   "Staff of the dead",
						Amount: 1,
					},
					{
						Name:   "Zamorak hilt",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 1",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 2",
						Amount: 1,
					},
					{
						Name:   "Godsword shard 3",
						Amount: 1,
					},
				},
				{
					{
						Name:   "Steam battlestaff",
						Amount: 1,
					},
					{
						Name:   "Zamorakian spear",
						Amount: 1,
					},
					{
						Name:   "Staff of the dead",
						Amount: 1,
					},
					{
						Name:   "Zamorak hilt",
						Amount: 1,
					},
				},
			},
		},
	},
}
