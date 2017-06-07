import { Injectable, Type } from '@angular/core';
import { Input } from './model'
import { TypeSelector } from './types/type-selector'

import { IntegerTypeComponent } from './types/integer-type/integer-type.component'
import { StringTypeComponent } from './types/string-type/string-type.component'
import { BoolTypeComponent } from './types/bool-type/bool-type.component'
import { CronTimeComponent} from './types/cron-time/cron-time.component'
import { FacebookTokenComponent } from './types/facebook-token/facebook-token.component'
import { FloatTypeComponent } from './types/float-type/float-type.component'
import { GoogleTokenComponent } from './types/google-token/google-token.component'

@Injectable()
export class TypeService {
  types: TypeSelector[] = []

  constructor() {
    this.add(IntegerTypeComponent, (i: Input) => i.isInteger())
    this.add(StringTypeComponent, (i: Input) => i.isString())
    this.add(StringTypeComponent, (i: Input) => i.isInterface())
    this.add(BoolTypeComponent, (i: Input) => i.isBool())
    this.add(FloatTypeComponent, (i: Input) => i.isFloat())
    this.add(CronTimeComponent, (i: Input) => i.type == "cron.Time")
    this.add(FacebookTokenComponent, (i: Input) => i.type == "facebook.Token")
    this.add(GoogleTokenComponent, (i: Input) => i.type == "google.Token")
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
