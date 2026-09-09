# Code review — ledger.go / ledger_test.go

Review focus: Go idiom, algorithm efficiency, correctness, test coverage.
Every correctness claim below was verified by running throwaway tests.

## Bugs

### 1. `ledger.go:61,65,69` — wrong variable in 3 error messages
`fromAccount` is `*AccountData`, not `AccountId`, but all three lines pass it to `%d`.

Verified output:
```
Insufficient funds on account &{5 0}
Cannot transfer to same account &{5 0}
Accounts with ids &{5 0} and 99 need to exist
```

`go vet` misses this because `%d` recursively accepts a struct of integers. Use `fromAccountId`.

### 2. `ledger.go:110` — `GetTopSpenders(0)` panics
With >= 1 account, `len(l.accounts) <= 0` is false, so `h` is built with capacity 0. The first
loop iteration takes the `else if` branch and indexes `h[0]`.

Verified: `runtime error: index out of range [0] with length 0`

### 3. `ledger.go:114` — negative `number` panics
`GetTopSpenders(-1)` with >= 1 account reaches `make([]idAndSpending, 0, -1)`.

Verified: `runtime error: makeslice: cap out of range`

Fix 2 and 3 with one guard at the top:

```go
if number <= 0 {
	return nil
}
```

### 4. `ledger.go:26` — `Withdraw` silently underflows
`MoneyAmount` is `uint`, so `ad.balance -= amount` wraps when `amount > balance`.

Verified: balance 10, withdraw 20 -> `18446744073709551606`

Today only `Transfer` calls it and it checks first, but the guard lives far from the subtraction
and the next caller can bypass it. Move the invariant into the method:

```go
func (ad *AccountData) Withdraw(amount MoneyAmount) error {
	if ad.balance < amount {
		return fmt.Errorf("insufficient funds: have %d, need %d", ad.balance, amount)
	}
	ad.balance -= amount
	ad.spendings += amount
	return nil
}
```

### 5. `ledger.go:109` — "top spenders" returns no usable order
- Heap path returns the raw heap array order. Verified: spendings 60/50/40 -> `[3 1 2]`.
- The `len(l.accounts) <= number` shortcut returns map iteration order and never reads
  `spendings` at all. Verified: `[5 6 1 2 3 4]`.

A caller asking for "top spenders" reasonably expects descending rank. Either sort the result
descending before returning, or rename/document that it returns an unordered set.

### 6. `ledger.go:1` — `package main` with no `func main`
Verified: `go build ./...` fails with
`runtime.main_main·f: function main is undeclared in the main package`.

Tests pass, so this is easy to miss, but any CI running `go build` is red. This is a library:
move it to `ledger/ledger.go` as `package ledger`, and add `cmd/` later if you want a binary.

## Idiom

### `ledger.go:78` — `(error, MoneyAmount)` return order is backwards
Go convention is values first, `error` last, so every caller reads wrong
(`err, balance := ...`). Worse, the failure case returns `1` instead of the zero value, so a
caller that ignores the error sees a phantom balance of 1.

Make it `(MoneyAmount, error)` returning `0, err`.

### `ledger.go:42` and 4 more — `errors.New(fmt.Sprintf(...))`
Use `fmt.Errorf(...)` (staticcheck S1028). Go error strings start lowercase and carry no
trailing punctuation (ST1005). Consider sentinel errors so callers and tests can use
`errors.Is` instead of only checking `!= nil`:

```go
var ErrNoSuchAccount = errors.New("no such account")
// return fmt.Errorf("%w: %d", ErrNoSuchAccount, id)
```

### `ledger.go:35` — `NewLedger` returns a value while all methods take `*Ledger`
`b := a` produces two `Ledger` values sharing one map — verified that mutating `a` is visible
through `b`. A footgun disguised as a copy. Return `*Ledger`.

### `ledger.go:39,48,57` — `timestamp` accepted and never used
`CreateAccount`, `Deposit`, and `Transfer` all carry a parameter the implementation drops, so
out-of-order or duplicate timestamps are silently accepted. Either store/validate it
(monotonicity check, or an event log) or remove it from the API until it's needed.

### `ledger.go:49` — double map lookup in `Deposit`
`_, exists := l.accounts[account]` followed by `l.accounts[account].Deposit(...)`. Use the
value from the first lookup: `acc, ok := l.accounts[account]; ...; acc.Deposit(amount)`.
`Transfer` already does this correctly for `fromAccount`.

### `ledger.go:137` — `for accountId, _ := range` should be `for accountId := range`
gosimple S1005.

### `ledger.go:44` — `&AccountData{0, 0}` unkeyed composite literal
Adding a third field to `AccountData` breaks this line at compile time for no benefit.
`&AccountData{}` already zeroes both fields.

### Tooling
Wire `staticcheck` into your loop. `gofmt` flags none of the last three; `staticcheck` catches
all of them plus the `errors.New(Sprintf)` pattern and the capitalized error strings.

### `ledger.go:31` — no synchronization
A ledger is the canonical thing people reach for from multiple goroutines. Concurrent
`Deposit` calls on one `Ledger` are a data race on the map and on `balance`. Either add a
`sync.Mutex` to `Ledger` and lock in every method, or document explicitly that callers must
serialize. Then run `go test -race`.

## Test suite

Decent coverage of the happy paths and the error branches that exist, but there is a blind spot
that maps exactly onto the bugs above: **every bug found is reachable and no current test
reaches it.**

- `assertTopSpenderIds` sorts `actual` before comparing, so it asserts set equality only and
  would pass on any permutation. The ordering contract is untested by construction.
- `TestTopNSpendersWhenSomeAccountsPresent` (2 accounts, `number=3`) takes the
  `getAllAccountIds` shortcut, so it never exercises ranking. Only one test actually runs the heap.
- No test for `number == 0`, negative `number`, or `number == len(accounts)` (the `<=`
  boundary). The first two panic.
- No tie test. Two accounts with equal spending currently give nondeterministic output —
  decide the contract and pin it.
- Error assertions are all `err == nil` / `err != nil`, so a test passes on the *wrong* error.
  Once sentinel errors exist, assert which one with `errors.Is`.
- No test that a **rejected** operation left state untouched: after a failed
  insufficient-funds transfer, assert both balances and both `spendings` are unchanged. That is
  the test that would have caught the underflow.
- Boundary: `Transfer` of exactly the full balance (`amount == balance`) is allowed and leaves
  0 — untested. `Transfer` of `0` is also allowed; intended?
- `ledger_test.go:188` — `assertTopSpenderIds` is missing `t.Helper()`. Failures report line
  191 inside the helper instead of the calling test, and with three call sites the output
  doesn't say which test failed.
- Setup calls mix `ledger.CreateAccount(1, 1)` and `_ = ledger.Deposit(...)` inconsistently and
  never check the result. If a setup call starts failing, the test fails with a confusing
  assertion rather than at the real cause. A small `mustCreate`/`mustDeposit` pair with
  `t.Helper()` and `t.Fatal` fixes that.
- The `TestTopNSpenders...` family is ideal for table-driven tests
  (`tests := []struct{ name string; ... }` plus `t.Run(tt.name, ...)`) — the standard Go idiom,
  and it makes adding the tie and boundary cases above nearly free.

## Suggested fix order

1. Panics (`number <= 0`) and the `%d` variable bug — small, mechanical.
2. `(MoneyAmount, error)` swap, `*Ledger` from `NewLedger`, `Withdraw` returning `error` —
   touches callers and tests.
3. Package move, sentinel errors, mutex.
4. Test additions: rollback-on-failure, ordering, ties, boundaries, table-driven rewrite.
