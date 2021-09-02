# Templite

Templite is a lite weight templating tool written in golang. It is intended to bridge the gap between kustomize and Helm in a relatively small and simple package.

## Building

```shell
make build
```

## Usage

```shell
# values from file
bin/templite --template test/basic/t.tmpl --file test/basic/t.yaml
# values from file (short flag)
bin/templite --template test/basic/t.tmpl -f test/basic/t.yaml

# values from stdin
cat test/basic/t.yaml | bin/templite --template test/basic/t.tmpl --file -
# values from stdin (short flag)
cat test/basic/t.yaml | bin/templite --template test/basic/t.tmpl -f -
```
