import { ReplaySubject } from 'rxjs/ReplaySubject';
import { Subject } from 'rxjs/Subject';
import { Unit, IUnit, Model, Event } from '../model'
import * as _ from "lodash";

export interface ITask {
  name: string
  uuid: string
  event: IUnit
  actions: IUnit[]
}

export class Task implements ITask, Model {
  // Model properties
  name: string = ""
  uuid: string = ""
  event: Event = null
  actions: Unit[] = []

  private eventRef: string

  // State properties
  units: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)
  _sources: Unit[] = []
  isSaved: boolean = false

  constructor(fields?: {
    name?: string
  }) {
    if (fields) Object.assign(this, fields)
  }

  static fromJson(model: ITask): Task {
    let task = new Task()
    Object.assign(task, model)
    task.eventRef = model.event as Object as string
    //task.event = Unit.fromJson(model.event)
    task.isSaved = true
    task.actions = model.actions.map(action => Unit.fromJson(action))
    task.bootstrap()
    return task
  }

  resolveEvent(events: Event[]) {
    let event = events.find(e => e.uuid == this.eventRef)

    if (!event) {
      throw new Error(`Event ${this.eventRef} not found for ${this.name}`)
    }

    this.event = event
    this.event.subscribeTask(this)
    this.updateUnitList()
    this.resolveIngredients()
  }

  checkCycles(): Promise<Unit[]> {
    var totalExplored: Set<Unit> = new Set<Unit>()

    function hasCycleFromUnit(unit: Unit): Unit[] {
      var explored: Set<Unit> = new Set<Unit>()
      var queue: Unit[] = []

      queue.push(unit)

      while (queue.length > 0) {
        var node: Unit = queue.pop()

        explored.add(node)
        totalExplored.add(node)

        var expanded = reachableUnits(node)
        for (let u of expanded) {
          if (explored.has(u)) {
            return extractCycle(u)
          }
          if (queue.find(q => q == u)) {
            continue
          }
          queue.push(u)
        }
      }
      return null
    }

    function extractCycle(u: Unit): Unit[] {
      var cycle: Unit[] = [u]
      var parent: Unit = u.parent
      while (parent) {
        if (parent === u) break
        cycle.push(parent)
        parent = parent.parent
      }
      return cycle
    }

    function reachableUnits(unit: Unit): Unit[] {
      var reachable: Unit[] = []
      unit.input.forEach(i => {
        i.recipe.forEach(r => {
          if (!r.isVariable() || !r.reference) return
          var other: Unit = r.reference.unit
          other.parent = unit
          reachable.push(other)
        })
      })
      return reachable
    }

    function getUnexploredAction(task: Task) {
      return task.actions.find(a => !totalExplored.has(a))
    }

    return new Promise((resolve, reject) => {
      var missing = getUnexploredAction(this)

      while (missing) {
        let cycle = hasCycleFromUnit(missing)
        if (cycle) reject(cycle)
        missing = getUnexploredAction(this)
      }
      resolve(null)
    })
  }

  private resolveIngredients() {
    this.actions.forEach(
      a => a.input.forEach(
        i => i.recipe.forEach(
          r => r.resolveReference(this._sources))))
  }

  copy(): Task {
    let copy = _.cloneDeep(this)
    return copy
  }

  public toJson(): ITask {
    return {
      name: this.name,
      uuid: this.uuid,
      event: this.event.toJson(),
      actions: this.actions.map(a => a.toJson())
    }
  }

  bootstrap() {
    this.updateUnitList()
    this.actions.forEach(a => a.bootstrap(this))
  }

  addAction(unit: Unit) {
    unit.bootstrap(this)
    this.actions.push(unit)
    this.updateUnitList()
  }

  updateUnitList() {
    let units: Unit[] = []
    this.event && units.push(this.event)
    this.actions.forEach(action => {
      action && units.push(action)
    })
    this._sources = units
    this.units.next(units)
  }

  deleteUnit(unit: Unit) {
    console.log("Deleting unit: ", unit)
    this.actions = this.actions.filter(u => u !== unit)
    this.updateUnitList()
  }

  deleteEvent(){
    console.log("Deleting event")
    this.event = null
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

  getSourceByName(name: string): Unit {
    if (this.event && this.event.name == name) return this.event
    return this.actions.find(unit => unit.name == name)
  }

}
