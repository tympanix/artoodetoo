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
  color: string
  icon: string

  // State properties
  parent: Unit = null
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

  package(): string {
    return this.id.split(".")[0]
  }

  lastname(): string {
    return this.id.split(".")[1]
  }

  isDark(): boolean {
    if (!this.color) {
      return false
    }
    
    var c = this.color.substring(1);      // strip #
    var rgb = parseInt(c, 16);   // convert rrggbb to decimal
    var r = (rgb >> 16) & 0xff;  // extract red
    var g = (rgb >> 8) & 0xff;  // extract green
    var b = (rgb >> 0) & 0xff;  // extract blue

    var luma = 0.2126 * r + 0.7152 * g + 0.0722 * b; // per ITU-R BT.709

    return  luma < 40
  }

  getStyle(): Object {
    return {
      "dark": this.isDark()
    }
  }

  bindToTask(task: Task) {
    this.task = task
  }

  copy(): Unit {
    let copy = _.cloneDeep(this)
    copy.bootstrap()
    return copy
  }

  bootstrap(task?: Task) {
    task && this.bindToTask(task)
    this.input.forEach(input => input.bootstrap(this))
    this.output.forEach(output => output.bootstrap(this))
  }

  static findByName(name: string): (u: Unit) => boolean {
    return function(unit: Unit) {
      return unit.name === name
    }
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
