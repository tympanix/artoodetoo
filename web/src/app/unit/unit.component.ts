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

  test() {
    console.log(this.unit)
  }

}
