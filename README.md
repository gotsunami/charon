## Amigo

The awesome, user-friendly admin interface.

### YAML specific language

#### global structure

```
---
!models:
  galaxie:
    nom:
      type: text
      quantity: 1
    position:
      type: point
      quantity: 1
    luminosité:
      type: floating number
      quantity: 2 or 4
      error:
        type: floating number
        quantity: 0 to 1
    images:
      type: image
      quantity: 0 to 5
  image:
    fichier:
      type: file
      quantity: 1
    bande:
      type: choice
      quantity: 1
      choices:
      - u
      - g
      - r
      - i
      - z
  amas:
    nom:
      type: text
      quantity: 1
    nombre_de_galaxies:
      type: integer
      constraint: positive
      quantity: 0 to 1
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

- `type`: possible values `text`, `floating number|float`, `integer|int`, `file`
- `constraint`: possible values `positive`, `negative`, `not null`, `not empty`
- `quantity`: possible values `<number>`, `<number> to <number>`, `<number> or <number> or ...`
- `choices`: YAML list
