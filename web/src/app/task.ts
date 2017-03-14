export class Task {
  name: string
  event: TaskUnit
  actions: TaskUnit[]
}

class TaskUnit {
  id: string
  name: string
  recipe: Recipe[]
}

class Recipe{
  type: number
  argument: string
  source: string
  value: string
}
