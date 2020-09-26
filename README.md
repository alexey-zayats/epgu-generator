# epgu-generator

Create folders for work path 

- /Users/alexis/workspace/rtlabs/epgu-generator/work
- .../artefact
- .../tmp

Copy `registry.xlsx` to work/artefact

Write `config.yaml` to work/config.yaml

### Config sample
```YAML
Log:
  Level: debug

Workers: 8
Registry: /work/registry.xlsx
Dir:
  Template: /app/templates
  Artefact: /work/artefact
  Tmp: /work/tmp
```

## Run

```
docker run --rm --name epgu-generator \
    -v /Users/alexis/workspace/rtlabs/epgu-generator/work:/work \
    aazayats/epgu-generator --config /work/config.yaml replication
```
