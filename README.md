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
      qty: '1'
    position:
      type: point
      qty: '1'
    luminosité:
      type: floating number
      qty: '0,1'
      error:
        type: floating number
        qty: '0,1'
    images:
      type: image
      qty: '0,5'
  image:
    fichier:
      type: file
      qty: '1'
    bande:
      type: choice
      qty: '1'
      choices:
      - u
      - g
      - r
      - i
      - z
  amas:
    nom:
      type: text
      qty: '1'
    nombre_de_galaxies:
      type: integer
      constraint: positive
      qty: '0,1'

#### structural keywords
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
```

### leaves for models

- type: possible values `text`, `floating number|float`, `integer|int`, `file`
- constraint: possible values `positive`, `negative`, `not null`, `not empty`
- qty: possible values `<number>`, `<number> to <number>`, `<number> or <number> or ...`
- choices: YAML list