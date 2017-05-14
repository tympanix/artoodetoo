import { Component, OnInit } from '@angular/core';
import { MdDialog, MdDialogRef } from '@angular/material';
import { ApiService } from '../../api.service'
import { Unit, Event } from '../../model'

@Component({
  selector: 'app-eventdialog',
  templateUrl: './eventdialog.component.html',
  styles: []
})
export class EventDialog implements OnInit {

  search: string = ""
  events: Event[]
  filtered: Event[]

  constructor(private api: ApiService, public dialogRef: MdDialogRef<EventDialog>) { }

  ngOnInit() {
    this.api.events.subscribe(u => this.events = u)
    this.filtered = this.events
  }

  doSearch(event) {
    console.log(this.events)
    this.filtered = this.events.filter(u => {
      return u.name.toLowerCase().includes(event.toLowerCase()) ||
        u.id.toLowerCase().includes(event.toLowerCase())
    })
  }

  chooseEvent(u: Event) {
    let event = u || this.filtered.length && this.filtered[0]
    this.dialogRef.close(event)
  }

  close() {
    this.dialogRef.close(undefined)
  }

}
