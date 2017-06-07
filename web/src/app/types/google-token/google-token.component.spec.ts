import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GoogleTokenComponent } from './google-token.component';

describe('GoogleTokenComponent', () => {
  let component: GoogleTokenComponent;
  let fixture: ComponentFixture<GoogleTokenComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GoogleTokenComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GoogleTokenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
