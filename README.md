# pokeapi-go
![Actions](https://github.com/mtslzr/pokeapi-go/workflows/Actions/badge.svg)
![Release](https://img.shields.io/github/v/release/mtslzr/pokeapi-go)
[![Codecov](https://img.shields.io/codecov/c/github/mtslzr/pokeapi-go.svg?style=flat)](https://codecov.io/gh/mtslzr/pokeapi-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtslzr/pokeapi-go?style=flat)](https://goreportcard.com/report/github.com/mtslzr/pokeapi-go)
[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat)](http://godoc.org/github.com/mtslzr/pokeapi-go)
[![License](https://img.shields.io/badge/license-mit-blue.svg?style=flat)](https://raw.githubusercontent.com/mtslzr/pokeapi-go/master/LICENSE)

Wrapper for [Poke API](https://pokeapi.co), written in Go. *Supports PokeAPI v2.*

- [pokeapi-go](#pokeapi-go)
  - [Documentation](#Documentation)
  - [Getting Started](#Getting-Started)
  - [Endpoints](#Endpoints)
    - [Berries](#Berries)
    - [Contests](#Contests)
    - [Encounters](#Encounters)
    - [Evolution](#Evolution)
    - [Games](#Games)
    - [Items](#Items)
    - [Locations](#Locations)
    - [Machines](#Machines)
    - [Moves](#Moves)
    - [Pokemon](#Pokemon)
    - [Utility](#Utility)
  - [Additional Options](#Additional-Options)
    - [Resource List Parameters](#Resource-List-Parameters)
    - [Resource List Filters](#Resource-List-Filters)
    - [Caching](#Caching)

## Documentation

Full API documentation can be found at [Poke API](https://pokeapi.co/docs/v2.html).

## Getting Started

```bash
go get github.com/mtslzr/pokeapi-go
```

```go
import "github.com/mtslzr/pokeapi-go"
```

## Endpoints

### Berries

<details>
  <summary>Berries</summary>
  
  #### Get Berries

  ```go
  b := pokeapi.Resource("berry")
  ```

  #### Get Berry

  *Must pass an ID (e.g. "1") or name (e.g. "cheri").*

  ```go
  b := pokeapi.Berry("cheri")
  ```
</details>

<details>
  <summary>Berry Firmness</summary>
  
  #### Get Berry Firmnesses

  ```go
  b := pokeapi.Resource("berry-firmness")
  ```

  #### Get Berry Firmness

  *Must pass an ID (e.g. "1") or name (e.g. "very-soft").*

  ```go
  b := pokeapi.BerryFirmness("very-soft")
  ```
</details>

<details>
  <summary>Berry Flavors</summary>
  
  #### Get Berry Flavors

  ```go
  b := pokeapi.Resource("berry-flavor")
  ```

  #### Get Berry Flavor

  *Must pass an ID (e.g. "1") or name (e.g. "spicy").*

  ```go
  b := pokeapi.BerryFlavor("spicy")
  ```
</details>

### Contests

<details>
  <summary>Contest Types</summary>
  
  #### Get Contest Types

  ```go
  c := pokeapi.Resource("berry")
  ```

  #### Get Contest Type

  *Must pass an ID (e.g. "1") or name (e.g. "cool").*

  ```go
  c := pokeapi.ContestType("cool")
  ```
</details>

<details>
  <summary>Contest Effects</summary>
  
  #### Get Contest Effects

  ```go
  c := pokeapi.Resource("contest-effect")
  ```

  #### Get Contest Effect

  *Must pass an ID (e.g. "1").*

  ```go
  c := pokeapi.ContestEffect("1")
  ```
</details>

<details>
  <summary>Super Contest Effects</summary>
  
  #### Get Super Contest Effects

  ```go
  c := pokeapi.Resource("super-contest-effect")
  ```

  #### Get Super Contest Effect

  *Must pass an ID (e.g. "1").*

  ```go
  c := pokeapi.SuperContestEffect("1")
  ```
</details>

### Encounters

<details>
  <summary>Encounter Methods</summary>
  
  #### Get Encounter Methods

  ```go
  e := pokeapi.Resource("encounter-method")
  ```

  #### Get Encounter Method

  *Must pass an ID (e.g. "1") or name (e.g. "walk").*

  ```go
  e := pokeapi.EncounterMethod("walk")
  ```
</details>

<details>
  <summary>Encounter Conditions</summary>
  
  #### Get Encounter Conditions

  ```go
  e := pokeapi.Resource("encounter-condition")
  ```

  #### Get Encounter Condition

  *Must pass an ID (e.g. "1") or name (e.g. "swarm").*

  ```go
  e := pokeapi.EncounterCondition("swarm")
  ```
</details>

<details>
  <summary>Encounter Condition Values</summary>
  
  #### Get Encounter Condition Values

  ```go
  e := pokeapi.Resource("encounter-condition-value")
  ```

  #### Get Encounter Condition Value

  *Must pass an ID (e.g. "1") or name (e.g. "swarm-yes").*

  ```go
  e := pokeapi.EncounterConditionValue("swarm-yes")
  ```
</details>

### Evolution

<details>
  <summary>Evolution Chains</summary>
  
  #### Get Evolution Chains

  ```go
  e := pokeapi.Resource("evolution-chain")
  ```

  #### Get Evolution Chain

  *Must pass an ID (e.g. "1").*

  ```go
  e := pokeapi.EvolutionChain("1")
  ```
</details>

<details>
  <summary>Evolution Triggers</summary>
  
  #### Get Evolution Triggers

  ```go
  e := pokeapi.Resource("evolution-trigger")
  ```

  #### Get Evolution Trigger

  *Must pass an ID (e.g. "1") or name (e.g. "level-up").*

  ```go
  e := pokeapi.EvolutionTrigger("level-up")
  ```
</details>

### Games

<details>
  <summary>Generations</summary>
  
  #### Get Generations

  ```go
  g := pokeapi.Resource("generation")
  ```

  #### Get Generation

  *Must pass an ID (e.g. "1") or name (e.g. "generation-i").*

  ```go
  g := pokeapi.Generation("generation-i")
  ```
</details>

<details>
  <summary>Pokedex</summary>
  
  #### Get All Pokedex

  ```go
  g := pokeapi.Resource("pokedex")
  ```

  #### Get Single Pokedex

  *Must pass an ID (e.g. "1") or name (e.g. "national").*

  ```go
  g := pokeapi.Pokedex("national")
  ```
</details>

<details>
  <summary>Versions</summary>
  
  #### Get Versions

  ```go
  g := pokeapi.Resource("version")
  ```

  #### Get Version

  *Must pass an ID (e.g. "1") or name (e.g. "red").*

  ```go
  g := pokeapi.Version("red")
  ```
</details>

<details>
  <summary>Version Groups</summary>
  
  #### Get Version Groups

  ```go
  g := pokeapi.Resource("version-group")
  ```

  #### Get Version Group

  *Must pass an ID (e.g. "1") or name (e.g. "red-blue").*

  ```go
  g := pokeapi.VersionGroup("red-blue")
  ```
</details>

### Items

<details>
  <summary>Items</summary>
  
  #### Get Items

  ```go
  i := pokeapi.Resource("item")
  ```

  #### Get Item

  *Must pass an ID (e.g. "1") or name (e.g. "master-ball").*

  ```go
  i := pokeapi.Item("master-ball")
  ```
</details>

<details>
  <summary>Item Attributes</summary>
  
  #### Get Item Attributes

  ```go
  i := pokeapi.Resource("item-attribute")
  ```

  #### Get Item Attribute

  *Must pass an ID (e.g. "1") or name (e.g. "countable").*

  ```go
  i := pokeapi.ItemAttribute("countable")
  ```
</details>

<details>
  <summary>Item Categories</summary>
  
  #### Get Item Ctegories

  ```go
  i := pokeapi.Resource("item-category")
  ```

  #### Get Item Category

  *Must pass an ID (e.g. "1") or name (e.g. "stat-boosts").*

  ```go
  i := pokeapi.ItemCategory("stat-boosts")
  ```
</details>

<details>
  <summary>Item Fling Effects</summary>
  
  #### Get Item Fling Effects

  ```go
  i := pokeapi.Resource("item-fling-effect")
  ```

  #### Get Item Fling Effect

  *Must pass an ID (e.g. "1") or name (e.g. "badly-poison").*

  ```go
  i := pokeapi.ItemFlingEffect("badly-poison")
  ```
</details>

<details>
  <summary>Item Pockets</summary>
  
  #### Get Item Pockets

  ```go
  i := pokeapi.Resource("item-pocket")
  ```

  #### Get Item Pocket

  *Must pass an ID (e.g. "1") or name (e.g. "misc").*

  ```go
  i := pokeapi.ItemPocket("misc")
  ```
</details>

### Locations

<details>
  <summary>Locations</summary>
  
  #### Get Locations

  ```go
  l := pokeapi.Resource("location")
  ```

  #### Get Location

  *Must pass an ID (e.g. "1") or name (e.g. "canalave-city").*

  ```go
  l := pokeapi.Location("canalave-city")
  ```
</details>

<details>
  <summary>Location Areas</summary>
  
  #### Get Location Areas

  ```go
  l := pokeapi.Resource("location-area")
  ```

  #### Get Location Area

  *Must pass an ID (e.g. "1") or name (e.g. "canalave-city-area").*

  ```go
  l := pokeapi.LocationArea("canalave-city-area")
  ```
</details>

<details>
  <summary>Pal Park Areas</summary>
  
  #### Get Pal Park Areas

  ```go
  l := pokeapi.Resource("pal-park-area")
  ```

  #### Get Pal Park Area

  *Must pass an ID (e.g. "1") or name (e.g. "forest").*

  ```go
  l := pokeapi.PalParkArea("forest")
  ```
</details>

<details>
  <summary>Regions</summary>
  
  #### Get Regions

  ```go
  l := pokeapi.Resource("region")
  ```

  #### Get Region

  *Must pass an ID (e.g. "1") or name (e.g. "kanto").*

  ```go
  l := pokeapi.Region("kanto")
  ```
</details>

### Machines

<details>
  <summary>Machines</summary>
  
  #### Get Machines

  ```go
  m := pokeapi.Resource("machine")
  ```

  #### Get Machine

  *Must pass an ID (e.g. "1").*

  ```go
  m := pokeapi.Machine("1")
  ```
</details>

### Moves

<details>
  <summary>Moves</summary>
  
  #### Get Moves

  ```go
  m := pokeapi.Resource("move")
  ```

  #### Get Move

  *Must pass an ID (e.g. "1") or name (e.g. "pound").*

  ```go
  m := pokeapi.Move("pound")
  ```
</details>

<details>
  <summary>Move Ailments</summary>
  
  #### Get Move Ailments

  ```go
  m := pokeapi.Resource("move-ailment")
  ```

  #### Get Move Ailment

  *Must pass an ID (e.g. "1") or name (e.g. "paralysis").*

  ```go
  m := pokeapi.MoveAilment("paralysis")
  ```
</details>

<details>
  <summary>Move Battle Styles</summary>
  
  #### Get Move Battle Styles

  ```go
  m := pokeapi.Resource("move-battle-style")
  ```

  #### Get Move Battle Style

  *Must pass an ID (e.g. "1") or name (e.g. "attack").*

  ```go
  m := pokeapi.MoveBattleStyle("attack")
  ```
</details>

<details>
  <summary>Move Categories</summary>
  
  #### Get Move Categories

  ```go
  m := pokeapi.Resource("move-catgory")
  ```

  #### Get Move Category

  *Must pass an ID (e.g. "1") or name (e.g. "ailment").*

  ```go
  m := pokeapi.MoveCategory("ailment")
  ```
</details>

<details>
  <summary>Move Damage Classes</summary>
  
  #### Get Move Damage Classes

  ```go
  m := pokeapi.Resource("move-damage-class")
  ```

  #### Get Move Damage Class

  *Must pass an ID (e.g. "1") or name (e.g. "status").*

  ```go
  m := pokeapi.MoveDamageClass("status")
  ```
</details>

<details>
  <summary>Move Learn Methods</summary>
  
  #### Get Move Learn Methods

  ```go
  m := pokeapi.Resource("move-learn-method")
  ```

  #### Get Move Learn Method

  *Must pass an ID (e.g. "1") or name (e.g. "level-up").*

  ```go
  m := pokeapi.MoveLearnMethod("level-up")
  ```
</details>

<details>
  <summary>Move Targets</summary>
  
  #### Get Move Targets

  ```go
  m := pokeapi.Resource("move-target")
  ```

  #### Get Move Target

  *Must pass an ID (e.g. "1") or name (e.g. "specific-move").*

  ```go
  m := pokeapi.MoveTarget("specific-move")
  ```
</details>

### Pokemon

<details>
  <summary>Abilities</summary>
  
  #### Get Abilities

  ```go
  p := pokeapi.Resource("ability")
  ```

  #### Get Ability

  *Must pass an ID (e.g. "1") or name (e.g. "stench").*

  ```go
  p := pokeapi.Ability("stench")
  ```
</details>

<details>
  <summary>Characteristics</summary>
  
  #### Get Characteristics

  ```go
  p := pokeapi.Resource("characteristic")
  ```

  #### Get Characteristic

  *Must pass an ID (e.g. "1").*

  ```go
  p := pokeapi.Characteristic("1")
  ```
</details>

<details>
  <summary>Egg Groups</summary>
  
  #### Get Egg Groups

  ```go
  p := pokeapi.Resource("egg-group")
  ```

  #### Get Egg Group

  *Must pass an ID (e.g. "1") or name (e.g. "monster").*

  ```go
  p := pokeapi.EggGroup("monster")
  ```
</details>

<details>
  <summary>Genders</summary>
  
  #### Get Genders

  ```go
  p := pokeapi.Resource("gender")
  ```

  #### Get Gender

  *Must pass an ID (e.g. "1") or name (e.g. "female").*

  ```go
  p := pokeapi.Gender("female")
  ```
</details>

<details>
  <summary>Growth Rates</summary>
  
  #### Get Growth Rates

  ```go
  p := pokeapi.Resource("growth-rate")
  ```

  #### Get Growth Rate

  *Must pass an ID (e.g. "1") or name (e.g. "slow").*

  ```go
  p := pokeapi.GrowthRate("slow")
  ```
</details>

<details>
  <summary>Natures</summary>
  
  #### Get Natures

  ```go
  p := pokeapi.Resource("nature")
  ```

  #### Get Nature

  *Must pass an ID (e.g. "1") or name (e.g. "hardy").*

  ```go
  p := pokeapi.Nature("hardy")
  ```
</details>

<details>
  <summary>Pokeathlon Stats</summary>
  
  #### Get Pokeathlon Stats

  ```go
  p := pokeapi.Resource("pokeathlon-stat")
  ```

  #### Get Pokeathlon Stat

  *Must pass an ID (e.g. "1") or name (e.g. "speed").*

  ```go
  p := pokeapi.PokeathlonStat("speed")
  ```
</details>

<details>
  <summary>Pokemon</summary>
  
  #### Get All Pokemon

  ```go
  l := pokeapi.Resource("pokemon")
  ```

  #### Get Single Pokemon

  *Must pass an ID (e.g. "1") or name (e.g. "bulbasaur").*

  ```go
  l := pokeapi.Pokemon("bulabsaur")
  ```
</details>

<details>
  <summary>Pokemon Colors</summary>
  
  #### Get Pokemon Colors

  ```go
  p := pokeapi.Resource("pokemon-color")
  ```

  #### Get Pokemon Color

  *Must pass an ID (e.g. "1") or name (e.g. "black").*

  ```go
  p := pokeapi.PokemonColor("black")
  ```
</details>

<details>
  <summary>Pokemon Forms</summary>
  
  #### Get Pokemon Forms

  ```go
  p := pokeapi.Resource("pokemon-form")
  ```

  #### Get Pokemon Form

  *Must pass an ID (e.g. "1") or name (e.g. "bulbasaur").*

  ```go
  p := pokeapi.PokemonForm("bulabsaur")
  ```
</details>

<details>
  <summary>Pokemon Habitats</summary>
  
  #### Get Pokemon Habitats

  ```go
  p := pokeapi.Resource("pokemon-habitat")
  ```

  #### Get Pokemon Habitat

  *Must pass an ID (e.g. "1") or name (e.g. "cave").*

  ```go
  p := pokeapi.PokemonHabitat("cave")
  ```
</details>

<details>
  <summary>Pokemon Shapes</summary>
  
  #### Get Pokemon Shapes

  ```go
  p := pokeapi.Resource("pokemon-shape")
  ```

  #### Get Pokemon Shape

  *Must pass an ID (e.g. "1") or name (e.g. "ball").*

  ```go
  p := pokeapi.PokemonShape("ball")
  ```
</details>

<details>
  <summary>Pokemon Species</summary>
  
  #### Get All Pokemon Species

  ```go
  p := pokeapi.Resource("pokemon-species")
  ```

  #### Get Single Pokemon Species

  *Must pass an ID (e.g. "1") or name (e.g. "bulbasaur").*

  ```go
  p := pokeapi.PokemonSpecies("bulabsaur")
  ```
</details>

<details>
  <summary>Stats</summary>
  
  #### Get Stats

  ```go
  p := pokeapi.Resource("stat")
  ```

  #### Get Stat

  *Must pass an ID (e.g. "1") or name (e.g. "hp").*

  ```go
  p := pokeapi.Stat("hp")
  ```
</details>

<details>
  <summary>Types</summary>
  
  #### Get Types

  ```go
  p := pokeapi.Resource("type")
  ```

  #### Get Type

  *Must pass an ID (e.g. "1") or name (e.g. "normal").*

  ```go
  p := pokeapi.Type("normal")
  ```
</details>

### Utility

<details>
  <summary>Languages</summary>
  
  #### Get Languages

  ```go
  u := pokeapi.Resource("language")
  ```

  #### Get Language

  *Must pass an ID (e.g. "1") or name (e.g. "en").*

  ```go
  u := pokeapi.Language("en")
  ```
</details>

## Additional Options

### Resource List Parameters

When calling `pokeapi.Resource()` for any resource list, you can optionally pass up to two integers. The first will be an offset (defaults to zero), and the second will be the limit (defaults two twenty).

<details>
  <summary>Default</summary>
  
  ```go
  r := pokeapi.Resource("pokemon")
  fmt.Println(len(r.Results)) // 20
  fmt.Println(r.Results[0].Name) // "bulbasaur"
  ```
</details>

<details>
  <summary>Offset</summary>

  ```go
  r := pokeapi.Resource("pokemon", 3)
  fmt.Println(len(r.Results)) // 20
  fmt.Println(r.Results[0].Name) // "charmander"
  ```
</details>

<details>
  <summary>Offset and Limit</summary>

  ```go
  r := pokeapi.Resource("pokemon", 6, 10)
  fmt.Println(len(r.Results)) // 10
  fmt.Println(r.Results[0].Name) // "squirtle"
  ```
</details>

### Resource List Filters

As an alternative to `pokeapi.Resource()`, you can use Search to filter resource lists. Pass the endpoint, followed by the search term. Or pass a string starting with `^` to search for items starting with the search team.

*result.Count is updated after the search with the new total (to get the full count, use `pokeapi.Resource()`.*

<details>
  <summary>Search</summary>

  ```go
  s := pokeapi.Search("pokemon", "saur")
  fmt.Println(len(s.Results)) // 4
  fmt.Println(s.Results[3].Name) // venusaur-mega
  ```
</details>

<details>
  <summary>Starts With</summary>

  ```go
  s := pokeapi.Search("pokemon", "^a")
  fmt.Println(len(s.Results)) // 44
  fmt.Println(s.Results[0].Name) // arbok
  ```

  ```go
  s := pokeapi.Search("pokemon", "^bulb")
  fmt.Println(len(s.Results)) // 1
  fmt.Println(s.Results[0].Name) // bulbasaur
  ```
</details>

### Caching

Calls are automatically cached to cut down on API traffic to PokeAPI, with subsequent calls returning local data.

#### Clearing Cache

```go
// Clear all existing cache entries.
pokeapi.ClearCache()
```

#### Custom Expiration

Custom cache expiration remains for all calls until changed or unset.

```go
// Set cache expiration to twenty minutes.
pokeapi.CacheSettings.CustomExpire = 20
// Turn custom expiration back off.
pokeapi.CacheSettings.CustomExpire = 0
```

#### Disable Cache

_Please be considerate of PokeAPI and be sure to always operate within this requested limits._

As with custom expiration, this setting remains for all calls until changed or unset.

```go
// Disable checking for cached data
pokeapi.CacheSettings.UseCache = false
// Re-enable checking for cached data
pokeapi.CacheSettings.UseCache = true
```
