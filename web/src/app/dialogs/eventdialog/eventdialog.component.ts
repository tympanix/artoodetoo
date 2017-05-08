import { Component, OnInit } from '@angular/core';
import { MdDialog, MdDialogRef } from '@angular/material';
import { ApiService } from '../../api.service'
import { Unit } from '../../model'

@Component({
  selector: 'eventdialog',
  templateUrl: './eventdialog.component.html'
})
export class EventDialog implements OnInit {

  search: string = ""
  events: Unit[]
  filtered: Unit[]

  constructor(private api: ApiService, public dialogRef: MdDialogRef<EventDialog>) {}

  ngOnInit() {
    this.api.templateEvents.subscribe(u => this.events = u)
    this.filtered = this.events
  }

  doSearch(event) {
    console.log(this.events)
    this.filtered = this.events.filter(u =>
      u.id.toLowerCase().includes(event.toLowerCase()))
  }

  addEvent(u: Unit) {
    let event
    let template = u || this.filtered[0]
    if (template) {
      event = template.copy()
      event.bootstrap()
    }
    this.dialogRef.close(event)
  }

  close() {
    this.dialogRef.close(undefined)
  }

}
