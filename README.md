# golang-overlay-experiment

There's a little-known `-overlay` option in `go build` and `go test`.
It instructs the tool to combine the *actual* filesystem content with content from elsewhere as specifyed by a JSON file.
It is possible to inject completely new files, replace existing files' content, and to hide existing files as if they weren't there.

A common use case for overlays is to allow editors to pass in a set of unsaved, modified files.

Another use case I am excited about is to enable code generation as a part of `go build`.
Imagine a wrapper command, let's call it `bpfgo`, which one invokes as usual, e.g. `bpfgo build`, `bpfgo test`.
The tool will *transparently* compile eBPF C code. 
It will invoke go toolchain and it will inject the compiled eBPF code via overlay.
Why is it exciting?
You won't have to add generated binary files produced by C compiler to the repository ever again.
The wrapper is as easy to use as `go build` command itself. No need to dabble in Makefiles!

## Repository content
There's a simple command line tool at `./cmd/main`.

```sh
% go run ./cmd/main
0
Hello, world!
```

We also have a `Makefile` that allows to conveniently invoke the toolchain with an overlay.
The output is now different, even though we didn't change `./cmd/main`.

```sh
% make run
go run -overlay overlay.json ./cmd/main
42
Hello, world!
```

**Note:** apparently, `go:embed` doesn't respect the overlay as the second line in the output didn't change.
It was supposed to change to `Hello, overlay!` as we have the following line in `overlay.json`:
```
"cmd/main/greeting": "overlay/cmd/main/greeting
```

## Testing
We can invoke `go test` with an overlay as well:
```sh
% make test
go test -overlay overlay.json ./cmd/main
ok  	example.com/overlay_test/cmd/main	0.158s
```

The really nice thing worth noting is that overlay doesn't interfere with caching.
If we invoke `make test` again, the results are coming from the cache.
```
% make test
go test -overlay overlay.json ./cmd/main
ok  	example.com/overlay_test/cmd/main	(cached)
```
