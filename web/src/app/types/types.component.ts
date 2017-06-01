import { Component, OnInit, Input } from '@angular/core';
import { IngredientType } from './ingredient-type'
import { Ingredient, Task, Input as UnitInput } from '../model'

@Component({
  styles: []
})
export abstract class TypesComponent implements OnInit, IngredientType {
  @Input() ingredient: Ingredient

  input: UnitInput
  task: Task

  constructor() { }

  ngOnInit() {
    this.input = this.ingredient.input
  }

}
