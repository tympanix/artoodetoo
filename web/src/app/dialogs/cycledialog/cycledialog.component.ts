import { Component, OnInit, Inject } from '@angular/core';
import { MD_DIALOG_DATA, MdDialog, MdDialogRef, MdSnackBar } from '@angular/material';
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

  static check(dialog: MdDialog, snack: MdSnackBar, task: Task): Promise<void> {
    var self = this
    if (!task) return
    return task.checkCycles().catch((cycle) => {
      let cycleSnack = snack.open("The is an cycle in your task!", "View", {duration: 8000, extraClasses: ["snackbar-error"]})
      cycleSnack.onAction().subscribe(() => CycleDialog.open(dialog, snack, task, cycle));
      return Promise.reject(cycle)
    })
  }

  static open(dialog: MdDialog, snack: MdSnackBar, task: Task, cycle: Unit[]) {
    let dialogRef = dialog.open(CycleDialog, {
      width: '750px',
      data: {
        task: task,
        cycle: cycle
      }
    })
    dialogRef.afterClosed().subscribe(() => CycleDialog.check(dialog, snack, task));
  }
}
