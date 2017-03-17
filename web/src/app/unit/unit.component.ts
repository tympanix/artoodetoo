import { Component, OnInit, Input } from '@angular/core';

import { ApiService } from '../api.service';
import { Task, Unit } from '../task';

@Component({
  selector: 'unit',
  templateUrl: './unit.component.html',
  styles: []
})
export class UnitComponent implements OnInit {
  @Input() unit: Unit;

  model: boolean = true
  model2: boolean = false

  constructor(private api: ApiService) {}

  ngOnInit() {
      console.log("Loading unit", this.unit)
  }

}
