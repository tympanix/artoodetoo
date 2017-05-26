import { Component, OnInit, Input } from '@angular/core';
import { Task, Unit, Ingredient, Input as UnitInput, Output as UnitOutput } from '../model';
import { ApiService } from '../api.service'
import { MdSnackBar, MdDialog } from '@angular/material';
import { CycleDialog } from '../dialogs'
import { SourceWarning } from '../model'
import { ErrorService } from '../error.service'

@Component({
  selector: 'ingredient',
  templateUrl: './ingredient.component.html',
  styles: []
})
export class IngredientComponent implements OnInit {
  @Input() model: Ingredient

  task: Task
  input: UnitInput
  unit: Unit

  sources: Unit[]
  source: Unit = new Unit()
  reference: UnitOutput

  warning: SourceWarning

  constructor(private api: ApiService, private errhub: ErrorService, private snackBar: MdSnackBar, public dialog: MdDialog) { }

  ngOnInit() {
    this.input = this.model.input
    this.unit = this.input.unit
    this.task = this.unit.task

    this.task.units.subscribe(units => this.sources = this.filterUnits(units))
    this.changeSource(this.model.source)

    this.errhub.errors
      .filter(e => e instanceof SourceWarning)
      .map(w => w as SourceWarning)
      .do(w => this.warning = w)
  }

  private filterUnits(units: Unit[]): Unit[] {
    return units.filter(unit =>
      unit != this.unit && unit.name && unit.name.length > 0
    )
  }

  changeSourceEvent(event) {
    //this.changeSource(event.value)
  }

  changeIngredientReference(event) {
    this.model.setVariable(event.value)
    this.checkCycles()

    if (!event.value.assignableTo(this.input)) {
      this.warning = new SourceWarning(this.input, event.value)
      this.errhub.push(this.warning)
    } else {
      this.warning = null
    }
  }

  checkCycles() {
    var self = this
    if (!this.model) return
    CycleDialog.check(this.dialog, this.snackBar, this.model.getTask())
  }

  changeSource(source: string) {
    let src = source ? source : this.model.source
    let found = this.sources.find(u => u.name == src)
    if (found) this.source = found
  }

  typeToNumber() {
    this.model.type = +this.model.type
  }

}
