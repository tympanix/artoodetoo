/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { UnitService } from './unit.service';

describe('UnitService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [UnitService]
    });
  });

  it('should ...', inject([UnitService], (service: UnitService) => {
    expect(service).toBeTruthy();
  }));
});
