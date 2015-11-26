# Amigo 1.0 Specifiaction

**Origin Date**: 2015-11-26

**Author**: GoTsunami

## Overview

AMIGO LIA (Langage Intermédiaire d’Abstraction) is a simple data model abstraction language. Primarily this specification defines several keywords and the structures to define models, constraints and relationships. It uses YAML as data format.
It is designed to be simple!

## Conventions

Since LIA utilizes YAML, it has the same basic syntax system (see).

## Keywords

### First level
Any LIA file contains only two first level keywords:
    models: contains the list of models
    admin (optional): defines the administration system to create for the models

### In `models`

Models are created directly under `models` using user-defined names.
Under each model, the fields are defined using also user-defined names.
And, under each field, some keywords can be used to define it further.

Example:
```
models:
    myfirstmodel:
          myfirstfield:
               <list of keywords>
          mysecondfield:
``` 

A given number of keywords are used to define models and fields:
- `type`: The type of the field. Possible values are `text`, `number`, `point`, and `file`. It can also be the name of another user-defined model in the file.
- `constraints`: a YAML list declaring constraints on the field. Possible values are `positive`, `negative`, `not null`, `not empty`, `in`, `float|floating`, `int|integer`
- `quantity`: the number of times the field can exist in the model. Possible values `<number>`, `<number> to <number>`, or a YAML list of the two previous definitions.
- the syntax to define range of numbers for `in` is the same than `quantity` but can also be a YAML list of user-defined values (text for instance).

### In `admin`

TBW

### Default Values

`type` = no default, required keyword
`quantity` = 1
`parent` = 0 to n ? 1 to n ?
`constraints`:
    for type `number` = `float`
    for type `text` = no default
    for type `point` = `cartesian`
    for type `file` = no default

### Disambiguation

If, for some reason, one wants to use a reserved keyword as a field or a model name, one just needs to prefix it with an exclamation mark. Example: `!type`

## Advanced types

Some advanced types are defined in LIA for an easy use of complex data.

`timeserie`

## Examples

```
models:
    topping:
        price:
            type: number
            constraints:
                - positive
        name:
            type: text
    pizza:
        toppings:
            type: topping
            quantity: 3 to n
```

See directory examples for more complexity.


## References

Copyright (C) 2015 by GoTsunami

This document and translations of it may be used to implement Amigo, it may be copied and furnished to others, and derivative works that comment on or otherwise explain it or assist in its implementation may be prepared, copied, published and distributed, in whole or in part, without restriction of any kind, provided that the above copyright notice and this paragraph are included on all such copies and derivative works. However, this document itself may not bemodified in any way.

The limited permissions granted above are perpetual and will not be revoked.

This document and the information contained herein is provided "AS IS" and ALL WARRANTIES, EXPRESS OR IMPLIED are DISCLAIMED, INCLUDING BUT NOT LIMITED TO ANY WARRANTY THAT THE USE OF THE INFORMATION HEREIN WILL NOT INFRINGE ANY RIGHTS OR ANY IMPLIED WARRANTIES OF MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE.