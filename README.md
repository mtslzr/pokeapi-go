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
  <summary>Get Berries</summary>
  
  #### All Berries

  ```go
  b := pokeapi.Resource("berry")
  ```

  #### Single Berry

  *Can pass an ID (e.g. "1") or name (e.g. "cheri").*

  ```go
  b := pokeapi.Berry("cheri")
  ```

</details>

<details>
  <summary>Get Berry Firmness</summary>
  
  #### All Berry Firmnesses

  ```go
  b := pokeapi.Resource("berry-firmness")
  ```

  #### Single Berry Firmness

  *Can pass an ID (e.g. "1") or name (e.g. "very-soft").*

  ```go
  b := pokeapi.BerryFirmness("very-soft")
  ```

</details>

<details>
  <summary>Get Berry Flavors</summary>
  
  #### All Berry Flavors

  ```go
  b := pokeapi.Resource("berry-flavor")
  ```

  #### Single Berry

  *Can pass an ID (e.g. "1") or name (e.g. "spicy").*

  ```go
  b := pokeapi.BerryFlavor("spicy")
  ```

</details>

### Contests

### Encounters

### Evolution

### Games

### Items

### Locations

### Machines

### Moves

### Pokemon

### Utility

## Not Implemented

Current progress on remaining endpoints. **Bold** are partially implemented.

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
### Utility
- [ ] GET /language/{id or name}/

## Documentation

Full API documentation can be found at [Poke API](https://pokeapi.co/docs/v2.html).