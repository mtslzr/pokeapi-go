# pokeapi-go
[![CircleCI](https://circleci.com/gh/mtslzr/pokeapi-go.svg?style=svg)](https://circleci.com/gh/mtslzr/pokeapi-go)

Wrapper for [Poke API](https://pokeapi.co), written in Go.

* [How To](#how-to)
* [Progress](#progress)
* [Endpoints](#endpoints)
* [Documentation](#documentation)

## How To

```go
import pokeapi "github.com/mtslzr/pokeapi-go"
```

## Progress

Current progress of endpoints. **Bold** are partially implemented.

### Berries
- [ ] GET /berry/{id or name}/
- [ ] GET /berry-firmness/{id or name}/
- [ ] GET /berry-flavor/{id or name}/
### Contests
- [ ] GET /contest-type/{id or name}/
- [ ] GET /contest-effect/{id}/
- [ ] GET /super-contest-effect/{id}/
### Encounters
- [ ] GET /encounter-method/{id or name}/
- [ ] GET /encounter-condition/{id or name}/
- [ ] GET /encounter-condition-value/{id or name}/
### Evolution
- [ ] GET /evolution-chain/{id}/
- [ ] GET /evolution-trigger/{id or name}/
### Games
- [x] GET /generation/{id or name}/
- [ ] GET /pokedex/{id or name}/
- [ ] GET /version/{id or name}/
- [ ] GET /version-group/{id or name}/
### Items
- [ ] GET /item/{id or name}/
- [ ] GET /item-attribute/{id or name}/
- [ ] GET /item-category/{id or name}/
- [ ] GET /item-fling-effect/{id or name}/
- [ ] GET /item-pocket/{id or name}/
### Locations
- [ ] GET /location/{id or name}/
- [ ] GET /location-area/{id or name}/
- [ ] GET /pal-park-area/{id or name}/
- [ ] GET /region/{id or name}/
### Machines
- [ ] GET /machine/{id}/
### Moves
- [ ] GET /move/{id or name}/
- [ ] GET /move-ailment/{id or name}/
- [ ] GET /move-battle-style/{id or name}/
- [ ] GET /move-category/{id or name}/
- [ ] GET /move-damage-class/{id or name}/
- [ ] GET /move-learn-method/{id or name}/
- [ ] GET /move-target/{id or name}/
### Pokemon
- [x] GET /ability/{id or name}/
- [x] GET /egg-group/{id or name}/
- [ ] **GET /gender/{id or name}/**
- [ ] **GET /growth-rate/{id or name}/**
- [ ] GET /nature/{id or name}/
- [ ] **GET /pokeathlon-stat/{id or name}/**
- [x] GET /pokemon/{id or name}/
- [ ] **GET /pokemon-color/{id or name}/**
- [ ] **GET /pokemon-form/{id or name}/**
- [ ] **GET /pokemon-habitat/{id or name}/**
- [ ] **GET /pokemon-species/{id or name}/**
- [ ] **GET /stat/{id or name}/**
- [ ] **GET /type/{id or name}/**
### Utility
- [ ] GET /language/{id or name}/

## Endpoints

### Games

#### Get Generations

Returns all generations.

```go
gens := pokeapi.Generations()

fmt.Println(gens.Count) // 7
fmt.Println(gens.Results[0].Name) // "generation-i"
```

#### Get Generation

Returns single generation (by name or ID).

```go
gen := pokeapi.Generation("1")
fmt.Println(gen.MainRegion.Name) // "kanto"

gen := pokeapi.GenerationByName("generation-i")
fmt.Println(gen.MainRegion.Name) // "kanto"
```

### Pokemon

#### Get Pokemons

Returns all Pokemon.

*NOTE: Pagination not yet implemented.*

```go
pokes := pokeapi.AllPokemon()

fmt.Println(pokes.Count) // 964
fmt.Println(pokes.Results[0].Name) // "bulbasaur"
```

#### Get Pokemon

Returns single Pokemon (by name or ID).

```go
poke := pokeapi.Pokemon("1")
fmt.Println(poke.Name) // "bulbasaur"

poke := pokeapi.Pokemon("bulbasaur")
fmt.Println(poke.Name) // "bulbasaur"
```

## Documentation

Full API documentation can be found at [Poke API](https://pokeapi.co/docs/v2.html).