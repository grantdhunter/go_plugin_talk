Go 1.8 lightning talk
====================
[https://golang.org/pkg/plugin/](https://golang.org/pkg/plugin/)

- A plugin is a shared object that can be loaded at run time.
- Plugins are cached after being initialized.
- Safe to use in multiple goroutines
- **Currently only works on linux** :(

### Load a plugin like this:
```go
p, err := plugin.Open("plugin_name.so")
if err != nil {
	panic(err)
}
```

### Access exported variables and functions (Symbols) like this:
```go
symbol, err := p.Lookup("Symbol")
if err != nil {
	panic(err)
}
```

### building a pluging:
```bash
go build -buildmode=plugin
```

Example
=======
Here is an example of an RPC server that loads up the correct procedure lazily.


```bash
go build -buildmode=plugin -o plugins/hello.so plugins/hello.go
go build -buildmode=plugin -o plugins/do_it.so plugins/do_it.go
go run main.go
```

Play around with removing the `.so` files and hitting these urls

```bash
curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -H "Postman-Token: 8c453945-6848-cf8c-4b27-eac133b10c3c" -d '{
	"Type": "do_it",
	"Args": [
		"12",
		"True",
		"wat is happens"
		]
	
}' "http://127.0.0.1:8080/"
```

```bash
curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -H "Postman-Token: 0e968e38-6fcb-3ab0-095d-d0d948befe91" -d '{
	"Type": "hello",
	"Args":["Edmonton Go!"]
}' "http://127.0.0.1:8080/"
```
