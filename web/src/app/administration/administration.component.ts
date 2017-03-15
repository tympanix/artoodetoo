import { Component, OnInit } from '@angular/core';

import { Unit } from '../unit';
import { Task } from '../task';
import { UnitService } from '../unit.service';
import { TaskService} from '../task.service';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-administration',
  templateUrl: './administration.component.html',
  styles: []
})
export class AdministrationComponent implements OnInit {
  units: Unit[]
  tasks: Task[]
  event: Unit
  actions: Unit[]
  task: Task

  constructor(private api: ApiService, private taskService: TaskService) {
    api.units.subscribe((units) => this.units = units)
    api.tasks.subscribe((tasks) => this.tasks = tasks)
  }

  ngOnInit() {}

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
