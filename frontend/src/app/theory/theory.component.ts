import { Component } from '@angular/core';
import {MatCard, MatCardContent} from '@angular/material/card';
import {MatButton} from '@angular/material/button';
import {AuthService} from '../services/auth.service';

@Component({
  selector: 'app-theory',
  imports: [
    MatCard,
    MatCardContent,
    MatButton
  ],
  templateUrl: './theory.component.html',
  standalone: true,
  styleUrl: './theory.component.scss'
})
export class TheoryComponent {


  constructor(private authService: AuthService) {
  }

  doSth() {

  }
}
