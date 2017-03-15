import { Component, OnInit, Input } from '@angular/core';

import { ApiService } from '../api.service';
import { Task, Unit } from '../task';
import { Meta} from '../meta';

@Component({
  selector: 'unit',
  templateUrl: './unit.component.html',
  styles: []
})
export class UnitComponent implements OnInit {
  @Input() unit: Unit;

  meta: Meta

  constructor(private api: ApiService) {
    api.metas.subscribe((metas) => {
      this.meta = this.findMetaById(this.unit.id, metas)
      console.log("Found meta", this.meta)
    })
  }

  ngOnInit() {}

  findMetaById(id: string, metas: Meta[]): Meta {
    return metas.find((meta) => meta.id == id)
  }

}
