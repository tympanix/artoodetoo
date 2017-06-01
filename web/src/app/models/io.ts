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
    } else if (other.isInterface()) {
      return true;
    } else if (this.isNumber() && other.isNumber()) {
      return true;
    } else {
      return false;
    }
  }

  isInterface(): boolean {
    return this.type == "interface {}"
  }

  isBool(): boolean {
    return this.type == "bool"
  }

  isNumber(): boolean {
    return this.isInteger() || this.isFloat()
  }

  isString(): boolean {
    return this.type == "string"
  }

  isInteger(): boolean {
    return this.type == "int32" || this.type == "int64"
  }

  isFloat(): boolean {
    return this.type == "float32" || this.type == "float64"
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

  private addDefaultRecipe() {
    if (this.isBool()) {
      this.recipe.push(new Ingredient(false))
    } else if (this.isString()) {
      this.recipe.push(new Ingredient(""))
    } else {
      this.recipe.push(new Ingredient(null))
    }
  }

  bootstrap(unit?: Unit) {
    unit && this.bindToUnit(unit)
    if (!this.recipe || !this.recipe.length) {
      this.addDefaultRecipe()
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
