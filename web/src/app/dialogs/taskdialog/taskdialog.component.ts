import { Component, OnInit } from '@angular/core';
import { MdDialog, MdDialogRef } from '@angular/material';
import { ApiService } from '../../api.service';
import { Task } from '../../model';

@Component({
  selector: 'app-taskdialog',
  templateUrl: './taskdialog.component.html',
  styles: []
})
export class TaskDialog implements OnInit {

  constructor(private api: ApiService, public dialogRef: MdDialogRef<TaskDialog>) { }

  ngOnInit() {
  }

  createTask(name) {
    this.dialogRef.close(name)
  }

  close() {
    this.dialogRef.close(undefined)
  }

}
