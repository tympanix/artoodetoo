import { Type } from '@angular/core';
import { Input } from '../model'

export class TypeSelector {
  constructor(public component: Type<any>, public selector: (Input) => boolean) {}
}
