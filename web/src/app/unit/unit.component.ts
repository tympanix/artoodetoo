import { Component, OnInit, Input } from '@angular/core';
import { MdSnackBar } from '@angular/material';

import { ApiService } from '../api.service';
import { Task, Unit } from '../model';

@Component({
  selector: 'unit',
  templateUrl: './unit.component.html',
  styles: []
})
export class UnitComponent implements OnInit {
  @Input() task: Task
  @Input() unit: Unit

  // Temporary placeholder when changing unit name
  unitname: string = ""

  constructor(private api: ApiService, private snackBar: MdSnackBar) {}

  ngOnInit() { }

  delete() {
    this.task.deleteUnit(this.unit)
  }

  moveDown() {
    this.task.moveUnitDown(this.unit)
  }

  moveUp() {
    this.task.moveUnitUp(this.unit)
  }

  changeUnitName(name: string) {
    if (!name) {
      this.snackBar.open("The unit must be given a name", "", { duration: 4000 })
      return
    }
    let unit = this.task.getSourceByName(name)
    if (unit) {
      this.snackBar.open("A unit already exists with that name", "", { duration: 4000 })
    } else {
      this.unit.name = name
      this.unitname = ""
      this.task.updateUnitList()
    }
  }
}
