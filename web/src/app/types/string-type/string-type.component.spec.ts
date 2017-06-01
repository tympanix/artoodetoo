/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { StringTypeComponent } from './string-type.component';

describe('StringTypeComponent', () => {
  let component: StringTypeComponent;
  let fixture: ComponentFixture<StringTypeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ StringTypeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(StringTypeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
