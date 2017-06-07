import { Component, OnInit } from '@angular/core';
import { Task, Unit, Log } from '../model';
import { ApiService } from '../api.service';
import { LogService } from '../log.service';

@Component({
  selector: 'app-debugger',
  templateUrl: './debugger.component.html',
  styles: []
})
export class DebuggerComponent implements OnInit {


  constructor(private api: ApiService, private log: LogService) {

  }

  ngOnInit() {
  }

}
