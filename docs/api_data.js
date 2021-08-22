define({ "api": [
  {
    "type": "get",
    "url": "http://localhost:8080/cryptocurrency/",
    "title": "",
    "name": "Просмотр_валюты",
    "group": "Криптовалюта",
    "description": "<p>Просмотр Криптовалюты, обноваляется в базе каждый цикл интервала</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "coin_id",
            "description": "<p>Вид просматриваемой валюты</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "rank",
            "description": "<p>Позиция валюты на мировой криптобирже</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "symbol",
            "description": "<p>символическое обожначение</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "interval",
            "description": "<p>Интервалы между обновлением валюты в секундах</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "priceUsd",
            "description": "<p>цена валюты в переводе в USD</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "        [\n\t\t\t{\n\t\t\t\t\"_id\": \"611a7074e1d840f625b58c92\",\n\t\t\t\t\"coin_id\": \"bitcoin\",\n\t\t\t\t\"interval\": 30,\n\t\t\t\t\"priceusd\": \"46439.5197486433777590\",\n\t\t\t\t\"rank\": \"1\",\n\t\t\t\t\"symbol\": \"BTC\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"_id\": \"611a70e5e1d840f625b58f1e\",\n\t\t\t\t\"coin_id\": \"ethereum\",\n\t\t\t\t\"interval\": 1,\n\t\t\t\t\"priceusd\": \"3228.3628716937351608\",\n\t\t\t\t\"rank\": \"2\",\n\t\t\t\t\"symbol\": \"ETH\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"_id\": \"611a71cad88a76ef6f00de5f\",\n\t\t\t\t\"coin_id\": \"bitcoin\",\n\t\t\t\t\"interval\": 30,\n\t\t\t\t\"priceusd\": \"46439.5197486433777590\",\n\t\t\t\t\"rank\": \"1\",\n\t\t\t\t\"symbol\": \"BTC\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"_id\": \"611a71d2c9a9566d5c98fb02\",\n\t\t\t\t\"coin_id\": \"bitcoin\",\n\t\t\t\t\"interval\": 30,\n\t\t\t\t\"priceusd\": \"46439.5197486433777590\",\n\t\t\t\t\"rank\": \"1\",\n\t\t\t\t\"symbol\": \"BTC\"\n\t\t\t}\n\t\t]",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "api/crypto_currency.go",
    "groupTitle": "Криптовалюта"
  },
  {
    "type": "POST",
    "url": "api/batches/",
    "title": "Создание партии",
    "name": "Создания_статистики_по_криптовалюте",
    "group": "Криптовалюта",
    "description": "<p>Создания статистики по криптовалюте</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "coin_id",
            "description": "<p>Вид просматриваемой валюты</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "interval",
            "description": "<p>Интервалы между обновлением валюты в секундах</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\n\t\"coin_id\": \"bitcoin\",\n\t\"interval\": 30\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "HTTP/1.1 200 OK\n{\n\t\"InsertedID\": \"611a8824e450d2183ab5f9a2\"\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "500 BAD REQUEST": [
          {
            "group": "500 BAD REQUEST",
            "type": "Object",
            "optional": false,
            "field": "errors",
            "description": "<p>List of errors</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ValidationErrors:",
          "content": "{\n\t{\n\t\t\"message\": \"Failed to decode JSON object: Expecting value: line 1 column 1 (char 0)\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "api/crypto_currency.go",
    "groupTitle": "Криптовалюта"
  },
  {
    "type": "delete",
    "url": "/api/batches/<batch_id>/",
    "title": "Удаление мониторинга за криптовалютой",
    "name": "Удаление_мониторинга_за_криптовалютой",
    "group": "Криптовалюта",
    "description": "<p>Удаление мониторинга за криптовалютой по ID</p>",
    "error": {
      "fields": {
        "404 NOT FOUND": [
          {
            "group": "404 NOT FOUND",
            "optional": false,
            "field": "ID",
            "description": "<p>обьекта некорректен</p>"
          },
          {
            "group": "404 NOT FOUND",
            "type": "string",
            "optional": false,
            "field": "errors.common",
            "description": "<p>Common message</p>"
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\tobjID: ObjectID(\"611a8824e450d2183ab5f9a2\")\n\tUserValue: 611a8824e450d2183ab5f9a2\n\t{\"DeletedCount\":0}\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "api/crypto_currency.go",
    "groupTitle": "Криптовалюта"
  }
] });
