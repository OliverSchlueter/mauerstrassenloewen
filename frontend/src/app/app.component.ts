import { Component } from '@angular/core';
import {HeaderComponent} from './header/header.component';
import {RouterOutlet} from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [HeaderComponent, RouterOutlet],
  templateUrl: './app.component.html',
  standalone: true,
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'frontend';
}
