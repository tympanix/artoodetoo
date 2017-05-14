/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { TaskResolver } from './task-resolver.service';

describe('TaskResolver', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [TaskResolver]
    });
  });

  it('should ...', inject([TaskResolver], (service: TaskResolver) => {
    expect(service).toBeTruthy();
  }));
});
