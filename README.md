# Packr Test

This repository was created to test out an error I got when using packr. 

## Usage

```Bash
go get -d github.com/yadunut/packr-test/...
```

or 

```
git clone https://github.com/yadunut/packr-test
```

## No Issues

```bash
cd packr-test
packr build cmd/server/main.go
./main
```

### Expected

```
Init Function
2 files in ./templates
ListFiles function
Files in box:
index.tmpl
register.tmpl
```

### Actual

```
Init Function
2 files in ./templates
ListFiles function
Files in box:
index.tmpl
register.tmpl
```

## Issues

Changed the path to the repository

```bash
cd ..
mv packr-test packr-tests
cd packr-tests
./main
cd ..
mv packr-tests packr-test
cd packr-test
```

### Expected

```
Init Function
2 files in ./templates
ListFiles function
Files in box:
index.tmpl
register.tmpl
```

### Actual

```
Init Function
0 files in ./templates
ListFiles function
Files in box:
index.tmpl
register.tmpl
```

## What Happened?

The init function doesn't detect files in the binary, but detects files when they're in the folders. 

Outside of the init function, there are no issues. 
