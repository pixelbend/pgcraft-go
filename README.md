# PgCraft Go

`pgcraft-go` is a PostgreSQL specification-compliant query builder written in Go. 
It simplifies the process of dynamically creating SQL queries while ensuring compliance with PostgreSQL standards.

## Features

- **PostgreSQL Compliant**: Ensures all queries adhere to PostgreSQL syntax and standards.
- **Dynamic Query Building**: Create SQL queries programmatically without writing raw SQL strings.
- **Type-Safe**: Leverages Go's strong typing to prevent common SQL injection vulnerabilities.
- **Extensible**: Easily adaptable for complex query requirements.

## Installation

Use `go get` to include `pgcraft-go` in your project:

```bash
go get github.com/arkamfahry/pgcraft-go
```

## Getting Started

### Import the Package

```go
import "github.com/arkamfahry/pgcraft-go"
```

### Example: Simple SELECT Query

Here's how to build a simple SQL `SELECT` query:

```go
package main

import (
	"fmt"
	"github.com/arkamfahry/pgcraft-go"
)

func main() {
	// Initialize the query builder
	qb := pgcraft.NewQueryBuilder()

	// Build a SELECT query
	query, args := qb.Select("id", "name").
		From("users").
		Where("age > ?", 18).
		OrderBy("name ASC").
		Build()

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)
}
```

**Output**:
```
Query: SELECT id, name FROM users WHERE age > $1 ORDER BY name ASC
Args: [18]
```

### Example: INSERT Query

```go
package main

import (
	"fmt"
	"github.com/arkamfahry/pgcraft-go"
)

func main() {
	// Initialize the query builder
	qb := pgcraft.NewQueryBuilder()

	// Build an INSERT query
	query, args := qb.InsertInto("users").
		Columns("name", "email").
		Values("John Doe", "john.doe@example.com").
		Build()

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)
}
```

**Output**:
```
Query: INSERT INTO users (name, email) VALUES ($1, $2)
Args: [John Doe john.doe@example.com]
```

### Example: UPDATE Query

```go
package main

import (
	"fmt"
	"github.com/arkamfahry/pgcraft-go"
)

func main() {
	// Initialize the query builder
	qb := pgcraft.NewQueryBuilder()

	// Build an UPDATE query
	query, args := qb.Update("users").
		Set("email = ?", "new.email@example.com").
		Where("id = ?", 1).
		Build()

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)
}
```

**Output**:
```
Query: UPDATE users SET email = $1 WHERE id = $2
Args: [new.email@example.com 1]
```

### Example: DELETE Query

```go
package main

import (
	"fmt"
	"github.com/arkamfahry/pgcraft-go"
)

func main() {
	// Initialize the query builder
	qb := pgcraft.NewQueryBuilder()

	// Build a DELETE query
	query, args := qb.DeleteFrom("users").
		Where("id = ?", 1).
		Build()

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)
}
```

**Output**:
```
Query: DELETE FROM users WHERE id = $1
Args: [1]
```
