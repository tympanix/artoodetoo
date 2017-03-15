import { Injectable } from '@angular/core';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

import { ReplaySubject } from 'rxjs/ReplaySubject';

import { ApiService } from '../api.service';
import { Task } from '../task';

@Injectable()
export class TaskResolver implements Resolve<Task> {

  constructor(private api: ApiService) { }

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    let obs: ReplaySubject<Task> = new ReplaySubject<Task>(1)
    let name = route.params['task']
    this.api.tasks.subscribe((tasks) => {
      obs.next(this.taskByName(name, tasks))
      obs.complete()
    })
    return obs
  }

  taskByName(name: string, tasks: Task[]): Task {
    return tasks.find((task) => task.name == name)
  }

}
