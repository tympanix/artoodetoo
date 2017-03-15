import { Component, OnInit } from '@angular/core';

import { Task } from '../task';
import { TaskService} from '../task.service';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styles: []
})
export class DashboardComponent implements OnInit {
  tasks: Task[]

  constructor(private taskService: TaskService, private api: ApiService) {
    api.tasks.subscribe((tasks) => this.tasks = tasks)
  }

  ngOnInit() {

  }

  getTasks(): void {
    this.taskService.getTasks().then(tasks => this.tasks = tasks);
  }

}
