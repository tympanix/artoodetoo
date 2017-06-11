import { Input, IInput, Output, IOutput } from './io'
import { Model } from './model'
import { ITask, Task } from './task'
import { IEvent, Event } from './event'
import { IUnit, Unit } from './unit'
import * as _ from "lodash";

export interface IData {
  events: IEvent[]
  eventtemplates: IUnit[]
  tasks: ITask[]
  actions: IUnit[]
}

export class Data implements IData, Model {
  events: Event[]
  eventtemplates: Event[]
  tasks: Task[]
  actions: Unit[]

  toJson(): IData {
    return {
      events: this.events.map(e => e.toJson()),
      eventtemplates: this.eventtemplates.map(e => e.toJson()),
      tasks: this.tasks.map(t => t.toJson()),
      actions: this.actions.map(a => a.toJson())
    }
  }

  static fromJson(model: IData) {
    let data = Object.create(Data.prototype)

    data.events = model.events.map(e => Event.fromJson(e))
    data.eventtemplates = model.eventtemplates.map(e => Unit.fromJson(e))
    data.tasks = model.tasks.map(t => Task.fromJson(t))
    data.actions = model.actions.map(a => Unit.fromJson(a))

    data.tasks.forEach(t => t.resolveEvent(data.events))
    console.log(data)
    return data
  }
}
