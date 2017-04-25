import { Input, IInput, Output, IOutput } from './io'
import { Model } from './model'
import { IUnit } from'./unit'
import * as _ from "lodash";

export interface IEvent extends IUnit {
  uuid: string
}

export class Event implements IEvent, Model {
  // Model properties
  id: string
  uuid: string
  name: string
  description: string
  input: Input[]
  output: Output[]

  static fromJson(model: IEvent): Event {
    let event = Object.create(Event.prototype)
    Object.assign(event, model)
    event.input = model.input.map(input => new Input(input))
    event.output = model.output.map(output => new Output(output))
    return event
  }

  copy(): Event {
    let copy = _.cloneDeep(this)
    return copy
  }

  bootstrap() {
    this.input.forEach(input => input.bootstrap())
  }

  toJson(): IEvent {
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
