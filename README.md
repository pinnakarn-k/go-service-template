# Development Checklist

Run these commands before every commit to keep the project healthy.

## 1. Run Unit Tests

```bash
go test ./...
```

Verify that all unit tests pass.

---

## 2. Build the Project

```bash
go build ./...
```

Ensure the project compiles successfully.

---

## 3. Run Linter

```bash
golangci-lint run
```

Check for:

- Unused code
- Dead code
- Common mistakes
- Style issues
- Static analysis warnings

---

## Daily Workflow

Run the following before committing:

```bash
go test ./...
go build ./...
golangci-lint run
```

If all commands pass, the project is ready to commit.

---

## Use Identity Context

```bash
identity, err := CurrentIdentity(c)
identity.CustCode
identity.AccountID
```