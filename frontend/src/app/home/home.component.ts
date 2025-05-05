import { Component } from '@angular/core';
import {MatCard, MatCardContent, MatCardHeader, MatCardSubtitle, MatCardTitle} from '@angular/material/card';

@Component({
  selector: 'app-home',
  imports: [
    MatCard,
    MatCardTitle,
    MatCardContent,
    MatCardSubtitle,
    MatCardHeader
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss'
})
export class HomeComponent {

}
