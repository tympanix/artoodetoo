import { Injectable, Type } from '@angular/core';
import { Input } from './model'
import { TypeSelector } from './types/type-selector'

import { IntegerTypeComponent } from './types/integer-type/integer-type.component'
import { StringTypeComponent } from './types/string-type/string-type.component'
import { BoolTypeComponent } from './types/bool-type/bool-type.component'

@Injectable()
export class TypeService {
  types: TypeSelector[] = []

  constructor() {
    this.add(new TypeSelector(IntegerTypeComponent, (input: Input) => input.isInteger()))
    this.add(new TypeSelector(StringTypeComponent, (input: Input) => input.isString()))
    this.add(new TypeSelector(BoolTypeComponent, (input: Input) => input.isBool()))
  }

  add(selector: TypeSelector) {
    this.types.push(selector)
  }

  getType(input: Input): Type<any> {
    let selector = this.types.find((s: TypeSelector) => s.selector(input))
    if (!selector) return
    return selector.component
  }
}
