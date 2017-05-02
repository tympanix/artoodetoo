import { Component, OnInit, Input } from '@angular/core';
import { Ingredient, Input as UnitInput } from '../model';

@Component({
  selector: 'event-ingredient',
  templateUrl: './event-ingredient.component.html',
  styles: []
})
export class EventIngredientComponent implements OnInit {
  @Input() input: UnitInput
  @Input() ingr: Ingredient

  minute: number[]
  hour: number[]
  selectedType: number
  selectedNumber: number

  constructor() {
    this.minute = [0,5,10,15,20,25,30,35,40,45,50,55]
    this.hour = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]
  }

  ngOnInit() {
  }

  typeToNumber() {
    this.selectedType = +this.selectedType
  }

  updateIngrValue(){
    this.ingr.value = this.selectedNumber + (this.selectedType == 0 ? "m" : "h")
  }

}
