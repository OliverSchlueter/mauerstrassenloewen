import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountLandingComponent } from './account-landing.component';

describe('AccountLandingComponent', () => {
  let component: AccountLandingComponent;
  let fixture: ComponentFixture<AccountLandingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AccountLandingComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AccountLandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
