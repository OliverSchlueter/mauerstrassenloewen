import {Component, Input, OnInit} from '@angular/core';
import {MatButton} from "@angular/material/button";
import {MatCard} from "@angular/material/card";
import {MatError, MatFormField, MatHint, MatInput, MatLabel, MatSuffix} from "@angular/material/input";
import {MatStep, MatStepLabel, MatStepper, MatStepperNext, MatStepperPrevious} from "@angular/material/stepper";
import {User} from '../models/User';
import {FormsModule} from '@angular/forms';
import {MatDatepicker, MatDatepickerInput, MatDatepickerToggle} from '@angular/material/datepicker';
import {MatOption, MatSelect} from '@angular/material/select';
import {Profile} from '../models/Profile';
import {MatIcon} from '@angular/material/icon';

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
    MatHint,
    MatError,
    MatIcon,
    MatSuffix
  ],
  templateUrl: './interview.component.html',
  standalone: true,
  styleUrl: './interview.component.scss'
})
export class InterviewComponent implements OnInit{
  @Input() user: User | undefined;
  userprofile: Profile = new Profile()

  constructor() {
  }

  ngOnInit() {
    if(this.user) {
      this.userprofile = this.user.profile;
    }
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
  }
}
