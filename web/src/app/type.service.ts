import { Injectable, Type } from '@angular/core';
import { Input } from './model'
import { TypeSelector } from './types/type-selector'

import { IntegerTypeComponent } from './types/integer-type/integer-type.component'
import { StringTypeComponent } from './types/string-type/string-type.component'
import { BoolTypeComponent } from './types/bool-type/bool-type.component'
import { CronTimeComponent} from './types/cron-time/cron-time.component'
import { FacebookTokenComponent } from './types/facebook-token/facebook-token.component'

@Injectable()
export class TypeService {
  types: TypeSelector[] = []

  constructor() {
    this.add(IntegerTypeComponent, (input: Input) => input.isInteger())
    this.add(StringTypeComponent, (input: Input) => input.isString())
    this.add(BoolTypeComponent, (input: Input) => input.isBool())
    this.add(CronTimeComponent, (input: Input) => input.type == "cron.Time")
    this.add(FacebookTokenComponent, (input: Input) => input.type == "facebook.Token")
  }

  add(type: Type<any>, selector: (Input) => boolean) {
    this.types.push(new TypeSelector(type, selector))
  }

  getType(input: Input): Type<any> {
    let selector = this.types.find((s: TypeSelector) => s.selector(input))
    if (!selector) return
    return selector.component
  }
}
