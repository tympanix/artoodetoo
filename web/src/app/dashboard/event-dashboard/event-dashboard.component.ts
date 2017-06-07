import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router'
import { ApiService } from '../../api.service'
import { ErrorService } from '../../error.service'
import { MdDialog, MdDialogRef } from '@angular/material';

import { Event, Task } from '../../model'
import { OptionDialog } from '../../dialogs';

@Component({
  selector: 'event-dashboard',
  templateUrl: './event-dashboard.component.html',
  styles: []
})
export class EventDashboardComponent implements OnInit {
  event: Event

  notfound: boolean

  constructor(private api: ApiService, private route: ActivatedRoute, private errhub: ErrorService, public dialog: MdDialog) {
    
  }

  subscribe() {
    this.api.events.map(events => events.find(e => e.uuid == this.event.uuid)).filter(e => !!e).subscribe(e => this.event = e)
  }

  ngOnInit() {
    this.route.data.subscribe((data: {event: Event}) => {
      if(data.event){
        this.event = data.event
        this.subscribe()
      } else {
        this.notfound = true
      }
    })
  }

  deleteTask(task: Task) {
    let dialogRef = this.dialog.open(OptionDialog, {
      width: '550px',
      data: {
        title: "Delete Task?",
        message: "Would you really like to delete " + task.name
      }
    });
    dialogRef.afterClosed().subscribe(ok => {
      if (ok) {
        this.api.deleteTask(task).subscribe()
      }
    });
  }

}
