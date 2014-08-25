## Specifications

Example:

    {
        "models": {
            "establishment": {
                "address": {
                    "type": "address"
                }
            },
            "address": {
                "zipcode": {
                    "type": "text",
                    "exact_lentgh": 5
                }
            },
            "firm": {
                "investors": {
                    "type": "investor",
                    "qty": "0,n"
                }
            },
            "investor": {
                "firms": {
                    "type": "firm",
                    "qty": "1,2"
                }
            }
        },
        "admin": [{
            "model": "firm",
            "list": {
                "display": [],
                "filter": [],
                "search_fields": [],
                "searchform_fields": [],
                "editable": [],
                "actions": ["print", "params"]
            },
            "form": {
                "fieldsets": [
                    {
                        "title": "A title",
                        "fields": "(c1,c2),c3,c4",
                        "classes": "myfield"
                    }
                ],
                "media": "",
                "template": "",
                "inlines": [],
                "actions": []
            }
        }]
    }
