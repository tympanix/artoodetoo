import { Injectable } from '@angular/core';

import { ReplaySubject } from 'rxjs/ReplaySubject';

import { Task } from './task';
import { TaskService} from './task.service';
import { Unit } from './unit';
import { UnitService } from './unit.service';

@Injectable()
export class ApiService {

  public tasks: ReplaySubject<Task[]> = new ReplaySubject<Task[]>(1)
  public units: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)

  constructor(private tasksService: TaskService, private unitsService: UnitService) {
    this.tasksService.getTasks().then(tasks => this.tasks.next(tasks))
    this.unitsService.getUnits().then(units => this.units.next(units))
  }

}
