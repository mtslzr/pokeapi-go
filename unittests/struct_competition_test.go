package unittests

import (
	"testing"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

type StructCompetitionTestCase struct {
	name string
	do   func(t *testing.T) // only test if linked successfully
}

type StructCompetitionTestCases []StructCompetitionTestCase

func (c StructCompetitionTestCases) Run(t *testing.T) {
	for _, c := range c {
		t.Run(c.name, func(t *testing.T) {
			assert.NotEmpty(t, c.do)
			// do not run test, only check if compilation is successful
			// c.do(t)
		})
	}
}

var berries = StructCompetitionTestCases{
	{
		name: "Berry",
		do: func(t *testing.T) {
			d := structs.Berry{}
			assert.Empty(t, d.Firmness.Name)
			assert.Empty(t, d.Firmness.URL)
			assert.Empty(t, d.Flavors)
			for _, f := range d.Flavors {
				assert.Empty(t, f.Flavor.Name)
				assert.Empty(t, f.Flavor.URL)
				assert.Empty(t, f.Potency)
			}
			assert.Empty(t, d.GrowthTime)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Item.Name)
			assert.Empty(t, d.Item.URL)
			assert.Empty(t, d.MaxHarvest)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.NaturalGiftPower)
			assert.Empty(t, d.NaturalGiftType.Name)
			assert.Empty(t, d.NaturalGiftType.URL)
			assert.Empty(t, d.Size)
			assert.Empty(t, d.Smoothness)
			assert.Empty(t, d.SoilDryness)
		},
	},
	{
		name: "BerryFirmness",
		do: func(t *testing.T) {
			d := structs.BerryFirmness{}
			assert.Empty(t, d.Berries)
			for _, b := range d.Berries {
				assert.Empty(t, b.Name)
				assert.Empty(t, b.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
		},
	},
	{
		name: "BerryFlavor",
		do: func(t *testing.T) {
			d := structs.BerryFlavor{}
			assert.Empty(t, d.Berries)
			for _, b := range d.Berries {
				assert.Empty(t, b.Berry.Name)
				assert.Empty(t, b.Berry.URL)
				assert.Empty(t, b.Potency)
			}
			assert.Empty(t, d.ContestType)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
		},
	},
}

var contests = StructCompetitionTestCases{
	{
		name: "ContestType",
		do: func(t *testing.T) {
			d := structs.ContestType{}
			assert.Empty(t, d.BerryFlavor.Name)
			assert.Empty(t, d.BerryFlavor.URL)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Color)
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
		},
	},
	{
		name: "ContestEffect",
		do: func(t *testing.T) {
			d := structs.ContestEffect{}
			assert.Empty(t, d.Appeal)
			assert.Empty(t, d.EffectEntries)
			for _, e := range d.EffectEntries {
				assert.Empty(t, e.Effect)
				assert.Empty(t, e.Language.Name)
				assert.Empty(t, e.Language.URL)
			}
			assert.Empty(t, d.FlavorTextEntries)
			for _, f := range d.FlavorTextEntries {
				assert.Empty(t, f.FlavorText)
				assert.Empty(t, f.Language.Name)
				assert.Empty(t, f.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Jam)
		},
	},
	{
		name: "SuperContestEffect",
		do: func(t *testing.T) {
			d := structs.SuperContestEffect{}
			assert.Empty(t, d.Appeal)
			assert.Empty(t, d.FlavorTextEntries)
			for _, f := range d.FlavorTextEntries {
				assert.Empty(t, f.FlavorText)
				assert.Empty(t, f.Language.Name)
				assert.Empty(t, f.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Moves)
			for _, m := range d.Moves {
				assert.Empty(t, m.Name)
				assert.Empty(t, m.URL)
			}
		},
	},
}

var encounters = StructCompetitionTestCases{
	{
		name: "EncounterMethod",
		do: func(t *testing.T) {
			d := structs.EncounterMethod{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Order)
		},
	},
	{
		name: "EncounterCondition",
		do: func(t *testing.T) {
			d := structs.EncounterCondition{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Values)
			for _, v := range d.Values {
				assert.Empty(t, v.Name)
				assert.Empty(t, v.URL)
			}
		},
	},
	{
		name: "EncounterConditionValue",
		do: func(t *testing.T) {
			d := structs.EncounterConditionValue{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			// assert.Empty(t, d.Values)
			// for _, v := range d.Values {
			// 	assert.Empty(t, v.Name)
			// 	assert.Empty(t, v.URL)
			// }
		},
	},
}

var evolution = StructCompetitionTestCases{
	{
		name: "EvolutionChain",
		do: func(t *testing.T) {
			d := structs.EvolutionChain{}
			assert.Empty(t, d.BabyTriggerItem)
			assert.Empty(t, d.Chain.EvolutionDetails)
			for _, e := range d.Chain.EvolutionDetails {
				assert.Empty(t, e)
			}
			assert.Empty(t, d.Chain.EvolvesTo)
			for _, e := range d.Chain.EvolvesTo {
				assert.Empty(t, e.EvolutionDetails)
				for _, e := range e.EvolutionDetails {
					assert.Empty(t, e.Gender)
					assert.Empty(t, e.HeldItem)
					assert.Empty(t, e.Item)
					assert.Empty(t, e.KnownMove)
					assert.Empty(t, e.KnownMoveType)
					assert.Empty(t, e.Location)
					assert.Empty(t, e.MinAffection)
					assert.Empty(t, e.MinBeauty)
					assert.Empty(t, e.MinHappiness)
					assert.Empty(t, e.MinLevel)
					assert.Empty(t, e.NeedsOverworldRain)
					assert.Empty(t, e.PartySpecies)
					assert.Empty(t, e.PartyType)
					assert.Empty(t, e.RelativePhysicalStats)
					assert.Empty(t, e.TimeOfDay)
					assert.Empty(t, e.TradeSpecies)
					assert.Empty(t, e.Trigger.Name)
					assert.Empty(t, e.Trigger.URL)
					assert.Empty(t, e.TurnUpsideDown)
				}
				assert.Empty(t, e.EvolvesTo)
				for _, e := range e.EvolvesTo {
					assert.Empty(t, e.EvolutionDetails)
					for _, e := range e.EvolutionDetails {
						assert.Empty(t, e.Gender)
						assert.Empty(t, e.HeldItem)
						assert.Empty(t, e.Item)
						assert.Empty(t, e.KnownMove)
						assert.Empty(t, e.KnownMoveType)
						assert.Empty(t, e.Location)
						assert.Empty(t, e.MinAffection)
						assert.Empty(t, e.MinBeauty)
						assert.Empty(t, e.MinHappiness)
						assert.Empty(t, e.MinLevel)
						assert.Empty(t, e.NeedsOverworldRain)
						assert.Empty(t, e.PartySpecies)
						assert.Empty(t, e.PartyType)
						assert.Empty(t, e.RelativePhysicalStats)
						assert.Empty(t, e.TimeOfDay)
						assert.Empty(t, e.TradeSpecies)
						assert.Empty(t, e.Trigger.Name)
						assert.Empty(t, e.Trigger.URL)
						assert.Empty(t, e.TurnUpsideDown)
					}
					assert.Empty(t, e.EvolvesTo)
					assert.Empty(t, e.IsBaby)
					assert.Empty(t, e.Species.Name)
					assert.Empty(t, e.Species.URL)
				}
				assert.Empty(t, e.IsBaby)
				assert.Empty(t, e.Species.Name)
				assert.Empty(t, e.Species.URL)
			}
			assert.Empty(t, d.Chain.IsBaby)
			assert.Empty(t, d.Chain.Species.Name)
			assert.Empty(t, d.Chain.Species.URL)
			assert.Empty(t, d.ID)
		},
	},
	{
		name: "EvolutionTrigger",
		do: func(t *testing.T) {
			d := structs.EvolutionTrigger{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonSpecies)
			for _, p := range d.PokemonSpecies {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
		},
	},
}

var games = StructCompetitionTestCases{
	{
		name: "Generation",
		do: func(t *testing.T) {
			d := structs.Generation{}
			assert.Empty(t, d.Abilities)
			for _, a := range d.Abilities {
				assert.Empty(t, a)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.MainRegion.Name)
			assert.Empty(t, d.MainRegion.URL)
			assert.Empty(t, d.Moves)
			for _, m := range d.Moves {
				assert.Empty(t, m.Name)
				assert.Empty(t, m.URL)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonSpecies)
			for _, p := range d.PokemonSpecies {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
			assert.Empty(t, d.Types)
			for _, v := range d.Types {
				assert.Empty(t, v.Name)
				assert.Empty(t, v.URL)
			}
			assert.Empty(t, d.VersionGroups)
			for _, v := range d.VersionGroups {
				assert.Empty(t, v.Name)
				assert.Empty(t, v.URL)
			}
		},
	},
	{
		name: "Pokedex",
		do: func(t *testing.T) {
			d := structs.Pokedex{}
			assert.Empty(t, d.Descriptions)
			for _, d := range d.Descriptions {
				assert.Empty(t, d.Description)
				assert.Empty(t, d.Language.Name)
				assert.Empty(t, d.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.IsMainSeries)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonEntries)
			for _, p := range d.PokemonEntries {
				assert.Empty(t, p.EntryNumber)
				assert.Empty(t, p.PokemonSpecies.Name)
				assert.Empty(t, p.PokemonSpecies.URL)
			}
			assert.Empty(t, d.Region)
			assert.Empty(t, d.VersionGroups)
		},
	},
	{
		name: "Version",
		do: func(t *testing.T) {
			d := structs.Version{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.VersionGroup.Name)
			assert.Empty(t, d.VersionGroup.URL)
		},
	},
	{
		name: "VersionGroup",
		do: func(t *testing.T) {
			d := structs.VersionGroup{}
			assert.Empty(t, d.Generation.Name)
			assert.Empty(t, d.Generation.URL)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.MoveLearnMethods)
			for _, m := range d.MoveLearnMethods {
				assert.Empty(t, m.Name)
				assert.Empty(t, m.URL)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Order)
			assert.Empty(t, d.Pokedexes)
			for _, p := range d.Pokedexes {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
			assert.Empty(t, d.Regions)
			for _, r := range d.Regions {
				assert.Empty(t, r.Name)
				assert.Empty(t, r.URL)
			}
			assert.Empty(t, d.Versions)
			for _, v := range d.Versions {
				assert.Empty(t, v.Name)
				assert.Empty(t, v.URL)
			}
		},
	},
}

var items = StructCompetitionTestCases{
	{
		name: "Item",
		do: func(t *testing.T) {
			d := structs.Item{}
			assert.Empty(t, d.Attributes)
			for _, a := range d.Attributes {
				assert.Empty(t, a.Name)
				assert.Empty(t, a.URL)
			}
			assert.Empty(t, d.BabyTriggerFor)
			assert.Empty(t, d.Category.Name)
			assert.Empty(t, d.Category.URL)
			assert.Empty(t, d.Cost)
			assert.Empty(t, d.EffectEntries)
			for _, e := range d.EffectEntries {
				assert.Empty(t, e.Effect)
				assert.Empty(t, e.Language.Name)
				assert.Empty(t, e.Language.URL)
				assert.Empty(t, e.ShortEffect)
			}
			assert.Empty(t, d.FlavorTextEntries)
			for _, f := range d.FlavorTextEntries {
				assert.Empty(t, f.Language.Name)
				assert.Empty(t, f.Language.URL)
				assert.Empty(t, f.Text)
				assert.Empty(t, f.VersionGroup.Name)
				assert.Empty(t, f.VersionGroup.URL)
			}
			assert.Empty(t, d.FlingEffect)
			assert.Empty(t, d.FlingPower)
			assert.Empty(t, d.GameIndices)
			for _, g := range d.GameIndices {
				assert.Empty(t, g.GameIndex)
				assert.Empty(t, g.Generation.Name)
				assert.Empty(t, g.Generation.URL)
			}
			assert.Empty(t, d.HeldByPokemon)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Machines)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Sprites.Default)
		},
	},
	{
		name: "ItemAttribute",
		do: func(t *testing.T) {
			d := structs.ItemAttribute{}
			assert.Empty(t, d.Descriptions)
			for _, b := range d.Descriptions {
				assert.Empty(t, b.Description)
				assert.Empty(t, b.Language.Name)
				assert.Empty(t, b.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Items)
			for _, i := range d.Items {
				assert.Empty(t, i.Name)
				assert.Empty(t, i.URL)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
		},
	},
	{
		name: "ItemCategory",
		do: func(t *testing.T) {
			d := structs.ItemCategory{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Items)
			for _, i := range d.Items {
				assert.Empty(t, i.Name)
				assert.Empty(t, i.URL)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Pocket.Name)
			assert.Empty(t, d.Pocket.URL)
		},
	},
	{
		name: "ItemFlingEffect",
		do: func(t *testing.T) {
			d := structs.ItemFlingEffect{}
			assert.Empty(t, d.EffectEntries)
			for _, e := range d.EffectEntries {
				assert.Empty(t, e.Effect)
				assert.Empty(t, e.Language.Name)
				assert.Empty(t, e.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Items)
			for _, i := range d.Items {
				assert.Empty(t, i.Name)
				assert.Empty(t, i.URL)
			}
			assert.Empty(t, d.Name)
		},
	},
	{
		name: "ItemPocket",
		do: func(t *testing.T) {
			d := structs.ItemPocket{}
			assert.Empty(t, d.Categories)
			for _, c := range d.Categories {
				assert.Empty(t, c.Name)
				assert.Empty(t, c.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
		},
	},
}

var location = StructCompetitionTestCases{
	{
		name: "Location",
		do: func(t *testing.T) {
			d := structs.Location{}
			assert.Empty(t, d.Areas)
			for _, a := range d.Areas {
				assert.Empty(t, a.Name)
				assert.Empty(t, a.URL)
			}
			assert.Empty(t, d.GameIndices)
			for _, g := range d.GameIndices {
				assert.Empty(t, g.GameIndex)
				assert.Empty(t, g.Generation.Name)
				assert.Empty(t, g.Generation.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Region.Name)
			assert.Empty(t, d.Region.URL)
		},
	},
	{
		name: "LocationArea",
		do: func(t *testing.T) {
			d := structs.LocationArea{}
			assert.Empty(t, d.EncounterMethodRates)
			for _, e := range d.EncounterMethodRates {
				assert.Empty(t, e.EncounterMethod.Name)
				assert.Empty(t, e.EncounterMethod.URL)
				assert.Empty(t, e.VersionDetails)
				for _, v := range e.VersionDetails {
					assert.Empty(t, v.Rate)
					assert.Empty(t, v.Version.Name)
					assert.Empty(t, v.Version.URL)
				}
			}
			assert.Empty(t, d.GameIndex)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Location.Name)
			assert.Empty(t, d.Location.URL)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonEncounters)
			for _, p := range d.PokemonEncounters {
				assert.Empty(t, p.Pokemon.Name)
				assert.Empty(t, p.Pokemon.URL)
				assert.Empty(t, p.VersionDetails)
				for _, v := range p.VersionDetails {
					assert.Empty(t, v.EncounterDetails)
					for _, e := range v.EncounterDetails {
						assert.Empty(t, e.Chance)
						assert.Empty(t, e.ConditionValues)
						for _, c := range e.ConditionValues {
							assert.Empty(t, c)
						}
						assert.Empty(t, e.MaxLevel)
						assert.Empty(t, e.Method.Name)
						assert.Empty(t, e.Method.URL)
						assert.Empty(t, e.MinLevel)
					}
					assert.Empty(t, v.MaxChance)
					assert.Empty(t, v.Version.Name)
					assert.Empty(t, v.Version.URL)
				}
			}
		},
	},
	{
		name: "PalParkArea",
		do: func(t *testing.T) {
			d := structs.PalParkArea{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonEncounters)
			for _, p := range d.PokemonEncounters {
				assert.Empty(t, p.BaseScore)
				assert.Empty(t, p.PokemonSpecies.Name)
				assert.Empty(t, p.PokemonSpecies.URL)
				assert.Empty(t, p.Rate)
			}
		},
	},
	{
		name: "Region",
		do: func(t *testing.T) {
			d := structs.Region{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Locations)
			for _, l := range d.Locations {
				assert.Empty(t, l.Name)
				assert.Empty(t, l.URL)
			}
			assert.Empty(t, d.MainGeneration.Name)
			assert.Empty(t, d.MainGeneration.URL)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Pokedexes)
			for _, p := range d.Pokedexes {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
			assert.Empty(t, d.VersionGroups)
			for _, v := range d.VersionGroups {
				assert.Empty(t, v.Name)
				assert.Empty(t, v.URL)
			}
		},
	},
}

var machines = StructCompetitionTestCases{
	{
		name: "Machine",
		do: func(t *testing.T) {
			d := structs.Machine{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Item.Name)
			assert.Empty(t, d.Item.URL)
			assert.Empty(t, d.Move.Name)
			assert.Empty(t, d.Move.URL)
			assert.Empty(t, d.VersionGroup.Name)
			assert.Empty(t, d.VersionGroup.URL)
		},
	},
}

var moves = StructCompetitionTestCases{
	{
		name: "Move",
		do: func(t *testing.T) {
			d := structs.Move{}
			assert.Empty(t, d.Accuracy)
			assert.Empty(t, d.ContestCombos.Normal.UseAfter)
			assert.Empty(t, d.ContestCombos.Normal.UseBefore)
			for _, b := range d.ContestCombos.Normal.UseBefore {
				assert.Empty(t, b.Name)
				assert.Empty(t, b.URL)
			}
			assert.Empty(t, d.ContestCombos.Super.UseAfter)
			assert.Empty(t, d.ContestCombos.Super.UseBefore)
			assert.Empty(t, d.ContestEffect.URL)
			assert.Empty(t, d.ContestType.Name)
			assert.Empty(t, d.ContestType.URL)
			assert.Empty(t, d.DamageClass.Name)
			assert.Empty(t, d.DamageClass.URL)
			assert.Empty(t, d.EffectChance)
			assert.Empty(t, d.EffectChanges)
			for _, e := range d.EffectChanges {
				assert.Empty(t, e)
			}
			assert.Empty(t, d.EffectEntries)
			for _, e := range d.EffectEntries {
				assert.Empty(t, e.Effect)
				assert.Empty(t, e.Language.Name)
				assert.Empty(t, e.Language.URL)
				assert.Empty(t, e.ShortEffect)
			}
			assert.Empty(t, d.FlavorTextEntries)
			for _, f := range d.FlavorTextEntries {
				assert.Empty(t, f.FlavorText)
				assert.Empty(t, f.Language.Name)
				assert.Empty(t, f.Language.URL)
				assert.Empty(t, f.VersionGroup.Name)
				assert.Empty(t, f.VersionGroup.URL)
			}
			assert.Empty(t, d.Generation.Name)
			assert.Empty(t, d.Generation.URL)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Machines)
			for _, m := range d.Machines {
				assert.Empty(t, m)
			}
			assert.Empty(t, d.Meta.Ailment.Name)
			assert.Empty(t, d.Meta.Ailment.URL)
			assert.Empty(t, d.Meta.AilmentChance)
			assert.Empty(t, d.Meta.Category.Name)
			assert.Empty(t, d.Meta.Category.URL)
			assert.Empty(t, d.Meta.CritRate)
			assert.Empty(t, d.Meta.Drain)
			assert.Empty(t, d.Meta.FlinchChance)
			assert.Empty(t, d.Meta.Healing)
			assert.Empty(t, d.Meta.MaxHits)
			assert.Empty(t, d.Meta.MaxTurns)
			assert.Empty(t, d.Meta.MinHits)
			assert.Empty(t, d.Meta.MinTurns)
			assert.Empty(t, d.Meta.StatChance)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PastValues)
			for _, p := range d.PastValues {
				assert.Empty(t, p)
			}
			assert.Empty(t, d.Power)
			assert.Empty(t, d.Pp)
			assert.Empty(t, d.Priority)
			assert.Empty(t, d.StatChanges)
			for _, s := range d.StatChanges {
				assert.Empty(t, s)
			}
			assert.Empty(t, d.SuperContestEffect.URL)
			assert.Empty(t, d.Target.Name)
			assert.Empty(t, d.Target.URL)
			assert.Empty(t, d.Type.Name)
			assert.Empty(t, d.Type.URL)
		},
	},
	{
		name: "MoveAilment",
		do: func(t *testing.T) {
			d := structs.MoveAilment{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Moves)
			for _, m := range d.Moves {
				assert.Empty(t, m.Name)
				assert.Empty(t, m.URL)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
		},
	},
	{
		name: "MoveBattleStyle",
		do: func(t *testing.T) {
			d := structs.MoveBattleStyle{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Name)
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
			}
		},
	},
	{
		name: "MoveCategory",
		do: func(t *testing.T) {
			d := structs.MoveCategory{}
			assert.Empty(t, d.Descriptions)
			for _, d := range d.Descriptions {
				assert.Empty(t, d.Description)
				assert.Empty(t, d.Language.Name)
				assert.Empty(t, d.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Moves)
			for _, m := range d.Moves {
				assert.Empty(t, m.Name)
				assert.Empty(t, m.URL)
			}
			assert.Empty(t, d.Name)
		},
	},
	{
		name: "MoveDamageClass",
		do: func(t *testing.T) {
			d := structs.MoveDamageClass{}
			assert.Empty(t, d.Descriptions)
			for _, d := range d.Descriptions {
				assert.Empty(t, d.Description)
				assert.Empty(t, d.Language.Name)
				assert.Empty(t, d.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Moves)
			for _, m := range d.Moves {
				assert.Empty(t, m.Name)
				assert.Empty(t, m.URL)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Name)
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
			}
		},
	},
	{
		name: "MoveLearnMethod",
		do: func(t *testing.T) {
			d := structs.MoveLearnMethod{}
			assert.Empty(t, d.Descriptions)
			for _, d := range d.Descriptions {
				assert.Empty(t, d.Description)
				assert.Empty(t, d.Language.Name)
				assert.Empty(t, d.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Name)
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
			}
			assert.Empty(t, d.VersionGroups)
			for _, v := range d.VersionGroups {
				assert.Empty(t, v.Name)
				assert.Empty(t, v.URL)
			}
		},
	},
	{
		name: "MoveTarget",
		do: func(t *testing.T) {
			d := structs.MoveTarget{}
			assert.Empty(t, d.Descriptions)
			for _, d := range d.Descriptions {
				assert.Empty(t, d.Description)
				assert.Empty(t, d.Language.Name)
				assert.Empty(t, d.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Moves)
			for _, m := range d.Moves {
				assert.Empty(t, m.Name)
				assert.Empty(t, m.URL)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Name)
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
			}
		},
	},
}

var pokemon = StructCompetitionTestCases{
	{
		name: "Ability",
		do: func(t *testing.T) {
			d := structs.Ability{}
			assert.Empty(t, d.EffectChanges)
			for _, e := range d.EffectChanges {
				assert.Empty(t, e.EffectEntries)
				for _, e := range e.EffectEntries {
					assert.Empty(t, e.Effect)
					assert.Empty(t, e.Language.Name)
					assert.Empty(t, e.Language.URL)
				}
				assert.Empty(t, e.VersionGroup.Name)
				assert.Empty(t, e.VersionGroup.URL)
			}
			assert.Empty(t, d.EffectEntries)
			for _, e := range d.EffectEntries {
				assert.Empty(t, e.Effect)
				assert.Empty(t, e.Language.Name)
				assert.Empty(t, e.Language.URL)
				assert.Empty(t, e.ShortEffect)
			}
			assert.Empty(t, d.FlavorTextEntries)
			for _, f := range d.FlavorTextEntries {
				assert.Empty(t, f.FlavorText)
				assert.Empty(t, f.Language.Name)
				assert.Empty(t, f.Language.URL)
				assert.Empty(t, f.VersionGroup.Name)
				assert.Empty(t, f.VersionGroup.URL)
			}
			assert.Empty(t, d.Generation.Name)
			assert.Empty(t, d.Generation.URL)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.IsMainSeries)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Pokemon)
			for _, p := range d.Pokemon {
				assert.Empty(t, p.IsHidden)
				assert.Empty(t, p.Pokemon.Name)
				assert.Empty(t, p.Pokemon.URL)
				assert.Empty(t, p.Slot)
			}
		},
	},
	{
		name: "Characteristic",
		do: func(t *testing.T) {
			d := structs.Characteristic{}
			assert.Empty(t, d.Descriptions)
			for _, d := range d.Descriptions {
				assert.Empty(t, d.Description)
				assert.Empty(t, d.Language.Name)
				assert.Empty(t, d.Language.URL)
			}
			assert.Empty(t, d.GeneModulo)
			assert.Empty(t, d.HighestStat.Name)
			assert.Empty(t, d.HighestStat.URL)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.PossibleValues)
			for _, p := range d.PossibleValues {
				assert.Empty(t, p)
			}
		},
	},
	{
		name: "EggGroup",
		do: func(t *testing.T) {
			d := structs.EggGroup{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonSpecies)
			for _, p := range d.PokemonSpecies {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
		},
	},
	{
		name: "Gender",
		do: func(t *testing.T) {
			d := structs.Gender{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.PokemonSpeciesDetails)
			for _, p := range d.PokemonSpeciesDetails {
				assert.Empty(t, p.PokemonSpecies.Name)
				assert.Empty(t, p.PokemonSpecies.URL)
				assert.Empty(t, p.Rate)
			}
			assert.Empty(t, d.RequiredForEvolution)
			for _, r := range d.RequiredForEvolution {
				assert.Empty(t, r.Name)
				assert.Empty(t, r.URL)
			}
		},
	},
	{
		name: "GrowthRate",
		do: func(t *testing.T) {
			d := structs.GrowthRate{}
			assert.Empty(t, d.Descriptions)
			for _, d := range d.Descriptions {
				assert.Empty(t, d.Description)
				assert.Empty(t, d.Language.Name)
				assert.Empty(t, d.Language.URL)
			}
			assert.Empty(t, d.Formula)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Levels)
			for _, l := range d.Levels {
				assert.Empty(t, l.Experience)
				assert.Empty(t, l.Level)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.PokemonSpecies)
			for _, p := range d.PokemonSpecies {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
		},
	},
	{
		name: "Nature",
		do: func(t *testing.T) {
			d := structs.Nature{}
			assert.Empty(t, d.DecreasedStat)
			assert.Empty(t, d.HatesFlavor)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.IncreasedStat)
			assert.Empty(t, d.LikesFlavor)
			assert.Empty(t, d.MoveBattleStylePreferences)
			for _, m := range d.MoveBattleStylePreferences {
				assert.Empty(t, m.LowHpPreference)
				assert.Empty(t, m.HighHpPreference)
				assert.Empty(t, m.MoveBattleStyle.Name)
				assert.Empty(t, m.MoveBattleStyle.URL)
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokeathlonStatChanges)
			for _, p := range d.PokeathlonStatChanges {
				assert.Empty(t, p.MaxChange)
				assert.Empty(t, p.PokeathlonStat.Name)
				assert.Empty(t, p.PokeathlonStat.URL)
			}
		},
	},
	{
		name: "PokeathlonStat",
		do: func(t *testing.T) {
			d := structs.PokeathlonStat{}
			assert.Empty(t, d.AffectingNatures.Decrease)
			for _, d := range d.AffectingNatures.Decrease {
				assert.Empty(t, d.MaxChange)
				assert.Empty(t, d.Nature.Name)
				assert.Empty(t, d.Nature.URL)
			}
			assert.Empty(t, d.AffectingNatures.Increase)
			for _, d := range d.AffectingNatures.Increase {
				assert.Empty(t, d.MaxChange)
				assert.Empty(t, d.Nature.Name)
				assert.Empty(t, d.Nature.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
		},
	},
	{
		name: "Pokemon",
		do: func(t *testing.T) {
			d := structs.Pokemon{}
			assert.Empty(t, d.Abilities)
			for _, a := range d.Abilities {
				assert.Empty(t, a.Ability.Name)
				assert.Empty(t, a.Ability.URL)
				assert.Empty(t, a.IsHidden)
				assert.Empty(t, a.Slot)
			}
			assert.Empty(t, d.BaseExperience)
			assert.Empty(t, d.Forms)
			for _, f := range d.Forms {
				assert.Empty(t, f.Name)
				assert.Empty(t, f.URL)
			}
			assert.Empty(t, d.GameIndices)
			for _, g := range d.GameIndices {
				assert.Empty(t, g.GameIndex)
				assert.Empty(t, g.Version.Name)
				assert.Empty(t, g.Version.URL)
			}
			assert.Empty(t, d.Height)
			assert.Empty(t, d.HeldItems)
			for _, h := range d.HeldItems {
				assert.Empty(t, h)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.IsDefault)
			assert.Empty(t, d.LocationAreaEncounters)
			assert.Empty(t, d.Moves)
			for _, m := range d.Moves {
				assert.Empty(t, m.Move.Name)
				assert.Empty(t, m.Move.URL)
				assert.Empty(t, m.VersionGroupDetails)
				for _, v := range m.VersionGroupDetails {
					assert.Empty(t, v.LevelLearnedAt)
					assert.Empty(t, v.MoveLearnMethod.Name)
					assert.Empty(t, v.MoveLearnMethod.URL)
					assert.Empty(t, v.VersionGroup.Name)
					assert.Empty(t, v.VersionGroup.URL)
				}
			}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Order)
			assert.Empty(t, d.Species.Name)
			assert.Empty(t, d.Species.URL)
			assert.Empty(t, d.Sprites.BackDefault)
			assert.Empty(t, d.Sprites.BackFemale)
			assert.Empty(t, d.Sprites.BackShiny)
			assert.Empty(t, d.Sprites.BackShinyFemale)
			assert.Empty(t, d.Sprites.FrontDefault)
			assert.Empty(t, d.Sprites.FrontFemale)
			assert.Empty(t, d.Sprites.FrontShiny)
			assert.Empty(t, d.Sprites.FrontShinyFemale)
			assert.Empty(t, d.Stats)
			for _, s := range d.Stats {
				assert.Empty(t, s.BaseStat)
				assert.Empty(t, s.Effort)
				assert.Empty(t, s.Stat.Name)
				assert.Empty(t, s.Stat.URL)
			}
			assert.Empty(t, d.Types)
			for _, v := range d.Types {
				assert.Empty(t, v.Slot)
				assert.Empty(t, v.Type.Name)
				assert.Empty(t, v.Type.URL)
			}
			assert.Empty(t, d.Weight)
		},
	},
	{
		name: "PokemonColor",
		do: func(t *testing.T) {
			d := structs.PokemonColor{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonSpecies)
			for _, p := range d.PokemonSpecies {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
		},
	},
	{
		name: "PokemonForm",
		do: func(t *testing.T) {
			d := structs.PokemonForm{}
			assert.Empty(t, d.FormName)
			assert.Empty(t, d.FormNames)
			for _, n := range d.FormNames {
				assert.Empty(t, n)
			}
			assert.Empty(t, d.FormOrder)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.IsBattleOnly)
			assert.Empty(t, d.IsDefault)
			assert.Empty(t, d.IsMega)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n)
			}
			assert.Empty(t, d.Order)
			assert.Empty(t, d.Pokemon.Name)
			assert.Empty(t, d.Pokemon.URL)
			assert.Empty(t, d.Sprites.BackDefault)
			assert.Empty(t, d.Sprites.BackShiny)
			assert.Empty(t, d.Sprites.FrontDefault)
			assert.Empty(t, d.Sprites.FrontShiny)
			assert.Empty(t, d.VersionGroup.Name)
			assert.Empty(t, d.VersionGroup.URL)
		},
	},
	{
		name: "PokemonHabitat",
		do: func(t *testing.T) {
			d := structs.PokemonHabitat{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonSpecies)
			for _, p := range d.PokemonSpecies {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
		},
	},
	{
		name: "PokemonShape",
		do: func(t *testing.T) {
			d := structs.PokemonShape{}
			assert.Empty(t, d.AwesomeNames)
			for _, a := range d.AwesomeNames {
				assert.Empty(t, a.AwesomeName)
				assert.Empty(t, a.Language.Name)
				assert.Empty(t, a.Language.URL)
			}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.PokemonSpecies)
			for _, p := range d.PokemonSpecies {
				assert.Empty(t, p.Name)
				assert.Empty(t, p.URL)
			}
		},
	},
	{
		name: "PokemonSpecies",
		do: func(t *testing.T) {
			d := structs.PokemonSpecies{}
			assert.Empty(t, d.BaseHappiness)
			assert.Empty(t, d.CaptureRate)
			assert.Empty(t, d.Color.Name)
			assert.Empty(t, d.Color.URL)
			assert.Empty(t, d.EggGroups)
			for _, e := range d.EggGroups {
				assert.Empty(t, e.Name)
				assert.Empty(t, e.URL)
			}
			assert.Empty(t, d.EvolutionChain.URL)
			assert.Empty(t, d.EvolvesFromSpecies)
			assert.Empty(t, d.FlavorTextEntries)
			for _, f := range d.FlavorTextEntries {
				assert.Empty(t, f.FlavorText)
				assert.Empty(t, f.Language.Name)
				assert.Empty(t, f.Language.URL)
				assert.Empty(t, f.Version.Name)
				assert.Empty(t, f.Version.URL)
			}
			assert.Empty(t, d.FormDescriptions)
			for _, f := range d.FormDescriptions {
				assert.Empty(t, f)
			}
			assert.Empty(t, d.FormsSwitchable)
			assert.Empty(t, d.GenderRate)
			assert.Empty(t, d.Genera)
			for _, g := range d.Genera {
				assert.Empty(t, g.Genus)
				assert.Empty(t, g.Language.Name)
				assert.Empty(t, g.Language.URL)
			}
			assert.Empty(t, d.Generation.Name)
			assert.Empty(t, d.Generation.URL)
			assert.Empty(t, d.GrowthRate.Name)
			assert.Empty(t, d.GrowthRate.URL)
			assert.Empty(t, d.Habitat.Name)
			assert.Empty(t, d.Habitat.URL)
			assert.Empty(t, d.HasGenderDifferences)
			assert.Empty(t, d.HatchCounter)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.IsBaby)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Order)
			assert.Empty(t, d.PalParkEncounters)
			for _, p := range d.PalParkEncounters {
				assert.Empty(t, p.Area.Name)
				assert.Empty(t, p.Area.URL)
				assert.Empty(t, p.BaseScore)
				assert.Empty(t, p.Rate)
			}
			assert.Empty(t, d.PokedexNumbers)
			for _, p := range d.PokedexNumbers {
				assert.Empty(t, p.EntryNumber)
				assert.Empty(t, p.Pokedex.Name)
				assert.Empty(t, p.Pokedex.URL)
			}
			assert.Empty(t, d.Shape.Name)
			assert.Empty(t, d.Shape.URL)
			assert.Empty(t, d.Varieties)
			for _, v := range d.Varieties {
				assert.Empty(t, v.IsDefault)
				assert.Empty(t, v.Pokemon.Name)
				assert.Empty(t, v.Pokemon.URL)
			}
		},
	},
	{
		name: "Stat",
		do: func(t *testing.T) {
			d := structs.Stat{}
			assert.Empty(t, d.AffectingMoves.Decrease)
			for _, a := range d.AffectingMoves.Decrease {
				assert.Empty(t, a)
			}
			assert.Empty(t, d.AffectingMoves.Increase)
			for _, a := range d.AffectingMoves.Increase {
				assert.Empty(t, a.Change)
				assert.Empty(t, a.Move.Name)
				assert.Empty(t, a.Move.URL)
			}
			assert.Empty(t, d.AffectingNatures.Decrease)
			for _, a := range d.AffectingNatures.Decrease {
				assert.Empty(t, a)
			}
			assert.Empty(t, d.AffectingNatures.Increase)
			for _, a := range d.AffectingNatures.Increase {
				assert.Empty(t, a)
			}
			assert.Empty(t, d.Characteristics)
			for _, c := range d.Characteristics {
				assert.Empty(t, c.URL)
			}
			assert.Empty(t, d.GameIndex)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.IsBattleOnly)
			assert.Empty(t, d.MoveDamageClass)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
		},
	},
	{
		name: "Type",
		do: func(t *testing.T) {
			d := structs.Type{}
			assert.Empty(t, d.DamageRelations.DoubleDamageFrom)
			for _, d := range d.DamageRelations.DoubleDamageFrom {
				assert.Empty(t, d.Name)
				assert.Empty(t, d.URL)
			}
			assert.Empty(t, d.DamageRelations.DoubleDamageTo)
			for _, d := range d.DamageRelations.DoubleDamageTo {
				assert.Empty(t, d)
			}
			assert.Empty(t, d.DamageRelations.HalfDamageFrom)
			for _, d := range d.DamageRelations.HalfDamageFrom {
				assert.Empty(t, d)
			}
			assert.Empty(t, d.DamageRelations.HalfDamageTo)
			for _, d := range d.DamageRelations.HalfDamageTo {
				assert.Empty(t, d.Name)
				assert.Empty(t, d.URL)
			}
			assert.Empty(t, d.DamageRelations.NoDamageFrom)
			for _, d := range d.DamageRelations.NoDamageFrom {
				assert.Empty(t, d.Name)
				assert.Empty(t, d.URL)
			}
			assert.Empty(t, d.DamageRelations.NoDamageTo)
			for _, d := range d.DamageRelations.NoDamageTo {
				assert.Empty(t, d.Name)
				assert.Empty(t, d.URL)
			}
			assert.Empty(t, d.GameIndices)
			for _, g := range d.GameIndices {
				assert.Empty(t, g.GameIndex)
				assert.Empty(t, g.Generation.Name)
				assert.Empty(t, g.Generation.URL)
			}
			assert.Empty(t, d.Generation.Name)
			assert.Empty(t, d.Generation.URL)
			assert.Empty(t, d.ID)
			assert.Empty(t, d.MoveDamageClass.Name)
			assert.Empty(t, d.MoveDamageClass.URL)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Pokemon)
			for _, p := range d.Pokemon {
				assert.Empty(t, p.Pokemon.Name)
				assert.Empty(t, p.Pokemon.URL)
				assert.Empty(t, p.Slot)
			}
		},
	},
}

var resource = StructCompetitionTestCases{
	// {
	// 	name: "Resource",
	// 	do: func(t *testing.T) {
	// 		d := structs.Resource{}
	// 		assert.Empty(t, d.Count)
	// 		assert.Empty(t, d.Next)
	// 		assert.Empty(t, d.Previous)
	// 		assert.Empty(t, d.Results)
	// 		for _, r := range d.Results {
	// 			assert.Empty(t, r.Name)
	// 			assert.Empty(t, r.URL)
	// 		}
	// 	},
	// },
	// {
	// 	name: "Result",
	// 	do: func(t *testing.T) {
	// 		d := structs.Result{}
	// 		assert.Empty(t, d.Name)
	// 		assert.Empty(t, d.URL)
	// 	},
	// },
	{
		name: "NamedAPIResource",
		do: func(t *testing.T) {
			d := structs.NamedApiResource{}
			assert.Empty(t, d.Name)
			assert.Empty(t, d.URL)
		},
	},
	{
		name: "NamedAPIResourceList",
		do: func(t *testing.T) {
			d := structs.NamedApiResourceList{}
			assert.Empty(t, d.Count)
			assert.Empty(t, d.Next)
			assert.Empty(t, d.Previous)
			assert.Empty(t, d.Results)
			for _, r := range d.Results {
				assert.Empty(t, r.Name)
				assert.Empty(t, r.URL)
			}
		},
	},
}

var utility = StructCompetitionTestCases{
	{
		name: "Language",
		do: func(t *testing.T) {
			d := structs.Language{}
			assert.Empty(t, d.ID)
			assert.Empty(t, d.Iso639)
			assert.Empty(t, d.Iso3166)
			assert.Empty(t, d.Name)
			assert.Empty(t, d.Names)
			for _, n := range d.Names {
				assert.Empty(t, n.Language.Name)
				assert.Empty(t, n.Language.URL)
				assert.Empty(t, n.Name)
			}
			assert.Empty(t, d.Official)
		},
	},
}

func TestStructCompetition(t *testing.T) {
	allCases := []StructCompetitionTestCases{
		berries,
		contests,
		encounters,
		evolution,
		games,
		items,
		location,
		machines,
		moves,
		pokemon,
		resource,
		utility,
	}

	for _, cases := range allCases {
		cases.Run(t)
	}
}
