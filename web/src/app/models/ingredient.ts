import { Model } from './model'
import { Input, Output } from './io'
import { Unit } from './unit'
import { Task } from './task'
import { ReplaySubject } from 'rxjs/ReplaySubject';

export interface IIngredient {
  type: number
  source: string
  value: Object
}

export class Ingredient implements IIngredient, Model{
  // Model properties
  type: number
  source: string
  value: Object

  input: Input
  reference: Output

  constructor(value?: Object) {
    this.type = 1
    this.value = value
  }

  static fromJson(model: IIngredient): Ingredient {
    let ingredient = new Ingredient()
    Object.assign(ingredient, model)
    return ingredient
  }

  resolveReference(units: Unit[]) {
    if (this.isStatic()) return
    var unit = units.find(Unit.findByName(this.source))
    if (!unit) console.error("Unknown ingredient source", this)
    var output = unit.output.find(Input.findByName(this.value.toString()))
    if (!output) console.error("Unknown ingredient variable", this)
    this.reference = output
  }

  setVariable(output: Output) {
    if (output.getTask() !== this.getTask()) {
      console.error("Setting invalid ingredient variable", output)
    }
    this.reference = output
    this.source = output.unit.name
    this.value = output.name
  }

  getTask(): Task {
    try {
      return this.input.unit.task
    } catch(e) {
      return null
    }
  }

  isVariable(): boolean {
    return this.type === 0
  }

  isStatic(): boolean {
    return this.type === 1
  }

  bindToInput(input: Input) {
    this.input = input
  }

  bootstrap(input?: Input) {
    input && this.bindToInput(input)
  }

  toJson(): IIngredient {
    return {
      type: this.type,
      source: this.source,
      value: this.value
    }
  }
}
