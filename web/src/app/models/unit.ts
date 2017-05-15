import { Input, IInput, Output, IOutput } from './io'
import { Model} from './model'
import { Task } from './task'
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

  // State properties
  task: Task = null

  static fromJson(model: IUnit): Unit {
    let unit = Object.create(Unit.prototype)
    Object.assign(unit, model)
    unit.input = model.input.map(input => new Input(input))
    unit.input.forEach(i => i.bindToUnit(unit))
    unit.output = model.output.map(output => new Output(output))
    unit.output.forEach(o => o.bindToUnit(unit))

    return unit
  }

  bindToTask(task: Task) {
    this.task = task
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
