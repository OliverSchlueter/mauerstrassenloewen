import {Component, OnInit} from '@angular/core';
import {MatButton} from "@angular/material/button";
import {MatCard} from "@angular/material/card";
import {MatFormField, MatInput, MatLabel, MatSuffix} from "@angular/material/input";
import {MatStep, MatStepLabel, MatStepper, MatStepperNext, MatStepperPrevious} from "@angular/material/stepper";
import {User} from '../models/User';
import {FormsModule} from '@angular/forms';
import {MatDatepicker, MatDatepickerInput, MatDatepickerToggle} from '@angular/material/datepicker';
import {MatOption, MatSelect} from '@angular/material/select';
import {Profile} from '../models/Profile';
import {MatSlider, MatSliderThumb} from '@angular/material/slider';
import {MatCheckbox} from '@angular/material/checkbox';
import {MatSlideToggle} from '@angular/material/slide-toggle';
import {UserService} from '../services/user.service';
import {Router} from '@angular/router';
import {AuthService} from '../services/auth.service';

@Component({
  selector: 'app-interview',
  imports: [
    MatButton,
    MatCard,
    MatFormField,
    MatInput,
    MatLabel,
    MatStep,
    MatStepLabel,
    MatStepper,
    MatStepperNext,
    MatStepperPrevious,
    FormsModule,
    MatDatepickerInput,
    MatDatepickerToggle,
    MatDatepicker,
    MatSelect,
    MatOption,
    MatSuffix,
    MatSlider,
    MatSliderThumb,
    MatCheckbox,
    MatSlideToggle
  ],
  templateUrl: './interview.component.html',
  standalone: true,
  styleUrl: './interview.component.scss'
})
export class InterviewComponent implements OnInit{
  user: User | undefined;
  userprofile: Profile = new Profile()


  constructor(private userService: UserService, private authService: AuthService, private router: Router) {
  }

  ngOnInit() {
    this.user = this.authService.user;
    if(this.user) {
      this.userprofile = this.user.profile;
    }
  }

  formatRiskAffinity(value: number): string {
    return value + "%"
  }

  save() {
    if(this.user) {
      this.user.profile = this.userprofile;
    }
    else {
      this.user = new User();
      this.user.profile = this.userprofile;
    }
    console.log(this.user)
    this.userService.updateUser(this.user).subscribe(response => {
      console.log("response: " + response);

    })
    this.router.navigate([''])
  }
}
