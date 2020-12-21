# Building ioquake3 on Windows

This guide will walk you through building a recent version of ioquake3 on Windows 10.

## Prerequisites

- A working [MSYS2](https://www.msys2.org/) installation

## Installing dependencies

We will be building ioquake3 using "MSYS2 MinGW 32-bit". Run "MSYS2 MinGW 32-bit" from your Start menu.

Install the following dependencies:
- `make`
- `git`
- `mingw-w64-i686-SDL`
- `mingw-w64-i686-curl`
- `mingw-w64-i686-gcc`
- `mingw-w64-i686-libjpeg-turbo`
- `mingw-w64-i686-openal`
- `bison`

```shell
$ pacman -S make git mingw-w64-i686-SDL mingw-w64-i686-curl mingw-w64-i686-gcc mingw-w64-i686-libjpeg-turbo mingw-w64-i686-openal bison
```

## Clone ioquake3

Clone the repository and optionally checkout to a known working commit. I used `d1b7ab6b22cc99205ac890910e286859e30df40e`.

```shell
$ git clone https://github.com/ioquake3/ioq3 && cd ioq3
$ git checkout d1b7ab6b22cc99205ac890910e286859e30df40e
```

## Build

We will build using 8 CPU cores, force the `mingw32` platform and `x86` architecture, and manually specify the path to `windres.exe`.

```shell
$ make PLATFORM=mingw32 ARCH=x86 WINDRES=windres -j8
```

Once the build is done, your executable should be in `build/release-mingw32-x86`.