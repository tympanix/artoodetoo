/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { TypeService } from './type.service';

describe('TypeService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [TypeService]
    });
  });

  it('should ...', inject([TypeService], (service: TypeService) => {
    expect(service).toBeTruthy();
  }));
});
