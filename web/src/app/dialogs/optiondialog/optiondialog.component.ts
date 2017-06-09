import { Component, OnInit, Inject } from '@angular/core';
import { MD_DIALOG_DATA, MdDialog, MdDialogRef } from '@angular/material';

@Component({
  selector: 'app-optiondialog',
  templateUrl: './optiondialog.component.html',
  styles: []
})
export class OptionDialog implements OnInit {

  constructor(
    public dialogRef: MdDialogRef<OptionDialog>,
    @Inject(MD_DIALOG_DATA) public data: any,
  ) { }

  ngOnInit() {
  }

  close() {
    this.dialogRef.close(false)
  }

  confirm() {
    this.dialogRef.close(true)
  }

}
