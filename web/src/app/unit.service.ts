import { Injectable } from '@angular/core';

import { Unit } from './unit';
import { UNITS } from './mock-units';

@Injectable()
export class UnitService {

  getUnits(): Promise<Unit[]> {
    return Promise.resolve(UNITS);
  }

  // See the "Take it slow" appendix
  getUnitsSlowly(): Promise<Unit[]> {
    return new Promise(resolve => {
      // Simulate server latency with 2 second delay
      setTimeout(() => resolve(this.getUnits()), 2000);
    });
  }
}
