import { Component, OnInit } from '@angular/core';
import { Task, Unit } from '../task';
import { UnitService } from '../unit.service';
import { TaskService} from '../task.service';
import { ApiService } from '../api.service';

import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-administration',
  templateUrl: './administration.component.html',
  styles: []
})
export class AdministrationComponent implements OnInit {
  tasks: Task[]
  units: Unit[]
  task: Task

  constructor(private api: ApiService, private taskService: TaskService, private route: ActivatedRoute) {
    api.units.subscribe((units) => this.units = units)
    api.tasks.subscribe((tasks) => this.tasks = tasks)
  }

  ngOnInit() {
    this.route.data.subscribe((data: {task: Task}) => {
      this.task = data.task
    })
  }

  // Return units with an input type mathcing the given argument
  getUnitsByType(type: string): Unit[] {
    let typeUnits: Unit[];
    typeUnits =  this.units.filter(unit => unit.input.find(x => x.type === type));
    return typeUnits;
  }

  createTask(): void {
    this.taskService.createTask(this.task);
  }

  // For test purpose only
  createMockTask():void {
    this.taskService.createMockTask()
  }

  runTask() {
      this.taskService.runTask(this.task)
  }


}
