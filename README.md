# cgo-play
playtesting cgo with a view to using it for abstracting an existing C/C++ API

## goldfish

demo example from blog - note doesn't work if leave hello.c in place

## pocketVNA

supplied with `libPocketVnaApi_x64.so.0` but renamed to `libPocketVnaApi_x64.so` in the local directory, for successful linking in with the emacs compile command. When running with `go run main.go` it needs `libPocketVnaApi_x64.so.0` in `/usr/lib` (but no need to update ldd).

So there  is some path stuff being done differently with each compilation method, that could be a trap for later when there are different versions of the shared library in the local directory and in `/usr/lib`. 



