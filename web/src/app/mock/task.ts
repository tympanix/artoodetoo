import { Task } from '../task';

export const TASK = {
  "name": "MyFirstTask",
  "event": {
    "id": "example.PersonEvent",
    "name": "Person",
		description: "string",
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
			description: "string",
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
			description: "string",
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
