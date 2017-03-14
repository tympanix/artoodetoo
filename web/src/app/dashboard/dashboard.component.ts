import { Component, OnInit } from '@angular/core';

import { Task } from '../task';
import { TaskService} from '../task.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styles: []
})
export class DashboardComponent implements OnInit {
  tasks: Task[]

  constructor(private taskService: TaskService) { }

  ngOnInit() {
    this.getTasks();
  }

  getTasks(): void {
    this.taskService.getTasks().then(tasks => this.tasks = tasks);
  }

}
