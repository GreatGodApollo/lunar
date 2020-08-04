<h1 align="center">Lunar</h1>
<p align="center"><i>Made with :heart: by <a href="https://github.com/GreatGodApollo">@GreatGodApollo</a></i></p>

GoLang Spacebin CLI

## Installation
### Standard Download
Just head on over to the [releases](https://github.com/GreatGodApollo/lunar/releases) page and download the latest release
for your platform. Extract it using something like [7-Zip](https://www.7-zip.org) for Windows or `tar` on other 
platforms (`tar -zxvf lunar*.tar.gz`).

That's it! Although you'll probably want to also add the binary to your path for ease of use.

### Scoop - Coming Soon

### Go Get
Do you have go installed? You can run just one simple command to install lunar!
```shell
$ go get -u github.com/GreatGodApollo/lunar
```

## Usage
```shell
$ lunar --help
    Lunar is a CLI for Spacebin that allows you to easily make documents.
    This application can be used in a couple of different ways
    to quickly create a document on an instance.
    
    You can either pipe a document into lunar by doing:
    'command | lunar'
    
    or upload a document directly:
    'lunar -f file.txt'
    
    Usage:
      lunar [flags]
    
    Flags:
          --config string       config file (default is $HOME/.lunar.yaml)
      -c, --copy                copy the url to your clipboard
      -f, --file string         the file to upload
      -h, --help                help for lunar
      -i, --instance string     the spacebin instance (default "https://api.spaceb.in")
      -r, --raw                 do you want the raw url
          --result-url string   the base url for response (default "https://spaceb.in")
      -v, --version             version for lunar
```

### Configuration
You can preset `instance` and `result-url` in `$HOME/.lunar.yaml`

## Built With
* [GoSpacebin](https://github.com/GreatGodApollo/gospacebin)
* [Cobra](https://github.com/spf13/cobra)
* [Viper](https://github.com/spf13/viper)
* [Chalk](https://github.com/ttacon/chalk)

## Licensing

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/)

## Authors

* [Brett Bender](https://github.com/GreatGodApollo)