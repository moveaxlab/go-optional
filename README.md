# Optionals in Go

This is a no nonsense zero dependency generic `Optional<T>` implementation for Go.

## Installation

```bash
go get github.com/moveaxlab/go-optional
```

## Usage

Return optionals from your repositories:

```go
package repository

import (
  "context"

  "github.com/moveaxlab/go-optional"
)

type Repository interface {
  FindById(ctx context.Context, id string) optional.Optional[Entity]
}

type repository struct {}

func (r *repository) FindById(ctx context.Context, id string) optional.Optional[Entity] {
  var entity Entity
  var entityWasFound bool

  // retrieve the entity in some way

  if !entityWasFound {
    return optional.EmptyOptional[Entity]()
  }

  return optional.OptionalOf(&entity)
}
```

The use the optionals in your application logic:

```go
package service

func (s *service) DoSomeStuff(ctx context.Context) {
  maybeEntity := s.repository.FindById(ctx, "1")

  if maybeEntity.IsPresent() {
    // we have the entity!
    entity := maybeEntity.Get()
  }
}
```

The `Optional` type offers these methods:

- `IsPresent` returns true if there is a value in the optional
- `Get` returns the value contained in the optional, and panics if it's empty
- `OrElseGet` accepts a function in input, and returns the value contained
  in the optional if present, or the result of the function otherwise
- `OrElsePanic` returns the value contained in the optional if present,
  or panics with a custom error passed in input
- `IfPresent` runs the function passed as argument if the value is present,
  passing it the value contained in the optional
- `IfPresentOrElse` behaves like `IfPresent` for the first argument,
  but calls the function passed as second argument if the optional is empty

