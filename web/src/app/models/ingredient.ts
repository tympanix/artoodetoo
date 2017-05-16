import { Model } from './model'
import { Input, Output } from './io'
import { Unit } from './unit'

export interface IIngredient {
  type: number
  source: string
  value: string
}

export class Ingredient implements IIngredient, Model{
  // Model properties
  type: number
  source: string
  value: string

  input: Input
  reference: Output

  constructor() {
    this.type = 1
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
    var output = unit.output.find(Input.findByName(this.value))
    if (!output) console.error("Unknown ingredient variable", this)
    this.reference = output
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
