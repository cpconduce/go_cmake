# go_cmake

Let's say you're a [CMake](http://www.cmake.org/) shop. Go on to
assert that you'd like to use [go](http://golang.org/) for a project
or two. Go has a very well thought out build system. Can it be made to
play nice with CMake? Here's an attempt.

## Instructions

In your top level [`CMakeLists.txt`](CMakeLists.txt) file, include
[`GolangSimple.cmake`](bld/cmake/GolangSimple.cmake):

    cmake_minimum_required(VERSION 3.0)
    include(bld/cmake/GolangSimple.cmake)
    add_subdirectory(some/subdirectory)

In `some/subdirectory` lies your go program. Let's say it's a simple
clint to list out the keys in a redis store. It depends upon a `go get`
provided module. Here's what you'd need in your subdirectory's
[`CMakeLists.txt`](some/subdirectory/CmakeLists.txt) file:

    GO_GET(go_redis github.com/hoisie/redis)
    ADD_GO_INSTALLABLE_PROGRAM(redis_lister # executable name
                               redis_lister.go # `package main` source file
                               go_redis) # everything else is a dependency
    
This will create a binary called
`$CMAKE_CURRENT_BINARY_DIR/some/subdirectory/redis_lister`. When you build
the `install` target it will put said binary in `$PREFIX/bin/redis_lister`.

If you want to set build flags, you can use the CMake variable
`CMAKE_GO_FLAGS`. If you want to include modules local to
`redis_lister.go`, include them in the `"./module"` style. Be sure
to put your top level source file in the package `main`.

