# epgu-generator

```
docker run --name epgu-generator -v /Users/alexis/workspace/rtlabs/epgu-generator/work:/work aazayats/epgu-generator \
    --config /work/config.yaml replication
```
# Config sample
```YAML
Log:
  Level: debug

Workers: 8
Registry: /work/registry.xlsx
Dir:
  Template: /work/templates
  Artefact: /work/artefact
  Tmp: /work/tmp
