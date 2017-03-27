import { Component, OnInit, Input } from '@angular/core';

import { ApiService } from '../api.service';
import { Task, Unit } from '../model';

@Component({
  selector: 'unit',
  templateUrl: './unit.component.html',
  styles: []
})
export class UnitComponent implements OnInit {
  @Input() task: Task
  @Input() unit: Unit;

  model: boolean = true
  model2: boolean = false

  constructor(private api: ApiService) {}

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

}
