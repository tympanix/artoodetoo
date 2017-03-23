import { Component, OnInit, Input } from '@angular/core';
import { Task, Ingredient, Input as UnitInput } from '../task';

@Component({
  selector: 'ingredient',
  templateUrl: './ingredient.component.html',
  styles: []
})
export class IngredientComponent implements OnInit {
  @Input() task: Task
  @Input() input: UnitInput
  @Input() model: Ingredient

  constructor() { }

  ngOnInit() { }

}
