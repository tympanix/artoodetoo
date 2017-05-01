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

  minute: number[]
  hour: number[]
  selectedType: number
  selectedNumber: number

  constructor(private api: ApiService, private snackBar: MdSnackBar) {
    this.minute = [0,5,10,15,20,25,30,35,40,45,50,55]
    this.hour = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]
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

  typeToNumber() {
    this.selectedType = +this.selectedType
  }

}
