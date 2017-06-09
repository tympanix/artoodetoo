import { Pipe, PipeTransform } from '@angular/core';

import { Unit, } from '../../model';

@Pipe({
  name: 'unitIdPipe'
})
export class UnitIdPipePipe implements PipeTransform {

  transform(units: Unit[]): Unit[] {
    if (units == undefined) {
      return units
    }
    
    return units.sort((a,b) => {
      if (a.id > b.id) {
          return 1;
      }

      if (a.id < b.id) {
          return -1;
      }

      return 0;
    });
  }
}
