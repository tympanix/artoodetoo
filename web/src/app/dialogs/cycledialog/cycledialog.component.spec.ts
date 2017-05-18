/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { CycleDialog } from './cycledialog.component';

describe('CycleDialog', () => {
  let component: CycleDialog;
  let fixture: ComponentFixture<CycleDialog>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CycleDialog ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CycleDialog);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
