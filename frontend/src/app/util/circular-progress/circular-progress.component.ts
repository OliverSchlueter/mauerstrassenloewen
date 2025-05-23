import {Component, Input} from '@angular/core';

@Component({
  selector: 'app-circular-progress',
  imports: [],
  templateUrl: './circular-progress.component.html',
  standalone: true,
  styleUrl: './circular-progress.component.scss'
})
export class CircularProgressComponent {
  @Input() progress = 70; // Wert in Prozent
  readonly radius = 54;
  readonly circumference = 2 * Math.PI * this.radius;
}
