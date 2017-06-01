import { Component, OnInit, Input, ComponentFactoryResolver, ViewChild, Type, AfterViewInit } from '@angular/core';
import { TypeService } from '../../type.service'
import { Task, Unit, Ingredient, Input as UnitInput, Output as UnitOutput } from '../../model';
import { TypeDirective } from '../../types/type.directive'
import { IngredientType } from '../../types/ingredient-type'
import { TypesComponent } from '../../types/types.component'

@Component({
  selector: 'typeeditor',
  templateUrl: './typeeditor.component.html',
  styles: []
})
export class TypeeditorComponent implements OnInit, AfterViewInit {
  @Input() ingredient: Ingredient
  @ViewChild(TypeDirective) typeHost: TypeDirective;

  constructor(private types: TypeService, private _componentFactoryResolver: ComponentFactoryResolver) { }

  ngOnInit() {
    this.loadType()
  }

  ngAfterViewInit() {
    //setTimeout(_=> this.loadType());
  }

  loadType() {
    let type: Type<any> = this.types.getType(this.ingredient.input)
    if (!type) {
      console.error("Could not find type for input", this.ingredient.input)
      return
    }
    let componentFactory = this._componentFactoryResolver.resolveComponentFactory(type)

    let view = this.typeHost.viewContainerRef
    view.clear()

    let typeComponent = view.createComponent(componentFactory)
    let ingrType = typeComponent.instance as TypesComponent

    if (!ingrType) {
      console.error("Can't cast type component")
      return
    }

    console.log(ingrType)

    ingrType.ingredient = this.ingredient
  }

}
