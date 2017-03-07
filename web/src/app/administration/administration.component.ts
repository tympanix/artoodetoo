import { Component, OnInit } from '@angular/core';

import { Unit } from '../unit';
import { UnitService } from '../unit.service';

@Component({
  selector: 'app-administration',
  templateUrl: './administration.component.html',
  styles: []
})
export class AdministrationComponent implements OnInit {
  units: Unit[]

  constructor(private unitService: UnitService) { }

  getUnits(): void {
    this.unitService.getUnits().then(units => this.units = units);
  }

  ngOnInit() {
    this.getUnits();
  }
}
