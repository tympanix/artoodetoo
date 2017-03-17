export class Task {
  name: string
  event: Unit
  actions: Unit[]
}

export class Unit {
  id: string
  name: string
  description: string
  input: Input[] = []
  output: Output[] = []
}

class Input {
  name: string;
  type: string;
  recipe: Ingredient[]

  public isArray(): boolean {
      return this.type.startsWith("[]")
  }
}

class Output {
  name: string;
  type: string;
}

class Ingredient{
  type: number
  source: string
  value: string
}
