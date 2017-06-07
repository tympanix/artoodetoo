import { Unit, IUnit, Model, Event } from '../model'

export class Log {
  type: string
  message: string
  task: string
  time: number

  static fromJson(model: Object): Log {
    let log = new Log()
    Object.assign(log, model)
    return log
  }

}