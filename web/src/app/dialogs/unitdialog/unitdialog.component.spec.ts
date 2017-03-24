/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { UnitDialog } from './unitdialog.component';

describe('UnitdialogComponent', () => {
  let component: UnitDialog;
  let fixture: ComponentFixture<UnitDialog>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UnitDialog ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UnitDialog);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
