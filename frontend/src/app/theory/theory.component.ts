
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatCard, MatCardContent, MatCardFooter, MatCardHeader, MatCardTitle} from '@angular/material/card';
import {MatIcon} from '@angular/material/icon';
import {Router} from '@angular/router';

@Component({
  selector: 'app-theory',
  standalone: true,
  imports: [
    CommonModule,
    MatCard,
    MatCardFooter,
    MatCardContent,
    MatIcon,
    MatCardHeader,
    MatCardTitle
  ],
  templateUrl: './theory.component.html',
  styleUrls: ['./theory.component.scss']
})
export class TheoryComponent {
  modules = [
    {
      index: "1",
      title: "The Basics"
    },
    {
      index: "2",
      title: "The Market"
    },
    {
      index: "3",
      title: "The First Step"
    },
    {
      index: "4",
      title: "The Choice"
    },
    {
      index: "5",
      title: "The Golden Rule"
    },
    {
      index: "6",
      title: "The Strategy"
    },
    {
      index: "7",
      title: "The Borders"
    },

  ]

  constructor(private router: Router) {
  }

  navigate(index: string) {
    this.router.navigate(["module/"+index])
  }
}
