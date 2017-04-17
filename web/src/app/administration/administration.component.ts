import { Component, OnInit } from '@angular/core';
import { Task, Unit } from '../model';
import { ApiService } from '../api.service';
import { MdSnackBar } from '@angular/material';

import { MdDialog, MdDialogRef } from '@angular/material';
import { UnitDialog } from '../dialogs/unitdialog/unitdialog.component';
import { TaskDialog } from '../dialogs/taskdialog/taskdialog.component';

import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-administration',
  templateUrl: './administration.component.html',
  styles: []
})
export class AdministrationComponent implements OnInit {
  tasks: Task[]
  units: Unit[]
  task: Task

  constructor(private api: ApiService, private route: ActivatedRoute, public dialog: MdDialog, private snackBar: MdSnackBar) {
    api.units.subscribe((units) => this.units = units)
    api.tasks.subscribe((tasks) => this.tasks = tasks)
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

  deleteTask(){
    this.api.deleteTask(this.task).subscribe()
  }

  test() {
    console.log(this.task)
  }

  openUnitDialog(type: string) {
    let dialogRef = this.dialog.open(UnitDialog, {
      height: '500px',
      width: '750px',
    });
    dialogRef.afterClosed().subscribe(unit => {
      if (unit) {
        console.log(type);
        switch(type){

          case "event" :
            this.task.event = unit;
            break;
          case "action":
            this.task.addAction(unit);
            break;
        }
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
