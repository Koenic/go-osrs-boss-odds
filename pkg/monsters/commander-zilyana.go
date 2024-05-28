package monsters

import "github.com/koenic/rs-odds/pkg"

var CommanderZilyana = pkg.DefaultMonster{
	Name:   "Commander Zilyana (including minions)",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Commander Zilyana",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet zilyana",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Saradomin sword",
						Odds: 1.0 / 127,
					},
					{
						Name: "Saradomin's light",
						Odds: 1.0 / 254,
					},
					{
						Name: "Armadyl crossbow",
						Odds: 1.0 / 508,
					},
					{
						Name: "Saradomin hilt",
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
						Name: "Saradomin sword",
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
						Name: "Saradomin sword",
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
						Name: "Saradomin sword",
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
						Name:   "Pet zilyana",
						Amount: 1,
					},
					{
						Name:   "Saradomin sword",
						Amount: 1,
					},
					{
						Name:   "Saradomin's light",
						Amount: 1,
					},
					{
						Name:   "Armadyl crossbow",
						Amount: 1,
					},
					{
						Name:   "Saradomin hilt",
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
						Name:   "Saradomin sword",
						Amount: 1,
					},
					{
						Name:   "Saradomin's light",
						Amount: 1,
					},
					{
						Name:   "Armadyl crossbow",
						Amount: 1,
					},
					{
						Name:   "Saradomin hilt",
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
						Name:   "Saradomin sword",
						Amount: 1,
					},
					{
						Name:   "Saradomin's light",
						Amount: 1,
					},
					{
						Name:   "Armadyl crossbow",
						Amount: 1,
					},
					{
						Name:   "Saradomin hilt",
						Amount: 1,
					},
				},
			},
		},
	},
}
