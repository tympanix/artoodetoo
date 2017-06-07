import { Component, OnInit } from '@angular/core';
import { Task, Unit, Log } from '../model';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-debugger',
  templateUrl: './debugger.component.html',
  styles: []
})
export class DebuggerComponent implements OnInit {
  logs: Log[] = []

  constructor(private api: ApiService) {
    this.api.logs.subscribe(l => this.logs.push(l))
  }

  ngOnInit() {
  }

  clear() {

  }

}
