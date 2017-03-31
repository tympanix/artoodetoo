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
    this.task.units.subscribe(units => this.sources = this.filterUnits(units))
    this.changeSource(this.model.source)
  }

  private filterUnits(units: Unit[]): Unit[] {
    return units.filter(unit =>
      unit.name && unit.name.length > 0
    )
  }

  changeSourceEvent(event) {
    this.changeSource(event.value)
  }

  changeSource(source: string) {
    let src = source ? source : this.model.source
    let found = this.sources.find(u => u.name == src)
    if (found) this.source = found
    console.log("Changed", this.source)
  }

}
