import { Input, IInput, Output, IOutput } from './io'
import { Model } from './model'
import { Task } from './task'
import { Unit, IUnit } from './unit'
import * as _ from "lodash";

export interface IEvent extends IUnit {
  running: boolean
}

export class Event extends Unit implements IEvent {
  subscribers: Task[] = []
  running: boolean
  id: string
  uuid: string = ""
  name: string
  description: string
  input: Input[]
  output: Output[]
  color: string
  icon: string


  static fromJson(model: IEvent): Event {
    let event = Object.create(Event.prototype)
    event.subscribers = []
    Object.assign(event, model)
    console.log("The event: ",model.running,", is currently running: ", model.running)
    event.running = model.running
    event.input = model.input.map(input => new Input(input))
    event.input.forEach(i => i.bindToUnit(event))
    event.output = model.output.map(output => new Output(output))
    event.output.forEach(o => o.bindToUnit(event))

    return event
  }

  toJson(): IEvent {
    return {
      id: this.id,
      uuid: this.uuid,
      name: this.name,
      running: this.running,
      description: this.description,
      input: this.input.map(i => i.toJson()),
      output: this.output.map(o => o.toJson())
    }
  }

  subscribeTask(task: Task) {
    let existing = this.subscribers.find(t => t.uuid == task.uuid)

    if (!existing) {
      this.subscribers.push(task)
    }
  }

  unsubscribeTask(task: Task) {
    this.subscribers = this.subscribers.filter(t => t.uuid != task.uuid)
  }
}
