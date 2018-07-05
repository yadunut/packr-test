# Packr Test

This repository was created to test out an error I got when using packr. 

## Usage

go get -d github.com/yadunut/packr-test/...

## To test issue

```bash
cd packr-test
packr build cmd/server/main.go
./main.go
```

### Expected

```
<h1>SMA Website</h1>

<h2>World</h1>
```

```bash
mv main ~/Desktop
cd ..
mv packr-test packr-tests
cd ~/Desktop
./main
```

### Expected

```
<h1>SMA Website</h1>

<h2>World</h1>
```

### Actual

```
0 files in ./templates
error parsing template: html/template: "index.tmpl" is undefined
```
