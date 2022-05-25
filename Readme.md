# conditional

Conditional is go/golang replacement for `ternary` and `nil coalescing` and `optional chaining` operators.

### Usage

For `ternary` just call `conditional.Ternary` with the values you want to return:

```go
val := conditional.Ternary(true, "true", "false")
// val == "true"
```

For `nil coalescing` just call `conditional.NilCoalescing` with the value as default you want to return:

```go
var val1 *string
val2 := "val2"
val := *conditional.NilCoalescing(val1, &val2)
// val == "val2"
```

For `optional chaining` just call `conditional.OptionalChaining` with the value and a function with access to a property:

```go
var st *struct {
    val2 string
}
val := conditional.OptionalChaining(st, func() *string {
    return &st.val2
})
// val == nil
```