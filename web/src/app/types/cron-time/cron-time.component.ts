import { Component, OnInit } from '@angular/core';
import { TypesComponent } from '../types.component'

@Component({
  templateUrl: './cron-time.component.html',
  styles: []
})
export class CronTimeComponent extends TypesComponent {
  timeType: number = 0
  selectedType: number
  selectedNumber: number

  second: number[] = [1, 5, 10, 20, 30, 60]
  minute: number[] = [1,5,10,15,20,25,30,35,40,45,50,55]
  hour: number[] = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]

  constructor() {
    super()
  }

  unitToNumber() {
    this.selectedType = +this.selectedType
  }

  typeToNumber(){
    this.timeType = +this.timeType
  }

  // Cron functions
  updateIngrValue(){
    this.ingredient.value = "@every " + this.selectedNumber + (this.selectedType == 0 ? "m" : "h")
  }

}
