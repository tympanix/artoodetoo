import { Component, OnInit } from '@angular/core';
import { Task, Unit} from '../model';
import { ApiService } from '../api.service';
import { MdSnackBar } from '@angular/material';

import { MdDialog, MdDialogRef } from '@angular/material';
import { UnitDialog, TaskDialog, EventTemplateDialog } from '../dialogs';

import { ErrorService } from '../error.service'
import { SourceWarning, EditorError } from '../model'

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

  warnings: EditorError[] = []

  constructor(private api: ApiService, private route: ActivatedRoute, private errhub: ErrorService) {
    api.units.subscribe((units) => this.units = units)
    api.tasks.subscribe((tasks) => this.tasks = tasks)
    api.events.subscribe((events) => this.events = events)
    api.templateEvents.subscribe((tevents) => this.templateEvents = tevents)

    this.eventActive = _.last(route.snapshot.url).path == 'event'
    this.taskActive = _.last(route.snapshot.url).path == 'task'
    this.editorState =  this.taskActive || this.eventActive

    this.errhub.errors
      .filter((e) => e instanceof EditorError)
      .map((e) => e as EditorError)
      .subscribe((e) => this.warnings.push(e))
  }

  ngOnInit() {
    this.route.data.subscribe((data: {task: Task}) => {
      if(data.task){
        this.task = data.task.copy()
        this.taskActive = true
        this.editorState = true
        console.log("Task", data.task)
      }
    })

  }

  // Return units with an input type mathcing the given argument
  getUnitsByType(type: string): Unit[] {
    let typeUnits: Unit[];
    typeUnits =  this.units.filter(unit => unit.input.find(x => x.type === type));
    return typeUnits;
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

}
