
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { OverlayComponent } from './overlay/overlay.component';
import { MatCard } from '@angular/material/card';

@Component({
  selector: 'app-theory',
  standalone: true,        
  imports: [
    CommonModule,
    OverlayComponent,      
  ],
  templateUrl: './theory.component.html',
  styleUrls: ['./theory.component.scss']
})
export class TheoryComponent { }
