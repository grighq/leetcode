# Agent Instructions

LeetCode practice repository in Go 1.26.2.

## Structure

```
easy/           # Solved problems (7)
medium/         # Empty
hard/           # Empty
```

Each problem directory contains:
- `description.md` — Problem description
- `<problem>.go` — Solution (function is lowercase/unexported by default)
- `<problem>_test.go` — Tests using `testify/assert`

## Conventions

- **Package names**: Derived from directory name (e.g., `twosum`, `mergedtwosortedlists`)
- **Function visibility**: Solutions are unexported (lowercase) unless tests need access
- **Test helpers**: Define in `*_test.go` (e.g., `sliceToList`, `listToSlice` for linked lists)

## Commands

```bash
# Run all tests
go test ./...

# Run tests for specific problem
go test ./easy/two_sum/

# Run single test
go test -run TestTwoSum ./easy/two_sum/

# Install/update dependencies
go mod tidy
```

## Adding New Problems

1. Create directory: `easy/<problem_name>/`
2. Package name: lowercase, no underscores (e.g., `mergedtwosortedlists`)
3. Copy `description.md` from LeetCode
4. Implement solution with unexported function
5. Write tests using `testify/assert`
6. Run `go test ./<difficulty>/<problem>/`
