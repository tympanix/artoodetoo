# Automato

## API
### `/api/units`, `GET`
Returns a list of available units

### `/api/tasks/{taskname}`, `POST`
Executes the task with the given `taskname`

### `/api/tasks/{taskname}`, `DELETE`
Deletes the task with `taskname` and removes it from the application

### `/api/tasks}`, `PUT`
Updates task. The html body should contain the complete task in json format, the name of the task can't be changed as it is used to identify the task. 

### `/api/tasks`, `GET`
Returns an array of the tasks registered in the application

### `/api/tasks`, `POST`
Adds a new task to the application.

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
