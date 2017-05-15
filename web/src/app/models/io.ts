import { Model } from './model'
import { Ingredient, IIngredient } from './ingredient'
import { Unit } from './unit'

export interface IInput {
  name: string;
  type: string;
  recipe: IIngredient[]
}

export class Input implements IInput, Model {
  // Model properties
  name: string;
  type: string;
  recipe: Ingredient[]

  // State properties
  unit: Unit

  constructor(model: IInput) {
    Object.assign(this, model)
    this.recipe = model.recipe ? model.recipe.map(r => Ingredient.fromJson(r)) : []
    this.recipe.forEach(i => i.bindToInput(this))
  }

  toJson(): IInput {
    return {
      name: this.name,
      type: this.type,
      recipe: this.recipe.map(i => i.toJson())
    }
  }

  bindToUnit(unit: Unit) {
    this.unit = unit
  }

  bootstrap() {
    if (!this.recipe || !this.recipe.length) {
      this.recipe = [new Ingredient()]
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

  // State properties
  unit: Unit

  constructor(model: IOutput) {
    Object.assign(this, model)
  }

  bindToUnit(unit: Unit) {
    this.unit = unit
  }

  toJson(): IOutput {
    return {
      name: this.name,
      type: this.type,
    }
  }
}
