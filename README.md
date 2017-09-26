# makedesc

Makedesc is a tool for getting a list of targets for a Makefile.

## Install

```
go get 4d63.com/makedesc
```

### Mac

```
curl -o /usr/local/bin/makedesc https://raw.githubusercontent.com/leighmcculloch/makedesc/binaries/mac/amd64/makedesc && chmod +x /usr/local/bin/makedesc
```

### Linux

```
curl -o /usr/local/bin/makedesc https://raw.githubusercontent.com/leighmcculloch/makedesc/binaries/linux/amd64/makedesc && chmod +x /usr/local/bin/makedesc
```

### Windows

[Download the executable](https://raw.githubusercontent.com/leighmcculloch/makedesc/binaries/windows/amd64/makedesc.exe), and save it to your path.

## Usage

```
$ makedesc -help
Makedesc is a tool for describing makefiles.

Usage:

  makedesc [-file=Makefile]

  Flags:

    -file string
          a Makefile (default "Makefile")
    -help
	  print this help list
```

