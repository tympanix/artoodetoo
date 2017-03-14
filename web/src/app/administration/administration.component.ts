import { Component, OnInit } from '@angular/core';

import { Unit } from '../unit';
import { Task } from '../task';
import { UnitService } from '../unit.service';
import { TaskService} from '../task.service';

@Component({
  selector: 'app-administration',
  templateUrl: './administration.component.html',
  styles: []
})
export class AdministrationComponent implements OnInit {
  units: Unit[]
  event: Unit
  actions: Unit[]
  task: Task

  constructor(private unitService: UnitService, private taskService: TaskService) { }

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

  createTask(): void {
    this.taskService.createTask(this.task);
  }

  // For test purpose only
  createMockTask():void {
    this.taskService.createMockTask()

  }


}
