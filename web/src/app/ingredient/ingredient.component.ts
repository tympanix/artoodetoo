import { Component, OnInit, Input } from '@angular/core';
import { Task, Unit, Ingredient, Input as UnitInput } from '../task';

@Component({
  selector: 'ingredient',
  templateUrl: './ingredient.component.html',
  styles: []
})
export class IngredientComponent implements OnInit {
  @Input() task: Task
  @Input() input: UnitInput
  @Input() model: Ingredient

  sources: Unit[]

  constructor() { }

  ngOnInit() {
    this.task.units.subscribe(units => this.sources = units)
  }

}
