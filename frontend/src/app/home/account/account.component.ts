import {Component, OnInit, signal} from '@angular/core';
import {MatDivider} from '@angular/material/divider';
import {Router, RouterOutlet} from '@angular/router';
import {MatButton, MatIconButton} from '@angular/material/button';
import {AuthService} from '../../services/auth.service';
import {User} from '../../models/User';
import {MatFormField, MatInput, MatLabel, MatSuffix} from '@angular/material/input';
import {FormsModule} from '@angular/forms';
import {PasswordDirective} from 'primeng/password';
import {MatIcon} from '@angular/material/icon';

@Component({
  selector: 'app-account',
  imports: [
    MatFormField,
    MatInput,
    FormsModule,
    MatLabel,
    PasswordDirective,
    MatIcon,
    MatIconButton,
    MatSuffix,
    MatButton,
    MatDivider
  ],
  templateUrl: './account.component.html',
  standalone: true,
  styleUrl: './account.component.scss'
})
export class AccountComponent implements OnInit{
  hideOld = signal(true);
  hideNew = signal(true);
  hideNewRepeat = signal(true);
  user: User | undefined;
  oldPassword: string = '';
  newPassword: string = '';
  newPasswordRepeat: string = '';

  constructor(private router: Router, private authService: AuthService) {
  }

  ngOnInit() {
    this.user = this.authService.user
  }

  navigate(location: string) {
    this.router.navigate([location]);
  }

  clickOld(event: MouseEvent) {
    this.hideOld.set(!this.hideOld());
    event.stopPropagation();
  }

  clickNew(event: MouseEvent) {
    this.hideNew.set(!this.hideNew());
    event.stopPropagation();
  }

  clickNewRepeat(event: MouseEvent) {
    this.hideNewRepeat.set(!this.hideNewRepeat());
    event.stopPropagation();
  }
}
