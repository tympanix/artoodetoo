import { Model } from './model'

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
