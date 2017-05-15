/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { OptionDialog } from './optiondialog.component';

describe('OptionDialog', () => {
  let component: OptionDialog;
  let fixture: ComponentFixture<OptionDialog>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OptionDialog ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OptionDialog);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
