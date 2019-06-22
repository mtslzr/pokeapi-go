# pokeapi-go

Wrapper for Poke API, written in Go.

## How To

```go
import poke "github.com/mtslzr/pokeapi-go"
```

## Endpoints

### Get All Pokemon

```go
all := poke.AllPokemon()

fmt.Println(all.Count) // 964
fmt.Println(all.Results[0].Name) // "bulbasaur"
```

### Get Pokemon (by ID)

```go
bulbasaur := poke.Pokemon(1)

fmt.Println(bulbasaur.Name) // "bulbasaur"
```