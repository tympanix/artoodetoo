import { Component, OnInit, Input } from '@angular/core';
import { Task, Unit } from '../../model'
import { ApiService } from '../../api.service'
import { MdDialog, MdDialogRef } from '@angular/material';
import { EventTemplateDialog } from '../../dialogs'

@Component({
  selector: 'eventeditor',
  templateUrl: './eventeditor.component.html',
  styles: []
})
export class EventeditorComponent implements OnInit {
  @Input() event: Unit

  constructor(private api: ApiService, public dialog: MdDialog) { }

  ngOnInit() {
  }

  saveEvent(){
    this.api.saveEvent(this.event).subscribe()
  }

  openEventTemplateDialog() {
    let dialogRef = this.dialog.open(EventTemplateDialog, {
      height: '500px',
      width: '750px',
    });
    dialogRef.afterClosed().subscribe(event => {
      if (event) {
        this.event = event;
      }
    });
  }

}
