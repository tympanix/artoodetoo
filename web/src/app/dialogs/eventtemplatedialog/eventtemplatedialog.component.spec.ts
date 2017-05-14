/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { EventTemplateDialog } from './eventtemplatedialog.component';

describe('EventTemplateDialogComponent', () => {
  let component: EventTemplateDialog;
  let fixture: ComponentFixture<EventTemplateDialog>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EventTemplateDialog ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EventTemplateDialog);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
