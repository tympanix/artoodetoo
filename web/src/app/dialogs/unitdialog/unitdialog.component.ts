import { Component, OnInit } from '@angular/core';
import { MdDialog, MdDialogRef } from '@angular/material';
import { ApiService } from '../../api.service'
import { Unit } from '../../model'

@Component({
  selector: 'unitdialog',
  templateUrl: './unitdialog.component.html'
})
export class UnitDialog implements OnInit {

  search: string = ""
  units: Unit[]
  filtered: Unit[]

  constructor(private api: ApiService, public dialogRef: MdDialogRef<UnitDialog>) {}

  ngOnInit() {
    this.api.units.subscribe(u => this.units = u)
    this.filtered = this.units
  }

  doSearch(event) {
    this.filtered = this.units.filter(u =>
      u.id.toLowerCase().includes(event.toLowerCase()))
  }

  addUnit() {
    this.dialogRef.close(this.filtered[0] || undefined)
  }

  close() {
    this.dialogRef.close(undefined)
  }

}
