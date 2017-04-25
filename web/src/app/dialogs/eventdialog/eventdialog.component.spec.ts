/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { EventDialog } from './eventdialog.component';

describe('EventdialogComponent', () => {
  let component: EventDialog;
  let fixture: ComponentFixture<EventDialog>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EventDialog ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EventDialog);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
