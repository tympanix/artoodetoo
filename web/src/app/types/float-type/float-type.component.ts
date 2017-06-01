import { Component, OnInit } from '@angular/core';
import { TypesComponent } from '../types.component'

@Component({
  templateUrl: './float-type.component.html',
  styles: []
})
export class FloatTypeComponent extends TypesComponent {
  stringFloat: string = ""

  constructor() {
    super()
  }

  stringToFloat(){
    this.ingredient.value = parseFloat(this.stringFloat)
  }

}
