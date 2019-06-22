# pokeapi-go
[![CircleCI](https://circleci.com/gh/mtslzr/pokeapi-go.svg?style=svg)](https://circleci.com/gh/mtslzr/pokeapi-go)

Wrapper for [Poke API](https://pokeapi.co), written in Go.

* [Getting Started](#getting-started)
* [Endpoints](#endpoints)
* [Not Implemented](#not-implemented)
* [Documentation](#documentation)

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
  c := pokeapi.ContestType("1")
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
  c := pokeapi.ContestType("1")
  ```

</details>

### Encounters

### Evolution

### Games

### Items

### Locations

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
  m := pokeapi.ContestType("1")
  ```

</details>

### Moves

### Pokemon

### Utility

<details>
  <summary>Languages</summary>
  
  #### Get Languages

  ```go
  m := pokeapi.Resource("language")
  ```

  #### Get Language

  *Must pass an ID (e.g. "1") or name (e.g. "en").*

  ```go
  m := pokeapi.ContestType("en")
  ```

</details>


## Not Implemented

Current progress on remaining endpoints. **Bold** are partially implemented.

### Encounters
- [ ] GET /encounter-method/{id or name}/
- [ ] GET /encounter-condition/{id or name}/
- [ ] GET /encounter-condition-value/{id or name}/
### Evolution
- [ ] GET /evolution-chain/{id}/
- [ ] GET /evolution-trigger/{id or name}/
### Games
- [ ] **GET /generation/{id or name}/**
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