# pokeapi-go
[![CircleCI](https://circleci.com/gh/mtslzr/pokeapi-go.svg?style=svg)](https://circleci.com/gh/mtslzr/pokeapi-go)

Wrapper for [Poke API](https://pokeapi.co), written in Go.

- [pokeapi-go](#pokeapi-go)
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
  - [Not Implemented](#Not-Implemented)
    - [Items](#Items-1)
    - [Moves](#Moves-1)
  - [Documentation](#Documentation)

## Getting Started

```bash
go get github.com/mtslzr/pokeapi-go
```

```go
import pokeapi "github.com/mtslzr/pokeapi-go"
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

### Pokemon

### Utility

<details>
  <summary>Languages</summary>
  
  #### Get Languages

  ```go
  l := pokeapi.Resource("language")
  ```

  #### Get Language

  *Must pass an ID (e.g. "1") or name (e.g. "en").*

  ```go
  l := pokeapi.Lanuage("en")
  ```
</details>


## Not Implemented

Current progress on remaining endpoints. **Bold** are partially implemented.

### Items
- [ ] GET /item/{id or name}/
- [ ] GET /item-attribute/{id or name}/
- [ ] GET /item-category/{id or name}/
- [ ] GET /item-fling-effect/{id or name}/
- [ ] GET /item-pocket/{id or name}/
### Moves
- [ ] GET /move/{id or name}/
- [ ] GET /move-ailment/{id or name}/
- [ ] GET /move-battle-style/{id or name}/
- [ ] GET /move-category/{id or name}/
- [ ] GET /move-damage-class/{id or name}/
- [ ] GET /move-learn-method/{id or name}/
- [ ] GET /move-target/{id or name}/

## Documentation

Full API documentation can be found at [Poke API](https://pokeapi.co/docs/v2.html).