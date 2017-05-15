import { Model } from './model'
import { Input } from './io'

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

  constructor() {
    this.type = 1
  }

  static fromJson(model: IIngredient): Ingredient {
    let ingredient = new Ingredient()
    Object.assign(ingredient, model)
    return ingredient
  }

  bindToInput(input: Input) {
    this.input = input
  }

  toJson(): IIngredient {
    return {
      type: this.type,
      source: this.source,
      value: this.value
    }
  }
}
