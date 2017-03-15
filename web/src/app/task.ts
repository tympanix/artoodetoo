export class Task {
  name: string
  event: Unit
  actions: Unit[]
}

export class Unit {
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
