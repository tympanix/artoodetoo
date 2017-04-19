import { Component, OnInit } from '@angular/core';

import { Task } from '../model';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styles: []
})
export class DashboardComponent implements OnInit {
  tasks: Task[]

  constructor(private api: ApiService) {
    api.tasks.subscribe((tasks) => this.tasks = tasks)
  }

  ngOnInit() { }

  runTask(task: Task) {
    console.log("Running", task.name)
    this.api.runTask(task).subscribe()

    console.log("Remove after refactor")
    task.running = true
  }

  stopTask(task: Task){
    console.log("Stopping", task.name)
    this.api.stopTask(task)

    console.log("Remove after refactor")
    task.running = false
  }

}
