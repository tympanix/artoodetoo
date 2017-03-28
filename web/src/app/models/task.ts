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
  name: string = ""
  event: Unit = null
  actions: Unit[] = []

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

  addAction(unit: Unit) {
    this.actions.push(unit)
    this.updateUnitList()
  }

  private updateUnitList() {
    let units: Unit[] = []
    this.event && units.push(this.event)
    this.actions.forEach(action => {
      action && units.push(action)
    })
    this.units.next(units)
  }

  deleteUnit(unit: Unit) {
    if (this.event === unit) {
      this.event === undefined
    }
    this.actions = this.actions.filter(u => u !== unit)
  }

  private swapActions(indexFrom: number, indexTo: number) {
    if (indexFrom < 0 || indexFrom >= this.actions.length) return
    if (indexTo < 0 || indexTo >= this.actions.length) return
    let temp = this.actions[indexTo]
    this.actions[indexTo] = this.actions[indexFrom]
    this.actions[indexFrom] = temp
  }

  moveUnitUp(unit: Unit) {
    let indexFrom = this.actions.indexOf(unit)
    this.swapActions(indexFrom, indexFrom - 1)
  }

  moveUnitDown(unit: Unit) {
    let indexFrom = this.actions.indexOf(unit)
    this.swapActions(indexFrom, indexFrom + 1)
  }

}
