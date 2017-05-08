import { Component, OnInit } from '@angular/core';
import { Task, Unit} from '../model';
import { ApiService } from '../api.service';
import { MdSnackBar } from '@angular/material';

import { MdDialog, MdDialogRef } from '@angular/material';
import { UnitDialog } from '../dialogs/unitdialog/unitdialog.component';
import { TaskDialog } from '../dialogs/taskdialog/taskdialog.component';
import { EventDialog} from '../dialogs/eventdialog/eventdialog.component';

import { Router, ActivatedRoute } from '@angular/router';
import * as _ from "lodash";

@Component({
  selector: 'app-administration',
  templateUrl: './administration.component.html',
  styles: []
})
export class AdministrationComponent implements OnInit {
  tasks: Task[]
  units: Unit[]
  events: Unit[]
  templateEvents: Unit[]
  task: Task
  event: Unit

  editorState: boolean
  eventActive: boolean
  taskActive: boolean

  // Event selection
  eventType: number

  constructor(private api: ApiService, private route: ActivatedRoute, private router: Router, public dialog: MdDialog, private snackBar: MdSnackBar) {
    api.units.subscribe((units) => this.units = units)
    api.tasks.subscribe((tasks) => this.tasks = tasks)
    api.events.subscribe((events) => this.events = events)
    api.templateEvents.subscribe((tevents) => this.templateEvents = tevents)

    this.eventActive = _.last(route.snapshot.url).path == 'event'
    this.taskActive = _.last(route.snapshot.url).path == 'task'
    this.editorState =  this.taskActive || this.eventActive

  }

  ngOnInit() {
    this.route.data.subscribe((data: {task: Task}) => {
      if(data.task){
        this.task = data.task.copy()
      }
      console.log("Task", data.task)
    })

  }

  // Return units with an input type mathcing the given argument
  getUnitsByType(type: string): Unit[] {
    let typeUnits: Unit[];
    typeUnits =  this.units.filter(unit => unit.input.find(x => x.type === type));
    return typeUnits;
  }

  createTask(): void {
    this.api.createTask(this.task).subscribe()
  }

  runTask() {
    this.api.runTask(this.task).subscribe()
  }

  updateTask() {
    this.api.updateTask(this.task).subscribe()
    //this.snackBar.open(this.task.name + " has been saved", "", {duration: 4000, extraClasses: ["snackbar-success"]})
  }

  saveEvent(){
    this.api.saveEvent(this.event).subscribe()
  }

  deleteTask(){
    this.api.deleteTask(this.task).subscribe()
  }

  test() {
    console.log(this.task)
  }

  eventTypeHandler(){
    this.eventType = +this.eventType

    if(this.eventType == 0){
      this.task.event.name = ""
    }
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

  openEventDialog() {
    let dialogRef = this.dialog.open(EventDialog, {
      height: '500px',
      width: '750px',
    });
    dialogRef.afterClosed().subscribe(event => {
      if (event) {
        this.event = event;
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
