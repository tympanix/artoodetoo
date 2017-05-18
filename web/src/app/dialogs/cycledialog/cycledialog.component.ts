import { Component, OnInit, Inject } from '@angular/core';
import { MD_DIALOG_DATA, MdDialog, MdDialogRef } from '@angular/material';
import { Response, Headers } from '@angular/http'
import { Unit, Task } from '../../model'

@Component({
  selector: 'cycledialog',
  templateUrl: './cycledialog.component.html',
  styles: []
})
export class CycleDialog implements OnInit {
  cycle: Unit[]
  task: Task

  constructor(
    public dialogRef: MdDialogRef<CycleDialog>,
    @Inject(MD_DIALOG_DATA) public data: {cycle: Unit[], task: Task}
  ) {
    this.cycle = data.cycle,
    this.task = data.task
  }

  ngOnInit() {
  }

  close() {
    this.dialogRef.close(false)
  }

}
