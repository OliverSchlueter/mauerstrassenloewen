import { Component } from '@angular/core';
import { MatCard, MatCardContent } from '@angular/material/card';
import { MatExpansionModule } from '@angular/material/expansion';
import { AuthService } from '../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-theory',
  templateUrl: './theory.component.html',
  styleUrl: './theory.component.scss',
  standalone: true,
  imports: [
    MatCard,
    MatCardContent,
    MatExpansionModule
  ],
})

export class TheoryComponent {
  constructor(private router: Router) {}

  navigate(path: string) {
    this.router.navigate([path]);
  }
}
