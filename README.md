# JIS X 0208 Converter in Go

## How to use

```go
ch, err := jisx0208.Rune(0x2422)
if err != nil {
       panic(err)
}
fmt.Printf("%c\n", ch) // あ
```
