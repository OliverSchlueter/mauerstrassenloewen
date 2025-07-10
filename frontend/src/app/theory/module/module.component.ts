import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {NgIf} from '@angular/common';
import {MatIconButton} from '@angular/material/button';
import {MatIcon} from '@angular/material/icon';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatCard, MatCardContent, MatCardHeader} from '@angular/material/card';
import {MatDivider} from '@angular/material/divider';

@Component({
  selector: 'app-module',
  imports: [
    NgIf,
    MatIconButton,
    MatIcon,
    MatExpansionModule,
    MatCardHeader,
    MatCardContent,
    MatDivider,
    MatCard,
  ],
  templateUrl: './module.component.html',
  standalone: true,
  styleUrl: './module.component.scss'
})
export class ModuleComponent implements OnInit{
  moduleNo = ""
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

  constructor(private route: ActivatedRoute, private router: Router) {
  }

  ngOnInit() {
    this.moduleNo = this.route.snapshot.paramMap.get("moduleNo")!
  }

  navigateBack() {
    this.router.navigate(["/theory"]);
  }

  getHeader(moduleNo: string) {
    const number = <number><unknown>moduleNo;
    return this.modules[number-1].title
  }
}
