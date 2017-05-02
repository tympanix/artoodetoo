import { Component, OnInit, Input } from '@angular/core';
import { MdSnackBar } from '@angular/material';

import { ApiService } from '../api.service';
import { Event } from '../model';

@Component({
  selector: 'app-event',
  templateUrl: './event.component.html',
  styles: []
})
export class EventComponent implements OnInit {
  @Input() event: Event

  // Temporary placeholder when changing event name
  eventname: string = ""

  constructor(private api: ApiService, private snackBar: MdSnackBar) {

  }

  ngOnInit() { }

  changeEventName(name: string) {
    if (!name) {
      this.snackBar.open("The event must be given a name", "", { duration: 4000 })
      return
    }

      this.event.name = name
      this.eventname = ""

  }

}
