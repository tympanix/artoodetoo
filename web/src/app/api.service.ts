import { Injectable } from '@angular/core';

import { ReplaySubject } from 'rxjs/ReplaySubject';

import { Task } from './task';
import { TaskService} from './task.service';
import { Meta } from './meta';
import { UnitService } from './unit.service';

@Injectable()
export class ApiService {

  public tasks: ReplaySubject<Task[]> = new ReplaySubject<Task[]>(1)
  public metas: ReplaySubject<Meta[]> = new ReplaySubject<Meta[]>(1)

  constructor(private tasksService: TaskService, private unitsService: UnitService) {
    this.tasksService.getTasks().then(tasks => this.tasks.next(tasks))
    this.unitsService.getMetas().then(metas => this.metas.next(metas))
  }

}
