import { Component, OnInit, Input } from '@angular/core';
import { Task, Unit } from '../../model'
import { ApiService } from '../../api.service'
import { MdDialog, MdDialogRef } from '@angular/material';
import { UnitDialog, TaskDialog } from '../../dialogs'

@Component({
  selector: 'taskeditor',
  templateUrl: './taskeditor.component.html',
  styles: []
})
export class TaskeditorComponent implements OnInit {
  @Input() task: Task

  events: Unit[]

  constructor(private api: ApiService, public dialog: MdDialog) { }

  ngOnInit() {
    this.api.events.subscribe(e => this.events = e)
  }

  eventChange(){
    this.task.updateUnitList()
  }

  openUnitDialog() {
    let dialogRef = this.dialog.open(UnitDialog, {
      height: '500px',
      width: '750px',
    });
    dialogRef.afterClosed().subscribe(unit => {
      if (unit) {
          this.task.addAction(unit);
      }
    });
  }

  openTaskDialog(){
    let dialogRef = this.dialog.open(TaskDialog, {
      width: '600px'
    });

    dialogRef.afterClosed().subscribe(name => {
      if(name != undefined && name != ""){
        this.task = new Task({name: name})
      }
    })
  }

}
