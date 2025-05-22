import {Component, inject} from '@angular/core';
import {InterviewComponent} from '../interview/interview.component';
import {User} from '../models/User';
import {MatTab, MatTabGroup} from '@angular/material/tabs';
import {MatCard, MatCardContent, MatCardFooter, MatCardHeader, MatCardTitle} from '@angular/material/card';
import {MatFormField, MatHint, MatInput, MatLabel} from '@angular/material/input';
import {FormsModule} from '@angular/forms';
import {MatButton, MatIconButton} from '@angular/material/button';
import {MatIcon} from '@angular/material/icon';
import {Router} from '@angular/router';
import {UserService} from '../services/user.service';
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-register',
  imports: [
    MatCard,
    MatCardTitle,
    MatCardHeader,
    MatCardContent,
    MatFormField,
    MatLabel,
    MatInput,
    FormsModule,
    MatCardFooter,
    MatButton,
    MatIcon,
    MatIconButton,
    MatHint
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.scss'
})
export class RegisterComponent {
  user= new User()
  repeatPassword = '';
  private _snackbar = inject(MatSnackBar)

  constructor(private router: Router, private userService: UserService) {
  }

  register() {
    this.userService.register(this.user).subscribe((response: any) => {
      console.log(response)
      if(response) {
        this.openSnackBar("Registration successfull", "close")
      }
      this.router.navigate(['login'])
    })
  }

  openSnackBar(message: string, action: string) {
    this._snackbar.open(message, action, {duration: 3000})
  }

  navigateLogin() {
    this.router.navigate(['login'])
  }
}
