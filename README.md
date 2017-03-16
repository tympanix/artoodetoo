# Automato

## API
### `/api/units`, `GET`
Returns a list of available units

### `/api/tasks/{taskname}`, `POST`
Executes the task with the given `taskname`

### `/api/tasks/{taskname}`, `DELETE`
Deletes the task with `taskname` and removes it from the application

### `/api/tasks`, `PUT`
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
    "input": [],
    "output": [
      {
        "name": "Name",
        "type": "string"
      },
      {
        "name": "Age",
        "type": "int"
      },
      {
        "name": "Height",
        "type": "float32"
      },
      {
        "name": "Married",
        "type": "bool"
      }
    ]
  },
  "actions": [
    {
      "id": "example.StringConverter",
      "name": "String",
      "input": [
        {
          "name": "Format",
          "type": "string",
          "recipe": [
            {
              "type": 1,
              "value": "Hello my name is %s"
            }
          ]
        },
        {
          "name": "Placeholder",
          "type": "interface {}",
          "recipe": [
            {
              "type": 0,
              "source": "Person",
              "value": "Name"
            }
          ]
        }
      ],
      "output": [
        {
          "name": "Formatted",
          "type": "string"
        }
      ]
    },
    {
      "id": "example.EmailAction",
      "name": "Email",
      "input": [
        {
          "name": "Message",
          "type": "string",
          "recipe": [
            {
              "type": 0,
              "source": "String",
              "value": "Formatted"
            }
          ]
        },
        {
          "name": "Subject",
          "type": "string",
          "recipe": [
            {
              "type": 1,
              "value": "A new friend"
            }
          ]
        },
        {
          "name": "Receiver",
          "type": "string",
          "recipe": [
            {
              "type": 1,
              "value": "johndoe@email.com"
            }
          ]
        }
      ]
    }
  ]
}
```
