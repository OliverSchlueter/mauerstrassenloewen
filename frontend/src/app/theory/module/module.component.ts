import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {NgIf} from '@angular/common';
import {MatIconButton} from '@angular/material/button';
import {MatIcon} from '@angular/material/icon';

@Component({
  selector: 'app-module',
  imports: [
    NgIf,
    MatIconButton,
    MatIcon
  ],
  templateUrl: './module.component.html',
  styleUrl: './module.component.scss'
})
export class ModuleComponent implements OnInit{
  moduleNo = ""

  constructor(private route: ActivatedRoute, private router: Router) {
  }

  ngOnInit() {
    this.moduleNo = this.route.snapshot.paramMap.get("moduleNo")!
  }

  navigateBack() {
    this.router.navigate(["/theory"]);
  }
}
