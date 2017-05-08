import { Input, IInput, Output, IOutput } from './io'
import { Model } from './model'
import * as _ from "lodash";

export interface IUnit {
  id: string
  uuid: string
  name: string
  description: string
  input: IInput[]
  output: IOutput[]
}

export class Unit implements IUnit, Model {
  // Model properties
  id: string
  uuid: string = ""
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

  copy(): Unit {
    let copy = _.cloneDeep(this)
    return copy
  }

  bootstrap() {
    this.input.forEach(input => input.bootstrap())
  }

  toJson(): IUnit {
    return {
      id: this.id,
      uuid: this.uuid,
      name: this.name,
      description: this.description,
      input: this.input.map(i => i.toJson()),
      output: this.output.map(o => o.toJson())
    }
  }
}
