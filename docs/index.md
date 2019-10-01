# 💉 Orninject – Helps to easily replace environment variables in file

## Getting started

### Simple replacement way 1

Pipe:

```bash
export TITLE=world
echo "Hello {{ title .TITLE }}" | orninject
$ Hello World
```

Pipe write:

```bash
export TITLE=Hello world
cat index.html
$ <h1>{{ upper .TITLE }}</h2>
cat index.html | orninject > index.html
cat index.html
$ <h1>HELLO WORLD</h2>
```

### Simple replacement way 2

File `index.html`:

```html
<h1>{{ upper .TITLE }}</h2>
```

Command:

```bash
export TITLE=Hello world
orninject replace index.html
```

New `index.html` content:

```html
<h1>HELLO WORLD</h2>
```

### Advanced replacement

File `index.html`:

```html
<h1>{{ .TITLE }}</h1>
<h2>{{ default .NO_VAR "default var value" | title }}</h2>
```

Command:

```bash
export TITLE=Hello world
orninject replace index.html
```

New `index.html` content:

```html
<h1>Hello world</h1>
<h2>Default Var Value</h2>
```

## Additional template features

Orninject use [text/template](https://golang.org/pkg/text/template/) and [Sprig](http://masterminds.github.io/sprig/) lib to provide more template features.
