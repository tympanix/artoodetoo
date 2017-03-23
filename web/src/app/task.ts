export class Task implements ITask {
  name: string
  event: Unit
  actions: Unit[]

  constructor(model: ITask) {
    Object.assign(this, model)
    this.event = new Unit(model.event)
    this.actions = model.actions.map(action => new Unit(action))
  }

  speak() {
    console.log("This is task", this.name)
  }
}

interface ITask {
  name: string
  event: Unit
  actions: Unit[]
}

export class Unit {
  id: string
  name: string
  description: string
  input: Input[]
  output: Output[]

  constructor(model: IUnit) {
    Object.assign(this, model)
    this.input = model.input.map(input => new Input(input))
    this.output = model.output.map(output => new Output(output))
  }
}

interface IUnit {
  id: string
  name: string
  description: string
  input: Input[]
  output: Output[]
}

export class Input {
  name: string;
  type: string;
  recipe: Ingredient[]

  constructor(model: Input) {
    Object.assign(this, model)
  }

  public isArray(): boolean {
      return this.type.startsWith("[]")
  }
}

class Output {
  name: string;
  type: string;

  constructor(model: Output) {
    Object.assign(this, model)
  }
}

export class Ingredient{
  type: number
  source: string
  value: string
}
