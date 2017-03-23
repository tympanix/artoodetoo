import { Component, OnInit, Input } from '@angular/core';
import { Task, Unit, Ingredient, Input as UnitInput, Output as UnitOutput } from '../model';
import { ApiService } from '../api.service'

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
  source: Unit = new Unit()

  constructor(private api: ApiService) { }

  ngOnInit() {
    this.task.units.subscribe(units => this.sources = units)
    this.changeSource(this.model.source)
  }

  changeSourceEvent(event) {
    this.changeSource(event.value)
  }

  changeSource(source: string) {
    let src = source ? source : this.model.source
    this.source = this.sources.find(u => u.name == src)
    console.log("Changed", this.source)
  }

}
