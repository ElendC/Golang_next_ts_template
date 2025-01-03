# PostgreSQL  SQL


## pgx package
- Native PostgreSQL driver, offers high performance.
- Can use custom PostgreSQL features. 
- Lightweight driver without much overhead.


## context methods
### context.Background()
- **Use**: For initialization, root-level, or standalone operations where no cancellation or timeout is needed.
- **Example**: Setting up database, creating tables

### context.TODO()
- **Use**: When you're unsure what context to use but need a placeholder to refactor later.
- **Example**: Temporarily during development

### context.WithCancel(parent)
- **Use**: When you need to cancel an operation manually 
- **Example**: Graceful shutdown in servers

`ctx, cancel := context.WithCancel(parent)
cancel()
`

### context.WithTimeout(parent, duration)
- **Use**: When you want a maximum duration for an operation
- **Example**: Database queries or HTTP requests

`ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()
`

### context.WithDeadline(parent, time)
- **Use**: When you want an operation to finish by a specific time
- **Example**: Fixed deadline for scheduled jobs

`ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
defer cancel()
`

### context.WithValue(parent, key, value)
- Use: To pass request-specific values through the context
- Example: Passing user IDs or auth tokens in HTTP handlers

`ctx := context.WithValue(parent, "key", "value")
`