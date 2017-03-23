import { Model } from './model'
import { Ingredient, IIngredient } from './ingredient'

export interface IInput {
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
    this.recipe = model.recipe ? model.recipe.map(r => Ingredient.fromJson(r)) : []
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

export interface IOutput {
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
