/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { TaskDialog } from './taskdialog.component';

describe('TaskDialog', () => {
  let component: TaskDialog;
  let fixture: ComponentFixture<TaskDialog>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TaskDialog ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TaskDialog);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
