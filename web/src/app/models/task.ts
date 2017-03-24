import { ReplaySubject } from 'rxjs/ReplaySubject';
import { Unit, IUnit } from './unit'
import { Model } from './model'

interface ITask {
  name: string
  event: IUnit
  actions: IUnit[]
}

export class Task implements ITask, Model {
  // Model properties
  name: string
  event: Unit
  actions: Unit[]

  // State properties
  units: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)

  constructor(fields?: {
    name?: string
  }) {
    if (fields) Object.assign(this, fields)
  }

  static fromJson(model: ITask): Task {
    let task = new Task()
    Object.assign(task, model)
    task.event = Unit.fromJson(model.event)
    task.actions = model.actions.map(action => Unit.fromJson(action))
    task.updateUnitList()
    return task
  }

  public toJson(): ITask {
    return {
      name: this.name,
      event: this.event.toJson(),
      actions: this.actions.map(a => a.toJson())
    }
  }

  private updateUnitList() {
    let units: Unit[] = []
    this.event && units.push(this.event)
    this.actions.forEach(action => {
      action && units.push(action)
    })
    this.units.next(units)
  }

}
