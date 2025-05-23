import { ComponentFixture, TestBed } from '@angular/core/testing';
import { TheoryComponent } from './theory.component';
import { OverlayComponent } from './overlay/overlay.component';

describe('TheoryComponent', () => {
  let component: TheoryComponent;
  let fixture: ComponentFixture<TheoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [TheoryComponent, OverlayComponent]
    }).compileComponents();

    fixture = TestBed.createComponent(TheoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
