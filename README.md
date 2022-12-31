# unicode2en
ASCII transliterations of Unicode text to English.


### Build 

- Windows
```
GOOS=windows GOARCH=amd64 go build .
```

- Linux
```
GOOS=linux GOARCH=amd64 go build .
```


### Usage

- names.lst
```
García
Núñez
Pérez
```

```
./unidecode2en -f names.lst
```

```
./unidecode2en -s 北京
```
