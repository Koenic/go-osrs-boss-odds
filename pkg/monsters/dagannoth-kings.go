package monsters

import "github.com/koenic/rs-odds/pkg"

var DagannothKings = pkg.DefaultMonster{
	Name:   "Dagannoth Kings",
	KcName: "kc",
	MonsterVariation: []pkg.MonsterVariation{
		{
			VariationName: "Dagannoth Kings",
			DropTables: [][]pkg.Drop{
				{
					{
						Name: "Pet dagannoth prime",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Pet dagannoth supreme",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Pet dagannoth rex",
						Odds: 1.0 / 5000,
					},
				},
				{
					{
						Name: "Beserker ring",
						Odds: 1.0 / 128,
					},
					{
						Name: "Warrior ring",
						Odds: 1.0 / 128,
					},
					{
						Name: "Dragon axe",
						Odds: 1.0 / 128,
					},
				},
				{
					{
						Name: "Archers ring",
						Odds: 1.0 / 128,
					},
					{
						Name: "Seercull",
						Odds: 1.0 / 128,
					},
					{
						Name: "Dragon axe",
						Odds: 1.0 / 128,
					},
				},
				{
					{
						Name: "Seers ring",
						Odds: 1.0 / 128,
					},
					{
						Name: "Mud battlestaff",
						Odds: 1.0 / 128,
					},
					{
						Name: "Dragon axe",
						Odds: 1.0 / 128,
					},
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Pet dagannoth prime",
						Amount: 1,
					},
					{
						Name:   "Pet dagannoth supreme",
						Amount: 1,
					},
					{
						Name:   "Pet dagannoth rex",
						Amount: 1,
					},
					{
						Name:   "Beserker ring",
						Amount: 1,
					},
					{
						Name:   "Warrior ring",
						Amount: 1,
					},
					{
						Name:   "Archers ring",
						Amount: 1,
					},
					{
						Name:   "Seercull",
						Amount: 1,
					},
					{
						Name:   "Seers ring",
						Amount: 1,
					},
					{
						Name:   "Mud battlestaff",
						Amount: 1,
					},
					{
						Name:   "Dragon axe",
						Amount: 1,
					},
				},
			},
		},
	},
}
