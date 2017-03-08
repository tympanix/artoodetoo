# Automato

## API
### `/api/units`, `GET`
Returns a list of available units

### `/api/newtask`, `POST`
Executes a new task immediately

Example post body:
```json
{
	"name": "MyFirstTask",
	"event": {
		"id": "example.PersonEvent",
		"name": "Person",
		"recipe": []
	},
	"actions": [{
		"id": "example.StringConverter",
		"name": "String",
		"recipe": [{
			"type": 1,
			"argument": "Format",
			"source": "",
			"value": "Person %s would like to say hello"
		}, {
			"type": 0,
			"argument": "Placeholder",
			"source": "Person",
			"value": "Name"
		}]
	},{
		"id": "example.EmailAction",
		"name": "Email",
		"recipe": [{
			"type": 0,
			"argument": "Message",
			"source": "String",
			"value": "Formatted"
		}, {
			"type": 1,
			"argument": "Subject",
			"source": "",
			"value": "A new friend"
		}, {
			"type": 1,
			"argument": "Receiver",
			"source": "",
			"value": "johndoe@email.com"
		}]
	}]
}
```
