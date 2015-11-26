## Amigo

The awesome, user-friendly admin interface.

### YAML specific language

#### global structure

```
---
models:
  galaxie:
    nom:
      type: text
    position:
      type: point
    luminosité:
      type: number
      quantity: 2 or 4
      constraints:
        - float
        - in:
            - 0 to 1
            - 5 to 6
      erreur:
        type: number
        constraints:
            - float
            - in: 0 to 1
        quantity: 0 to 1
    images:
      type: image
      quantity: 0 to 5
  image:
    fichier:
      type: file
    bande:
      type: text
      constraints:
        - in:
            - u
            - g
            - r
            - i
            - z
  amas:
    nom:
      type: text
    nombre_de_galaxies:
      type: number
      constraints:
        - positive
        - integer
      quantity: 0 or 1
    galaxies:
        type: galaxie
        quantity: 0 to n
```

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
