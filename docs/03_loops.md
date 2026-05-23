# Loops

Go has **one** loop keyword: `for`. It covers every looping pattern.

---

## 1. C-style `for`

```go
for init; condition; post {
}
```

```go
for i := 0; i < 5; i++ { }    // 0 1 2 3 4
for i := 5; i > 0; i-- { }    // 5 4 3 2 1
for i := 0; i <= 10; i += 2 { } // 0 2 4 6 8 10
```

- `init` — runs once before the loop starts
- `condition` — checked before each iteration; loop stops when false
- `post` — runs after each iteration (`i++`, `i--`, `i += 2`)

---

## 2. While-style

Go has no `while` keyword. Drop `init` and `post`, keep only the condition.

```go
n := 1
for n < 64 {
    n *= 2
}
// n = 64
```

---

## 3. Infinite Loop

Omit everything. Must exit with `break` or `return`.

```go
for {
    if done {
        break
    }
}
```

---

## 4. `continue` — Skip an Iteration

Jumps straight to the next iteration, skipping the rest of the loop body.

```go
for i := 0; i < 10; i++ {
    if i%2 != 0 {
        continue  // skip odd numbers
    }
    fmt.Print(i, " ")  // prints: 0 2 4 6 8
}
```

---

## 5. `range` over Slice

`range` yields the **index** and **value** for each element.

```go
animals := []string{"cat", "dog", "eagle"}

for i, animal := range animals { }  // both index and value
for i := range animals { }          // index only
for _, animal := range animals { }  // value only (discard index with _)
```

---

## 6. `range` over String

Yields the **byte offset** (not character number) and the **rune** (Unicode code point).

```go
for i, ch := range "hello,世界" {
    fmt.Printf("byte index=%d  char=%c\n", i, ch)
}
```

Output:
```
byte index=0  char=h
byte index=1  char=e
...
byte index=6  char=世   ← jumps by 3 (UTF-8 multibyte)
byte index=9  char=界
```

> The index is a **byte position**, not a character position. Multibyte characters (like Chinese) cause index jumps. Use `[]rune(s)` if you need character positions.

---

## 7. `range` over Map

```go
legs := map[string]int{"cat": 4, "bird": 2}
for animal, count := range legs {
    fmt.Printf("%s has %d legs\n", animal, count)
}
```

> Map iteration order is **random** every run by design. Never rely on order.

---

## 8. Nested Loops

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("%4d", i*j)
    }
    fmt.Println()
}
```

---

## 9. Labeled `break` — Exit an Outer Loop

Without a label, `break` only exits the innermost loop. A label lets you target an outer one.

```go
outer:
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            if i+j == 4 {
                break outer  // exits the outer loop entirely
            }
        }
    }
```

---

## 10. Labeled `continue` — Continue an Outer Loop

```go
loop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue loop  // skips rest of inner, continues outer
            }
        }
    }
```

---

## Quick Reference

| Pattern | Syntax |
|---|---|
| C-style | `for i := 0; i < n; i++` |
| While | `for condition` |
| Infinite | `for { }` |
| Range (slice) | `for i, v := range slice` |
| Range (string) | `for i, ch := range str` — i is byte offset |
| Range (map) | `for k, v := range m` — random order |
| Skip iteration | `continue` |
| Exit loop | `break` |
| Exit outer loop | `break label` |
| Continue outer loop | `continue label` |
