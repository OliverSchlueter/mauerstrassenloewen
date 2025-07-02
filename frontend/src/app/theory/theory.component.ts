
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatCard, MatCardContent, MatCardFooter, MatCardHeader, MatCardTitle} from '@angular/material/card';
import {MatIcon} from '@angular/material/icon';

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
      index: "Module 1",
      title: "The Basics"
    },
    {
      index: "Module 2",
      title: "The Market"
    },
    {
      index: "Module 3",
      title: "The First Step"
    },
    {
      index: "Module 4",
      title: "The Choice"
    },
    {
      index: "Module 5",
      title: "The Golden Rule"
    },
    {
      index: "Module 6",
      title: "The Strategy"
    },
    {
      index: "Module 7",
      title: "The Borders"
    },

  ]


}
