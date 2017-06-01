import { Component, OnInit, Input } from '@angular/core';
import { TypesComponent } from '../types.component'

@Component({
  templateUrl: './integer-type.component.html',
  styles: []
})
export class IntegerTypeComponent extends TypesComponent {

  constructor() {
    super()
  }

}
