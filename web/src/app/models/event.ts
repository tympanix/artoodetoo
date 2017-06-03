import { Input, IInput, Output, IOutput } from './io'
import { Model } from './model'
import { Task } from './task'
import { Unit, IUnit } from './unit'
import * as _ from "lodash";

export class Event extends Unit {
  subscribers: Task[] = []

  constructor() {
    super()
  }

  static fromJson(model: IUnit): Event {
    let event = Object.create(Event.prototype)
    event.subscribers = []
    Object.assign(event, Unit.fromJson(model))
    return event
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