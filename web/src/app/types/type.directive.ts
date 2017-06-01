import { Directive, ViewContainerRef} from '@angular/core';

@Directive({
  selector: '[type-host]'
})
export class TypeDirective {

  constructor(public viewContainerRef: ViewContainerRef) { }

}
