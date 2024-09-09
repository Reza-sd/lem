# lem
https://cdn.intra.42.fr/pdf/pdf/947/lem-in.en.pdf

gofmt -l -w .
go test -v -failfast
go test -failfast
golint ./...

L1 =[1-3-4-0]
L2 =[1-2-5-6-0]
L3 = [""-1-3-4-0]
# test mio miooo


# strings.Builder
-------------------------------------------------------
var me1_test = map[string]any{
	"return 3 if...": func() []any {
		skip := false
		inp1 := 1
		inp2 := 5
		//inp := Inp(inp1, inp2)
		got := (inp2 - inp1)
		exp := 4
		return []any{skip, got, exp, inp1, inp2}
	},
}
---------------------------------------------------