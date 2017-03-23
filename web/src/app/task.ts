import { ReplaySubject } from 'rxjs/ReplaySubject';

interface Model {
  toJson()
}

interface ITask {
  name: string
  event: IUnit
  actions: IUnit[]
}

export class Task implements ITask, Model {
  // Model properties
  name: string
  event: Unit
  actions: Unit[]

  // State properties
  units: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)

  constructor(fields?: {
    name?: string
  }) {
    if (fields) Object.assign(this, fields)
  }

  static fromJson(model: ITask): Task {
    let task = new Task()
    Object.assign(task, model)
    task.event = Unit.fromJson(model.event)
    task.actions = model.actions.map(action => Unit.fromJson(action))
    task.updateUnitList()
    return task
  }

  public toJson(): ITask {
    return {
      name: this.name,
      event: this.event.toJson(),
      actions: this.actions.map(a => a.toJson())
    }
  }

  private updateUnitList() {
    let units: Unit[] = []
    this.event && units.push(this.event)
    this.actions.forEach(action => {
      action && units.push(action)
    })
    this.units.next(units)
  }

  speak() {
    console.log("This is task", this.name)
  }
}

export class Unit implements Model {
  id: string
  name: string
  description: string
  input: Input[]
  output: Output[]

  static fromJson(model: IUnit): Unit {
    let unit = Object.create(Unit.prototype)
    Object.assign(unit, model)
    unit.input = model.input.map(input => new Input(input))
    unit.output = model.output.map(output => new Output(output))
    return unit
  }

  toJson(): IUnit {
    return {
      id: this.id,
      name: this.name,
      description: this.description,
      input: this.input.map(i => i.toJson()),
      output: this.output.map(o => o.toJson())
    }
  }
}

interface IUnit {
  id: string
  name: string
  description: string
  input: IInput[]
  output: IOutput[]
}

interface IInput {
  name: string;
  type: string;
  recipe: IIngredient[]
}

export class Input implements IInput, Model {
  name: string;
  type: string;
  recipe: Ingredient[]

  constructor(model: IInput) {
    Object.assign(this, model)
    this.recipe = model.recipe.map(r => Ingredient.fromJson(r))
  }

  toJson(): IInput {
    return {
      name: this.name,
      type: this.type,
      recipe: this.recipe.map(i => i.toJson())
    }
  }

  public isArray(): boolean {
      return this.type.startsWith("[]")
  }
}

interface IOutput {
  name: string;
  type: string;
}

export class Output implements IOutput, Model {
  name: string;
  type: string;

  constructor(model: IOutput) {
    Object.assign(this, model)
  }

  toJson(): IOutput {
    return {
      name: this.name,
      type: this.type,
    }
  }
}

interface IIngredient {
  type: number
  source: string
  value: string
}

export class Ingredient implements IIngredient, Model{
  type: number
  source: string
  value: string

  static fromJson(model: IIngredient): Ingredient {
    let ingredient = new Ingredient()
    Object.assign(ingredient, model)
    return ingredient
  }

  toJson(): IIngredient {
    return {
      type: this.type,
      source: this.source,
      value: this.value
    }
  }
}
