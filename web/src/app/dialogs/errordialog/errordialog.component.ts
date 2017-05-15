import { Component, OnInit, Inject } from '@angular/core';
import { MD_DIALOG_DATA, MdDialog, MdDialogRef } from '@angular/material';
import { Response, Headers } from '@angular/http'

@Component({
  selector: 'errordialog',
  templateUrl: './errordialog.component.html',
  styles: []
})
export class ErrorDialog implements OnInit {

  constructor(
    public dialogRef: MdDialogRef<ErrorDialog>,
    @Inject(MD_DIALOG_DATA) public data: Response
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
