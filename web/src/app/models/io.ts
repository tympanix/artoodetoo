import { Model } from './model'
import { Ingredient, IIngredient } from './ingredient'
import { Unit } from './unit'
import { Task } from './task'

function hasPrefix(prefix: string, ...str: string[]) {
  return str.every((s) => s.startsWith(prefix))
}

export class IO {
  name: string
  type: string

  // State properties
  unit: Unit

  bindToUnit(unit: Unit) {
    this.unit = unit
  }

  assignableTo(other: IO): boolean {
    if (this.type == other.type) {
      return true;
    } else if (this.type == "interface{}") {
      return true;
    } else if (hasPrefix("int", this.type, other.type)) {
      return true;
    } else if (hasPrefix("float", this.type, other.type)) {
      return true;
    } else {
      return false;
    }
  }

  getTask(): Task {
    try {
      return this.unit.task
    } catch(e) {
      return null
    }
  }

  static findByName(name: string): (i: Input) => boolean {
    return function(i: Input) {
      return i.name === name
    }
  }



}

export interface IInput {
  name: string;
  type: string;
  recipe: IIngredient[]
}

export class Input extends IO implements IInput, Model {
  // Model properties
  name: string;
  type: string;
  recipe: Ingredient[]

  constructor(model: IInput) {
    super()
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

  bootstrap(unit?: Unit) {
    unit && this.bindToUnit(unit)
    if (!this.recipe || !this.recipe.length) {
      this.recipe = [new Ingredient()]
    }
    this.recipe.forEach(i => i.bootstrap(this))
  }

  public isArray(): boolean {
      return this.type.startsWith("[]")
  }
}

export interface IOutput {
  name: string;
  type: string;
}

export class Output extends IO implements IOutput, Model {
  name: string;
  type: string;

  constructor(model: IOutput) {
    super()
    Object.assign(this, model)
  }

  bootstrap(unit?: Unit) {
    unit && this.bindToUnit(unit)
  }

  toJson(): IOutput {
    return {
      name: this.name,
      type: this.type,
    }
  }
}
