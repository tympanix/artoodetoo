import { Injectable } from '@angular/core';

import { Subject } from 'rxjs/Subject';

import { Task } from './task';
import { TaskService} from './task.service';
import { Unit } from './unit';
import { UnitService } from './unit.service';

@Injectable()
export class ApiService {

  public tasks: Subject<Task[]> = new Subject<Task[]>()
  public units: Subject<Unit[]> = new Subject<Unit[]>()

  constructor(private tasksService: TaskService, private unitsService: UnitService) {
    this.tasksService.getTasks().then(tasks => this.tasks.next(tasks))
    this.unitsService.getUnits().then(units => this.units.next(units))
  }

}
