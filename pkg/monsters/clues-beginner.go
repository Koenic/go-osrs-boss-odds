package monsters

import (
	"github.com/koenic/rs-odds/pkg"
)

var BeginnerClues = pkg.CluesMonster{
	Name:   "Beginner Clues",
	KcName: "Loot rolls",
	MonsterVariation: []pkg.CluesVariation{
		{
			VariationName: "Beginner Clues",
			DropTable: []pkg.Drop{
				{
					Name: "Mole slippers",
					Odds: 1.0 / 360,
				},
				{
					Name: "Frog slippers",
					Odds: 1.0 / 360,
				},
				{
					Name: "Bear feet",
					Odds: 1.0 / 360,
				},
				{
					Name: "Demon feet",
					Odds: 1.0 / 360,
				},
				{
					Name: "Jester cape",
					Odds: 1.0 / 360,
				},
				{
					Name: "Shoulder parrot",
					Odds: 1.0 / 360,
				},
				{
					Name: "Monk's robe top (t)",
					Odds: 1.0 / 360,
				},
				{
					Name: "Monk's robe (t)",
					Odds: 1.0 / 360,
				},
				{
					Name: "Amulet of defence (t)",
					Odds: 1.0 / 360,
				},
				{
					Name: "Sandwich lady hat",
					Odds: 1.0 / 360,
				},
				{
					Name: "Sandwich lady top",
					Odds: 1.0 / 360,
				},
				{
					Name: "Sandwich lady bottom",
					Odds: 1.0 / 360,
				},
				{
					Name: "Rune scimitar ornament kit (guthix)",
					Odds: 1.0 / 360,
				},
				{
					Name: "Rune scimitar ornament kit (saradomin)",
					Odds: 1.0 / 360,
				},
				{
					Name: "Rune scimitar ornament kit (zamorak)",
					Odds: 1.0 / 360,
				},
			},
			WantedDrops: [][]pkg.DropGoal{
				{
					{
						Name:   "Mole slippers",
						Amount: 1,
					},
					{
						Name:   "Frog slippers",
						Amount: 1,
					},
					{
						Name:   "Bear feet",
						Amount: 1,
					},
					{
						Name:   "Demon feet",
						Amount: 1,
					},
					{
						Name:   "Jester cape",
						Amount: 1,
					},
					{
						Name:   "Shoulder parrot",
						Amount: 1,
					},
					{
						Name:   "Monk's robe top (t)",
						Amount: 1,
					},
					{
						Name:   "Monk's robe (t)",
						Amount: 1,
					},
					{
						Name:   "Amulet of defence (t)",
						Amount: 1,
					},
					{
						Name:   "Sandwich lady hat",
						Amount: 1,
					},
					{
						Name:   "Sandwich lady top",
						Amount: 1,
					},
					{
						Name:   "Sandwich lady bottom",
						Amount: 1,
					},
					{
						Name:   "Rune scimitar ornament kit (guthix)",
						Amount: 1,
					},
					{
						Name:   "Rune scimitar ornament kit (saradomin)",
						Amount: 1,
					},
					{
						Name:   "Rune scimitar ornament kit (zamorak)",
						Amount: 1,
					},
				},
			},
		},
	},
}
