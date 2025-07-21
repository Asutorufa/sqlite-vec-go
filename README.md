# Building `sqlite-vec` for `ncruces/go-sqlite3`

This demonstrates building [`sqlite-vec`](https://github.com/asg017/sqlite-vec) for use with [`ncruces/go-sqlite3`](https://github.com/ncruces/go-sqlite3).

To build and test `sqlite-vec.wasm`:
```sh
./build.sh
go run .
```

See [`main.go`](main.go) for how to use `sqlite-vec.wasm`.

Update [`build.sh`](build.sh) to use specific versions of `sqlite-vec` or `ncruces/go-sqlite3`.