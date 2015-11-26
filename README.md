## Amigo

The awesome, user-friendly admin interface.

### Specifications

Read the [specs](SPECS.md).

### YAML specific language

#### Examples

- [artist.yaml](examples/artist.yaml)
- [earthquake.yaml](examples/earthquake.yaml)
- [galaxy.yaml](examples/galaxy.yaml)

#### structural keywords

Field names may start with an exclamation mark `!` to disambiguate the meaning: they will be used a field names, not structural keywords.

- models
  
- admin
  - model
  - list
    - display
    - filter
    - search
    - editable
    - actions
  - form
    - fieldsets
      - title
      - fields
    - inlines
    - actions

### Leaves for Models

- `type`: possible values `text`, `number`, `point`, `file`
- `constraints`: possible values `positive`, `negative`, `not null`, `not empty`, `in`, `float|floating number`, `int|integer`
- `quantity` and `in`: possible values `<number>`, `<number> to <number>` == `<range>`, `<number|range> or <number|range> or ...`
- `choices`: YAML list

### Rules

#### Defaults values

- `type` has no default value
    - for `number`, default `constraints` is `float`
- `quantity` default value is 1
