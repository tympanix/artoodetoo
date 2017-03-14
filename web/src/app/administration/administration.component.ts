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

  // Return units with an input type mathcing the given argument
  getUnitsByType(type: string) {
    let typeUnits: Unit[];
    typeUnits =  this.units.filter(unit => unit.input.find(x => x.type === type));
    console.log(typeUnits);
    return typeUnits;
  }
}
