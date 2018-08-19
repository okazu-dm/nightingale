# Nightingale
Notify something to somewhere.

# Usage
```
cat sample.json | nightingale [flags]
```

## Flags:
```
Flags:
  -c, --config string          config file (default "config.yml")
  -h, --help                   help for nightingale
      --slack_url string       Slack Webhook URL to notify something.
  -p, --template_path string   Template file path (default "default")
  -t, --type string            Channel to notify something.(stdout/slack) (default "stdout")
```

## Sample
`sample.json`
```json
{
    "Name": "Test",
    "Slice": [3, 5, 7]
}
```
`sample.tmpl`
```gotemplate
Hello {{.Name}}
length: {{len .Slice}}
```

```sh
> cat sample.json | ./nightingale -p sample.tmpl  -t stdout
Hello Test
length: 3
```

or you use yaml/json instead of command line options
```yaml
template_path: sample.tmpl
type: stdout
```